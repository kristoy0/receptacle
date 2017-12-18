package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kristoy0/receptacle-worker/container"
	"github.com/kristoy0/receptacle-worker/image"

	"github.com/docker/docker/client"
)

var (
	ctx    = context.Background()
	cli, _ = client.NewEnvClient()
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/image/", ListImage).Methods("GET")
	r.HandleFunc("/api/containers/", ListCont).Methods("GET")
	r.HandleFunc("/api/containers/create/", RunCont).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func RunCont(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	conv := new(container.Container)
	err := json.Unmarshal(body, &conv)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	conv.Run(ctx, cli)
	w.Write(body)
}

func ListImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	out, err := image.List(ctx, cli)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	outConv, err := json.MarshalIndent(out, "", "	")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	w.Write([]byte(outConv))
}

func ListCont(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	out, err := container.List(ctx, cli)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	outConv, err := json.MarshalIndent(out, "", "	")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	w.Write([]byte(outConv))
}
