package main

import (
	"algorithms/artifacts"
	"algorithms/utils"
	"fmt"
	"net/http"         // for profiling: LINE1/3
	_ "net/http/pprof" // for profiling: LINE2/3
)

func main() {
	fmt.Println("hello algorithm")
	utils.Dummy()

	arr := []int{0, 34, 23, 45, 234, 15, 87, 20}

	fmt.Println(artifacts.HeapSort(arr))

	http.ListenAndServe(":8080", nil) // for profiling: LINE3/3
}
