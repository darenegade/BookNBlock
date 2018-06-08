package logserver

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Body []string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	filename := "testlogfile.log"
	p := &Page{Body: readLogFile(filename)}
	renderTemplate(w, "logging", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func readLogFile(filename string) (logs []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}
	return
}

func SetLogging(filename string) (f *os.File, err error) {
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return
}

func StartWebserver() {
	http.HandleFunc("/logging", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
