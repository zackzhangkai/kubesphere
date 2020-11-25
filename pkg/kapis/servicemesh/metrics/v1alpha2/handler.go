/*
Copyright 2020 KubeSphere Authors

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

package v1alpha2

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/kiali/kiali/handlers"
	"io/ioutil"
	"k8s.io/klog"
	"kubesphere.io/kubesphere/pkg/api"
	"net/http"
)

// default jaeger query api endpoint address
var JaegerQueryUrl = "http://jaeger-query.istio-system.svc:16686"

// default kiali query api endpoint address
// To be compatible with the old settings,
// use the API of kiali POD when kialiQueryHost has been set in options,
// otherwise use the API of kiali code.
// var KialiQueryUrl = "http://kiali.istio-system:20001"
var KialiQueryUrl string

// Get app metrics
func getAppMetrics(request *restful.Request, response *restful.Response) {
	if len(KialiQueryUrl) == 0 {
		handlers.AppMetrics(request, response)
	} else {
		namespace := request.PathParameter("namespace")
		app := request.PathParameter("app")
		url := fmt.Sprintf("%s/kiali/api/namespaces/%s/apps/%s/metrics?%s", KialiQueryUrl, namespace, app, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy appMetrics request to %s", url)

		if err != nil {
			klog.Errorf("query kiali appMetrics failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

// Get workload metrics
func getWorkloadMetrics(request *restful.Request, response *restful.Response) {
	namespace := request.PathParameter("namespace")
	workload := request.PathParameter("workload")

	if len(namespace) > 0 && len(workload) > 0 {
		request.Request.URL.RawQuery = fmt.Sprintf("%s&namespaces=%s&workload=%s", request.Request.URL.RawQuery, namespace, workload)
	}

	if len(KialiQueryUrl) == 0 {
		handlers.WorkloadMetrics(request, response)
	} else {
		url := fmt.Sprintf("%s/kiali/api/namespaces/%s/workloads/%s/metrics?%s", KialiQueryUrl, namespace, workload, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy workload metrics request to %s", url)

		if err != nil {
			klog.Errorf("query kiali workload metrics failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

// Get service metrics
func getServiceMetrics(request *restful.Request, response *restful.Response) {
	if len(KialiQueryUrl) == 0 {
		handlers.ServiceMetrics(request, response)
	} else {
		namespace := request.PathParameter("namespace")
		service := request.PathParameter("service")
		url := fmt.Sprintf("%s/kiali/api/namespaces/%s/services/%s/metrics?%s", KialiQueryUrl, namespace, service, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy service metrics request to %s", url)

		if err != nil {
			klog.Errorf("query kiali service metrics failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

// Get namespace metrics
func getNamespaceMetrics(request *restful.Request, response *restful.Response) {
	if len(KialiQueryUrl) == 0 {
		handlers.NamespaceMetrics(request, response)
	} else {
		namespace := request.PathParameter("namespace")
		url := fmt.Sprintf("%s/kiali/api/namespaces/%s/metrics?%s", KialiQueryUrl, namespace, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy namespace metrics request to %s", url)

		if err != nil {
			klog.Errorf("query kiali namespace metrics failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

// Get service graph for namespace
func getNamespaceGraph(request *restful.Request, response *restful.Response) {
	namespace := request.PathParameter("namespace")

	if len(namespace) > 0 {
		request.Request.URL.RawQuery = fmt.Sprintf("%s&namespaces=%s", request.Request.URL.RawQuery, namespace)
	}

	if len(KialiQueryUrl) == 0 {
		handlers.GetNamespaceGraph(request, response)
	} else {
		url := fmt.Sprintf("%s/kiali/api/namespaces/graph?%s", KialiQueryUrl, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy namespace graph request to %s", url)

		if err != nil {
			klog.Errorf("query kiali namespace graph failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

// Deprecated
// It will cause erros as follows, when this API is called
//{
//    "error": "At least one namespace must be specified via the namespaces query parameter."
//}
// Get service graph for namespaces
func getNamespacesGraph(request *restful.Request, response *restful.Response) {
	handlers.GraphNamespaces(request, response)
}

// Get namespace health
func getNamespaceHealth(request *restful.Request, response *restful.Response) {
	if len(KialiQueryUrl) == 0 {
		handlers.NamespaceHealth(request, response)
	} else {
		namespace := request.PathParameter("namespace")
		url := fmt.Sprintf("%s/kiali/api/namespaces/%s/health?%s", KialiQueryUrl, namespace, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy namespace health request to %s", url)

		if err != nil {
			klog.Errorf("query kiali namespace health failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

// Get workload health
func getWorkloadHealth(request *restful.Request, response *restful.Response) {
	if len(KialiQueryUrl) == 0 {
		handlers.WorkloadHealth(request, response)
	} else {
		namespace := request.PathParameter("namespace")
		workload := request.PathParameter("workload")
		url := fmt.Sprintf("%s/kiali/api/namespaces/%s/workloads/%s/health?%s", KialiQueryUrl, namespace, workload, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy workload health request to %s", url)

		if err != nil {
			klog.Errorf("query kiali workload health failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

// Get app graph
func getAppGraph(request *restful.Request, response *restful.Response) {
	namespace := request.PathParameter("namespace")
	app := request.PathParameter("app")
	url := fmt.Sprintf("%s/kiali/api/namespaces/%s/applications/%s/graph?%s", KialiQueryUrl, namespace, app, request.Request.URL.RawQuery)
	resp, err := http.Get(url)
	klog.V(4).Infof("Proxy app graph request to %s", url)

	if err != nil {
		klog.Errorf("query kiali app graph failed with err %v", err)
		api.HandleInternalError(response, nil, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		klog.Errorf("read response error : %v", err)
		api.HandleInternalError(response, nil, err)
		return
	}

	// need to set header for proper response
	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(body)

	if err != nil {
		klog.Errorf("write response failed %v", err)
	}
}

// Get app health
func getAppHealth(request *restful.Request, response *restful.Response) {
	if len(KialiQueryUrl) == 0 {
		handlers.AppHealth(request, response)
	} else {
		namespace := request.PathParameter("namespace")
		app := request.PathParameter("app")
		url := fmt.Sprintf("%s/kiali/api/namespaces/%s/apps/%s/health?%s", KialiQueryUrl, namespace, app, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy app health request to %s", url)

		if err != nil {
			klog.Errorf("query kiali app health failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

// Get service health
func getServiceHealth(request *restful.Request, response *restful.Response) {
	if len(KialiQueryUrl) == 0 {
		handlers.ServiceHealth(request, response)
	} else {
		namespace := request.PathParameter("namespace")
		service := request.PathParameter("service")
		url := fmt.Sprintf("%s/kiali/api/namespaces/%s/services/%s/health?%s", KialiQueryUrl, namespace, service, request.Request.URL.RawQuery)
		resp, err := http.Get(url)
		klog.V(4).Infof("Proxy service health request to %s", url)

		if err != nil {
			klog.Errorf("query kiali service health failed with err %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			klog.Errorf("read response error : %v", err)
			api.HandleInternalError(response, nil, err)
			return
		}

		// need to set header for proper response
		response.Header().Set("Content-Type", "application/json")
		_, err = response.Write(body)

		if err != nil {
			klog.Errorf("write response failed %v", err)
		}
	}
}

func getServiceTracing(request *restful.Request, response *restful.Response) {
	namespace := request.PathParameter("namespace")
	service := request.PathParameter("service")

	serviceName := fmt.Sprintf("%s.%s", service, namespace)

	url := fmt.Sprintf("%s/api/traces?%s&service=%s", JaegerQueryUrl, request.Request.URL.RawQuery, serviceName)

	resp, err := http.Get(url)
	klog.V(4).Infof("Proxy trace request to %s", url)

	if err != nil {
		klog.Errorf("query jaeger failed with err %v", err)
		api.HandleInternalError(response, nil, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		klog.Errorf("read response error : %v", err)
		api.HandleInternalError(response, nil, err)
		return
	}

	// need to set header for proper response
	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(body)

	if err != nil {
		klog.Errorf("write response failed %v", err)
	}
}
