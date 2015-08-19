package main

import (
	"io/ioutil"
	"net/http"
)

func handleHostname(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content-Type", "application/json")
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		err := ioutil.WriteFile("/proc/sys/kernel/hostname", body, 0644)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		hostnameBytes, err := ioutil.ReadFile("/proc/sys/kernel/hostname")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(hostnameBytes)
	}
}
