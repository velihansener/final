package main

import (
	"mux-master"
	"net/http"
)

func api() {

	r := mux.NewRouter()

	r.HandleFunc("/", anasayfa)
	r.HandleFunc("/giris", giris)
	r.HandleFunc("/kayit", kayit)
	r.HandleFunc("/bilgiler", bilgiler)

	//http.HandleFunc("/img", serveResource)
	//http.HandleFunc("/js", serveResource)
	//http.HandleFunc("/css", serveResource)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)

	/*s := http.NewServeMux()
	s.Handle("/", http.FileServer(http.Dir("./static")))
	s.HandleFunc("/giris", login)

	http.ListenAndServe(":8080", s)*/
}
