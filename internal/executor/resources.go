package executor

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type Resources map[string]resource.Quantity

func FromResourceList(list v1.ResourceList) Resources {
	resources := make(Resources)
	for k, v := range list {
		resources[string(k)] = v.DeepCopy()
	}
	return resources
}

func (a Resources) Add(b Resources) {
	for k, v := range b {
		existing, ok := a[k]
		if ok {
			existing.Add(v)
			a[k] = existing
		} else {
			a[k] = v.DeepCopy()
		}
	}
}

func (a Resources) Sub(b Resources) {
	if b == nil {
		return
	}
	for k, v := range b {
		existing, ok := a[k]
		if ok {
			existing.Sub(v)
			a[k] = existing
		} else {
			cpy := v.DeepCopy()
			cpy.Neg()
			a[k] = cpy
		}
	}
}
