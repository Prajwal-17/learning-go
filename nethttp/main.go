package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// r -> a pointer to http.Request struct
	// "/" -> catch all routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("req struct", r)
		fmt.Fprintf(w, "Hello You are at: %s\n", r.URL.Path)
	})

	// http.FileServer - inbuilt handler to serve static files
	fs := http.FileServer(http.Dir("./static"))

	// http.StripPrefix - to strips out "static" in route
	// else the FileServer looks at ./static/static/index.html
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Raw body pointer(mem address)", r.Body)

		// in go body is a stream, io.ReadAll -> to read into byte slice
		// This is reads data from Network Buffer and store to ram, if the user sends 100mb then there is a 100mb spike in ram
		// If the data is not used from Network Buffer the os clears when the connection closes
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Failed to read Body", http.StatusInternalServerError)
		}

		// defer - close the body as soon it finishes executing
		defer r.Body.Close()

		fmt.Printf("Body %s", string(body))
		fmt.Fprintf(w, "You are successfully registered %s\n", r.URL.Path)
	})

	fmt.Println("Listening on port 80")
	http.ListenAndServe(":8080", nil)
}
