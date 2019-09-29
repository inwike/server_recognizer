package main

import (
	"fmt"
)

type Answer struct {
	emotion string
}

func init() {
	initCV()
	fmt.Println("init")
}

func main() {
	Start()
}
