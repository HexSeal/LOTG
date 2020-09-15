package main

import (
	"fmt"
	"log"
	"lotgPackage/translate"
)

func main() {
	txt, err := translate.TranslateText("es", "My name is Inigo Montoya, you killed my father, prepare to die!")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(txt)
}
