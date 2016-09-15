#1 Hello World!
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
 
### Step #2:  Hello World!
File: **src/main/web.go**  
```go
package main
import "fmt"
func main() {
	fmt.Println("Hello, World!")
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
Hello, World!
```
```sh
#  Windows
d:\> set GOPATH=d:\GoCode\gUser
d:\> set GOBIN=d:\GoCode\gUser\bin
d:\> go run src/main/web.go
Hello, World!
```
