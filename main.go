package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const devexRate = 0.0038

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type Resp struct {
	Robux float64 `json:"robux"`
	USD   float64 `json:"usd"`
	Rate  float64 `json:"rate"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/convert", apiConvertHandler)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, map[string]any{
		"Rate": devexRate,
	})
}

func apiConvertHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	var robux, usd float64

	if v := q.Get("robux"); v != "" {
		robux, _ = strconv.ParseFloat(v, 64)
		usd = robux * devexRate
	}

	if v := q.Get("usd"); v != "" {
		usd, _ = strconv.ParseFloat(v, 64)
		robux = usd / devexRate
	}

	resp := Resp{robux, usd, devexRate}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
