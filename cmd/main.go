package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrthoabby/content-management-service-ck/pkg/core"
	"github.com/mrthoabby/content-management-service-ck/pkg/middlewares"
	"github.com/sirupsen/logrus"
)

func main() {
	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := mux.NewRouter()

	router.Use(middlewares.GlobalRecoveryPanic)

	core.RunIoc(router)
	defer core.IOCCleanUp(context)

	logrus.Debug("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
