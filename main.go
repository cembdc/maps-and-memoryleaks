package main

import (
	"fmt"
	"runtime"
)

func main() {
	map_with_value()
	//map_with_pointer()
}

func map_with_value() {

	fmt.Println("Map with value")
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc("After m is allocated")

	for i := 0; i < n; i++ {
		m[i] = [128]byte{}
	}
	printAlloc("After we add 1 million elements")

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc("After we remove 1 million elements")
	runtime.KeepAlive(m)
}

func map_with_pointer() {

	fmt.Println("Map with pointer")
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc("After m is allocated")

	for i := 0; i < n; i++ {
		m[i] = &[128]byte{}
	}
	printAlloc("After we add 1 million elements")

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc("After we remove 1 million elements")
	runtime.KeepAlive(m)
}

func printAlloc(str string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB   <-- %s\n", m.Alloc/1024, str)
}
