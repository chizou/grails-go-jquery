package main

import (
	"math/rand"
)
var eusips []string

func askSwarmForEUSIPs() {
	//external user service IPs
	eusips = []string{`52.88.140.188`,`54.68.64.105`}
}

func randomExternalUserService() string {
	askSwarmForEUSIPs()
	r := rand.Intn(len(eusips))
	return eusips[r]
}

