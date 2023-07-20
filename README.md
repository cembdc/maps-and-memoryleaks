# maps-and-memoryleaks
Maps and Memory Leaks in Go

| Step      | `map[int][128]byte` | `map[int]*[128]byte` |
| ----------- | ----------- | ----------- | 
| Allocate an empty map      | 0 MB       | 0 MB |
| Add 1 million elements  | 461 MB        | 182 MB |
| Remove all the elements and run a GC   | 293 MB        | 38 MB |