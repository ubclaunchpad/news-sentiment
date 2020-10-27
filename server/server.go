package main


import (
    "fmt"
    "log"
    "net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
	}
	

	fmt.Fprintf(w, "News App")

}




func main() {

	fileServer := http.FileServer(http.Dir("./static")) // New code
    http.Handle("/", fileServer) // New code
	http.HandleFunc("/hello", helloHandler)



	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}


