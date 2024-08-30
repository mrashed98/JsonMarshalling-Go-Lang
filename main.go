package main

import (
	"fmt"
	m "jsonMarshalling/marshall"
	unm "jsonMarshalling/unmarshall"
)

func main() {
	fmt.Printf("UnMarshalling API\n\n")
	unm.TryUnmarshall()
	fmt.Printf("\n\nMarhsalling Struct\n\n")
	m.TryMarshall()
}
