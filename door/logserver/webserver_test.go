package logserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestRunWebserver(t *testing.T) {
	filename := "testlogfile.log"
	removeLogFile(filename)

	f, err := SetLogging(filename)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	log.Println("Log Zeile 1")
	mockLogging()

	go StartWebserver()

	resp, err := http.Get("http://localhost:8080/logging")

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		words := strings.Split(string(bodyBytes), "\n")

		if len(words) < 9 {
			t.Error("Serverantwort ist zu kurz")
		}
		if strings.Compare(words[0], "<h1>Logging Tür</h1>") == 0 {
			t.Error("<h1> ist falsch")
		}
		if !strings.Contains(words[4], "Log Zeile 1") {
			t.Error("Zeile 1 entspricht nicht der log-nachricht: Log Zeile 1")
		}
		if !strings.Contains(words[6], "Log Zeile 2") {
			t.Error("Zeile 2 entspricht nicht der log-nachricht: Log Zeile 2")
		}

	} else {
		t.Error("Anfrage hat nicht den Status HTTP_200")
	}
}

func TestLoggingToFile(t *testing.T) {
	filename := "testlogfile.log"
	removeLogFile(filename)

	f, err := SetLogging(filename)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	log.Println("Log Zeile 1")
	mockLogging()

	text := readLogFile(filename)
	if len(text) != 2 {
		t.Error("Logdatei hat nicht die richtige Anzahl an Zeilen")
	}

	if !strings.Contains(text[0], "Log Zeile 1") {
		t.Error(text[0])
		t.Error("Zeile 1 entspricht nicht der log-nachricht: Log Zeile 1")
	}

	if !strings.Contains(text[1], "Log Zeile 2") {
		t.Error(text[1])
		t.Error("Zeile 2 entspricht nicht der log-nachricht: Log Zeile 2")
	}

}

func mockLogging() {
	log.Println("Log Zeile 2")
}

func removeLogFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println(err)
	}
}
