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

func splitPath(path string) (pathlist []string) {
	pathlist = append(pathlist, "/")
	tmp := strings.Split(path, "/")
	for i := 0; i < len(tmp); i++ {
		p := tmp[i]
		if p == "." {
			continue
		} else if p == ".." {
			if len(pathlist) > 0 {
				pathlist = pathlist[:len(pathlist)-1]
			}
		} else if p != "" {
			pathlist = append(pathlist, "/"+p)
		}
	}
	return pathlist
}

func joinPath(pathlist []string) (path string) {
	if len(pathlist) == 1 {
		return "/"
	}
	for _, p := range pathlist[1:] {
		path += p
	}
	return path
}

func errorPrint(err string) {
	fmt.Printf("[error!!!] %s\n", err)
	panic("error!")
}

func debugPrint(format string, a ...any) {
	fmt.Printf("[MYGIN-debug] "+format+"\n", a...)
}

func normalPrint(format string, a ...any) {
	fmt.Printf("[MYGIN] "+format+"\n", a...)
}
