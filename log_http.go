package log_http

import (
	"github.com/fatih/color"
)


func TestFuntion() string {
	return "Hello, world."
}



func Begin(name string){
	color.Blue("-------------------- Begin Service  %s ---------------------", name)
}

func end(){
	color.Blue("-------------------- End Service ---------------------")
}

func Request(json string){
	color.Blue("-------------------- Request Service ---------------------")
	color.Blue(" Request:  %s ", json)
	color.Blue("----------------------------------------------------------")
}


func Reponse(statusCode int , json string){


	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------")
		color.Green(" Status:  %s ", statusCode)
		color.Green(" Reponse:  %s ", json)
		color.Green("-------------------- %s ---------------------", json)
		color.Green("----------------------------------------------------------")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------")
		color.Magenta(" Status:  %s ", statusCode)
		color.Magenta(" Reponse:  %s ", json)
		color.Magenta("-------------------- %s ---------------------", json)
		color.Magenta("----------------------------------------------------------")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------")
		color.Magenta(" Status:  %s ", statusCode)
		color.Magenta(" Reponse:  %s ", json)
		color.Magenta("-------------------- %s ---------------------", json)
		color.Magenta("----------------------------------------------------------")
	}

	end()

}



func ReponseService(statusCode int ,json string,  url string, name string){

	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------")
		color.Green(" Name Service:  %s ", name)
		color.Green(" Url:  %s ", url)
		color.Green(" Status:  %s ", statusCode)
		color.Green(" Reponse:  %s ", json)
		color.Green("-------------------- %s ---------------------", json)
		color.Green("----------------------------------------------------------")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------")
		color.Magenta(" Name Service:  %s ", name)
		color.Magenta(" Url:  %s ", url)
		color.Magenta(" Status:  %s ", statusCode)
		color.Magenta(" Reponse:  %s ", json)
		color.Magenta("-------------------- %s ---------------------", json)
		color.Magenta("----------------------------------------------------------")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------")
		color.Magenta(" Name Service:  %s ", name)
		color.Magenta(" Url:  %s ", url)
		color.Magenta(" Status:  %s ", statusCode)
		color.Magenta(" Reponse:  %s ", json)
		color.Magenta("-------------------- %s ---------------------", json)
		color.Magenta("----------------------------------------------------------")
	}

}
