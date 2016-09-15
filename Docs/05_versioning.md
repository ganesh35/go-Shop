# #05: API Versioning
Authentication via Json Web Tokens included with API Versioning.  
It defines a middleware that parses the version, checks a min and a max, and makes it available in the request.Env.
### Dependencies:
+ [Go-Json-Rest v3.3.1](https://github.com/ant0ine/go-json-rest)
+ [JWT Middleware for Go-Json-Rest](https://github.com/StephanDollberg/go-json-rest-middleware-jwt)
+ [go-semver - Semantic Versioning Library](https://github.com/ant0ine/go-json-rest#api-versioning)

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
curl -i http://127.0.0.1:8080/api/1.0.0/message
curl -i http://127.0.0.1:8080/api/2.0.0/message
curl -i http://127.0.0.1:8080/api/2.0.1/message
curl -i http://127.0.0.1:8080/api/0.0.1/message
curl -i http://127.0.0.1:8080/api/4.0.1/message
```

#### Example: 
```sh
Ganesh@myPC MINGW64 ~
$ curl -i http://127.0.0.1:8080/api/1.0.0/message
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    28  100    28    0     0     28      0  0:00:01 --:--:--  0:00:01 28000HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Thu, 15 Sep 2016 21:12:17 GMT
Content-Length: 28

{
  "Body": "Hello World!"
}

Ganesh@myPC MINGW64 ~
$ curl -i http://127.0.0.1:8080/api/2.0.0/message
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    35  100    35    0     0     35      0  0:00:01 --:--:--  0:00:01 35000HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Thu, 15 Sep 2016 21:12:18 GMT
Content-Length: 35

{
  "Body": "Hello broken World!"
}

Ganesh@myPC MINGW64 ~
$ curl -i http://127.0.0.1:8080/api/2.0.1/message
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    35  100    35    0     0     35      0  0:00:01 --:--:--  0:00:01 35000HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Thu, 15 Sep 2016 21:12:18 GMT
Content-Length: 35

{
  "Body": "Hello broken World!"
}

Ganesh@myPC MINGW64 ~
$ curl -i http://127.0.0.1:8080/api/0.0.1/message
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    47  100    47    0     0     47      0  0:00:01 --:--:--  0:00:01 47000HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Thu, 15 Sep 2016 21:12:18 GMT
Content-Length: 47

{
  "Error": "Min supported version is 1.0.0"
}

Ganesh@myPC MINGW64 ~
$ curl -i http://127.0.0.1:8080/api/4.0.1/message
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    47  100    47    0     0     47      0  0:00:01 --:--:--  0:00:01 47000HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Thu, 15 Sep 2016 21:12:23 GMT
Content-Length: 47

{
  "Error": "Max supported version is 3.0.0"
}
```
---