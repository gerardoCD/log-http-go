package log_http

import (
	"github.com/fatih/color"
	"encoding/json"
)


func TestFuntion() string {
	return "Hello, world."
}



func Begin(name string){
	color.Blue("-------------------- Begin Service  %s --------------------- \n\n\n\n", name)
}

func End(){
	color.Blue("-------------------- End Service ---------------------\n\n\n\n")
}

func Request(object interface{}){

	json,_ := json.Marshal(object)

	color.Blue("-------------------- Request Service ---------------------\n\n")
	color.Blue(" Request:  %s \n\n", string(json))
	color.Blue("----------------------------------------------------------\n\n\n\n")
}


func Reponse(statusCode int , object interface{}){

	json,_ := json.Marshal(object)
	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------\n\n")
		color.Green(" Status:  %d \n\n", statusCode)
		color.Green(" Reponse:  %s \n\n", string(json))
		color.Green("----------------------------------------------------------\n\n\n\n")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", string(json))
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", string(json))
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	}

	End()

}

func Info(info string){
	color.Blue(" --------------------  %s  --------------------\n\n", info)
}

func Error(err string){
	color.Magenta(" --------------------  %s  --------------------\n\n", err)
}

func RequestService(object interface{},  url string, name string){
	json,_ := json.Marshal(object)
		color.Green("-------------------- Request Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Request:  %s \n\n", string(json))
		color.Green("----------------------------------------------------------\n\n\n\n")

}



func RequestServiceString(json string,  url string, name string){
		color.Green("-------------------- Request Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Request:  %s \n\n", json)
		color.Green("----------------------------------------------------------\n\n\n\n")

}


func ReponseService(statusCode int ,object interface{},  url string, name string){
	json,_ := json.Marshal(object)
	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Status:  %d \n\n", statusCode)
		color.Green(" Reponse:  %s \n\n", string(json))
		color.Green("----------------------------------------------------------\n\n\n\n")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Name Service:  %s \n\n", name)
		color.Magenta(" Url:  %s \n\n", url)
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", string(json))
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Name Service:  %s \n\n", name)
		color.Magenta(" Url:  %s \n\n", url)
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", string(json))
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	}

}


func ReponseServiceString(statusCode int ,json string,  url string, name string){
	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Status:  %d \n\n", statusCode)
		color.Green(" Reponse:  %s \n\n", json)
		color.Green("----------------------------------------------------------\n\n\n\n")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Name Service:  %s \n\n", name)
		color.Magenta(" Url:  %s \n\n", url)
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", json)
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Name Service:  %s \n\n", name)
		color.Magenta(" Url:  %s \n\n", url)
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", json)
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	}

}


