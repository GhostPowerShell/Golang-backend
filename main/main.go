package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func rotN(input string, key int) string {
	output := ""
	for _, i := range input {
		if i >= 'a' && i <= 'z' {
			i = 'a' + (i-'a'+rune(key))%26
		} else if i >= 'A' && i <= 'Z' {
			i = 'A' + (i-'A'+rune(key))%26
		}
		output += string(i)
	}
	return output
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		plaintext := r.FormValue("plaintext")
		keyStr := r.FormValue("key")

		key, err := strconv.Atoi(keyStr)
		if err != nil || key < 0 {
			http.Error(w, "Invalid key input", http.StatusBadRequest)
			return
		}

		output := rotN(plaintext, key)

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, struct{ Output string }{Output: output})
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, struct{ Output string }{})
}

func main() {
	http.HandleFunc("/", rootHandler)

	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
