package main 
import(
    "net/http"
)
func main(){
    http.HandleFunc("/", serveHome)						// Route #1 for Home Page
    http.HandleFunc("/contact", serveContact)			// Route #2 for Contact Page
    http.ListenAndServe(":8080", nil)
}
func serveHome(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello World!"))
}
func serveContact(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello World!  This is a contact page"))
}