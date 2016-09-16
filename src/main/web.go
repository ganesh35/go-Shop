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
        Condition: func(request *rest.Request) bool {

        	publicRoutes := []string{ "login", "message", "countries" }
        	
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

    )
    api.SetApp(api_router)

    http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

    log.Fatal(http.ListenAndServe(":8080", nil))
}




