package main

import (
	"fmt"
	lgs "github.com/gerardoCD/log-http-go"
)

func main() {
	fmt.Println("Hello, playground")

	lgs.Begin("Hola")
	lgs.Request("{\"hola\":4}")
}
