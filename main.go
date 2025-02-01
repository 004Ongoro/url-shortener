package main

import (
	"fmt"
	"log"
	"net/http"
	"crypto/rand"
	"encoding/base64"
	"strings"
)

var urlStore = make(map[string] string)

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to URL Shortener!")
}

func shortenHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Ony POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	longUrl := r.FormValue("url")
	if longUrl== ""{
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := GenerateShortCode()
	urlStore[shortCode] = longUrl

	fmt.Fprintf(w, "https://gurulo/%s", shortCode)
}

func redirectHandler(w http.ResponseWriter, r *http.Request){
	shortCode := r.URL.Path[1:]

	longUrl, exists := urlStore[shortCode]

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longUrl, http.StatusFound)
}

func GenerateShortCode() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	code := base64.URLEncoding.EncodeToString(b)
	return strings.TrimRight(code, "=")
}

func main(){
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server is running on http:gurulo.onrender.com:10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}
