// Go import path support for Atlassian Stash
//
// Installation and usage instructions here: https://github.com/rubenv/stash-go-import
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
		<a href="/">Redirecting to Stash soon</a>
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

	project := parts[1]
	repo := parts[2]
	importPath := fmt.Sprintf("%s/%s/%s", req.Host, project, repo)
	cloneUrl := fmt.Sprintf("ssh://git@%s:%d/%s/%s.git", req.Host, *sshPort, project, repo)

	d := &templateData{
		ImportPath: importPath,
		CloneUrl:   cloneUrl,
	}

	err := tpl.Execute(w, d)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
