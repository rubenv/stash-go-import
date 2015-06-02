package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var tpl = template.Must(template.New("redirect").Parse(`<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="go-import" content="{{.ImportPath}} git {{.CloneUrl}}">
		<meta http-equiv="refresh" content="0; url=/">
	</head>
	<body>
		<a href="/">Redirecting to stash soon</a>
	</body>
</html>
`))

var (
	port    = flag.Int("port", 80, "listen on port")
	sshPort = flag.Int("sshPort", 7999, "use SSH port")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: stash-go-import\n\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

type templateData struct {
	ImportPath string
	CloneUrl   string
}

func redirect(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 3 {
		http.NotFound(w, req)
		return
	}

	importPath := fmt.Sprintf("%s%s", req.Host, req.URL.Path)
	cloneUrl := fmt.Sprintf("ssh://git@%s:%d/%s/%s.git", req.Host, *sshPort, parts[1], parts[2])

	d := &templateData{
		ImportPath: importPath,
		CloneUrl:   cloneUrl,
	}

	err := tpl.Execute(w, d)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
