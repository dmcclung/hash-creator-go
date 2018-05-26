package main

import (
	"net/http"
	"strings"
	"strconv"
	"log"
	"io/ioutil"
	"encoding/json"
	"context"
)

func hashGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		parts := strings.Split(r.RequestURI, "/")
		idpart := parts[len(parts) - 1]
		id, err := strconv.ParseUint(idpart, 10, 64)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hash := store.GetHash(id)
		w.Write([]byte(hash))
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func hashPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		kv, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		pw := strings.Split(string(kv), "=")[1]
		w.Write([]byte(strconv.FormatUint(store.StoreHash(pw), 10)))
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func stats(w http.ResponseWriter, r *http.Request) {
	total, avg := store.GetStats()
	stats := map[string]string{"total": strconv.FormatUint(total, 10),
		"average": strconv.FormatFloat(avg, 'f', 4, 64)}
	stats_json, _ := json.Marshal(stats)
	w.Write(stats_json)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	srv.Shutdown(context.Background())
	store.Stop()
}
