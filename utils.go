package mygin

import "fmt"

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

func errorPrint(err string) {
	fmt.Printf("error!!!: %s\n", err)
	panic("error!")
}
