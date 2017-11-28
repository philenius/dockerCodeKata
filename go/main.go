package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"os"
)

var counter = 0

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	http.Handle("/", router)
	log.Println("listerning on :8080")
	http.ListenAndServe(":8080", router)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

	headerMap := request.Header
	acc := headerMap["Accept"]
	enc := headerMap["Accept-Encoding"]
	lang := headerMap["Accept-Language"]
	cache := headerMap["Cache-Control"]
	con := headerMap["Connection"]
	host := headerMap["Host"]
	pragma := headerMap["Pragma"]
	userAgent := headerMap["User-Agent"]
	upgradeSSL := headerMap["Upgrade-Insecure-Requests"]
	log.Println("received request", acc, enc, lang, cache, con, host, pragma, upgradeSSL, userAgent)

	counter++
	log.Println("visitor count:", counter)

	envVar := os.Getenv("NAME")
	if len(envVar) == 0 {
		writer.Write([]byte("environment variable $NAME not set \n"))
	} else {
		writer.Write([]byte(fmt.Sprintf("hello, %s \n", envVar)))
	}
	writer.Write([]byte(fmt.Sprintf("visitor count: %d", counter)))
}
