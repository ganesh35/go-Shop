# #02: Basic Server 
Basic server with basic routing.  This program provides 2 routes to browse. One for Home Page and other for a Contact Page.
### Step #1: File and Folder Structure
-- gUser
+ src     :  # Source code goes here
    - main  :  # main Package
        - web.go  :  # main entry file
+ bin     :  # Generated binary files and configuration settings goes here
+ pkg
+ Docs       :  # Documentation goes here
  - 01_hello.md
+ .gitignore
+ LICENCE
+ README.md
 
### Step #2:  Basic Server
File: **src/main/web.go**  
```go
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
```

### Step #3:  Setting paths and running the program
- Open command prompt / terminal (in windows : Start -> Run -> cmd )
```sh
$ cd /GoCode/gUser
```
```sh
#  Linux / Unix / gitbash
$ export GOPATH= /GoCode/gUser
$ export GOBIN= /GoCode/gUser/bin
$ go run src/main/web.go
```
```sh
#  Windows
d:\> set GOPATH=d:\GoCode\gUser
d:\> set GOBIN=d:\GoCode\gUser\bin
d:\> go run src/main/web.go
```

### Step #4: Output over Browser
- Open the browser (Firefox / IE / Chrome ...)
- Enter http://localhost:8080 in the URL
```html
http://localhost:8080
Hello World!
```
```html
http://localhost:8080/contact
Hello World!  This is a contact page
```
