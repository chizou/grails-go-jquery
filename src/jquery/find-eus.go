package main

import (
	"math/rand"
)

func askSwarmForEUSIPs() []string {
	//external user service IPs
	return []string{`52.88.140.188`,`54.68.64.105`}
}

func randomExternalUserService() string {
	var eus = askSwarmForEUSIPs()
	r := rand.Intn(len(eus))
	return eus[r]
}

