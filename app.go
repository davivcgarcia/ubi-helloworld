package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	prefixText := `<!DOCTYPE html>
<html lang="en">
<head>
<title>Simple Demo App</title>
</head>
<body>
<h1>Simple Demo App</h1>
<p>This app dumps container runtime information:</p>
`
	sufixText := `</body>
</html>`

	var content string

	content += "<h2>Request Headers</h2>\n"
	for key, value := range r.Header {
		content += fmt.Sprintf("<p><b>%q:</b> %q</p>\n", key, value)
	}

	content += "<h2>Environment Variables</h2>\n"
	for _, text := range os.Environ() {
		text := strings.Split(text, "=")
		content += fmt.Sprintf("<p><b>%q:</b> %q</p>\n", text[0], text[1])
	}

	fmt.Fprint(w, prefixText+content+sufixText)
	log.Println("Request received!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Simple demo app started!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
