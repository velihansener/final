package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func giris(w http.ResponseWriter, r *http.Request) {
	fmt.Println("giris sayfası")
	staticPage := staticPages.Lookup("giris.html")
	staticPage.Execute(w, nil)

	if r.Method == "POST" {
		mail := r.FormValue("myEmail")
		pass := r.FormValue("sifrem")
		pin := r.FormValue("pinkod")
		log.Println("mail: ", mail)
		log.Println("şifre: ", pass)
		log.Println("pin: ", pin)

		dataGet(mail, pass, pin)

	}
}

func kayit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("kayıt ol")
	staticPage := staticPages.Lookup("kayit.html")
	staticPage.Execute(w, nil)

	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		mail := r.FormValue("myEmail")
		pass := r.FormValue("sifrem")
		pin := r.FormValue("pinkod")
		log.Println("mail: ", mail)
		log.Println("şifre: ", pass)
		log.Println("pin: ", pin)

		if kontrol(mail) {
			fmt.Println("Girilen email adresi doğrudur.")
			addUsers(mail, pass, pin)
		} else {
			fmt.Println("Girilen email adresi yanlıştır.")
			fmt.Println("Tekrar deneyiniz.")
		}
	}
}

func kontrol(str string) bool {
	var chk bool
	var a, b, c, d bool
	var indx int
	var before, after string

	if strings.Contains(str, "@") && strings.Contains(str, ".") && strings.Count(str, "@") == 1 {
		a = strings.HasSuffix(str, "@gmail.com")
		b = strings.HasSuffix(str, "@hotmail.com")
		c = strings.HasSuffix(str, "@outlook.com")
		d = strings.HasSuffix(str, "@yahoo.com")
		indx = strings.Index(str, "@")
		before = string(str[indx-1])
		after = string(str[indx+1])
		if before != "." && after != "." {
			if a || b || c || d {
				chk = true
			}
		}
	}
	return chk
}

func anasayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ana sayfa")
	staticPage := staticPages.Lookup("index.html")
	staticPage.Execute(w, nil)
}
