// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/sjenning/client-go/quota/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterResourceQuotas returns a ClusterResourceQuotaInformer.
	ClusterResourceQuotas() ClusterResourceQuotaInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterResourceQuotas returns a ClusterResourceQuotaInformer.
func (v *version) ClusterResourceQuotas() ClusterResourceQuotaInformer {
	return &clusterResourceQuotaInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
