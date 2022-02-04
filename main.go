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
// in a browser type: localhost:8080 to see the webpage after entering (go run .) in the terminal
func main() {
	http.HandleFunc("/", process)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		if r.URL.Path != "/ascii-art" {
			http.Error(w, "404 Status not found", http.StatusNotFound)
			return
		}
	}
	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "templates/form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // if there is an error it returns bad request 400
			return
		}

		input := r.FormValue("input")
		banner := r.FormValue("Banner")
		if len(input) == 0 {
			http.Error(w, "Input should not be empty: Bad Request 400", http.StatusBadRequest) // if there is an error it returns bad request 400
			return
		}
		response, err := AsciiArt(input, banner) // return the error if the banner not found.
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Error(w, "No such file or directory: Internal Server Error 500", http.StatusInternalServerError)
			return
		}

		_, _ = w.Write([]byte(response)) // Write returns the response with a 200 status code in the header as this is built into the Write function

	default:
		http.Error(w, "Sorry, only GET and POST methods are supported.", http.StatusUnsupportedMediaType) // this will return the right status code
	}
}

func AsciiArt(str string, banner string) (string, error) {
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
					x, err := ReturnLine((1 + int(char-' ')*9 + i), banner)
					if err != nil {
						return "", err
					}
					result += x
				}
				result = result + "\n"
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for _, char := range string {
				x, err := ReturnLine((1 + int(char-' ')*9 + i), banner)
				if err != nil {
					return "", err
				}
				result += x
			}
			result = result + "\n"
		}
	}
	return result, nil
}

func ReturnLine(num int, banner string) (string, error) {
	entry := ""

	f, err := os.Open(banner)
	if err != nil {
		return "", err
	}
	defer f.Close()

	f.Seek(0, 0)
	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		entry, _ = content.ReadString('\n')
	}
	entry = strings.TrimSuffix(entry, "\n")
	return entry, nil
}
