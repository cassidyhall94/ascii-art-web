package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// https://git.learn.01founders.co/root/public/src/branch/master/subjects/ascii-art-web/audit
// Are all the pages working? Does the project implement 404 status?
// Does the project handle HTTP status 400 - Bad Request?
// Does the project handle HTTP status 500 - Internal Server Errors?
// https://www.restapitutorial.com/httpstatuscodes.html
// in a browser type: localhost:8080 to see the webpage after entering (go run .) in the terminal
func main() {
	http.HandleFunc("/", process)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		if r.URL.Path != "/ascii-art" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}
	}
	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		input := r.FormValue("input")
		banner := r.FormValue("Font")

		fmt.Fprintf(w, "%s\n", AsciiArt(input, banner))

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func AsciiArt(str string, banner string) string {
	string := str

	previous := 'a'
	manylines := false
	for _, v := range string {
		if v == 'n' && previous == '\\' {
			manylines = true
		}
		previous = v
	}

	result := ""
	if manylines {
		args := strings.Split(string, "\\n")
		for _, word := range args {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					result += ReturnLine((1 + int(char-' ')*9 + i), banner)
				}
				result = result + "\n"
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for _, char := range string {
				result += ReturnLine((1 + int(char-' ')*9 + i), banner)
			}
			result = result + "\n"
		}
	}
	return result
}

func ReturnLine(num int, banner string) string {
	string := ""

	f, e := os.Open(banner) // add variable for banner
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()

	f.Seek(0, 0)
	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		string, _ = content.ReadString('\n')
	}
	string = strings.TrimSuffix(string, "\n")
	return string
}
