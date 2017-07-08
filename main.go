package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

var (
	host         string             = "http://localhost:8888/"
	themeName    string             = "default"
	themePath    string             = "theme/" + themeName + "/"
	templatePath string             = "theme/" + themeName + "/template/"
	tmpl         *template.Template = parseTemplateFiles(templatePath)
)

type T struct {
	Foo string `json:"foo`
}

func main() {
	serveWeb()
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	tmpl.ExecuteTemplate(w, "404", nil)
}

func serveWeb() {

	router := mux.NewRouter()

	// Serve pages dinamyc by path
	router.HandleFunc("/", serveContent)
	router.HandleFunc("/{path}", serveContent)

	// Serve resources
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/img/", serveResource)

	http.Handle("/", router)

	// set handler for not found
	router.NotFoundHandler = http.HandlerFunc(notFound)

	http.ListenAndServe(":8888", nil)
}

// -------------------------------------------------------------
// Get url vars and serve associated file
func serveContent(w http.ResponseWriter, r *http.Request) {

	urlVars := mux.Vars(r)

	//	In this way we can use queryString in golang
	//	queryString := r.URL.Query()
	//	fmt.Println(queryString)

	file := urlVars["path"]
	if file == "" {
		file = "home"
	}

	// file must be in root of template dir
	_, e := os.Open(templatePath + file + ".html")
	if e != nil {
		file = "404"
	}

	//tmpl := parseTemplateFiles(templatePath)
	err := tmpl.ExecuteTemplate(w, file, nil)
	if err != nil {
		log.Println(err)
	}
}

//Parse all html files from /template dir
func parseTemplateFiles(dir string) *template.Template {
	templ := template.New("").Funcs(
		template.FuncMap{
			"showLink": func(target string) string {
				return host + target
			},
		})

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//log.Println(path)
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}
	return templ
}

// -------------------------------------------------------------
// Serve Resources like js, img, css files
func serveResource(w http.ResponseWriter, req *http.Request) {

	path := themePath + strings.Trim(req.URL.Path, "/")

	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charser=utf-8"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png; charser=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg; charser=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript; charser=utf-8"
	} else {
		contentType = "text/plain; charser=utf-8"
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
