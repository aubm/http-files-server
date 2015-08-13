package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

var filesDir string
var token string

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}
	filesDir = os.Args[1]

	token = os.Getenv("TOKEN")

	r := mux.NewRouter()
	r.HandleFunc("/listFiles", listFiles)
	r.HandleFunc("/downloadFile", downloadFile)
	r.HandleFunc("/deleteFile", deleteFile).Methods("DELETE")
	http.Handle("/", r)

	fmt.Println("Server started :8888")
	http.ListenAndServe(":8888", nil)
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	if err := checkRequestToken(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filename := r.URL.Query().Get("filename")
	filename = getAbsoluteFilePathname(filename)
	if err := checkFilePath(filename); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	http.ServeFile(w, r, filename)
}

func listFiles(w http.ResponseWriter, r *http.Request) {
	if err := checkRequestToken(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	availableFiles := scanDir(filesDir)
	b, err := json.Marshal(availableFiles)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b[:]))
}

func scanDir(dirToScan string) []string {
	var availableFiles []string
	cleanedFilesDir := filepath.Clean(filesDir)
	filepath.Walk(dirToScan, func(path string, f os.FileInfo, err error) error {
		path = filepath.Clean(path)
		// Check if file is a directory
		if f.IsDir() {
			return nil
		}
		// Check if file is a hidden file
		if validId := regexp.MustCompile(`^\.`); validId.MatchString(f.Name()) {
			return nil
		}
		// Clear file path prefix
		validId := regexp.MustCompile(`^` + cleanedFilesDir + `\/`)
		availableFiles = append(availableFiles, validId.ReplaceAllString(path, ""))
		return nil
	})
	return availableFiles
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	if err := checkRequestToken(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filename := r.URL.Query().Get("filename")
	filename = getAbsoluteFilePathname(filename)
	if err := checkFilePath(filename); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	os.Remove(filename)
	w.WriteHeader(http.StatusNoContent)
}

func checkRequestToken(r *http.Request) error {
	requestToken := r.URL.Query().Get("token")
	if requestToken != token {
		return errors.New("Invalid token")
	}
	return nil
}

func getAbsoluteFilePathname(filename string) string {
	filename = filesDir + "/" + filename
	filename, _ = filepath.Abs(filename)
	filename = filepath.Clean(filename)
	return filename
}

func checkFilePath(absFilePathname string) error {
	fileInfo, err := os.Stat(absFilePathname)
	if err != nil {
		return errors.New("File does not exist")
	}
	if fileInfo.IsDir() {
		return errors.New("File is a directory")
	}
	return nil
}
