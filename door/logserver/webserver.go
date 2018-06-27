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

// SetLogging sets the global logging to a file with the filename that's passed
// Warning: When using this you have to care about the close() Operation of the file
func SetLogging(filename string) (f *os.File, err error) {
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return
}

// StartWebserver is running a webserver where the logs are listed the url is '/logging'
// Warning: Run this in a goroutine otherwise it's blocking
func StartWebserver() {
	http.HandleFunc("/logging", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
