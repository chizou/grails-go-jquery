package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type healthArray struct {
	myip     string
	port8079 bool
	port8081 bool
	port8083 bool
	port8084 bool
	port8085 bool
	port8088 bool
}
func newHealthArray() *healthArray {
	p := new(healthArray)
	updateHealthArray(*p)
	return p
}
func updateHealthArray(h healthArray) *healthArray {
	h.port8079 = test8079()
	h.port8081 = test8081()
	h.port8083 = test8083()
	h.port8084 = test8084()
	h.port8085 = test8085()
	h.port8088 = test8088()
	return &h
}
func test8079() bool { return testPort(`8079`,`Nexus`) }
func test8081() bool { return testPort(`8081`,`pz-gateway`) }
func test8083() bool { return testPort(`8083`,`pz-jobmanager`) }
func test8084() bool { return testPort(`8084`,`Loader`) }
func test8085() bool { return testPort(`8085`,`pz-access`) }
func test8088() bool { return testPort(`8088`,`Piazza Service Controller`) }

func testPort(port string, contains string) bool {
	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("http://localhost:"+port)
	if err != nil {
		fmt.Println("Something went wrong for "+port)
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return strings.Contains(string(body), contains)
}
func statusString(qhealth *healthArray) string {
	s := ""
	if qhealth != nil {
		if qhealth.port8079 == false {
			s += "nexus?"
		}
		if qhealth.port8081 == false {
			s += "gateway?"
		}
		if qhealth.port8083 == false {
			s += "jobmanager?"
		}
		if qhealth.port8084 == false {
			s += "ingest?"
		}
		if qhealth.port8085 == false {
			s += "access?"
		}
		if qhealth.port8088 == false {
			s += "servicecontroller?"
		}
	}
	return s
}
