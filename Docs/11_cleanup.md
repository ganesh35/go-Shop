# #11: Code cleanup

### Folders moved
**/src/lib**  ->   **/src/main/lib**
### Files moved
**/src/main/users.go**  ->   **/src/main/components/users/controller.go**

### Code changed
```go
    ...
    "strings"
    "main/lib"
    GUsers "main/components/users"
)
var gConfig lib.GConfig
var gLog lib.GLog
...
func main() {
	defer close()
    users := GUsers.Users{
        Store: map[string]* GUsers.User{},
    }
    ...
```
### Dependencies:
+ [Go-Json-Rest v3.3.1](https://github.com/ant0ine/go-json-rest)
+ [JWT Middleware for Go-Json-Rest](https://github.com/StephanDollberg/go-json-rest-middleware-jwt)
+ [go-semver - Semantic Versioning Library](https://github.com/ant0ine/go-json-rest#api-versioning)

---
### Step #1: File and Folder Structure
-- gUser
+ src     :  # Source code goes here
    - main  :  # main Package
        - web.go  :  # main entry file
        - countries.go : # to get the countries list
        - SemVerMiddleware.go :  # SemVer Middleware implementation 
        - components
            - users
                - controller.go : #  User actions
        - lib : # necessary library files goes here
            - gconfig.go : #  json parser and Config Structure
            - glog.go : #  json parser and Config Structure
            - ghelpers.go : #  Necessary helper functions
            - gmail.go : #  SendMail using GMail smtp server
+ bin     :  # Generated binary files and configuration settings goes here
    - etc
        - config.json  : # JSON configuration file
+ pkg
+ Docs       :  # Documentation goes here
  - 01_hello.md
  - ...
  - 10_cleanup.md
+ .gitignore
+ LICENCE
+ README.md
---

