/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	quotav1alpha2 "kubesphere.io/kubesphere/pkg/apis/quota/v1alpha2"
	versioned "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	internalinterfaces "kubesphere.io/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "kubesphere.io/kubesphere/pkg/client/listers/quota/v1alpha2"
)

// ResourceQuotaInformer provides access to a shared informer and lister for
// ResourceQuotas.
type ResourceQuotaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.ResourceQuotaLister
}

type resourceQuotaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceQuotaInformer constructs a new informer for ResourceQuota type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceQuotaInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceQuotaInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceQuotaInformer constructs a new informer for ResourceQuota type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceQuotaInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuotaV1alpha2().ResourceQuotas().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuotaV1alpha2().ResourceQuotas().Watch(options)
			},
		},
		&quotav1alpha2.ResourceQuota{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceQuotaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceQuotaInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceQuotaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&quotav1alpha2.ResourceQuota{}, f.defaultInformer)
}

func (f *resourceQuotaInformer) Lister() v1alpha2.ResourceQuotaLister {
	return v1alpha2.NewResourceQuotaLister(f.Informer().GetIndexer())
}
