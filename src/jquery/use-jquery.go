package main

 import (
 	"fmt"
 	"net/http"
 )

 func Home(w http.ResponseWriter, r *http.Request) {
 	html := `<head>	
                    <script src="http://cdnjs.cloudflare.com/ajax/libs/raphael/2.1.0/raphael-min.js"></script>
                    <script src='http://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js'></script>
              </head>    
                  <html><body>
                  <h1>Golang Jquery AJAX example</h1>

                  <div id='result'><h3>before</h3></div><br><br>

                  <input id='ajax_btn' type='button' value='POST via AJAX to Golang server'>


                  </body></html>

                   <script>
 var paper = Raphael(0, 0, 1200, 1200);

 paper.circle(100, 100, 20).attr({ "stroke": "blue", "fill": "red"});
                   $(document).ready(function () { 
                         $('#ajax_btn').click(function () {
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { ajax_post_data: 'hello'},
                               success : function(data) {
                                 alert('ajax data posted');
                                 $('#result').html(data);
                               },
                             });
                          });
                    });
                    </script>`

 	w.Write([]byte(fmt.Sprintf(html)))

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

 	http.ListenAndServe(":8080", mux)
 }

