package main

import {
"fmt"
"net/http"
} 

type Message struct{
	 Name string `jason:"name"`
	 Content string `jason:"content"`
}
func addint(a inr, b int) int{
	return a+b
}
func gethandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go API")

}
func main (){
	http.HandleFunc("/", gethandler)
    pNumber :=8091
	fmt.Printf("Server is running on the port: %d\n", pNumber)
	http.ListenAndServe(pNumber, Mad)
}