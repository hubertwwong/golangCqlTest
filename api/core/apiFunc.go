package core

import(
  "github.com/hubertwwong/golangCqlTest/db/tables"

  "encoding/json"
  "net/http"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(tables.HeartbeatResponse{Status: "OK", Code: 200})
}