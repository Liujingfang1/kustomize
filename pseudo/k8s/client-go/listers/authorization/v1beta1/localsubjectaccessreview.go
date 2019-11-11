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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "sigs.k8s.io/kustomize/pseudo/k8s/api/authorization/v1beta1"
	"sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/api/errors"
	"sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/labels"
	"sigs.k8s.io/kustomize/pseudo/k8s/client-go/tools/cache"
)

// LocalSubjectAccessReviewLister helps list LocalSubjectAccessReviews.
type LocalSubjectAccessReviewLister interface {
	// List lists all LocalSubjectAccessReviews in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.LocalSubjectAccessReview, err error)
	// LocalSubjectAccessReviews returns an object that can list and get LocalSubjectAccessReviews.
	LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewNamespaceLister
	LocalSubjectAccessReviewListerExpansion
}

// localSubjectAccessReviewLister implements the LocalSubjectAccessReviewLister interface.
type localSubjectAccessReviewLister struct {
	indexer cache.Indexer
}

// NewLocalSubjectAccessReviewLister returns a new LocalSubjectAccessReviewLister.
func NewLocalSubjectAccessReviewLister(indexer cache.Indexer) LocalSubjectAccessReviewLister {
	return &localSubjectAccessReviewLister{indexer: indexer}
}

// List lists all LocalSubjectAccessReviews in the indexer.
func (s *localSubjectAccessReviewLister) List(selector labels.Selector) (ret []*v1beta1.LocalSubjectAccessReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.LocalSubjectAccessReview))
	})
	return ret, err
}

// LocalSubjectAccessReviews returns an object that can list and get LocalSubjectAccessReviews.
func (s *localSubjectAccessReviewLister) LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewNamespaceLister {
	return localSubjectAccessReviewNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LocalSubjectAccessReviewNamespaceLister helps list and get LocalSubjectAccessReviews.
type LocalSubjectAccessReviewNamespaceLister interface {
	// List lists all LocalSubjectAccessReviews in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.LocalSubjectAccessReview, err error)
	// Get retrieves the LocalSubjectAccessReview from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.LocalSubjectAccessReview, error)
	LocalSubjectAccessReviewNamespaceListerExpansion
}

// localSubjectAccessReviewNamespaceLister implements the LocalSubjectAccessReviewNamespaceLister
// interface.
type localSubjectAccessReviewNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LocalSubjectAccessReviews in the indexer for a given namespace.
func (s localSubjectAccessReviewNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.LocalSubjectAccessReview, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.LocalSubjectAccessReview))
	})
	return ret, err
}

// Get retrieves the LocalSubjectAccessReview from the indexer for a given namespace and name.
func (s localSubjectAccessReviewNamespaceLister) Get(name string) (*v1beta1.LocalSubjectAccessReview, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("localsubjectaccessreview"), name)
	}
	return obj.(*v1beta1.LocalSubjectAccessReview), nil
}
