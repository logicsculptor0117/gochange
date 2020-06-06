/*
This file is under GNU AFFERO GENERAL PUBLIC LICENSE

Permissions of this strongest copyleft license are conditioned
on making available complete source code of licensed works and
modifications, which include larger works using a licensed work,
under the same license. Copyright and license notices must be preserved.
Contributors provide an express grant of patent rights.
When a modified version is used to provide a service over a network,
the complete source code of the modified version must be made available.

Edoardo Ottavianelli, https://edoardoottavianelli.it

*/

package webserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// TODO
func StartListen() {
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/save/", handlerSave)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO
func handlerHome(w http.ResponseWriter, r *http.Request) {

	setContentType(w, r)

	page := ""
	if r.RequestURI == "/" {
		page, _ = loadPage("fe/home.html")
	} else {
		page, _ = loadPage("." + r.RequestURI)
	}

	fmt.Fprintf(w, "%s", page)
}

// TODO
func handlerSave(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	telegram := r.FormValue("telegram")
	website := r.FormValue("website")
	interval := r.FormValue("interval")
	fmt.Println(email)
	fmt.Println(telegram)
	fmt.Println(website)
	fmt.Println(interval)
	fmt.Fprintf(w, "%s %s %s %s", email, telegram, website, interval)
}

// TODO
func loadPage(filename string) (string, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func setContentType(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	contentType := "text/html"

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	}

	w.Header().Set("Content-Type", contentType)

}
