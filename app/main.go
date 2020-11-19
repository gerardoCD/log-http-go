package main

import (
	"fmt"
	lgs "github.com/gerardoCD/log-http-go"
)

func main() {
	fmt.Println("Hello, playground")

	lgs.Begin("Hola")
	lgs.Request("{\"hola\":4}")
	lgs.ReponseService(201,"{\"hola\":4}", "http://app/getElement", "Get Element")
	lgs.Reponse(200,"{\"hola\":4}")
}
