/*
Copyright The Flagger Authors.

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

package v1

import (
	v1 "github.com/weaveworks/flagger/pkg/apis/gloo/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UpstreamGroupLister helps list UpstreamGroups.
type UpstreamGroupLister interface {
	// List lists all UpstreamGroups in the indexer.
	List(selector labels.Selector) (ret []*v1.UpstreamGroup, err error)
	// UpstreamGroups returns an object that can list and get UpstreamGroups.
	UpstreamGroups(namespace string) UpstreamGroupNamespaceLister
	UpstreamGroupListerExpansion
}

// upstreamGroupLister implements the UpstreamGroupLister interface.
type upstreamGroupLister struct {
	indexer cache.Indexer
}

// NewUpstreamGroupLister returns a new UpstreamGroupLister.
func NewUpstreamGroupLister(indexer cache.Indexer) UpstreamGroupLister {
	return &upstreamGroupLister{indexer: indexer}
}

// List lists all UpstreamGroups in the indexer.
func (s *upstreamGroupLister) List(selector labels.Selector) (ret []*v1.UpstreamGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.UpstreamGroup))
	})
	return ret, err
}

// UpstreamGroups returns an object that can list and get UpstreamGroups.
func (s *upstreamGroupLister) UpstreamGroups(namespace string) UpstreamGroupNamespaceLister {
	return upstreamGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UpstreamGroupNamespaceLister helps list and get UpstreamGroups.
type UpstreamGroupNamespaceLister interface {
	// List lists all UpstreamGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.UpstreamGroup, err error)
	// Get retrieves the UpstreamGroup from the indexer for a given namespace and name.
	Get(name string) (*v1.UpstreamGroup, error)
	UpstreamGroupNamespaceListerExpansion
}

// upstreamGroupNamespaceLister implements the UpstreamGroupNamespaceLister
// interface.
type upstreamGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all UpstreamGroups in the indexer for a given namespace.
func (s upstreamGroupNamespaceLister) List(selector labels.Selector) (ret []*v1.UpstreamGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.UpstreamGroup))
	})
	return ret, err
}

// Get retrieves the UpstreamGroup from the indexer for a given namespace and name.
func (s upstreamGroupNamespaceLister) Get(name string) (*v1.UpstreamGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("upstreamgroup"), name)
	}
	return obj.(*v1.UpstreamGroup), nil
}
