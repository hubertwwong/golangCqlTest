package api

import (
  "github.com/gorilla/mux"

  //"github.com/hubertwwong/golangCqlTest/db/cassandra"
  "github.com/hubertwwong/golangCqlTest/config"

  "fmt"
  "log"
  "net/http"
  "strconv"
)

// Create a router object, define all of the routes, and returns a mux.Router object.
func createRoutes() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  
  router.HandleFunc("/", Heartbeat)
  router.HandleFunc("/monkey", Monkey)

  return router
}

func Run() {
  // config
  cfg := config.GetConfig()
  fmt.Println(cfg)
  
  // router
  router := createRoutes()

  // run the server
  fmt.Println(cfg.Port)
  portSuffix := ":" + strconv.Itoa(cfg.Port)
  log.Fatal(http.ListenAndServe(portSuffix, router))
}