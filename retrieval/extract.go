package main

import (
	"fmt"
	"log"
	"net/http"
)

func getZipFile() {
	fileObject, err := http.Get("https://www.kaggle.com/c/outbrain-click-prediction/download/page_views_sample.csv.zip")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileObject)
}

func extractZip() {

}

func main() {
	getZip()
}