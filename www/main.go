package main

import (
	"fmt"
	"html/template"
	"ifth"
	"log"
	"net/http"
)

type UrlResponse struct {
	Ok   bool
	Tips string
	Url  string
}

func main() {
	ifth.InitSlotGenerator()
	_, err := ifth.InitMgo("localhost")
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr: ":8081",
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
	http.HandleFunc("/create.html", createHandle)
	log.Println("Start listening...")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	t := template.Must(template.ParseFiles("./templates/index.html"))
	t.Execute(w, nil)
}

func createHandle(w http.ResponseWriter, r *http.Request) {
	var response UrlResponse
	r.ParseForm()
	uri := fmt.Sprintf("%s://%s", r.FormValue("protocol"), r.FormValue("url"))
	if !IsUrl(uri) {
		response.Ok = false
		response.Tips = "url不符合规范"
	} else {
		url := ifth.NewUrl(uri, false)
		if url == nil {
			response.Ok = false
			response.Tips = "创建失败，请联系管理员"
		} else {
			response.Ok = true
			response.Url = fmt.Sprintf("http://localhost:8080/%s", url.Slot)
		}
	}
	w.Header().Set("Content-type", "text/html")
	t := template.Must(template.ParseFiles("./templates/index.html"))
	t.Execute(w, response)
}

func IsUrl(url string) bool {
	if len(url) < 12 {
		return false
	}
	return true
	response, err := http.Get(url)
	if err != nil {
		// log.Println(err)
		return false
	}
	// defer response.Body.Close()
	// body, _ := ioutil.ReadAll(response.Body)
	// log.Println(string(body))
	if response.StatusCode >= 200 && response.StatusCode < 400 {
		return true
	}
	return false
}
