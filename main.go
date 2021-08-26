package main

import (
	"flag"
	"fmt"
	buildTools "leetcode/buildTools"
)

func main() {
	buildPtr := flag.String("build", "performance-table", "a string")
	flag.Parse()

	// performance table build
	fmt.Println("build: ", *buildPtr)
	switch *buildPtr {
	case "performance-table":
		buildTools.BuildPerformanceTable()
	default:
	}
}
