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
			if path=="/ascci/style.css"{
				provide_style(w,"style_ascci.css")
			}else{
				provide_style(w,"style.css")
			}
		}else if strings.HasSuffix(path,".html") || path=="/"{
			html:=request(path)
			fmt.Fprint(w,html+"<meta name='viewport' content='width=device-widht, initial-scale=1.0'>")
		}else{
			html:=request(path)
			fmt.Fprint(w,html)
		}
	})

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

 func provide_style(w http.ResponseWriter,style_path string){
	w.Header().Set("Content-Type", "text/css")
	data, err := ioutil.ReadFile(style_path)
	if err != nil {
	    fmt.Println(err)
	}
	style:=string(data)
	fmt.Fprint(w,style)
 }
