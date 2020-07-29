package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var v int64 = 0
var lastSync int64 = 0
var mux = &sync.Mutex{}

var filename string = "/data/counter.txt"

func writeValue() {
	seqVStr := fmt.Sprintf("%d", v+1)
	ioutil.WriteFile(filename, []byte(seqVStr), 0600)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	seqV := v
	v++
	lastSync++
	if lastSync == 10000 {
		writeValue()
		lastSync = 0
	}
	mux.Unlock()
	fmt.Fprintf(w, "%d", seqV)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Great!")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readValue() {
	if fileExists(filename) {
		body, _ := ioutil.ReadFile(filename)
		i, _ := strconv.ParseInt(string(body), 10, 64)
		v = i + 20000 + 10
	}
}

func main() {
	readValue()
	writeValue()
	http.HandleFunc("/", handler)
	http.HandleFunc("/how/you/doing", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
