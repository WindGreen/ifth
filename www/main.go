package main

import (
	"fmt"
	"html/template"
	"ifth"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"

	"gopkg.in/yaml.v2"
)

type UrlResponse struct {
	Ok      bool
	Tips    string
	Url     string
	Hottest []ifth.Url
	Newest  []ifth.Url
}

var config Config

func main() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	ifth.InitSlotGenerator(config.Url.Length, config.Url.Algorithm, config.Url.Humanity)

	_, err = ifth.InitMgo(config.MongoDB.Host)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr: fmt.Sprintf(":%d", config.WWW.Port),
	}

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
		data.Url = config.WWW.Home
		data.Hottest, _ = ifth.FindHottestUrls(10)
		for i, url := range data.Hottest {
			data.Hottest[i].Url = fmt.Sprintf(config.Url.Base, url.Slot)
		}
		data.Newest, _ = ifth.FindNewestUrls(10)
		for i, url := range data.Newest {
			data.Newest[i].Url = fmt.Sprintf(config.Url.Base, url.Slot)
		}
		t := template.Must(template.ParseFiles("./templates/index.html"))
		err := t.Execute(w, data)
		if err != nil {
			log.Println(err)
		}
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
		url := ifth.NewUrl(uri, config.Url.Unique)
		if url == nil {
			response.Ok = false
			response.Tips = "创建失败，请联系管理员"
		} else {
			response.Ok = true
			response.Url = fmt.Sprintf(config.Url.Base, url.Slot)
		}
	}
	w.Header().Set("Content-type", "text/html")
	t := template.Must(template.ParseFiles("./templates/index.html"))
	t.Execute(w, response)
}

func IsUrl(url string) bool {
	//根据dns判断可用
	reg := regexp.MustCompile(`http[s]?:\/\/([\w\.-]+)`)
	match := reg.FindAllStringSubmatch(url, 2)
	if len(match) > 0 && len(match[0]) > 1 {
		domain := match[0][1]
		_, err := net.LookupIP(domain)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
