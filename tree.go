package mygin

import (
	"strings"
)

type methodTree struct {
	method string
	root   *node
}

func newMethodTree(method string) methodTree {
	tree := methodTree{
		method: method,
		root: &node{
			path:     "/",
			children: make([]*node, 0),
		},
	}
	return tree
}

type methodTrees []methodTree

func (trees *methodTrees) getMethodTree(method string) (root *node) {
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
	nodeType int
}

const (
	STATIC = iota
	PARAM
)

func newNode(path string, nodeType int) *node {
	n := &node{
		path:     path,
		children: nil,
		nodeType: nodeType,
	}
	return n
}

// search
func (n *node) search(c *Context, index int) *node {
	subpath := c.Pathlist[index]
	if subpath == n.path || n.path[1] == ':' {
		// dynamic route
		if len(n.path) > 3 {
			if n.path[1] == ':' {
				param := Param{
					Key:   n.path[2:],
					Value: subpath[1:],
				}
				c.Params = append(c.Params, param)
			}
		}

		if index == len(c.Pathlist)-1 {
			return n
		} else {
			if index+1 < len(c.Pathlist) {
				for _, child := range n.children {
					tmp := child.search(c, index+1)
					if tmp != nil {
						return tmp
					}
				}
			}
		}
	}

	if index == len(c.Pathlist)-1 {
		c.Params = nil
	}
	return nil
}

// insert builds the tree by path and returns the final child node
func (n *node) insert(path string) *node {
	pathlist := make([]string, 0)
	for _, p := range strings.Split(path, "/") {
		if p != "" {
			pathlist = append(pathlist, p)
		}
	}
	for _, p := range pathlist {
		subpath := "/" + p
		isFound := false
		for _, child := range n.children {
			if child.path == subpath {
				isFound = true
				n = child
				break
			}
		}
		if !isFound {
			switch p[0] {
			case ':':
				child := newNode(subpath, PARAM)
				n.children = append(n.children, child)
				n = child
			default:
				child := newNode(subpath, STATIC)
				n.children = append(n.children, child)
				n = child
			}

		}
	}
	return n
}
