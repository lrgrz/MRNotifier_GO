package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/v4/merge_requests/", getMergeRequestsFile)
	http.HandleFunc("/api/v4/projects/", getApprovals)
	http.ListenAndServe(":8080", nil)
}

func getMergeRequestsFile(w http.ResponseWriter, req *http.Request) {
	streamFile(w, "mergeRequest.json")
}

func getApprovals(w http.ResponseWriter, req *http.Request) {
	streamFile(w, "approvals.json")
}

func streamFile(w http.ResponseWriter, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File not found: ", err)
	}
	defer file.Close()

	w.Header().Add("Content-Type", "application/json")

	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatal("IO error: ", err)
	}
}
