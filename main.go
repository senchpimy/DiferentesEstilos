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
		//fmt.Fprint(w,request("/"))
		path := r.URL.Path
		if strings.HasSuffix(path,".css"){
			w.Header().Set("Content-Type", "text/css")
			data, err := ioutil.ReadFile("style.css")
			if err != nil {
			    fmt.Println(err)
			}
			style:=string(data)
			fmt.Fprint(w,style)
		}else{
		html:=request(path)
		fmt.Fprint(w,html)
		}
	})

        http.ListenAndServe(":3000", nil)
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
