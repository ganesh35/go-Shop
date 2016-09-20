package main

import (
    "log"
    "net/http"
    "time"
    "github.com/StephanDollberg/go-json-rest-middleware-jwt"
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/coreos/go-semver/semver"
    "strings"
    "main/lib"
    GUsers "main/components/users"
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

	lib.SendEmail(gConfig.SmtpSettings, gConfig.MailSettings, "ganesh.35@gmail.com", "Test mail fro GO", "this is a sample body message")
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
    users := GUsers.Users{
        Store: map[string]* GUsers.User{},
    }


    svmw := SemVerMiddleware{
        MinVersion: "1.0.0",
        MaxVersion: "3.0.0",
    }

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

    // CORS 
 	api.Use(&rest.CorsMiddleware{
        RejectNonCorsRequests: false,
        OriginValidator: func(origin string, request *rest.Request) bool {
            return origin == "*"
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
        Condition: func(request *rest.Request) bool {

        	publicRoutes := []string{ "login", "message", "countries", "users", "user" }
        	
			urlParts := strings.Split(request.URL.Path, "/")
        	result, _ :=  lib.In_array_strings(urlParts[2], publicRoutes) ;
        	if result == true {
        		return false;
        	}

        	return true;

            //return request.URL.Path != "/login"
        },
        IfTrue: jwt_middleware,
    })
    api_router, _ := rest.MakeRouter(
        rest.Post("/login", jwt_middleware.LoginHandler),
        rest.Get("/auth_test", handle_auth),
        rest.Get("/refresh_token", jwt_middleware.RefreshHandler),

         rest.Get("/#version/message", svmw.MiddlewareFunc(
            func(w rest.ResponseWriter, req *rest.Request) {
                version := req.Env["VERSION"].(*semver.Version)
                if version.Major == 2 {
                    // https://en.wikipedia.org/wiki/Second-system_effect
                    w.WriteJson(map[string]string{
                        "Body": "Hello broken World!",
                    })
                } else {
                    w.WriteJson(map[string]string{
                        "Body": "Hello World!",
                    })
                }
            },
        )),

        // CORS
        rest.Get("/#version/countries", svmw.MiddlewareFunc( 
            func(w rest.ResponseWriter, req *rest.Request) {
                version := req.Env["VERSION"].(*semver.Version)
                if version.Major == 2 {
                    w.WriteJson(map[string]string{
                        "Body": "Hello broken World!",
                    })
                } else {
                	GetAllCountries(w, req);
                }
            },
        )),
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

    http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

    log.Fatal(http.ListenAndServe(":8080", nil))
}




