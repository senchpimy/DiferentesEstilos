package main

import (
	"fmt"
	"io"
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
			html:=request(path)
			fmt.Fprint(w,html)
		}
		html:=request(path)
		fmt.Fprint(w,html)
	})

        http.ListenAndServe(":3000", nil)
}

 func request(url string)(foo string){
	resp, err := http.Get(domain+url)
	if err != nil {
		fmt.Println("erro get")
	    // handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("erro read")
	    // handle error
	}
	return string(body)
 }
