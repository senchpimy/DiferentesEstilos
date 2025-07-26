package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const domain = "http://localhost:1313"

func main() {
	fmt.Println("Server Start")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasSuffix(path, ".css") {
			if path == "/ascci/style.css" {
				provideStyle(w, "style_ascci.css")
			} else {
				provideStyle(w, "style.css")
			}
		} else if strings.HasSuffix(path, ".html") || path == "/" {
			html, err := request(path)
			if err != nil {
				http.Error(w, "Error fetching HTML", http.StatusInternalServerError)
				return
			}
			fmt.Fprint(w, html+"<meta name='viewport' content='width=device-width, initial-scale=1.0'>")
		} else {
			html, err := request(path)
			if err != nil {
				http.Error(w, "Error fetching resource", http.StatusInternalServerError)
				return
			}
			fmt.Fprint(w, html)
		}
	})

	log.Fatal(http.ListenAndServe(":3002", nil))
}

func request(url string) (string, error) {
	resp, err := http.Get(domain + url)
	if err != nil {
		log.Println("Error making GET request:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return "", err
	}
	return string(body), nil
}

func provideStyle(w http.ResponseWriter, stylePath string) {
	w.Header().Set("Content-Type", "text/css")
	data, err := os.ReadFile(stylePath)
	if err != nil {
		log.Println("Error reading CSS file:", err)
		http.Error(w, "CSS not found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, string(data))
}
