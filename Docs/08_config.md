# #08: Configuration File
In this example we will create config.json file and retrieve its values using Go structs and JSON Decoder.
### Dependencies:
+ [Go-Json-Rest v3.3.1](https://github.com/ant0ine/go-json-rest)
+ [JWT Middleware for Go-Json-Rest](https://github.com/StephanDollberg/go-json-rest-middleware-jwt)
+ [go-semver - Semantic Versioning Library](https://github.com/ant0ine/go-json-rest#api-versioning)

**Referance :** https://golang.org/pkg/encoding/json/

---
### Step #1: File and Folder Structure
-- gUser
+ src     :  # Source code goes here
    - main  :  # main Package
        - web.go  :  # main entry file
        - countries.go : # to get the countries list
        - helpers.go :  # helper functions goes here such as in_array ...
        - SemVerMiddleware.go :  # SemVer Middleware implementation 
        - users.go : #  User actions
    - **lib** : # necessary library files goes here
        - **config.go** : #  json parser and Config Structure
+ bin     :  # Generated binary files and configuration settings goes here
    - **etc**
        - **config.json**  : # JSON configuration file
+ pkg
+ Docs       :  # Documentation goes here
  - 01_hello.md
  - 02_basicServer.md
  - 03_gorillaroute.md
  - 04_jwt.md
  - 05_versioning.md
  - 06_cors.md
  - 07_users.md
  - 08_config.md
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
Please check the file **src/lib/gconfig.go**
```go
    ...
      "strings"
    "lib"
)
var gConfig lib.GConfig
func init(){
	log.Println("Loading etc/config.json")
	gConfig.LoadFile("etc/config.json")
	log.Println("Loading completed")
	log.Println("HttpSettings from config.json file " + gConfig.HttpSettings.Domain + " : " +  gConfig.HttpSettings.Port )
}
func handle_auth(w rest.ResponseWriter, r *rest.Request) {
    ...
```
Please check the file **bin/etc/gconfig.json**
```json
{
	"DbSettings":{
		"Domain" : "localhost",
		"Port": "27017",
		"Username": "dbUserName",
		"Password": "",
		"Database": "gs01"
	},
	"HttpSettings":{
		"Domain": "localhost",
		"Port": "8080"
	},  
	"LogSettings":{
		"LogFolder": "logs/",
    	"LogFile": "log",
    	"LogFormat": "json"
	},
	"SystemSettings":{
		"DefaultRole": "ROLE_USER_L1",
		"DefaultManager": ""
	},
	"SmtpSettings":{
      "Smtp_enabled": true,
      "Host":"smtp.mydomain.com",
      "Username":"mysmtp@mydomain.com",
      "Password":"",
      "Port":587,
      "Secure":""
  	},
	"MailSettings":{
		"Alert_email":"alert-me@mydomain.com",
	   	"Sender_name": "Support Team",
	   	"Sender_email": "Me@mydomain.org"
	}
}

```

Please check the file **src/lib/gconfig.go**

---
### Step #3:  Running the program
- Open command prompt / terminal (in windows : Start -> Run -> cmd )
```sh
#  Linux / Unix / gitbash
$ cd GoCode/gUser/bin
$ go install main
$ ./main.exe
2016/09/19 02:36:31 Loading etc/config.json
2016/09/19 02:36:32 Loading completed
2016/09/19 02:36:32 HttpSettings from config.json file localhost : 8080
```
```sh
#  Windows
d:\>cd GoCode\gUser\bin
D:\GoCode\gUser\bin>go install main
D:\GoCode\gUser\bin>main.exe
2016/09/19 02:37:59 Loading etc/config.json
2016/09/19 02:37:59 Loading completed
2016/09/19 02:37:59 HttpSettings from config.json file localhost : 8080
```
---

