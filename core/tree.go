// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package core

import (
	"fmt"
)

type (
	kind uint8
	// node is tree node
	node struct {
		kind     kind   // kind of node: common, param, wild
		fragment string // HTTP URL fragment
		handlers HandlersChain
		parent   *node
		children []*node
	}
	// tree is an alias for router
	tree struct {
		method string
		root   *node
	}
	// MethodForest is alias for MethodTrees
	MethodForest []*tree
)

const (
	// root kind
	root kind = iota
	// common kind
	cKind
	// param kind
	pKind
	// wild kind
	wKind
)

const (
	nullString = ""
	charSlash  = '/'
	strSlash   = "/"
	charColon  = ':'
	strColon   = ":"
	charStar   = '*'
	strStar    = "*"
)

// get method tree according to the HTTP method
func (forest MethodForest) get(method string) (router *tree, ok bool) {
	for _, tree := range forest {
		if tree.method == method {
			router = tree
			ok = true
			return
		}
	}
	return
}

// addRoute adds a node with the given handle to the path
func (t *tree) addRoute(path string, handlers HandlersChain) {
	validatePath(path)
	// purePath is path without any handle
	purePath := path
	if handlers == nil {
		panic(fmt.Sprintf("Adding route without handler function: %v", purePath))
	}
	t.insert(purePath, handlers)
}

// insert into trie tree
func (t *tree) insert(path string, handlers HandlersChain) {
	if t.root == nil {
		t.root = &node{
			kind:     root,
			fragment: strSlash,
		}
	}
	currNode := t.root
	fragments := splitPath(path)
	for i, fragment := range fragments {
		child := currNode.matchChild(fragment)
		if child == nil {
			child = &node{
				kind:     matchKind(fragment),
				fragment: fragment,
				parent:   currNode,
			}
			currNode.children = append(currNode.children, child)
		}
		if i == len(fragments)-1 {
			child.handlers = handlers
		}
		currNode = child
	}
}

// search matched node in trie tree, return nil when no matched
func (t *tree) search(path string, params *Params) *node {
	fragments := splitPath(path)
	var matchedNode *node
	currNode := t.root
	for i, fragment := range fragments {
		child := currNode.matchChildWithParam(fragment, params)
		if child == nil {
			return nil
		}
		if i == len(fragments)-1 {
			matchedNode = child
		}
		currNode = child
	}
	return matchedNode
}

// matchChild for the node children field with HTTP URL fragment
func (n *node) matchChild(fragment string) *node {
	for _, child := range n.children {
		if child.fragment == fragment {
			return child
		}
	}
	return nil
}

// matchChildWithParam for the param node
func (n *node) matchChildWithParam(fragment string, params *Params) *node {
	for _, child := range n.children {
		switch child.kind {
		case pKind:
			params.Set(child.fragment[1:], fragment)
			return child
		case wKind:
			return child
		case cKind:
			if child.fragment == fragment {
				return child
			}
		default:
			return nil
		}
	}
	return nil
}

// matchKind according to the HTTP URL fragment
func matchKind(fragment string) kind {
	switch fragment[0] {
	case charColon:
		return pKind
	case charStar:
		return wKind
	default:
		return cKind
	}
}
