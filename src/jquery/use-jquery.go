package main

 import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strconv"
  "strings"
  "time"
 )

type HealthArray struct {
    myip string
    port8079 bool
    port8081 bool
    port8083 bool
    port8084 bool
    port8085 bool
    port8088 bool
}

func test8079() bool {
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
    resp, err := client.Get("http://localhost:8079")
    if err != nil {
        fmt.Println("Something went wrong for 8079")
        return false
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("Body: %s\n", body)
    //fmt.Printf("Error: %v\n", err)
    return strings.Contains(string(body), `Nexus`)
}

func test8081() bool {
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
    resp, err := client.Get("http://localhost:8081")
    if err != nil {
        fmt.Println("Something went wrong for 8081")
        return false
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("Body: %s\n", body)
    //fmt.Printf("Error: %v\n", err)
    return strings.Contains(string(body), `pz-gateway`)
}
func test8083() bool { timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
    resp, err := client.Get("http://localhost:8083")
    if err != nil {
        fmt.Println("Something went wrong for 8083")
        return false
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("Body: %s\n", body)
    //fmt.Printf("Error: %v\n", err)
    return strings.Contains(string(body), `pz-jobmanager`)
}
func test8084() bool {
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
    resp, err := client.Get("http://localhost:8084")
    if err != nil {
        //fmt.Println("Something went wrong for 8084")
        return false
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("Body: %s\n", body)
    //fmt.Printf("Error: %v\n", err)
    return strings.Contains(string(body), `Loader`)
}
func test8085() bool {
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
    resp, err := client.Get("http://localhost:8085")
    if err != nil {
        fmt.Println("Something went wrong for 8085")
        return false
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("Body: %s\n", body)
    //fmt.Printf("Error: %v\n", err)
    return strings.Contains(string(body), `pz-access`)
}
func test8088() bool {
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
    resp, err := client.Get("http://localhost:8088")
    if err != nil {
        fmt.Println("Something went wrong for 8088")
        return false
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("Body: %s\n", body)
    //fmt.Printf("Error: %v\n", err)
    return strings.Contains(string(body), `Piazza Service Controller`)
}
func myIPWithTimeout() string {
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
    resp, err := client.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
    if err != nil {
        fmt.Println("Something went wrong in myIPWithTimeout")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("Body: %s\n", body)
    //fmt.Printf("Error: %v\n", err)
    return string(body)
}

 func Work(w http.ResponseWriter, r *http.Request) {
  piazzaBox := /*(params.containers) ?: */myIPWithTimeout()
  externalUserService := myIPWithTimeout()
  workers := 0 /* for now, don't use this or multiple invocations (no go global var) */
  const NUM_COLORFUL_DISPLAY_DOTS = 16 //100
  var q [NUM_COLORFUL_DISPLAY_DOTS] *TestVector


  //work() is invoked from the gsp page each time the browser is refreshed.
  //we don't want multiple workers, so each worker gets a worker number.
  //if the worker number doesn't match the number of workers, then
  //the worker exits.
  workers++
  iamworker := workers
  for i:=0; i<NUM_COLORFUL_DISPLAY_DOTS; i++ {
    fmt.Println("0th pass:i=" + strconv.Itoa(i))
    q[i] = NewTestVector(piazzaBox, externalUserService)
  }
  //This is an arbitrarily large number. We might as well loop forever.
  //MAX_ITERATION_TO_CALL_TEST_VECTOR := 4 //64000

  time.Sleep(10 * time.Second)
  for i:=0; i<NUM_COLORFUL_DISPLAY_DOTS; i++ {
    fmt.Println("1st pass:i=" + strconv.Itoa(i))
    q[i] = nextStep(*q[i])
  }

  time.Sleep(10 * time.Second)
  for i:=0; i<NUM_COLORFUL_DISPLAY_DOTS; i++ {
    fmt.Println("2nd pass:i=" + strconv.Itoa(i))
    q[i] = nextStep(*q[i])
  }

  time.Sleep(10 * time.Second)
  for i:=0; i<NUM_COLORFUL_DISPLAY_DOTS; i++ {
    fmt.Println("3rd pass:i=" + strconv.Itoa(i))
    q[i] = nextStep(*q[i])
  }

  time.Sleep(10 * time.Second)
  for i:=0; i<NUM_COLORFUL_DISPLAY_DOTS; i++ {
    fmt.Println("4th pass:i=" + strconv.Itoa(i))
    q[i] = nextStep(*q[i])
  }

  time.Sleep(10 * time.Second)
  for i:=0; i<NUM_COLORFUL_DISPLAY_DOTS; i++ {
    fmt.Println("5th pass:i=" + strconv.Itoa(i))
    q[i] = nextStep(*q[i])
  }

  fmt.Printf("Work(%s,%s,%d)", piazzaBox, externalUserService, iamworker)
/*
 time.Sleep(n * time.Second
        //s/m: this can be made more groovy, right? spread operator?
        for (int i=0; i<NUM_COLORFUL_DISPLAY_DOTS; i++) {
            q[i] = new TestVector(piazzaBox, externalUserService)
        }

        //This is an arbitrarily large number. We might as well loop forever.
        def MAX_ITERATION_TO_CALL_TEST_VECTOR = 64000

        (0..MAX_ITERATION_TO_CALL_TEST_VECTOR).each {
            if ((iamworker == workers) && (!booleanOfDotCompletion())) {
                def HEALTH_CHECK_SERVICES_EVERY_SO_OFTEN = 10
                if (it % HEALTH_CHECK_SERVICES_EVERY_SO_OFTEN == 0) {
                    qhealth[0] = new HealthArray()
                }
                for (int i=0; i<NUM_COLORFUL_DISPLAY_DOTS; i++) {
                    q[i].nextstep()
                }
            }
        }
        render "work()#$iamworker exits (num workers is $workers)"
*/
 }

 func Home(w http.ResponseWriter, r *http.Request) {
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
   w = foox.dotCompletion
   paper.text(200, 400, w);
   var THROWAWAY_BR_CHARS = 0
   var ROWS_PER_SQUARE = Math.sqrt(r.length) //10
   var COLS_PER_SQUARE = ROWS_PER_SQUARE
   var ORIGINX = 40
   var ORIGINY = 100
   var STOPLIGHT_YELLOW = '#FAD201'
   var STOPLIGHT_GREEN = '#27E833'
   var MEDIUM_GREEN =    '#27C833'
   var DARK_GREEN =      '#27A833'
   for (row = 0; row < ROWS_PER_SQUARE; row++) {
    for (col = 0; col < COLS_PER_SQUARE; col++) {
       var stroke
       var fill
       var filltext = ''
       var substring_start = THROWAWAY_BR_CHARS+row*ROWS_PER_SQUARE+col
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
           if ((tchar == '0') ||
               (tchar == '1')) {
               fill = DARK_GREEN
               filltext = 'S'
           }
           if (tchar == '2') {
               fill = MEDIUM_GREEN
               filltext = 'M'
           }
           if (tchar == '3') {
               fill = STOPLIGHT_GREEN
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
   byu := []byte(`{"dotStatus":"1124012411241124", ` +
       `"dotDuration":"1122331122331122", ` + 
       `"squareHealth":"pz-jobmanager?", ` +
       `"dotCompletion":"` + myIPWithTimeout() + `\n` +
       `nexus: ` + strconv.FormatBool(test8079()) +  ` \n` +
       `gateway: ` + strconv.FormatBool(test8081()) +  ` \n` +
       `jobmanager: ` + strconv.FormatBool(test8083()) +  ` \n` +
       `Loader: ` + strconv.FormatBool(test8084()) +  ` \n` +
       `access: ` + strconv.FormatBool(test8085()) +  ` \n` +
       `servicecontroller: ` + strconv.FormatBool(test8088()) +  ` \n` +
       `"}`)
   w.Write(byu)
 }

 func receiveAjax(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
   ajax_post_data := r.FormValue("ajax_post_data")
   fmt.Println("Receive ajax post data string ", ajax_post_data)
   w.Write([]byte("<h2>after<h2>"))
  }
 }

type TestVector struct {
    EXTERNAL_USER_SERVICE string
    PIAZZA_PRIME_BOX string
    id1 string
    id2 string
    id3 string
    id4 string
}

func NewTestVector(piazzaBox string, externalUserService string) *TestVector {
    p := new(TestVector)
    p.EXTERNAL_USER_SERVICE = "http://" + externalUserService + ":8078/green/timer/external"
    p.PIAZZA_PRIME_BOX = piazzaBox
    p.id1 = ""
    p.id2 = ""
    p.id3 = ""
    p.id4 = ""
    //fmt.Println(p.EXTERNAL_USER_SERVICE + p.PIAZZA_PRIME_BOX)
    return p
}

func nextStep(p TestVector) *TestVector {
    if p.id1 == "" {
        p.id1 = pz1(p)
        if p.id1 == "" { return &p }
    }

    if p.id2 == "" {
        p.id2 = pz2(p)
        if p.id2 == "" { return &p }
    }

    if p.id3 == "" {
        p.id3 = pz3(p)
        if p.id3 == "" { return &p }
    }

    if p.id4 == "" {
        p.id4 = pz4(p)
        if p.id4 == "" { return &p }
    }

    return &p
}

func oldNextStep(p TestVector) *TestVector {
/* someday
    if p.id5 == "" && p.id4 != "" {
        p.id5 = pz5(p)
    }
    if p.id4 == "" && p.id3 != "" {
        p.id4 = pz4(p)
    }
*/
    if p.id3 == "" && p.id2 != "" {
        p.id3 = pz3(p)
        if p.id3 != "" {
            p.id4 = pz4(p)
        }
    }

    if p.id2 == "" && p.id1 != "" {
        p.id2 = pz2(p)
        if p.id2 != "" {
            p.id3 = pz3(p)
        }
        if p.id3 != "" {
            p.id4 = pz4(p)
        }
    }
    if p.id1 == "" {
        p.id1 = pz1(p)

        if p.id1 != "" {
            p.id2 = pz2(p)
        }
        if p.id2 != "" {
            p.id3 = pz3(p)
        }
        if p.id3 != "" {
            p.id4 = pz4(p)
        }
    }

    return &p
}

func pz1curl(p TestVector, body2 string) string {
/* s/m: reinstate these four lines 
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
*/

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

body := strings.NewReader(body2)
req, err := http.NewRequest("POST", "http://" + p.PIAZZA_PRIME_BOX + ":8081/service", body)
if err != nil {
	// handle err
}
req.Header.Set("Content-Type", "application/json")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()
    body3, err8 := ioutil.ReadAll(resp.Body)
    if err8 != nil {
        // handle err
    }
    var dat map[string]interface{}
    if err3 := json.Unmarshal(body3, &dat); err3 != nil {
        panic(err3)
    }
    data := dat["data"].(map[string]interface{})
    foo := data["serviceId"]
    return foo.(string) 
}
/*
{
  "type" : "service-id",
  "data" : {
    "serviceId" : "7d2f1c07-0238-433a-bf72-b5f20a215b6c"
  }
}
*/
func pz1(p TestVector) string {
    body2 := `{"url":"REPLACEME","method":"GET","contractUrl":"REPLACEME/","resourceMetadata":{"name":"Hello World Test","description":"Hello world test","classType":{"classification":"UNCLASSIFIED"}}}`
    body2 = strings.Replace(body2, "REPLACEME", p.EXTERNAL_USER_SERVICE, -1)
    pz1curlresult := pz1curl(p, body2)
    fmt.Println(pz1curlresult)
    return pz1curlresult 
}

func pz2curl(p TestVector, body2 string) string {
/* s/m: reinstate these four lines 
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
*/

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

body := strings.NewReader(body2)
req, err := http.NewRequest("POST", "http://" + p.PIAZZA_PRIME_BOX + ":8081/job", body)
if err != nil {
	// handle err
}
req.Header.Set("Content-Type", "application/json")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()

    body3, err8 := ioutil.ReadAll(resp.Body)
    if err8 != nil {
        // handle err
    }
    var dat map[string]interface{}
    if err3 := json.Unmarshal(body3, &dat); err3 != nil {
        panic(err3)
    }
    data := dat["data"].(map[string]interface{})
    foo := data["jobId"]
    return foo.(string) 
}

func pz2(p TestVector) string {
    body2 := `{"type":"execute-service","data":{"serviceId":"REPLACEME","dataInputs":{},"dataOutput":[{"mimeType":"application/json","type":"text"}]}}`
    body2 = strings.Replace(body2, "REPLACEME", p.id1, -1)
    pz2curlresult := pz2curl(p, body2)
    fmt.Println(" " + pz2curlresult)
    return pz2curlresult
}

func pz3curl(p TestVector, body2 string) string {
/* s/m: reinstate these four lines
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
*/

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

req, err := http.NewRequest("GET", "http://" + p.PIAZZA_PRIME_BOX + ":8081/job/" + p.id2, nil)
if err != nil {
	// handle err
}
req.Header.Set("Content-Type", "application/json")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()

    body3, err8 := ioutil.ReadAll(resp.Body)
    if err8 != nil {
        // handle err
    }
    var dat map[string]interface{}
    if err3 := json.Unmarshal(body3, &dat); err3 != nil {
        panic(err3)
    }
    //fmt.Println("***dat***:")
    //fmt.Println(dat)
    data := dat["data"].(map[string]interface{})
    //fmt.Println("***data***:")
    //fmt.Println(data)
    if data["result"] == nil {
       //fmt.Println("The result is nil")
       return ""
    } 

    result := data["result"].(map[string]interface{})
    //fmt.Println("***result***:")
    //fmt.Println(result)
    foo := result["dataId"]
    return foo.(string)
}

func pz3(p TestVector) string {
    body2 := `{"type":"execute-service","data":{"serviceId":"REPLACEME","dataInputs":{},"dataOutput":[{"mimeType":"application/json","type":"text"}]}}`
    body2 = strings.Replace(body2, "REPLACEME", p.id2, -1)
    pz3curlresult := pz3curl(p, body2)
    fmt.Println(" +" + pz3curlresult)
    return pz3curlresult
}

func pz4curl(p TestVector, body2 string) string {
/* s/m: reinstate these four lines
    timeout := time.Duration(3 * time.Second)
    client := http.Client {
        Timeout: timeout,
    }
*/

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
//http://$PIAZZA_PRIME_BOX:8081/data/${id3.data.result.dataId
temp := "http://" + p.PIAZZA_PRIME_BOX + ":8081/data/" + p.id3
//fmt.Println(temp)
req, err := http.NewRequest("GET", temp, nil)
if err != nil {
	// handle err
}
req.Header.Set("Content-Type", "application/json")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()

    body3, err8 := ioutil.ReadAll(resp.Body)
    if err8 != nil {
        // handle err
    }
    var dat map[string]interface{}
    if err3 := json.Unmarshal(body3, &dat); err3 != nil {
        panic(err3)
    }
    //fmt.Println("***dat***:")
    //fmt.Println(dat)
/*
map[type:error message:Data not found: 2690beab-3beb-459d-9401-d034f6e8be9e origin:Access]
*/
    if dat["data"] == nil {
        return ""
    }
//id4?.data?.dataType?.content
// http://$PIAZZA_PRIME_BOX:8081/data/${id3.data.result.dataId}
    data := dat["data"].(map[string]interface{})
    //fmt.Println("***data***:")
    //fmt.Println(data)
    if data["dataType"] == nil {
       fmt.Println("pz4curl The dataType is nil")
       return ""
    } 

    result := data["dataType"].(map[string]interface{})
    //fmt.Println("***result***:")
    //fmt.Println(result)
    foo := result["content"]
    return foo.(string)
}

func pz4(p TestVector) string {
    body2 := `{"type":"execute-service","data":{"serviceId":"REPLACEME","dataInputs":{},"dataOutput":[{"mimeType":"application/json","type":"text"}]}}`
    body2 = strings.Replace(body2, "REPLACEME", p.id3, -1)
    pz4curlresult := pz4curl(p, body2)
    fmt.Println("  -" + pz4curlresult)
    return pz4curlresult
}

func NewHealthArray() *HealthArray {
    p := new(HealthArray)
    p.port8079 = test8079()
    p.port8081 = test8081()
    p.port8083 = test8083()
    p.port8084 = test8084()
    p.port8085 = test8085()
    p.port8088 = test8088()
    return p
}

 func main() {

  ha := NewHealthArray()
  fmt.Printf("test8079: %s\n", ha.port8079)
  // http.Handler
  mux := http.NewServeMux()
  mux.HandleFunc("/", Home)
  mux.HandleFunc("/greentimerwork", Work)
  mux.HandleFunc("/receive", receiveAjax)
  mux.HandleFunc("/status", greenstatus)
  http.ListenAndServe(":8077", mux)
 }
/*
    def pz3() {
        assert id2

        def returnval = null
        def myprocess4 = [ 'bash', '-c', "curl -v -k -X GET -H \"Content-Type: application/json\" http://$PIAZZA_PRIME_BOX:8081/job/${id2?.data?.jobId}" ].execute()
        def myprocess4AsText
        try {
            myprocess4AsText = myprocess4.text
            def result4AsJson = new JsonSlurper().parseText(myprocess4AsText)
            result4AsJson.data.result.dataId
            returnval = result4AsJson
        } catch(e) {}
        returnval
    }

    def pz4() {
        assert id3
        if (id3?.data) {
            def myprocess5 = [ 'bash', '-c', "curl -v -k -X GET -H \"Content-Type: application/json\" http://$PIAZZA_PRIME_BOX:8081/data/${id3.data.result.dataId}" ].execute()
            String myprocess5AsText = myprocess5.text
            def result5AsJson = new JsonSlurper().parseText(myprocess5AsText)
            id5 = result5AsJson?.data?.dataType?.content //this eliminates the need for an extra iteration to populate id5
            result5AsJson
        }
    }

    def pz5() {
        assert id4
        id4?.data?.dataType?.content
    }
*/
