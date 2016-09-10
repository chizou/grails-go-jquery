package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func simulatedExternalUserService(w http.ResponseWriter, r *http.Request) {
	t := simulateComputation()
	ip := myIPWithTimeout()
	writeReturnedJSON(w, t, ip)
}

func simulateComputation() int {
	//rand.Seed(time.Now().Unix())
	const maxSimulatedComputation = 16
	t := rand.Intn(maxSimulatedComputation)
	time.Sleep(time.Second * time.Duration(t))
	return t
}

type returnedJSON struct {
	Results int
	Status  string
}

func writeReturnedJSON(w http.ResponseWriter, t int, ip string) {
	p := returnedJSON{t, ip}
	b, err := json.Marshal(p)
	if err != nil {
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
