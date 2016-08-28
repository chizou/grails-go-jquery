package main

 import (
  "fmt"
  "io/ioutil"
  "net/http"
 )

func mycurl() {
    resp, err := http.Get("http://52.42.17.198:8080/alabanza/praise")
    if err != nil {
        fmt.Println("Something went wrong")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("Body: %s\n", body)
    fmt.Printf("Error: %v\n", err)
}

 func Home(w http.ResponseWriter, r *http.Request) {
  html := `<head>	
 <script src="http://cdnjs.cloudflare.com/ajax/libs/raphael/2.1.0/raphael-min.js"></script>

 <script src='http://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js'></script>

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

//  });
 }
 function getMessages1() {

  $.post('/green/timer/status', {}, function(r) {
   var foox = JSON.parse(r)
   console.log(foox);
   r = foox.dotStatus
   t = foox.dotDuration
   v = foox.squareHealth
   w = foox.dotCompletion
   $('#controllerresults').html(w + '<br>' + v);
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

 $.post('/green/timer/work', {}, function(r) {
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
   //byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
   byt := []byte(`{"dotStatus":"1124012411241124","dotDuration":"1122331122331122", "squareHealth":"pz-jobmanager?", "dotCompletion":"unused"}`)
   w.Write(byt)
 }

 func receiveAjax(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
   ajax_post_data := r.FormValue("ajax_post_data")
   fmt.Println("Receive ajax post data string ", ajax_post_data)
   w.Write([]byte("<h2>after<h2>"))
  }
 }

 func main() {
  // http.Handler
  mux := http.NewServeMux()
  mux.HandleFunc("/", Home)
  mux.HandleFunc("/receive", receiveAjax)
  mux.HandleFunc("/status", greenstatus)
  mycurl()
  http.ListenAndServe(":8080", mux)
 }


