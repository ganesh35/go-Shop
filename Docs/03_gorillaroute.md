# #03: Advanced Routing with Gorilla
Advanced routing with Gorilla/mux.  This program is able to take the url params and shows the appropriate pages.
### Dependency:
gorilla/mux [v1.1 "Long Overdue"](https://github.com/gorilla/mux)

---
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
---
### Step #2:  Setting paths and installing necessary dependencies
- Open command prompt / terminal (in windows : Start -> Run -> cmd )
```sh
$ cd /GoCode/gUser
```
```sh
#  Linux / Unix / gitbash
$ export GOPATH=/GoCode/gUser
$ export GOBIN=/GoCode/gUser/bin
$ go get "github.com/gorilla/mux"
```
```sh
#  Windows
d:\> set GOPATH=d:\GoCode\gUser
d:\> set GOBIN=d:\GoCode\gUser\bin
d:\> go get "github.com/gorilla/mux"
```
---
### Step #3:  Gorilla/mux routing
File: **src/main/web.go**  
```go
package main 
import(
	"net/http"
	"github.com/gorilla/mux"
)
func main(){
	serveWeb()
}
func serveWeb(){
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/", serveContent)
	gorillaRoute.HandleFunc("/{page_alias}", serveContent)  // Dynamic url values
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}
func serveContent(w http.ResponseWriter, r *http.Request){
	urlParams := mux.Vars(r)
	page_alias := urlParams["page_alias"]
	if page_alias == ""{
		page_alias = "home"
	}
	w.Write([]byte("Hello World!" + page_alias))
}
```
---
### Step #4:  Running the program
- Open command prompt / terminal (in windows : Start -> Run -> cmd )
```sh
#  Linux / Unix / gitbash
$ go run src/main/web.go
```
```sh
#  Windows
d:\> go run src/main/web.go
```
---
### Step #5: Output over Browser
- Open the browser (Firefox / IE / Chrome ...)
- Enter http://localhost:8080 in the URL
```html
http://localhost:8080
Hello World!
```
```html
http://localhost:8080/SomeExample
Hello World!  SomeExample
```
---