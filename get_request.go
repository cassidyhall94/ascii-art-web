package main

import (
	"fmt"
	"log"
	"net/http"
)

// https://git.learn.01founders.co/root/public/src/branch/master/subjects/ascii-art-web/audit
// Are all the pages working? Does the project implement 404 status?
// Does the project handle HTTP status 400 - Bad Request?
// Does the project handle HTTP status 500 - Internal Server Errors?
// https://www.restapitutorial.com/httpstatuscodes.html

// in a browser type: localhost:8080 to see the webpage after entering (go run .) in the terminal

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "form.html")
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		font := r.FormValue("Font")
		input := r.FormValue("Enter Text:")

		fmt.Fprintf(w, "%s is a %s\n", input, font)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", process)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
