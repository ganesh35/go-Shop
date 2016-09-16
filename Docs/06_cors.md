# #06: CORS
Cross-origin resource sharing (CORS) is a mechanism that allows restricted resources (e.g. fonts) on a web page to be requested from another domain outside the domain from which the resource originated.
### Dependencies:
+ [Go-Json-Rest v3.3.1](https://github.com/ant0ine/go-json-rest)
+ [JWT Middleware for Go-Json-Rest](https://github.com/StephanDollberg/go-json-rest-middleware-jwt)
+ [go-semver - Semantic Versioning Library](https://github.com/ant0ine/go-json-rest#api-versioning)

**Referance :** https://github.com/ant0ine/go-json-rest#cors


---
### File and Folder Structure
-- gUser
+ src     :  # Source code goes here
    - main  :  # main Package
        - web.go  :  # main entry file
        - countries.go : # to get the countries list
        - helpers.go :  # helper functions goes here such as in_array ...
        - SemVerMiddleware.go :  # SemVer Middleware implementation 
+ bin     :  # Generated binary files and configuration settings goes here
+ pkg
+ Docs       :  # Documentation goes here
  - 01_hello.md
  - 02_basicServer.md
  - 03_gorillaroute.md
  - 04_jwt.md
  - 05_versioning.md
  - 06_cors.md
+ .gitignore
+ LICENCE
+ README.md
---
### Step #1:  Setting paths and installing necessary dependencies
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
$ go get github.com/coreos/go-semver/semver
```
```sh
#  Windows
d:\> set GOPATH=d:\GoCode\gUser
d:\> set GOBIN=d:\GoCode\gUser\bin
d:\> go get github.com/ant0ine/go-json-rest/rest
d:\> go get github.com/StephanDollberg/go-json-rest-middleware-jwt
d:\> go get github.com/coreos/go-semver/semver
```
---
### Step #2:  Make the program
Please check the file **src/main/web.go**
```go
    ...
    api.Use(rest.DefaultDevStack...)

    // CORS 
    api.Use(&rest.CorsMiddleware{
        RejectNonCorsRequests: false,
        OriginValidator: func(origin string, request *rest.Request) bool {
            return origin == "http://my.other.host"
        },
        AllowedMethods: []string{"GET", "POST", "PUT"},
        AllowedHeaders: []string{
            "Accept", "Content-Type", "X-Custom-Header", "Origin"},
        AccessControlAllowCredentials: true,
        AccessControlMaxAge:           3600,
    })
    // CORS /

    // we use the IfMiddleware to remove certain paths from needing authentication
    api.Use(&rest.IfMiddleware{
    ...
```
  
---
### Step #3:  Running the program
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
### Step #4: curl demo:
```sh
curl -i http://127.0.0.1:8080/countries
```

#### Example: 
```sh
Ganesh@myPC MINGW64 ~
$ curl -i http://127.0.0.1:8080/api/1.0.0/countries
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Fri, 16 Sep 2016 06:49:23 GMT
Content-Length: 105
[
  {
    "Code": "FR",
    "Name": "France"
  },
  {
    "Code": "US",
    "Name": "United States"
  }
]
```
---