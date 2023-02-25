package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	ip := struct {
		Ip string `json:"ip"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&ip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, ip.Ip)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
