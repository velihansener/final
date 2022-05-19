package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func serveContent(w http.ResponseWriter, r *http.Request) {
	u := mux.Vars(r)

	page_alias := u["page_alias"]

	if page_alias == "" {
		page_alias = "i_home"
	}

	staticPage := staticPages.Lookup(page_alias + "html")

	if staticPage == nil {
		staticPage = staticPage.Lookup("404.html")
	}
	staticPage.Execute(w, nil)
}

func populateStaticPage() *template.Template {
	result := template.New("templates")

	templatePaths := new([]string)
	//basePath := GetAppPath() + "/" + "pages/ui"
	basePath := "C:/Users/Velihan/Desktop/final/bin/pages/ui"

	templateFolder, err := os.Open(basePath)

	if err != nil {
		log.Fatal("Unable to parse from template:", err)
	}

	templatePathsRaw, err := templateFolder.ReadDir(-1)

	if err != nil {
		log.Fatal("Unable to parse from template:", err)
	}

	for _, pathInfo := range templatePathsRaw {
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}

	/*basePath = GetAppPath() + "/" + "pages/kayit"

	templateFolder, _ = os.Open(basePath)
	templatePathsRaw, _ = templateFolder.ReadDir(-1)

	for _, pathInfo := range templatePathsRaw {
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}*/

	result.ParseFiles(*templatePaths...)

	defer templateFolder.Close()
	return result

}

func serveResource(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	var path string
	/*var a []string

	if strings.Contains(r.URL.Path, "/") {
		a = strings.Split(r.URL.Path, "/")
	}
	*/
	path = GetAppPath() + "/" + "public" + r.URL.Path

	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}

}

func GetAppPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)
	//dir = strings.Replace(dir, string(filepath.Separator), "/", -1)

	return dir
}
