package cacheclient

import (
	"gotest.tools/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"kubesphere.io/kubesphere/pkg/apis"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"testing"
)

var cfg *rest.Config

var sch *runtime.Scheme

func TestClient(t *testing.T) {
	e := &envtest.Environment{}

	if err := apis.AddToScheme(scheme.Scheme); err != nil {
		t.Fatalf("unable add APIs to scheme: %v", err)
	}

	cfg, err := e.Start()
	if err != nil {
		t.Fatal(err)
	}

	stopCh := make(chan struct{})
	if err := New(cfg, scheme.Scheme, stopCh); err != nil {
		t.Fatal(err)
	}

	assert.Assert(t, Client != nil)
}
