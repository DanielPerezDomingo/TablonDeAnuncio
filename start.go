package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Anuncio struct {
	Texto string
	tmpl  *template.Template
}

func (a *Anuncio) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("method: %s, path: %s,", r.Method, r.URL.Path)

	if r.URL.Path == "/edit" {

		if r.Method == "POST" {
			a.Texto = r.FormValue("texto")
			fmt.Printf("se ha guardado:%s\n", a.Texto)
		}

		err := a.tmpl.ExecuteTemplate(w, "index", a)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	err := a.tmpl.ExecuteTemplate(w, "display", a)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {

	tmpl, err := template.New("index").ParseFiles("web/tmpl/index.tmpl", "web/tmpl/display.tmpl")

	if err != nil {
		fmt.Println(err)
		return
	}

	Anun := Anuncio{
		tmpl:  tmpl,
		Texto: "Escribe algo antes",
	}

	fs := http.FileServer(http.Dir("./web/"))
	http.Handle("/js/", fs)
	http.Handle("/css/", fs)
	http.Handle("/", &Anun)

	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
