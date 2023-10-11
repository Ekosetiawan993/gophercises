package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// func (us *urlShortenerMap) handleRedirect(w http.ResponseWriter, r *http.Request) {
// 	shortKey := r.URL.Path

// 	fmt.Printf("%s", shortKey)

// 	originalURL, found := us.urls[shortKey]

// 	if !found {
// 		http.Error(w, "Not found short link", http.StatusNotFound)
// 		return
// 	}
// 	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)

// }

func myHandler(w http.ResponseWriter, r *http.Request) {
	// Logika penanganan permintaan HTTP di sini
	fmt.Fprintln(w, "Hello, World!")
	fmt.Fprintln(w, r.URL.Path)
}

type urlShortenerMap struct {
	urls map[string]string
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.New(rand.NewSource(time.Now().UnixNano()))

	shortkey := make([]byte, keyLength)

	for i := range shortkey {
		shortkey[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortkey)
}

func (us *urlShortenerMap) handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is missing", http.StatusBadRequest)
		return
	}

	shortKey := generateShortKey()
	us.urls[shortKey] = originalURL

	shortenedURL := fmt.Sprintf("http://localhost:8084/short/%s", shortKey)

	w.Header().Set("Content-Type", "text/html")
	responseHTML := fmt.Sprintf(`
	<h2>URL Shortener</h2>
	<p>Original URL: %s</p>
	<p>Shortened URL: <a href="%s">%s</a></p>
        <form method="post" action="/shorten">
            <input type="text" name="url" placeholder="Enter a URL">
            <input type="submit" value="Shorten">
        </form>
	`, originalURL, shortenedURL, shortenedURL)

	fmt.Fprintf(w, responseHTML)

}

func (us *urlShortenerMap) handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]

	if shortKey == "" {
		http.Error(w, "Shortkey is missing", http.StatusBadRequest)
		return
	}
	originalURL, found := us.urls[shortKey]

	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>URL Shortener</title>
		</head>
		<body>
			<h2>URL Shortener</h2>
			<form method="post" action="/shorten">
				<input type="url" name="url" placeholder="Enter a URL" required>
				<input type="submit" value="Shorten">
			</form>
		</body>
		</html>
	`)
}

func main() {
	shortener := &urlShortenerMap{
		urls: make(map[string]string),
	}

	http.HandleFunc("/shorten", shortener.handleShorten)
	http.HandleFunc("/short/", shortener.handleRedirect)

	// pathToUrl := map[string]string{
	// 	"/go/eko": "https://ekosetiawan993.github.io/",
	// }
	// urlMaps := &urlShortenerMap{
	// 	urls: pathToUrl,
	// }

	http.HandleFunc("/", handleForm)
	// http.HandleFunc("/go/", urlMaps.handleRedirect)
	// handler.TestHandler()

	fmt.Println("serve http on port: 8084")
	http.ListenAndServe(":8084", nil)
}
