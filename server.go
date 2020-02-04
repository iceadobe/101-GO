package main

import (
	handlers "./mypackage"
	"log"
	"net/http"
)

func main() {
	/* // Reading with file and returning to std.out
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.Save()
	p2, _ := LoadPage("TestPage")
	fmt.Println(string(p2.Body))
	*/
	/*	// Simple Handler func
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	*/
	/*	// View Handler
	http.HandleFunc("/views/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
*/

	// Create a wiki page
	// HTTP path: http://localhost:8080/view/testpage GET
	http.HandleFunc("/view/", handlers.ViewHandler)
	// HTTP path: http://localhost:8080/save/testpage POST (with form data)
	http.HandleFunc("/edit/", handlers.EditHandler)
	// HTTP /save is automatically called from /edit/ page to save the file.
	http.HandleFunc("/save/", handlers.SaveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
