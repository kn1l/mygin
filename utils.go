package mygin

import (
	"fmt"
	"strings"
)

func resolveAddress(addr ...string) string {
	switch len(addr) {
	case 0:
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

func splitPath(path string) []string {
	pathlist := make([]string, 0)
	pathlist = append(pathlist, "/")
	for _, p := range strings.Split(path, "/") {
		if p != "" {
			pathlist = append(pathlist, "/"+p)
		}
	}
	return pathlist
}

func errorPrint(err string) {
	fmt.Printf("error!!!: %s\n", err)
	panic("error!")
}
