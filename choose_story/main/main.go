package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	choosestory "github.com/Ekosetiawan993/choose_story"
)

func tryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Eko......")
}

func main() {
	// flag for optional variable
	port := flag.Int("port", 8084, "The port to start this application")
	filename := flag.String("file", "gopher.json", "The story's JSON filename")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// open the json file
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := choosestory.JsonStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := choosestory.NewHandler(story,
		choosestory.WithTemplate(tpl),
		choosestory.WithPathFunc(pathFn))

	mux := http.NewServeMux()

	mux.Handle("/story/", h)

	mux.Handle("/", choosestory.NewHandler(story))

	fmt.Printf("Starting server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`

// mux := http.DefaultServeMux
// mux.HandleFunc("/", tryHandler)

// fmt.Println("Serve on port 8084")

// http.ListenAndServe(":8084", nil)
