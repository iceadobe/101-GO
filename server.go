package main

import (
	handlers "./mypackage"
	"log"
	"net/http"
)

func main() {
	// HTTP path: http://localhost:8080/view/testpage GET
	http.HandleFunc("/view/", handlers.ViewHandler)
	// HTTP path: http://localhost:8080/save/testpage POST (with form data)
	http.HandleFunc("/edit/", handlers.EditHandler)
	// HTTP /save is automatically called from /edit/ page to save the file.
	http.HandleFunc("/save/", handlers.SaveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
