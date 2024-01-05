package main

import (
	_ "embed"
	"errors"
	"flag"
	"html/template"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path"
)

//go:embed index.gtpl
var index string

// Catalogs represents a collection of file names.
type Catalogs struct {
	Dir       string
	ParentDir string
	Files     []File
}

type File struct {
	Name string
	Path string
}

type handler struct {
	dir string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.list(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		http.Error(w, "unsupported method "+r.Method, http.StatusBadRequest)
	}
}

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	name := path.Join(h.dir, r.URL.Path)
	stat, err := os.Stat(name)
	if errors.Is(err, os.ErrNotExist) {
		http.Error(w, r.URL.Path+" not found", http.StatusNotFound)
		return
	}

	if stat.IsDir() {
		tpl := template.Must(template.New("index").Parse(index))
		catalog := Catalogs{
			Dir:       name,
			ParentDir: path.Dir(name),
		}
		entries, err := os.ReadDir(name)
		if err != nil {
			slog.Error("read %s entry failed: %s", name, err.Error())
			http.Error(w, "", http.StatusServiceUnavailable)
			return
		}
		for _, e := range entries {
			catalog.Files = append(catalog.Files, File{
				Name: e.Name(),
				Path: path.Join(name, e.Name()),
			})
		}
		if err := tpl.Execute(w, catalog); err != nil {
			slog.Error("execute template failed: %s", err.Error())
		}
		return
	}

	http.ServeFile(w, r, name)
}

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		slog.Error("read form file failed: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	f, err := os.Create(path.Join(h.dir, r.URL.Path, header.Filename))
	if err != nil {
		slog.Error("create file %s failed: %s", path.Join(h.dir, header.Filename), err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		slog.Error("write file failed: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/"+r.URL.Path, http.StatusSeeOther)
}

func main() {
	var directory, bind string
	flag.StringVar(&directory, "directory", "./", "root directory")
	flag.StringVar(&bind, "bind", "127.0.0.1:8080", "bind address")
	flag.Parse()

	http.Handle("/", http.StripPrefix("/", &handler{dir: directory}))
	log.Println("Listen on", bind, "server directory", directory)
	log.Fatal(http.ListenAndServe(bind, nil))
}
