# #12: Pages
Managing Pages / Articles, Data will be saved in the MongoDB and retrieved when needed.
### ToDo:
- Attachment support
- HTML and Plaintext body support

### Dependencies:
+ [Go-Json-Rest v3.3.1](https://github.com/ant0ine/go-json-rest)
+ [JWT Middleware for Go-Json-Rest](https://github.com/StephanDollberg/go-json-rest-middleware-jwt)
+ [go-semver - Semantic Versioning Library](https://github.com/ant0ine/go-json-rest#api-versioning)
+ [MGO - Rich MongoDB driver for Go](http://labix.org/mgo)
  go get gopkg.in/mgo.v2
### References:
- : 

---
### Step #1: File and Folder Structure
-- go-Shop
+ src     :  # Source code goes here
    - main  :  # main Package
        - web.go  :  # main entry file
        - countries.go : # to get the countries list
        - SemVerMiddleware.go :  # SemVer Middleware implementation 
        - components
            - pages
                - controller.go : #  Page actions
                - modal.go : #  Logic and database transactions
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
### Step #1:  Setting paths and installing necessary dependencies
- Open command prompt / terminal (in windows : Start -> Run -> cmd )
```sh
$ cd /GoCode/go-Shop
```
```sh
#  Linux / Unix / gitbash
$ export GOPATH=/GoCode/go-Shop
$ export GOBIN=/GoCode/go-Shop/bin
$ go get github.com/ant0ine/go-json-rest/rest
$ go get github.com/StephanDollberg/go-json-rest-middleware-jwt
$ go get github.com/coreos/go-semver/semver
$ go get gopkg.in/mgo.v2
```
```sh
#  Windows
d:\> set GOPATH=d:\go-Shop\gUser
d:\> set GOBIN=d:\go-Shop\gUser\bin
d:\> go get github.com/ant0ine/go-json-rest/rest
d:\> go get github.com/StephanDollberg/go-json-rest-middleware-jwt
d:\> go get github.com/coreos/go-semver/semver
d:\> go get gopkg.in/mgo.v2
```
---

### Step #2: Code changes
Change domain and port based config.json  
Changes in **/src/main/web.go**
```go
// **Replace** 
// log.Fatal(http.ListenAndServe(":8080", nil))
// **with**
domainAndPort := gConfig.HttpSettings.Domain + ":" + gConfig.HttpSettings.Port
    log.Fatal(http.ListenAndServe(domainAndPort, nil))
```
New file **/src/lib/mongodb.go**
```go
package lib
import (
	"gopkg.in/mgo.v2"
	"net/http"
	)
var (
    clt *http.Client
	Fullsession *mgo.Session
	Readsession *mgo.Session
	FullDatabase *mgo.Database
	ReadDatabase *mgo.Database
)
func ConnectDB(conf *GConfig) {
	var err error
	domainAndPort := conf.DbSettings.Domain + ":" + conf.DbSettings.Port
	Fullsession, err = mgo.Dial(domainAndPort)
	check(err)
	Fullsession.SetSafe(&mgo.Safe{})
	Fullsession.SetMode(mgo.Monotonic, true)
	FullDatabase = Fullsession.DB(conf.DbSettings.Database)
	err = FullDatabase.Login(conf.DbSettings.Username, conf.DbSettings.Password)
	check(err)
}

func CloseDB() {
	FullDatabase.Logout()
	Fullsession.Close()
}
```
Initiate MongoDB  
Changes in **/src/main/web.go**
```go
import (
    "main/lib"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
var gFullDatabase *mgo.Database
func init(){
    lib.ConnectDB(&gConfig)
    gFullDatabase = lib.FullDatabase
}
```

Close MongoDB  
Changes in **/src/main/web.go**
```go
func close(){
    gsLog.Info("Application ended ----- ");
    gsLog.Close(gsConfig.LogSettings.LogFolder, gsConfig.LogSettings.LogFile, gsConfig.LogSettings.LogFormat)
    lib.CloseDB()
}
```
Test MongoDB connection  
```go
    var result []interface{}
    var err error
    err = gFullDatabase.C("pages").Find(nil).Select(bson.M{"_id": 1,"Title": 1,"Alias": 1,"Lang": 1 }) .All(&result) 
    log.Println(result)
    log.Println(err)
```

---
### Step #3: Testing
```sh
$ ./main.exe
2016/09/21 00:42:55 Info|| ---------------------------------------
2016/09/21 00:42:55 Info|| Application started
2016/09/21 00:42:55 Info|| Loading etc/config.json file
2016/09/21 00:42:55 Info|| Loading config.json completed
2016/09/21 00:42:55 Warning|| HttpSettings from config.json file localhost : 8080
2016/09/21 00:42:55 Error|| Testing Error log entry
2016/09/21 00:42:55 Critical|| Testing Critical log entry
2016/09/21 00:42:56 [map[Alias:my-page-title-3 Lang:en-GB _id:ObjectIdHex("574ecba1b607ccf0d194a955") Title:My Page Title #3] map[Lang:en-GB _id:ObjectIdHex("574f5b8cb607ccf0d194a96f") Title:My Page Title #3 Alias:ykcb4MJqXW]]
2016/09/21 00:42:56 <nil>
```





