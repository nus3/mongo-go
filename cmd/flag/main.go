package main

import (
	"flag"
	"fmt"
)

func main() {
	count := flag.Uint("count", 0, "Count(>= 0)")

	flag.Parse()
	fmt.Printf("param -count -> %d\n", *count)
}
