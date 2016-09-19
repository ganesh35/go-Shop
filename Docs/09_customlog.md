# #09: Custom Log
Using in-memory log system.  All the log entries will be saved in the memory.
### ToDo:
- Save entries to disk on exit
- Save entries to disk on grace shutdown
- Save entries to disk at specified intermidiate times
- Save entries to database

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
        - users.go : #  User actions
    - lib : # necessary library files goes here
        - gconfig.go : #  json parser and Config Structure
        - **glog.go** : #  json parser and Config Structure
        - **ghelpers.go** : #  Necessary helper functions
+ bin     :  # Generated binary files and configuration settings goes here
    - etc
        - config.json  : # JSON configuration file
+ pkg
+ Docs       :  # Documentation goes here
  - 01_hello.md
  - ...
  - 09_log.md
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
    "lib"
)
var gConfig lib.GConfig
var gLog lib.GLog
func init(){
    gLog.Info("---------------------------------------");
    gLog.Info("Application started");   
    gLog.Info("Loading etc/config.json file ");
    gConfig.LoadFile("etc/config.json")
    gLog.Info("Loading config.json completed")
    gLog.Warning("HttpSettings from config.json file " + gConfig.HttpSettings.Domain + " : " +  gConfig.HttpSettings.Port )
    gLog.Error("Testing Error log entry ")
    gLog.Critical("Testing Critical log entry ")
}
func close(){
    gLog.Info("Closing Logger");
    gLog.Info("Application ended ----- ");
    gLog.Close(gConfig.LogSettings.LogFolder, gConfig.LogSettings.LogFile, gConfig.LogSettings.LogFormat)
}

func handle_auth(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}
func main() {
    defer close()
    users := Users{
    ...
```

Please check the file **src/lib/glog.go**

---
### Step #3:  Running the program
- Open command prompt / terminal (in windows : Start -> Run -> cmd )
```sh
#  Linux / Unix / gitbash
$ cd GoCode/gUser/bin
$ go install main
$ ./main.exe
2016/09/19 02:59:21 Info|| ---------------------------------------
2016/09/19 02:59:21 Info|| Application started
2016/09/19 02:59:21 Info|| Loading etc/config.json file
2016/09/19 02:59:21 Info|| Loading config.json completed
2016/09/19 02:59:21 Warning|| HttpSettings from config.json file localhost : 8080
2016/09/19 02:59:21 Error|| Testing Error log entry
2016/09/19 02:59:21 Critical|| Testing Critical log entry
```
```sh
#  Windows
d:\>cd GoCode\gUser\bin
D:\GoCode\gUser\bin>go install main
D:\GoCode\gUser\bin>main.exe
2016/09/19 03:08:10 Info|| ---------------------------------------
2016/09/19 03:08:10 Info|| Application started
2016/09/19 03:08:10 Info|| Loading etc/config.json file
2016/09/19 03:08:10 Info|| Loading config.json completed
2016/09/19 03:08:10 Warning|| HttpSettings from config.json file localhost : 8080
2016/09/19 03:08:10 Error|| Testing Error log entry
2016/09/19 03:08:10 Critical|| Testing Critical log entry
```
---

