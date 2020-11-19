package log_http

import (
	"github.com/fatih/color"
)


func TestFuntion() string {
	return "Hello, world."
}



func Begin(name string){
	color.Blue("-------------------- Begin Service  %s --------------------- \n\n\n\n", name)
}

func end(){
	color.Blue("-------------------- End Service ---------------------\n\n\n\n")
}

func Request(json string){
	color.Blue("-------------------- Request Service ---------------------\n\n")
	color.Blue(" Request:  %s \n\n", json)
	color.Blue("----------------------------------------------------------\n\n\n\n")
}


func Reponse(statusCode int , json string){


	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------\n\n")
		color.Green(" Status:  %d \n\n", statusCode)
		color.Green(" Reponse:  %s \n\n", json)
		color.Green("----------------------------------------------------------\n\n\n\n")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", json)
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", json)
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	}

	end()

}



func ReponseService(statusCode int ,json string,  url string, name string){

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
