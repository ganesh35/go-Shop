# #04: JWT
Authentication via Json Web Tokens
### Dependency:
+ [Go-Json-Rest v3.3.1](https://github.com/ant0ine/go-json-rest)
+ [JWT Middleware for Go-Json-Rest](https://github.com/StephanDollberg/go-json-rest-middleware-jwt)

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
$ go get github.com/ant0ine/go-json-rest/rest
$ go get github.com/StephanDollberg/go-json-rest-middleware-jwt
```
```sh
#  Windows
d:\> set GOPATH=d:\GoCode\gUser
d:\> set GOBIN=d:\GoCode\gUser\bin
d:\> go get github.com/ant0ine/go-json-rest/rest
d:\> go get github.com/StephanDollberg/go-json-rest-middleware-jwt
```
---
### Step #3:  Example JWT Authentication
File: **src/main/web.go**  
```go
package main
import (
    "log"
    "net/http"
    "time"
    "github.com/StephanDollberg/go-json-rest-middleware-jwt"
    "github.com/ant0ine/go-json-rest/rest"
)
func handle_auth(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}
func main() {
    jwt_middleware := &jwt.JWTMiddleware{
        Key:        []byte("secret key"),
        Realm:      "jwt auth",
        Timeout:    time.Hour,
        MaxRefresh: time.Hour * 24,
        Authenticator: func(userId string, password string) bool {
            return userId == "admin" && password == "admin"
        }}

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    // we use the IfMiddleware to remove certain paths from needing authentication
    api.Use(&rest.IfMiddleware{
        Condition: func(request *rest.Request) bool {
            return request.URL.Path != "/login"
        },
        IfTrue: jwt_middleware,
    })
    api_router, _ := rest.MakeRouter(
        rest.Post("/login", jwt_middleware.LoginHandler),
        rest.Get("/auth_test", handle_auth),
        rest.Get("/refresh_token", jwt_middleware.RefreshHandler),
    )
    api.SetApp(api_router)

    http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

    log.Fatal(http.ListenAndServe(":8080", nil))
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
### Step #5: curl demo:
```sh
curl -d '{"username": "admin", "password": "admin"}' -H "Content-Type:application/json" http://localhost:8080/api/login
curl -H "Authorization:Bearer TOKEN_RETURNED_FROM_ABOVE" http://localhost:8080/api/auth_test
curl -H "Authorization:Bearer TOKEN_RETURNED_FROM_ABOVE" http://localhost:8080/api/refresh_token
```

#### Example: 
```sh
Ganesh@myPC MINGW64 ~
$ curl curl -d '{"username": "admin", "password": "admin"}' -H "Content-Type:application/json" http://localhost:8080/api/login
curl: (6) Couldn't resolve host 'curl'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   211  100   169  100    42  11266   2800 --:--:-- --:--:-- --:--:--  165k{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzM5NzQzNjUsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTQ3Mzk3MDc2NX0.9ENjB-h5d4XtlrQJEPUw-KQFm2y5Figh0M391Nmo4xc"
}

Ganesh@myPC MINGW64 ~
$ curl -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzM5NzQzNjUsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTQ3Mzk3MDc2NX0.9ENjB-h5d4XtlrQJEPUw-KQFm2y5Figh0M391Nmo4xc" http://localhost:8080/api/auth_test
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    23  100    23    0     0   1437      0 --:--:-- --:--:-- --:--:-- 23000{
  "authed": "admin"
}

Ganesh@myPC MINGW64 ~
$ curl -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzM5NzQzNjUsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTQ3Mzk3MDc2NX0.9ENjB-h5d4XtlrQJEPUw-KQFm2y5Figh0M391Nmo4xc" http://localhost:8080/api/refresh_token
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   169  100   169    0     0    169      0  0:00:01 --:--:--  0:00:01  165k{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzM5NzQ0MzMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTQ3Mzk3MDc2NX0.JTULyYLzTZZ0n9k7OD2ZkUnPMsW-EIRiAHwPdhaUP3E"
}
```
---