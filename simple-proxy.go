package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func imageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL)

	url := r.URL.Query()["url"][0]

	if url != "" {
		response, err := http.Get(string(url))
		if err != nil {
			fmt.Print("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Print("%s", err)
				os.Exit(1)
			}
			fmt.Fprintf(w, "%s", string(contents))
		}
	} else {
		fmt.Fprintf(w, "URL param missing")
	}
}

func main() {
	http.HandleFunc("/i/", imageHandler)
	http.ListenAndServe(":8080", nil)
}
