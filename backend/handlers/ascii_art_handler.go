package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"ascii/backend/internal"
	"ascii/backend/pkg"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		text := strings.TrimSpace(r.Form.Get("text-input"))
		fmt.Println(text)
		if len(text) > 100 {

			http.Error(w, "ERROR", http.StatusBadRequest)
			// http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		banner := r.Form.Get("banner-select")

		lines, err := pkg.ReadLine(banner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		maps, err := internal.Create(lines)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		text = strings.ReplaceAll(text, "\r", "")

		words := strings.Split(text, "\n")
		// fmt.Println("qaz", []byte(words[0]), string(words[0]), "qwer", []byte(words[1]), string(words[1]))
		// fmt.Println(words, len(words))
		result := ""
		for _, word := range words {
			if word == "" {
				result += "\n"
				continue
			}
			for i := 0; i < 8; i++ {
				for _, value := range word {
					result += maps[value][i]
				}
				result += "\n"

			}
		}

		tmpl := template.Must(template.ParseFiles("frontend/templates/index.html"))
		err = tmpl.Execute(w, result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
