package main

import (
	"encoding/json"
	"strings"
)

func booleanOfDotCompletion() bool {
	return stringOfDotStatusEachRepresentsAPiazzaJob() == strings.Repeat("4", numColorfulDisplayDots)
}

func stringOfDotStatusEachRepresentsAPiazzaJob() string {
	s := ""
	for i := 0; i < numColorfulDisplayDots; i++ {
		s += testVectorStatus(q[i])
	}
	return s
}

func testVectorStatus(p *testVector) string {
	if p.id4 != "" {
		return "4"
	}
	if p.id3 != "" {
		return "3"
	}
	if p.id2 != "" {
		return "2"
	}
	if p.id1 != "" {
		return "1"
	}
	return "0"
}
func stringOfDotDurationEachRepresentsAPiazzaJob() string {
	s := ""
	for i := 0; i < numColorfulDisplayDots; i++ {
		temp := "5"
		if q[i].id4 != "" {
			var dat map[string]interface{}
			if err3 := json.Unmarshal([]byte(q[i].id4), &dat); err3 != nil {
				panic(err3)
			}
			if dat["Results"] != nil {
				if dat["Results"].(float64) <= 16 {
					temp = "3"
				}
				if dat["Results"].(float64) <= 12 {
					temp = "2"
				}
				if dat["Results"].(float64) <= 8 {
					temp = "1"
				}
				if dat["Results"].(float64) <= 4 {
					temp = "0"
				}
				s += temp
			} else {
				s += temp
			}
		} else {
			s += temp
		}
	}
	return s
}
func stringOfDotColor() string {
	s := ""
	for i := 0; i < numColorfulDisplayDots; i++ {
		s += translateIPToColor(extractIPFromID4ResultsString(q[i].id4))
	}
	return s
}
func extractIPFromID4ResultsString(s string) string {
	//Temporary function to read a field (Status) from a JSON string, e.g.
	//  {"Results":1,"Status":"52.88.226.0"}
	//Long term would definitely prefer to receive a JSON object instead.
	var r string

	ix := strings.Index(s, `"Status":"`)
	if ix != -1 {
		r = s[ix+len(`"Status":"`) : len(s)-len(`"}`)]
	}
	return r
}


func translateIPToColor(s string) string {
	if s == eusips[0] {
		return "G"
	}
	if s == eusips[1] {
		return "B"
	}
	return "X"
}

func stringOfDotCompletion() string {
	if booleanOfDotCompletion() {
		return "active... completed"
	}
	return "active"
}
