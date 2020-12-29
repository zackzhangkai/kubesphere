package cacheclient

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var Client client.Client

func New(config *rest.Config, sch *runtime.Scheme, stopCh <-chan struct{}) error {

	// setup controller-runtime cache
	ce, err := cache.New(config, cache.Options{Scheme: sch})
	if err != nil {
		klog.Fatalf("unable to set up controller runtime cache: %v", err)
		return err
	}

	// cache resources
	go ce.Start(stopCh)
	ce.WaitForCacheSync(stopCh)

	// set up controller-runtime client
	c, err := client.New(config, client.Options{Scheme: sch})
	if err != nil {
		klog.Fatalf("unable to set up controller runtime client: %v", err)
		return err
	}

	Client = &client.DelegatingClient{
		Reader: &client.DelegatingReader{
			CacheReader:  ce,
			ClientReader: c,
		},
		Writer:       c,
		StatusClient: c,
	}

	return nil
}
