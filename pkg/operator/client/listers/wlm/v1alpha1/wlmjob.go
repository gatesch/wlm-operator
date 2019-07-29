// Copyright (c) 2019 Sylabs, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WlmJobLister helps list WlmJobs.
type WlmJobLister interface {
	// List lists all WlmJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WlmJob, err error)
	// WlmJobs returns an object that can list and get WlmJobs.
	WlmJobs(namespace string) WlmJobNamespaceLister
	WlmJobListerExpansion
}

// wlmJobLister implements the WlmJobLister interface.
type wlmJobLister struct {
	indexer cache.Indexer
}

// NewWlmJobLister returns a new WlmJobLister.
func NewWlmJobLister(indexer cache.Indexer) WlmJobLister {
	return &wlmJobLister{indexer: indexer}
}

// List lists all WlmJobs in the indexer.
func (s *wlmJobLister) List(selector labels.Selector) (ret []*v1alpha1.WlmJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WlmJob))
	})
	return ret, err
}

// WlmJobs returns an object that can list and get WlmJobs.
func (s *wlmJobLister) WlmJobs(namespace string) WlmJobNamespaceLister {
	return wlmJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WlmJobNamespaceLister helps list and get WlmJobs.
type WlmJobNamespaceLister interface {
	// List lists all WlmJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WlmJob, err error)
	// Get retrieves the WlmJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WlmJob, error)
	WlmJobNamespaceListerExpansion
}

// wlmJobNamespaceLister implements the WlmJobNamespaceLister
// interface.
type wlmJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WlmJobs in the indexer for a given namespace.
func (s wlmJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WlmJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WlmJob))
	})
	return ret, err
}

// Get retrieves the WlmJob from the indexer for a given namespace and name.
func (s wlmJobNamespaceLister) Get(name string) (*v1alpha1.WlmJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wlmjob"), name)
	}
	return obj.(*v1alpha1.WlmJob), nil
}
