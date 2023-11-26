package main

import (
	"flag"
	"github.com/golangcollege/sessions"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type errors map[string][]string
type Form struct {
	url.Values
	Errors errors
}
type ImagePair struct {
	FrontImagePath string
	FrontTitle     string
	BackImagePath  string
	BackTitle      string
}
type PageVariables struct {
	Title  string
	Images []ImagePair
	Form   *Form
}
type application struct {
	errorlog *log.Logger
	session  *sessions.Session
	infolog  *log.Logger
}

//func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
//	data := PageVariables{
//		Title: "Book",
//		Images: []string{
//			"static/img/dog.jpg", // Replace with the actual paths to your images
//			"static/img/dog.jpg",
//			"static/img/dog.jpg",
//			"static/img/dog.jpg",
//			"static/img/dog.jpg",
//		},
//	}
//
//	tmpl, err := template.ParseFiles("ui/html/index.layout.tmpl")
//	if err != nil {
//		log.Fatal(err)
//		http.Error(w, "Error parsing template file", http.StatusInternalServerError)
//		return
//	}
//
//	err = tmpl.Execute(w, data)
//	if err != nil {
//		log.Fatal(err)
//		http.Error(w, "Error executing template", http.StatusInternalServerError)
//		return
//	}
//}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	secret := flag.String("secret", "mynameisanthonygolandservice@123", "secret key")
	flag.Parse() //CLF ko command line me bhi input kar sakte h nhi dete h toh ye apni default value lete h
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	session := sessions.New([]byte(*secret)) //New initializes a new Session object to hold the configuration settings for your sessions.
	// The key parameter is the secret you want to use to authenticate and encrypt session cookies. It should be exactly 32 bytes long.
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	app := &application{
		errorlog: errorlog,
		infolog:  infolog,
		session:  session,
	}
	//tlsConfig := &tls.Config{
	//	PreferServerCipherSuites: true,
	//	CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	//}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  app.routes(),
	}
	err := srv.ListenAndServe()
	errorlog.Fatal(err)
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", indexHandler)
	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//log.Println("Starting server on :4000")
	//err := http.ListenAndServe(":4000", mux)
	//log.Fatal(err)
}
