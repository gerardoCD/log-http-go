package log_http

import (
	"github.com/fatih/color"
)


func TestFuntion() string {
	return "Hello, world."
}



func Begin(name string){
	color.Blue("-------------------- Begin Service  %s --------------------- \n", name)
}

func end(){
	color.Blue("-------------------- End Service ---------------------\n")
}

func Request(json string){
	color.Blue("-------------------- Request Service ---------------------\n")
	color.Blue(" Request:  %s \n", json)
	color.Blue("----------------------------------------------------------\n")
}


func Reponse(statusCode int , json string){


	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------\n")
		color.Green(" Status:  %d ", statusCode)
		color.Green(" Reponse:  %s ", json)
		color.Green("----------------------------------------------------------\n")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------\n")
		color.Magenta(" Status:  %d ", statusCode)
		color.Magenta(" Reponse:  %s ", json)
		color.Magenta("----------------------------------------------------------\n")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------\n")
		color.Magenta(" Status:  %d ", statusCode)
		color.Magenta(" Reponse:  %s ", json)
		color.Magenta("----------------------------------------------------------\n")
	}

	end()

}



func ReponseService(statusCode int ,json string,  url string, name string){

	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------\n")
		color.Green(" Name Service:  %s ", name)
		color.Green(" Url:  %s ", url)
		color.Green(" Status:  %d ", statusCode)
		color.Green(" Reponse:  %s ", json)
		color.Green("----------------------------------------------------------\n")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------\n")
		color.Magenta(" Name Service:  %s ", name)
		color.Magenta(" Url:  %s ", url)
		color.Magenta(" Status:  %d ", statusCode)
		color.Magenta(" Reponse:  %s ", json)
		color.Magenta("----------------------------------------------------------\n")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------\n")
		color.Magenta(" Name Service:  %s ", name)
		color.Magenta(" Url:  %s ", url)
		color.Magenta(" Status:  %d ", statusCode)
		color.Magenta(" Reponse:  %s ", json)
		color.Magenta("----------------------------------------------------------\n")
	}

}
