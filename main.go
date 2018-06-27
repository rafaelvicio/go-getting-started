package main

import (
	"log"
	"net/http"
	"os"
	"pipenator-backend/api/auth"
	"pipenator-backend/api/pipelines"
	"pipenator-backend/common"
	"pipenator-backend/middleware"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func withDB(db *mgo.Session) middleware.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			dbsession := db.Copy()
			defer dbsession.Close()

			context.Set(r, "database", dbsession)

			h.ServeHTTP(w, r)
		})
	}
}

func main() {

	db, err := common.CreateDB()

	if err != nil {
		log.Fatal("Não foi possível conectar no MongoDB", err)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.SetHeadersDefault)

	routes := append(pipelines.Routes)
	routes = append(auth.Routes)

	for _, route := range routes {
		h := middleware.Adapt(http.HandlerFunc(route.Handler), withDB(db))

		if route.Security {
			h = middleware.Adapt(http.HandlerFunc(route.Handler), middleware.VerifyToken)
		}

		router.Handle(route.Path, context.ClearHandler(h)).Methods(route.Method)
	}

	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, router))

}
