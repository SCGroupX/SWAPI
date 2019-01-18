package service

import (
	"fmt"
    "net/http"
    "text/template"
    "github.com/unrolled/render"
)

const LocalURL = "http://localhost:8080/api/"

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method:", r.Method) //获取请求的方法
		if r.Method == "GET" {
			t := template.Must(template.ParseFiles("templates/home.html"))
			t.Execute(w, map[string]string{
				"Result": "",
			})
		} else {
			r.ParseForm()
			path :=  r.Form["path"][0]
			url := path
			if path[:4] != "http" {
				url = LocalURL + path
				fmt.Println(url)
			}
			http.Redirect(w, r, url, http.StatusFound)
		}
	}

}