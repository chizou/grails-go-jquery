package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const numColorfulDisplayDots = 100 //16
var q [numColorfulDisplayDots] *testVector
var qhealth = newHealthArray()
var containers []string

func myIPWithTimeout() string {
	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
	if err != nil {
		fmt.Println("Something went wrong in myIPWithTimeout")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func work(w http.ResponseWriter, r *http.Request) {
	var piazzaBox string
	if containers != nil {
		piazzaBox = containers[0]
	} else {
		piazzaBox = myIPWithTimeout()
	}

	workers := 0 /* for now, don't use this or multiple invocations (no go global var) */

	//work() is invoked from the gsp page each time the browser is refreshed.
	//we don't want multiple workers, so each worker gets a worker number.
	//if the worker number doesn't match the number of workers, then
	//the worker exits.
	workers++
	iamworker := workers
	for i := 0; i < numColorfulDisplayDots; i++ {
		q[i] = newTestVector(piazzaBox, randomExternalUserService())
	}

	var maxIterationToCallTestVector = 64000
	for j := 0; j < maxIterationToCallTestVector; j++ {
		if iamworker == workers && !booleanOfDotCompletion() {
			var healthCheckServicesEverySoOften = 10
			if j%healthCheckServicesEverySoOften == 0 {
				qhealth = updateHealthArray(*qhealth)
			}
			for i := 0; i < numColorfulDisplayDots; i++ {
				q[i] = nextStep(*q[i])
			}
		}
	}
}

func stringOfDotStatusEachRepresentsAPiazzaJob() string {
	s := ""
	for i := 0; i < numColorfulDisplayDots; i++ {
		s += testVectorStatus(q[i])
	}
	return s
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

func translateIPToColor(s string) string {
	if s == eusips[0] {
		return "G"
	}
	if s == eusips[1] {
		return "B"
	}
	return "X"
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

func stringOfDotColor() string {
	s := ""
	for i := 0; i < numColorfulDisplayDots; i++ {

		s += translateIPToColor(extractIPFromID4ResultsString(q[i].id4))
	}
	return s
}

func stringOfDotCompletion() string {
	if booleanOfDotCompletion() {
		return "active... completed"
	}
	return "active"
}

func booleanOfDotCompletion() bool {
	return stringOfDotStatusEachRepresentsAPiazzaJob() == strings.Repeat("4", numColorfulDisplayDots)
}

func home(w http.ResponseWriter, r *http.Request) {
	containers = r.URL.Query()["containers"]

	html := `<head>	
 <script src="http://cdnjs.cloudflare.com/ajax/libs/raphael/2.1.0/raphael-min.js">
 </script>

 <script src='http://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js'>
 </script>

 </head>
 <body>
  <script>
 var si1

 function getMessages2() {
  $.post('status', {}, function(r) {
   var foox = JSON.parse(r)
   r = foox.dotStatus
   t = foox.dotDuration
   v = foox.squareHealth
   paper.text(800, 400, v);
   x = foox.dotColor
   var THROWAWAY_BR_CHARS = 0
   var ROWS_PER_SQUARE = Math.sqrt(r.length) //10
   var COLS_PER_SQUARE = ROWS_PER_SQUARE
   var ORIGINX = 40
   var ORIGINY = 100
   var STOPLIGHT_YELLOW = '#FAD201'
   var STOPLIGHT_GREEN = '#27E833'
   var MEDIUM_GREEN =    '#27C833'
   var DARK_GREEN =      '#27A833'
   var STOPLIGHT_BLUE = '#CCE5FF'
   var MEDIUM_BLUE =    '#66B2FF'
   var DARK_BLUE =      '#0080FF'
   for (row = 0; row < ROWS_PER_SQUARE; row++) {
    for (col = 0; col < COLS_PER_SQUARE; col++) {
       var stroke
       var fill
       var filltext = ''
       var substring_start = THROWAWAY_BR_CHARS+row*ROWS_PER_SQUARE+col
       //console.log(substring_start)
       var rchar = r.substring(substring_start, substring_start+1)
       if (rchar == '0') {
           stroke = "white"
           fill = "white"
       }
       if (rchar == '1') {
           stroke = STOPLIGHT_YELLOW
           fill = "white"
       }
       if (rchar == '2') {
           stroke = STOPLIGHT_YELLOW
           fill = STOPLIGHT_YELLOW
       }
       if (rchar == '4') {
           fill = 'pink'
           filltext = 'error'
           var tchar = t.substring(substring_start, substring_start+1)
           var xchar = x.substring(substring_start, substring_start+1)
           if ((tchar == '0') ||
               (tchar == '1')) {
               if (xchar == 'G') {
                   fill = DARK_GREEN
               }
               if (xchar == 'B') {
                   fill = DARK_BLUE
               }
               filltext = 'S'
           }
           if (tchar == '2') {
               if (xchar == 'G') {
                   fill = MEDIUM_GREEN
               }
               if (xchar == 'B') {
                   fill = MEDIUM_BLUE
               }
               filltext = 'M'
           }
           if (tchar == '3') {
               if (xchar == 'G') {
                   fill = STOPLIGHT_GREEN
               }
               if (xchar == 'B') {
                   fill = STOPLIGHT_BLUE
               }
               filltext = 'L'
           }
           stroke = fill
       }
       paper.circle(ORIGINX+46*col, ORIGINY+46*row, 20).attr({ "stroke": stroke, "fill": fill });
       paper.text(ORIGINX+46*col, ORIGINY+46*row, filltext)
    }
   }
  },'html');
 }

 $.post('greentimerwork', {}, function(r) {
  //$('#workresults').html('This is urlBar.gsps work plus work result: ' + r);
 },'html');

 var paper = Raphael(0, 0, 1200, 1200);
 console.log('xxx');
 si1 = setInterval(getMessages2, 200);

  </script>
  <!--
  <div id="workresults" >This is #workresults in urlBar.gsp...</div>
  -->
  <div id="controllerresults" ></div>
 </body>`

	w.Write([]byte(fmt.Sprintf(html)))

}

func greenstatus(w http.ResponseWriter, r *http.Request) {
	var s1 string
	var s2 string
	var s4 string
	var s5 string
	if q[0] == nil {
		s1 = `1124012411241124`
		s2 = `1122331122331122`
		s4 = `hello`
		s5 = `hello`
	} else {
		s1 = stringOfDotStatusEachRepresentsAPiazzaJob()
		s2 = stringOfDotDurationEachRepresentsAPiazzaJob()
		s4 = statusString(qhealth)
		s5 = stringOfDotColor()
	}
	s := `{"dotStatus":"` + s1 + `",` +
		`"dotDuration":"` + s2 + `",` +
		`"squareHealth":"` + s4 + `",` +
		`"dotColor":"` + s5 + `"}`

	byu := []byte(s)
	w.Write(byu)
}

type testVector struct {
	externalUserService string
	piazzaPrimeBox      string
	id1                 string
	id2                 string
	id3                 string
	id4                 string
}

func newTestVector(piazzaBox string, externalUserService string) *testVector {
	p := new(testVector)
	p.externalUserService = "http://" + externalUserService + ":8077/external"
	p.piazzaPrimeBox = piazzaBox
	p.id1 = ""
	p.id2 = ""
	p.id3 = ""
	p.id4 = ""
	return p
}

func nextStep(p testVector) *testVector {
	if p.id1 == "" {
		p.id1 = pz1(p)
		if p.id1 == "" {
			return &p
		}
	}

	if p.id2 == "" {
		p.id2 = pz2(p)
		if p.id2 == "" {
			return &p
		}
	}

	if p.id3 == "" {
		p.id3 = pz3(p)
		if p.id3 == "" {
			return &p
		}
	}

	if p.id4 == "" {
		p.id4 = pz4(p)
		if p.id4 == "" {
			return &p
		}
	}

	return &p
}

func pz1curl(p testVector, body2 string) string {

	body := strings.NewReader(body2)
	req, err := http.NewRequest("POST", "http://"+p.piazzaPrimeBox+":8081/service", body)
	if err != nil {
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body3, err8 := ioutil.ReadAll(resp.Body)
	if err8 != nil {
	}
	var dat map[string]interface{}
	if err3 := json.Unmarshal(body3, &dat); err3 != nil {
		panic(err3)
	}
	data := dat["data"].(map[string]interface{})
	foo := data["serviceId"]
	return foo.(string)
}

func pz1(p testVector) string {
	body2 := `{"url":"REPLACEME","method":"GET","contractUrl":"REPLACEME/","resourceMetadata":{"name":"Hello World Test","description":"Hello world test","classType":{"classification":"UNCLASSIFIED"}}}`
	body2 = strings.Replace(body2, "REPLACEME", p.externalUserService, -1)
	pz1curlresult := pz1curl(p, body2)
	return pz1curlresult
}

func pz2curl(p testVector, body2 string) string {
	body := strings.NewReader(body2)
	req, err := http.NewRequest("POST", "http://"+p.piazzaPrimeBox+":8081/job", body)
	if err != nil {
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()

	body3, err8 := ioutil.ReadAll(resp.Body)
	if err8 != nil {
	}
	var dat map[string]interface{}
	if err3 := json.Unmarshal(body3, &dat); err3 != nil {
		panic(err3)
	}
	data := dat["data"].(map[string]interface{})
	foo := data["jobId"]
	return foo.(string)
}

func pz2(p testVector) string {
	body2 := `{"type":"execute-service","data":{"serviceId":"REPLACEME","dataInputs":{},"dataOutput":[{"mimeType":"application/json","type":"text"}]}}`
	body2 = strings.Replace(body2, "REPLACEME", p.id1, -1)
	pz2curlresult := pz2curl(p, body2)
	return pz2curlresult
}

func pz3curl(p testVector, body2 string) string {
	req, err := http.NewRequest("GET", "http://"+p.piazzaPrimeBox+":8081/job/"+p.id2, nil)
	if err != nil {
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()

	body3, err8 := ioutil.ReadAll(resp.Body)
	if err8 != nil {
	}
	var dat map[string]interface{}
	if err3 := json.Unmarshal(body3, &dat); err3 != nil {
		panic(err3)
	}
	data := dat["data"].(map[string]interface{})
	if data["result"] == nil {
		return ""
	}

	result := data["result"].(map[string]interface{})
	foo := result["dataId"]
	return foo.(string)
}

func pz3(p testVector) string {
	body2 := `{"type":"execute-service","data":{"serviceId":"REPLACEME","dataInputs":{},"dataOutput":[{"mimeType":"application/json","type":"text"}]}}`
	body2 = strings.Replace(body2, "REPLACEME", p.id2, -1)
	pz3curlresult := pz3curl(p, body2)
	return pz3curlresult
}

func pz4curl(p testVector, body2 string) string {
	temp := "http://" + p.piazzaPrimeBox + ":8081/data/" + p.id3
	req, err := http.NewRequest("GET", temp, nil)
	if err != nil {
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()

	body3, err8 := ioutil.ReadAll(resp.Body)
	if err8 != nil {
	}
	var dat map[string]interface{}
	if err3 := json.Unmarshal(body3, &dat); err3 != nil {
		panic(err3)
	}
	if dat["data"] == nil {
		return ""
	}
	data := dat["data"].(map[string]interface{})
	if data["dataType"] == nil {
		fmt.Println("pz4curl The dataType is nil")
		return ""
	}

	result := data["dataType"].(map[string]interface{})
	foo := result["content"]
	return foo.(string)
}

func pz4(p testVector) string {
	body2 := `{"type":"execute-service","data":{"serviceId":"REPLACEME","dataInputs":{},"dataOutput":[{"mimeType":"application/json","type":"text"}]}}`
	body2 = strings.Replace(body2, "REPLACEME", p.id3, -1)
	pz4curlresult := pz4curl(p, body2)
	return pz4curlresult
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/greentimerwork", work)
	mux.HandleFunc("/status", greenstatus)
	mux.HandleFunc("/external", simulatedExternalUserService)
	http.ListenAndServe(":8077", mux)
}

/*
   def dots() {
   }
   def zdots() {
       zwork()
       string()
   }
   def string() {
       render stringOfDotStatusEachRepresentsAPiazzaJob() + '\n'
   }
*/
/* s/m: reinstate these four lines into pz1curl,...pz4curl
   timeout := time.Duration(3 * time.Second)
   client := http.Client {
       Timeout: timeout,
   }
*/
