package main

import (
	"fmt"
	"html/template"
	"ifth"
	"log"
	"net"
	"net/http"
	"regexp"
)

type UrlResponse struct {
	Ok   bool
	Tips string
	Url  string
}

func main() {
	ifth.InitSlotGenerator()
	_, err := ifth.InitMgo("mongo")
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr: ":80",
	}
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../templates/js"))))
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../templates/css"))))
	// http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("../templates/images"))))
	// http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../templates/favicon.ico")
	// })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index.html", http.StatusSeeOther)
	})
	http.HandleFunc("/index.html", homeHandle)
	log.Println("Start listening...")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if r.Method == "POST" {
		createHandle(w, r)
	} else {
		var data UrlResponse
		data.Ok = true
		data.Url = "http://www.ifth.net"
		t := template.Must(template.ParseFiles("./templates/index.html"))
		t.Execute(w, data)
	}
}

func createHandle(w http.ResponseWriter, r *http.Request) {
	var response UrlResponse
	r.ParseForm()
	uri := fmt.Sprintf("%s://%s", r.FormValue("protocol"), r.FormValue("url"))
	if !IsUrl(uri) {
		response.Ok = false
		response.Tips = "url不符合规范或域名无法解析"
	} else {
		url := ifth.NewUrl(uri, false)
		if url == nil {
			response.Ok = false
			response.Tips = "创建失败，请联系管理员"
		} else {
			response.Ok = true
			response.Url = fmt.Sprintf("http://ifth.net/%s", url.Slot)
		}
	}
	w.Header().Set("Content-type", "text/html")
	t := template.Must(template.ParseFiles("./templates/index.html"))
	t.Execute(w, response)
}

func IsUrl(url string) bool {
	//根据dns判断可用
	regexp.MatchString(`(?<=://)[a-zA-Z\.0-9]+(?=\/)`, url)
	_, err := net.LookupIP(domain)
	if err != nil {
		return false
	}
	return true
}
