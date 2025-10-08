package executor

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	informer "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
)

type ClusterStateManager struct {
	podInformer  informer.PodInformer
	nodeInformer informer.NodeInformer
	stopCh       chan struct{}
}

func NewClusterStateManager(ctx context.Context, kubernetesClient kubernetes.Interface) *ClusterStateManager {
	factory := informers.NewSharedInformerFactoryWithOptions(kubernetesClient, 0)

	stopCh := make(chan struct{})

	manager := &ClusterStateManager{
		podInformer:  factory.Core().V1().Pods(),
		nodeInformer: factory.Core().V1().Nodes(),
		stopCh:       stopCh,
	}

	manager.podInformer.Lister()
	manager.nodeInformer.Lister()

	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)

	go func() {
		select {
		case <-ctx.Done():
			// context expired or cancelled
			stopCh <- struct{}{} // write signal
		}
	}()

	return manager
}

func (cm *ClusterStateManager) GetClusterState() (*ClusterState, error) {
	nodes, err := cm.GetNodes()
	if err != nil {
		return nil, fmt.Errorf("getting nodes: %v", err)
	}

	fmt.Println(fmt.Sprintf("found %d nodes", len(nodes)))

	nodeMap := make(map[string]*v1.Node)
	for _, node := range nodes {
		nodeMap[node.Name] = node
	}

	pods, err := cm.podInformer.Lister().List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %v", err)
	}

	fmt.Println(fmt.Sprintf("found %d pods", len(pods)))

	nodeToPods := make(map[string][]*v1.Pod)
	for _, node := range nodes {
		nodeToPods[node.Name] = []*v1.Pod{}
	}

	for _, pod := range pods {
		if pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
			continue // Ignore terminal pods
		}

		node, ok := nodeMap[pod.Spec.NodeName]
		if !ok {
			fmt.Println(fmt.Sprintf("pod %s does not belong to a known node", pod.Name))
			continue
		}

		nodeToPods[node.Name] = append(nodeToPods[node.Name], pod)
	}

	totalAvailable := Resources{}

	for _, node := range nodes {
		nodeAvailable := FromResourceList(node.Status.Allocatable)
		nodePods, _ := nodeToPods[node.Name]
		nodeUsed := cm.SumAllocatedResources(nodePods)

		nodeAvailable.Sub(nodeUsed)
		totalAvailable.Add(nodeAvailable)
	}

	fmt.Println(fmt.Sprintf("total available cpu: %+v", totalAvailable["cpu"]))

	return &ClusterState{totalAvailable: totalAvailable}, nil
}

func (cm *ClusterStateManager) GetNodes() ([]*v1.Node, error) {
	nodes, err := cm.nodeInformer.Lister().List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("listing nodes: %v", err)
	}

	validNodes := make([]*v1.Node, 0)

	for _, node := range nodes {
		if node.Spec.Unschedulable {
			continue
		}

		for _, taint := range node.Spec.Taints {
			if taint.Effect == v1.TaintEffectNoSchedule {
				continue
			}
		}

		validNodes = append(validNodes, node)
	}

	return validNodes, nil
}

func (cm *ClusterStateManager) SumAllocatedResources(pods []*v1.Pod) Resources {
	allocations := map[string]resource.Quantity{}
	for _, pod := range pods {
		for _, initContainer := range pod.Spec.InitContainers {
			for k, v := range initContainer.Resources.Requests {
				existing, exists := allocations[k.String()]

				if exists {
					existing.Add(v)
					allocations[k.String()] = existing
				} else {
					allocations[k.String()] = v.DeepCopy()
				}
			}
		}
		for _, container := range pod.Spec.Containers {
			for k, v := range container.Resources.Requests {
				existing, exists := allocations[k.String()]

				if exists {
					existing.Add(v)
					allocations[k.String()] = existing
				} else {
					allocations[k.String()] = v.DeepCopy()
				}
			}
		}

	}
	return allocations
}
