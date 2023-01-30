package client

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

func templateFile(dir string, inject injection) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		buf := bytes.NewBuffer(nil)

		filename := req.URL.Path

		if filename == "/" {
			filename = "/index.html"
		}

		switch filepath.Ext(filename) {
		case ".css":
			writer.Header().Set("content-type", "text/css")
		case ".html":
			writer.Header().Set("content-type", "text/html")
		case ".js":
			writer.Header().Set("content-type", "text/javascript")
		}

		tmplFile, err := os.Open(filepath.Clean(path.Join(dir, filename)))
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		data, err := io.ReadAll(tmplFile)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		tmpl, err := template.New("name").Parse(string(data))
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		if err = tmpl.Execute(buf, inject); err != nil {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		if _, err := io.Copy(writer, buf); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}
