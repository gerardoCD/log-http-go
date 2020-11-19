package main

import (
	"fmt"
	"encoding/json"
	"reflect"
	lgs "github.com/gerardoCD/log-http-go"
)

func main() {
	fmt.Println("Hello, playground")

	lgs.Begin("Hola")
	lgs.Request("{\"hola\":4}")
	lgs.ReponseService(201,"{\"hola\":4}", "http://app/getElement", "Get Element")
	lgs.Reponse(200,"{\"hola\":4}")


	mapD := map[string]interface{}{"apple": 5, "lettuce": "Hola"}
	mapB, _ := json.Marshal(mapD)
	lgs.Reponse(200,string(mapB))

	fmt.Println(string(mapB))
	fmt.Println(reflect.TypeOf(mapD))

}
