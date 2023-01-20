package main

import ("fmt"
	"net/http"
	"io"
	"strings"
	"regexp"
)

const domain = "https://senchpimy.xyz"
//const domain = "https://senchpimy.xyz"

func main() {
	fmt.Println("Server Start")
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprint(w,request("/"))
		path := r.URL.Path
		if path!="/"{
			fmt.Println("hola"+path)
			html:=request(path)
			//re := regexp.MustCompile(`src="(.*?)"`)
			re := regexp.MustCompile(`src=".*."`)
			matches := re.FindAllStringSubmatch(html, -1)
			for _, match := range matches {
				fmt.Println(match)
			}
			if strings.HasSuffix(path,".html"){
			fmt.Fprint(w,html)
			}
			if strings.HasSuffix(path,".jpg"){
				html:=request(path)
				fmt.Fprint(w,html)
			}
		}else {
			html:=request(path)
			fmt.Fprint(w,html)

		}
	})

        http.ListenAndServe(":3000", nil)
}

 func request(url string)(foo string){
	resp, err := http.Get(domain+url)
	if err != nil {
	    // handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
	    // handle error
	}
	return string(body)
 }
