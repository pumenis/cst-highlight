package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pumenis/source-highlight/highlight"
)

//go:embed css/*.css
var cssFS embed.FS

func main() {
	preview := flag.Bool("preview", false, "Embed CSS into HTML output")
	cssOnly := flag.Bool("css", false, "Output only the CSS for the detected language")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: prog [options] <codefile>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	codePath := flag.Arg(0)

	content, err := os.ReadFile(codePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	sourceCode := string(content)

	var result, cssFile string
	ext := filepath.Ext(codePath)
	switch ext {
	case ".go":
		result = highlight.GetGoHighlighted(sourceCode)
		cssFile = "css/gosyntax.css"
	case ".sh":
		result = highlight.GetBashHighlighted(sourceCode)
		cssFile = "css/bashsyntax.css"
	case ".html":
		result = highlight.GetHTMLHighlighted(sourceCode)
		cssFile = "css/htmlsyntax.css"
	case ".js":
		result = highlight.GetJSHighlighted(sourceCode)
		cssFile = "css/htmlsyntax.css"
	case ".sql":
		result = highlight.GetSQLHighlighted(sourceCode)
		cssFile = "css/sqlsyntax.css"
	case ".css":
		result = highlight.GetCSSHighlighted(sourceCode)
		cssFile = "css/csssyntax.css"
	default:
		fmt.Fprintf(os.Stderr, "Unsupported file type: %s\n", ext)
		os.Exit(1)
	}

	cssContent, err := cssFS.ReadFile(cssFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading embedded CSS: %v\n", err)
		os.Exit(1)
	}

	if *cssOnly {
		fmt.Print(string(cssContent))
		return
	}

	if *preview {
		fmt.Println(`<!DOCTYPE html>
<html>
<head>
<style>
			` + string(cssContent) + `</style>
</head>
<body><div><pre>` + result + `</pre></div></body>
</html>`)
	} else {
		fmt.Println(`<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" href="` + cssFile + `" />
</head>
<body><div><pre>` + result + `</pre></div></body>
</html>`)
	}
}
