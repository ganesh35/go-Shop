package main

import (
    "github.com/coreos/go-semver/semver"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    )

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