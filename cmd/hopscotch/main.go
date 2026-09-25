package main

import (
	"log"

	"atomicptr.dev/hopscotch/pkg/hopscotch"
)

func main() {
	err := hopscotch.Run()
	if err != nil {
		log.Fatal(err)
	}
}
