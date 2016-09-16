# #07: Basic API
Demonstrate how to use Method Values.
Method Values have been introduced in Go 1.1.
This shows how to map a Route to a method of an instantiated object (i.e: receiver of the method)
### Dependencies:
+ [Go-Json-Rest v3.3.1](https://github.com/ant0ine/go-json-rest)
+ [JWT Middleware for Go-Json-Rest](https://github.com/StephanDollberg/go-json-rest-middleware-jwt)
+ [go-semver - Semantic Versioning Library](https://github.com/ant0ine/go-json-rest#api-versioning)

**Referance :** https://github.com/ant0ine/go-json-rest#users


---
### Step #1: File and Folder Structure
-- gUser
+ src     :  # Source code goes here
    - main  :  # main Package
        - web.go  :  # main entry file
        - countries.go : # to get the countries list
        - helpers.go :  # helper functions goes here such as in_array ...
        - SemVerMiddleware.go :  # SemVer Middleware implementation 
        - **users.go** : #  User actions
+ bin     :  # Generated binary files and configuration settings goes here
+ pkg
+ Docs       :  # Documentation goes here
  - 01_hello.md
  - 02_basicServer.md
  - 03_gorillaroute.md
  - 04_jwt.md
  - 05_versioning.md
  - 06_cors.md
  - 07_users.md
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
      // CORS /

        // Users
        rest.Get("/#version/users", svmw.MiddlewareFunc(
            func(w rest.ResponseWriter, req *rest.Request) {
                version := req.Env["VERSION"].(*semver.Version)
                if version.Major >= 2 {
                    // https://en.wikipedia.org/wiki/Second-system_effect
                    w.WriteJson(map[string]string{
                        "Body": "Not supported version!",
                    })
                } else {
                	users.GetAllUsers(w, req);
                }
            },
        )),
        rest.Post("/#version/users", svmw.MiddlewareFunc(
            func(w rest.ResponseWriter, req *rest.Request) {
                version := req.Env["VERSION"].(*semver.Version)
                if version.Major >= 2 {
                    // https://en.wikipedia.org/wiki/Second-system_effect
                    w.WriteJson(map[string]string{
                        "Body": "Not supported version!",
                    })
                } else {
                	users.PostUser(w, req);
                }
            },
        )),
        rest.Get("/#version/users/:id", svmw.MiddlewareFunc(
            func(w rest.ResponseWriter, req *rest.Request) {
                version := req.Env["VERSION"].(*semver.Version)
                if version.Major >= 2 {
                    // https://en.wikipedia.org/wiki/Second-system_effect
                    w.WriteJson(map[string]string{
                        "Body": "Not supported version!",
                    })
                } else {
                	users.GetUser(w, req);
                }
            },
        )),

        rest.Put("/#version/users/:id", svmw.MiddlewareFunc(
            func(w rest.ResponseWriter, req *rest.Request) {
                version := req.Env["VERSION"].(*semver.Version)
                if version.Major >= 2 {
                    // https://en.wikipedia.org/wiki/Second-system_effect
                    w.WriteJson(map[string]string{
                        "Body": "Not supported version!",
                    })
                } else {
                	users.PutUser(w, req);
                }
            },
        )),


        rest.Delete("/#version/users/:id", svmw.MiddlewareFunc(
            func(w rest.ResponseWriter, req *rest.Request) {
                version := req.Env["VERSION"].(*semver.Version)
                if version.Major >= 2 {
                    // https://en.wikipedia.org/wiki/Second-system_effect
                    w.WriteJson(map[string]string{
                        "Body": "Not supported version!",
                    })
                } else {
                	users.DeleteUser(w, req);
                }
            },
        )),

        // Users /

    )
    api.SetApp(api_router)    
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
curl -i -H 'Content-Type: application/json' \
    -d '{"Name":"Antoine"}' http://127.0.0.1:8080/users
curl -i http://127.0.0.1:8080/users/0
curl -i -X PUT -H 'Content-Type: application/json' \
    -d '{"Name":"Antoine Imbert"}' http://127.0.0.1:8080/users/0
curl -i -X DELETE http://127.0.0.1:8080/users/0
curl -i http://127.0.0.1:8080/users
```

#### Users POST Example : 
```sh
Ganesh@myPC MINGW64 ~
$ curl -i -H 'Content-Type: application/json' -d '{"Name": "Ganesh"}' \
   http://localhost:8080/api/1.0.0/users
{
  "Id": "0",
  "Name": "Ganesh"
}

Ganesh@myPC MINGW64 ~
$ curl -i -H 'Content-Type: application/json' -d '{"Name": "Antoine"}' \
  http://localhost:8080/api/1.0.0/users
{
  "Id": "1",
  "Name": "Antoine"
}
Ganesh@myPC MINGW64 ~
$ curl -i -H 'Content-Type: application/json' -d '{"Name": "Al vero"}' \ 
  http://localhost:8080/api/1.0.0/users
{
  "Id": "2",
  "Name": "Al vero"
}
Ganesh@myPC MINGW64 ~
$ curl -i -H 'Content-Type: application/json' -d '{"Name": "Stuva Grundlig"}' \
    http://localhost:8080/api/1.0.0/users
{
  "Id": "3",
  "Name": "Stuva Grundlig"
}
```
#### Users GET Example : 
```sh
Ganesh@myPC MINGW64 ~
$ curl -i http://127.0.0.1:8080/api/1.0.0/users
[
  {
    "Id": "0",
    "Name": "Ganesh"
  },
  {
    "Id": "1",
    "Name": "Antoine"
  },
  {
    "Id": "2",
    "Name": "Al vero"
  },
  {
    "Id": "3",
    "Name": "Stuva Grundlig"
  }
]
```
#### User GET Example with input field: id 
```sh
$ curl -i http://127.0.0.1:8080/api/1.0.0/users/1
{
  "Id": "1",
  "Name": "Antoine"
}
```
#### User PUT Example with input field: id 
```sh
$ curl -i -X PUT -H 'Content-Type: application/json' \
 -d '{"name": "Antoine Imbert"}' http://127.0.0.1:8080/api/1.0.0/users/1
{
  "Id": "1",
  "Name": "Antoine Imbert"
}
```
#### User DELETE Example with input field: id 
```sh
$ curl -i -X DELETE http://127.0.0.1:8080/api/1.0.0/users/1
```

---
