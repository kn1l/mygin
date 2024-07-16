package mygin

import "strings"

type methodTree struct {
	method string
	root   *node
}

func newMethodTree(method string) methodTree {
	tree := methodTree{
		method: method,
		root: &node{
			path:     "",
			children: make([]*node, 0),
		},
	}
	return tree
}

type methodTrees []methodTree

func (trees *methodTrees) getMethodTree(method string) *node {
	for _, tree := range *trees {
		if tree.method == method {
			return tree.root
		}
	}
	newTree := newMethodTree(method)
	*trees = append(*trees, newTree)
	return newTree.root
}

type node struct {
	path     string
	children []*node
	handlers HandlerFuncChain
}

func newNode(path string) *node {
	n := &node{
		path:     path,
		children: nil,
	}
	return n
}

// search searchs the node by path
func (n *node) search(path string) *node {
	pathlist := strings.Split(path[1:], "/")
	subpath := "/" + pathlist[0]
	for _, child := range n.children {
		if child.path == subpath {
			if len(pathlist) == 1 {
				return child
			} else if len(pathlist) > 1 {
				return child.search("/" + strings.Join(pathlist[1:], "/"))
			}
		}
	}
	return nil
}

// insert builds the tree by path and returns the final child node
func (n *node) insert(path string) *node {
	pathlist := strings.Split(path[1:], "/")
	subpath := "/" + pathlist[0]

	child := newNode(subpath)
	n.children = append(n.children, child)
	if len(pathlist) == 1 {
		return child
	} else if len(pathlist) > 1 {
		return child.insert("/" + strings.Join(pathlist[1:], "/"))
	}
	return nil
}

func (n *node) setHandlers(handlers HandlerFuncChain) {

}
