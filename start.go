package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Anuncio struct {
	Texto string
}

func (a *Anuncio) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.New("index").ParseFiles("index.tmpl", "display.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.URL.Path == "/display" {
		err = tmpl.ExecuteTemplate(w, "display", a)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	if r.Method == "POST" {
		a.Texto = r.FormValue("texto")
		fmt.Printf("se ha guardado:%s\n", a.Texto)
	}

	err = tmpl.ExecuteTemplate(w, "index", a)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	Anun := Anuncio{}
	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", &Anun)
}
