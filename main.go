package main

import "net/http"



func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/plain; charset=utf-8")

		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	mux.Handle("/app/",http.StripPrefix("/app", http.FileServer(http.Dir("."))) )
	mux.Handle("/assests", http.FileServer(http.Dir("/logo.png")))

	port := "8080"

	server := http.Server{
		Addr: ":" + port,
		Handler: mux,
	}
	server.ListenAndServe()
}