# #10: SMTP Mailer
Using GMail as smtp server to send emails from the Go program.
Be sure to enter Gmail username and password in the config.json file
### ToDo:
- Attachment support
- HTML and Plaintext body support
### Dependencies:
+ [Go-Json-Rest v3.3.1](https://github.com/ant0ine/go-json-rest)
+ [JWT Middleware for Go-Json-Rest](https://github.com/StephanDollberg/go-json-rest-middleware-jwt)
+ [go-semver - Semantic Versioning Library](https://github.com/ant0ine/go-json-rest#api-versioning)
### References:
- SMTP mail: https://golang.org/pkg/net/smtp/
- Email: https://golang.org/pkg/net/mail/

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
        - glog.go : #  json parser and Config Structure
        - ghelpers.go : #  Necessary helper functions
        - **gmail.go** : #  SendMail using GMail smtp server
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
    gLog.Critical("Testing Critical log entry ")
    lib.SendEmail(
        gConfig.SmtpSettings, 
        gConfig.MailSettings, 
        "some_email@gmail.com", 
        "Test mail fro GO", 
        "this is a sample body message" )
}
func close(){
    ...
```

Please check the file **src/lib/gmail.go**
```go
package lib
import (
    "fmt"
    "net/smtp"
    "net/mail"
)
func SendEmail(
    smtpSettings TypeSmtpSettings, 
    mailSettings TypeMailSettings, 
    toEmail string, 
    subject string, 
    body string) (err error) {
    defer CatchPanic(&err, "sendEmail")
    emailauth := smtp.PlainAuth("", smtpSettings.Username, smtpSettings.Password, smtpSettings.Host)
    sender :=  mailSettings.Sender_email// 
    receivers := []string{
        toEmail,
    }
    from := mail.Address{mailSettings.Sender_name, mailSettings.Sender_email}
    to   := mail.Address{"", toEmail}
    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subject
    headers["MIME-version"] = "1.0"
    headers["Content-Type"] = `text/html; charset="UTF-8"`
    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body
    // send out the email
    err = smtp.SendMail( fmt.Sprintf("%s:%d", smtpSettings.Host, smtpSettings.Port), //convert port number from int to string
        emailauth,
        sender,
        receivers,
        []byte(message),
    )
    return err 
}
```

---
### Step #3:  Running the program
- Open command prompt / terminal (in windows : Start -> Run -> cmd )
```sh
#  Linux / Unix / gitbash
$ cd GoCode/gUser/bin
$ go install main
$ ./main.exe
```
```sh
#  Windows
d:\>cd GoCode\gUser\bin
D:\GoCode\gUser\bin>go install main
D:\GoCode\gUser\bin>main.exe
```
---

