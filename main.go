package main

import (
	"fmt"
	"io"
	    "io/ioutil"
	"net/http"
	"strings"
)

const domain = "https://senchpimy.xyz"

func main() {
	fmt.Println("Server Start")
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasSuffix(path,".css"){
			w.Header().Set("Content-Type", "text/css")
			data, err := ioutil.ReadFile("style.css")
			if err != nil {
			    fmt.Println(err)
			}
			style:=string(data)
			fmt.Fprint(w,style)
		}else if strings.HasSuffix(path,".html") || path=="/"{
			html:=request(path)
			fmt.Fprint(w,html+"<meta name='viewport' content='width=device-widht, initial-scale=1.0'>")
		}else{
			html:=request(path)
			fmt.Fprint(w,html)
		}
	})
	fmt.Println("Server Start")

        http.ListenAndServe(":3002", nil)
}

 func request(url string)(foo string){
	resp, err := http.Get(domain+url)
	if err != nil {
		fmt.Println("erro get")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("erro read")
	}
	return string(body)
 }
