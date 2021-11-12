package errorHandle

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func ServerStart() {
	http.HandleFunc("/static/",
		errWrapper(HandleFile))

	http.ListenAndServe("localhost:8888", nil)
}

func HandleFile(rw http.ResponseWriter, r *http.Request) error {

	myPath, e := os.Getwd()
	if e != nil {
		return e
	}
	fmt.Println("path", myPath)

	fileToRead := r.URL.Path
	f, e := os.ReadFile(filepath.Join(filepath.Dir(myPath), fileToRead))
	if e != nil {
		return e
	}
	fmt.Fprintf(rw, "=========\nFile Content\n========= \n\n%s\n", string(f))
	return nil
}

// handle func returns error, so wrap it
type handler func(rw http.ResponseWriter, r *http.Request) error

func errWrapper(handler) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := HandleFile(rw, r)
		if err != nil {
			fmt.Fprintf(rw, "error: %s", err)
		}

	}
}
