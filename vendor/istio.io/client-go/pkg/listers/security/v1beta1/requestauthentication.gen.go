// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RequestAuthenticationLister helps list RequestAuthentications.
type RequestAuthenticationLister interface {
	// List lists all RequestAuthentications in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.RequestAuthentication, err error)
	// RequestAuthentications returns an object that can list and get RequestAuthentications.
	RequestAuthentications(namespace string) RequestAuthenticationNamespaceLister
	RequestAuthenticationListerExpansion
}

// requestAuthenticationLister implements the RequestAuthenticationLister interface.
type requestAuthenticationLister struct {
	indexer cache.Indexer
}

// NewRequestAuthenticationLister returns a new RequestAuthenticationLister.
func NewRequestAuthenticationLister(indexer cache.Indexer) RequestAuthenticationLister {
	return &requestAuthenticationLister{indexer: indexer}
}

// List lists all RequestAuthentications in the indexer.
func (s *requestAuthenticationLister) List(selector labels.Selector) (ret []*v1beta1.RequestAuthentication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.RequestAuthentication))
	})
	return ret, err
}

// RequestAuthentications returns an object that can list and get RequestAuthentications.
func (s *requestAuthenticationLister) RequestAuthentications(namespace string) RequestAuthenticationNamespaceLister {
	return requestAuthenticationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RequestAuthenticationNamespaceLister helps list and get RequestAuthentications.
type RequestAuthenticationNamespaceLister interface {
	// List lists all RequestAuthentications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.RequestAuthentication, err error)
	// Get retrieves the RequestAuthentication from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.RequestAuthentication, error)
	RequestAuthenticationNamespaceListerExpansion
}

// requestAuthenticationNamespaceLister implements the RequestAuthenticationNamespaceLister
// interface.
type requestAuthenticationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RequestAuthentications in the indexer for a given namespace.
func (s requestAuthenticationNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.RequestAuthentication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.RequestAuthentication))
	})
	return ret, err
}

// Get retrieves the RequestAuthentication from the indexer for a given namespace and name.
func (s requestAuthenticationNamespaceLister) Get(name string) (*v1beta1.RequestAuthentication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("requestauthentication"), name)
	}
	return obj.(*v1beta1.RequestAuthentication), nil
}