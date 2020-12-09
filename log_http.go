package lgs

import (
	"github.com/fatih/color"
	"encoding/json"
	"encoding/xml"
	"github.com/go-xmlfmt/xmlfmt"
	"bytes"
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

	json,_ := json.MarshalIndent(object, "", "\t")

	color.Blue("-------------------- Request Service ---------------------\n\n")
	color.Blue(" Request:  %s \n\n", string(json))
	color.Blue("----------------------------------------------------------\n\n\n\n")
}


func Reponse(statusCode int , object interface{}){

	json,_ := json.MarshalIndent(object, "", "\t")
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
	json,_ := json.MarshalIndent(object, "", "\t")
		color.Green("-------------------- Request Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Request:  %s \n\n", string(json))
		color.Green("----------------------------------------------------------\n\n\n\n")

}



func RequestServiceXML(object interface{},  url string, name string){
	json,_ := xml.MarshalIndent(object, "", "\t")
		color.Green("-------------------- Request Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Request:  %s \n\n", string(json))
		color.Green("----------------------------------------------------------\n\n\n\n")

}


func RequestServiceString(json string,  url string, name string){
	// jsonBytes := []byte(json)
	// jsonPretty,_ := prettyprint(jsonBytes)
		color.Green("-------------------- Request Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Request:  %s \n\n", json)
		color.Green("----------------------------------------------------------\n\n\n\n")

}

func RequestServiceStringXML(xml string,  url string, name string){
	xmlPretty := xmlfmt.FormatXML(xml, "\t", "  ")
		color.Green("-------------------- Request Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Request:  %s \n\n", xmlPretty)
		color.Green("----------------------------------------------------------\n\n\n\n")

}


func ReponseService(statusCode int ,object interface{},  url string, name string){
	json,_ := json.MarshalIndent(object, "", "\t")
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


func ReponseServiceXML(statusCode int ,object interface{},  url string, name string){
	json,_ := xml.MarshalIndent(object, "", "\t")
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


func ReponseServiceStringXML(statusCode int ,xml string,  url string, name string){
	xmlPretty := xmlfmt.FormatXML(xml, "\t", "  ")
	if (statusCode == 200){
		color.Green("-------------------- Reponse Service ---------------------\n\n")
		color.Green(" Name Service:  %s \n\n", name)
		color.Green(" Url:  %s \n\n", url)
		color.Green(" Status:  %d \n\n", statusCode)
		color.Green(" Reponse:  %s \n\n", xmlPretty)
		color.Green("----------------------------------------------------------\n\n\n\n")
	} else if ( statusCode >= 201 && statusCode <= 299 ) {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Name Service:  %s \n\n", name)
		color.Magenta(" Url:  %s \n\n", url)
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", xmlPretty)
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	} else {
		color.Magenta("-------------------- Reponse Service ---------------------\n\n")
		color.Magenta(" Name Service:  %s \n\n", name)
		color.Magenta(" Url:  %s \n\n", url)
		color.Magenta(" Status:  %d \n\n", statusCode)
		color.Magenta(" Reponse:  %s \n\n", xmlPretty)
		color.Magenta("----------------------------------------------------------\n\n\n\n")
	}

}


func prettyprint(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
}


