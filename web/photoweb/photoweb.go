package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"

	"html/template"
)

const (
	UPLOAD_DIR   = "uploads"
	TEMPLATE_DIR = "views"
	STATIC_DIR   = "public"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderHTML(w, "views/upload.html", nil)
		check(err)
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		name := h.Filename
		size := h.Size
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + name)
		check(err)
		defer t.Close()
		n, err := io.Copy(t, f)
		check(err)
		if n != size {
			http.Error(w, "save fail", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+name, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	path := UPLOAD_DIR + "/" + id
	_, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, path)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, info := range files {
		images = append(images, info.Name())
	}
	locals["images"] = images
	err = renderHTML(w, "views/list.html", locals)
	check(err)
	return
}

var templates = make(map[string]*template.Template)

func renderHTML(w http.ResponseWriter, tpl string, locals map[string]interface{}) (err error) {
	err = templates[tpl].Execute(w, locals)
	return
}

func loadViewTpl() {
	files, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		log.Fatal("load view tpl failed")
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if ext := path.Ext(name); ext != ".html" {
			continue
		}
		tplPath := TEMPLATE_DIR + "/" + name
		log.Println("loading template: ", tplPath)
		t := template.Must(template.ParseFiles(tplPath))
		templates[tplPath] = t
	}
}

func init() {
	loadViewTpl()
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Panicf("CRIT: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/list", safeHandler(listHandler))
	log.Printf("listent 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
