package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/index.html"))
		temp.Execute(w, nil)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/fregments/results.html"))
		data := map[string][]Stock{"Results": SearchStocksResult(r.URL.Query().Get("key"))}
		temp.Execute(w, data)
	})

	http.HandleFunc("/stock/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			ticker := r.PostFormValue("ticker")
			println("ticker", ticker)
			stk := SearchStocksResult(ticker)[0]
			val := getDailyValue(ticker)
			temp := template.Must(template.ParseFiles("./templates/index.html"))
			temp.ExecuteTemplate(w, "stock-element", Stock{Ticker: stk.Ticker, Name: stk.Name, Price: val.Open})
		}
	})

	log.Println("App runing on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
