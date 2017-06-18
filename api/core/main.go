package core

import (
  "github.com/gorilla/mux"

  "github.com/hubertwwong/golangCqlTest/db/cassandra"
//  "github.com/hubertwwong/golangCqlTest/db/tables"
  "github.com/hubertwwong/golangCqlTest/config"

  "fmt"
//  "encoding/json"
  "log"
  "net/http"
  "strconv"
)

// type heartbeatResponse struct {
//   Status string `json:"status"`
//   Code int `json:"code"`
// }

// func heartbeat(w http.ResponseWriter, r *http.Request) {
//   json.NewEncoder(w).Encode(tables.HeartbeatResponse{Status: "OK", Code: 200})
// }

func createRoutes() *mux.Router {
  // router
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Heartbeat)

  return router
}

func Run() {
  // config
  cfg := config.GetConfig()
  fmt.Println(cfg)

  // cassandra
  CassandraSession := Cassandra.Session
  defer CassandraSession.Close()
  //router := mux.NewRouter().StrictSlash(true)

  // router
  router := createRoutes()

  // run the server
  //log.Fatal(http.ListenAndServe(":8080", router))
  fmt.Println(cfg.Port)
  portSuffix := ":" + strconv.Itoa(cfg.Port)
  log.Fatal(http.ListenAndServe(portSuffix, router))
}