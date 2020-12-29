/*
Copyright 2019 The KubeSphere Authors.

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

package application

import (
	"context"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"
	"kubesphere.io/kubesphere/pkg/apiserver/query"
	"kubesphere.io/kubesphere/pkg/simple/client/cacheclient"
	"path/filepath"
	appv1beta1 "sigs.k8s.io/application/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"testing"
)

var c client.Client

func createNamespace(name string, ctx context.Context) {
	namespace := &core.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	err := c.Create(ctx, namespace)
	if err != nil {
		klog.Error(err)
	}
}

func TestListApplications(t *testing.T) {
	e := &envtest.Environment{CRDDirectoryPaths: []string{filepath.Join("..", "..", "..", "..", "..", "config", "crds")}}

	if err := appv1beta1.AddToScheme(scheme.Scheme); err != nil {
		t.Fatalf("unable add APIs to scheme: %v", err)
	}

	cfg, err := e.Start()
	if err != nil {
		t.Fatal(err)
	}

	stopCh := make(chan struct{})
	if err := cacheclient.New(cfg, scheme.Scheme, stopCh); err != nil {
		t.Fatal(err)
	}
	ctx := context.TODO()
	c = cacheclient.Client
	createNamespace("bar", ctx)

	instance := &appv1beta1.Application{ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"}, Spec: appv1beta1.ApplicationSpec{}}

	if err := c.Create(ctx, instance); err != nil {
		t.Fatal(err)
	}

	getter := New()

	_, err = getter.List("bar", &query.Query{})
	if err != nil {
		t.Fatal(err)
	}

	_, err = getter.Get("bar", "foo")
	if err != nil {
		t.Fatal(err)
	}

}
