package service

import (
	"fmt"
	"strings"
	"strconv"
    "net/http"
    "encoding/json"
    "text/template"
    "github.com/unrolled/render"
    "github.com/SCGroupX/SWAPI/swapi"
)

func filmsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method:", r.Method) //获取请求的方法
		if r.Method == "GET" {
			path := r.URL.Path
			fmt.Println("path: ", path)
			pos := strings.LastIndex(path, "/")
			fmt.Println("pos: ", pos)
			fmt.Println("substr: ", path[pos+1:])
			id, _ := strconv.Atoi(path[pos+1:])
			fmt.Println("id: ", id)

			s, _ := swapi.GetFilm(id)
			jsIndent,_ := json.MarshalIndent(s, "", "\t")
			fmt.Println(string(jsIndent))


			t := template.Must(template.ParseFiles("templates/home.html"))
			t.Execute(w, map[string]string{
				"Result": string(jsIndent),
			})
		}
	}

}