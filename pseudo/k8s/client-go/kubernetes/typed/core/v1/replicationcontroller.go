/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	autoscalingv1 "sigs.k8s.io/kustomize/pseudo/k8s/api/autoscaling/v1"
	v1 "sigs.k8s.io/kustomize/pseudo/k8s/api/core/v1"
	metav1 "sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/apis/meta/v1"
	types "sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/types"
	watch "sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/watch"
	scheme "sigs.k8s.io/kustomize/pseudo/k8s/client-go/kubernetes/scheme"
	rest "sigs.k8s.io/kustomize/pseudo/k8s/client-go/rest"
)

// ReplicationControllersGetter has a method to return a ReplicationControllerInterface.
// A group's client should implement this interface.
type ReplicationControllersGetter interface {
	ReplicationControllers(namespace string) ReplicationControllerInterface
}

// ReplicationControllerInterface has methods to work with ReplicationController resources.
type ReplicationControllerInterface interface {
	Create(*v1.ReplicationController) (*v1.ReplicationController, error)
	Update(*v1.ReplicationController) (*v1.ReplicationController, error)
	UpdateStatus(*v1.ReplicationController) (*v1.ReplicationController, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ReplicationController, error)
	List(opts metav1.ListOptions) (*v1.ReplicationControllerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ReplicationController, err error)
	GetScale(replicationControllerName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(replicationControllerName string, scale *autoscalingv1.Scale) (*autoscalingv1.Scale, error)

	ReplicationControllerExpansion
}

// replicationControllers implements ReplicationControllerInterface
type replicationControllers struct {
	client rest.Interface
	ns     string
}

// newReplicationControllers returns a ReplicationControllers
func newReplicationControllers(c *CoreV1Client, namespace string) *replicationControllers {
	return &replicationControllers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the replicationController, and returns the corresponding replicationController object, and an error if there is any.
func (c *replicationControllers) Get(name string, options metav1.GetOptions) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ReplicationControllers that match those selectors.
func (c *replicationControllers) List(opts metav1.ListOptions) (result *v1.ReplicationControllerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ReplicationControllerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested replicationControllers.
func (c *replicationControllers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a replicationController and creates it.  Returns the server's representation of the replicationController, and an error, if there is any.
func (c *replicationControllers) Create(replicationController *v1.ReplicationController) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Body(replicationController).
		Do().
		Into(result)
	return
}

// Update takes the representation of a replicationController and updates it. Returns the server's representation of the replicationController, and an error, if there is any.
func (c *replicationControllers) Update(replicationController *v1.ReplicationController) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(replicationController.Name).
		Body(replicationController).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *replicationControllers) UpdateStatus(replicationController *v1.ReplicationController) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(replicationController.Name).
		SubResource("status").
		Body(replicationController).
		Do().
		Into(result)
	return
}

// Delete takes name of the replicationController and deletes it. Returns an error if one occurs.
func (c *replicationControllers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *replicationControllers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched replicationController.
func (c *replicationControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("replicationcontrollers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

// GetScale takes name of the replicationController, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *replicationControllers) GetScale(replicationControllerName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(replicationControllerName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *replicationControllers) UpdateScale(replicationControllerName string, scale *autoscalingv1.Scale) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(replicationControllerName).
		SubResource("scale").
		Body(scale).
		Do().
		Into(result)
	return
}
