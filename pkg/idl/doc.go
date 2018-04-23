/*
Copyright 2018 The Kubernetes Authors.

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

/*
The IDL package describes comment directives that may be applied to API types.
*/
package idl

// ListType annotates a list to further describe its topology. It may
// have 3 possible values: "atomic", "map", or "set". Note that there is
// no default, and the generation step will fail if a list is found that
// is missing the tag.
//
// This tag MUST only be used on lists, or the generation step will
// fail.
//
// Atomic
//
// Example:
//  +listType=atomic
//
// Atomic lists will be entirely replaced when updated. This tag may be
// used on any type of list (struct, scalar, ...).
//
// Using this tag will generate the following OpenAPI extension:
//  "x-kubernetes-list-type": "atomic"
//
// Map
//
// Example:
//  +listType=map
//
// These lists are like maps in that their elements have a non-index key
// used to identify them. Order is preserved upon merge. Using the map
// tag on a list with non-struct elements will result in an error during
// the generation step.
//
// Using this tag will generate the following OpenAPI extension:
//  "x-kubernetes-list-type": "map"
//
// Set
//
// Example:
//  +listType=set
//
// Sets are lists that must not have multiple times the same value. Each
// value must be a scalar.
//
// Using this tag will generate the following OpenAPI extension:
//  "x-kubernetes-list-type": "set"
type ListType string
