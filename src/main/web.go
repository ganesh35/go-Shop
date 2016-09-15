package main

import (
    "log"
    "net/http"
    "time"

    "github.com/StephanDollberg/go-json-rest-middleware-jwt"
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/coreos/go-semver/semver"
    "strings"
)

func handle_auth(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}

func main() {
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
    // we use the IfMiddleware to remove certain paths from needing authentication
    api.Use(&rest.IfMiddleware{
        Condition: func(request *rest.Request) bool {

        	publicRoutes := []string{ "login", "message" }
        	
			urlParts := strings.Split(request.URL.Path, "/")
        	result, _ :=  in_array_strings(urlParts[2], publicRoutes) ;
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

    )
    api.SetApp(api_router)

    http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

    log.Fatal(http.ListenAndServe(":8080", nil))
}




type SemVerMiddleware struct {
    MinVersion string
    MaxVersion string
}

func (mw *SemVerMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

    minVersion, err := semver.NewVersion(mw.MinVersion)
    if err != nil {
        panic(err)
    }

    maxVersion, err := semver.NewVersion(mw.MaxVersion)
    if err != nil {
        panic(err)
    }

    return func(writer rest.ResponseWriter, request *rest.Request) {

        version, err := semver.NewVersion(request.PathParam("version"))
        if err != nil {
            rest.Error(
                writer,
                "Invalid version: "+err.Error(),
                http.StatusBadRequest,
            )
            return
        }

        if version.LessThan(*minVersion) {
            rest.Error(
                writer,
                "Min supported version is "+minVersion.String(),
                http.StatusBadRequest,
            )
            return
        }

        if maxVersion.LessThan(*version) {
            rest.Error(
                writer,
                "Max supported version is "+maxVersion.String(),
                http.StatusBadRequest,
            )
            return
        }

        request.Env["VERSION"] = version
        handler(writer, request)
    }
}

func in_array_strings(val string, array []string) (ok bool, i int) {  // Only for string array elements
    for i = range array {
        if ok = array[i] == val; ok {
            return
        }
    }
    return
}
