package main

import (
	"fmt"      // formatting and printing values to the console.
	"log"      // logging messages to the console.
	"net/http" // Used for build HTTP servers and clients.
)

/*
Create a simple go lang HTTP server in a few lines. Register one path that
serves an html file and another path that will show some string!
*/

// Port we listen on.
const portNum string = ":8080"

// Handler functions.
/*
when we make a request to a HTTP server, the server must process
this and send back a response back to the client. Our http.ResponseWriter
is used to write this response back to the client, and the http.Request
object contains information about the incoming request.
*/
func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./cv_resume_template.html")
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info Page")
}

func main() {
	log.Println("Starting our simple http server.")

	//Registering our handler functions, and creating paths
	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	/*
		Using the http.ListenAndServe function we can start a
		HTTP server and "listen" for incoming requests on a specified
		network address and port. Once a reuqest is received, it is
		passed to the appropriate handler function for processing.

		Now let's listen on port 8080 and for our ServeMux we are just
		using the default http.ServeMux() by using nil.
	*/
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
