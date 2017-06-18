package api

import(
  "github.com/gocql/gocql"

  "github.com/hubertwwong/golangCqlTest/db/tables"
  "github.com/hubertwwong/golangCqlTest/db/cassandra"

  "encoding/json"
  "fmt"
  "net/http"
)

// init

func Heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(tables.HeartbeatResponse{Status: "OK", Code: 200})
}

func Monkey(w http.ResponseWriter, r *http.Request) {
  // cassandra. Don't like this session.
  CassandraSession := Cassandra.Session
  defer CassandraSession.Close()

  // generic inetface
  var rowList []tables.MonkeyResponse
  m := map[string]interface{}{}

  // query
  query := "SELECT * FROM monkey"
  iterable := Cassandra.Session.Query(query).Iter()

  for iterable.MapScan(m) {
    curRow := tables.MonkeyResponse{
      Identifier: m["identifier"].(gocql.UUID),
      Species: m["species"].(string),
      Nickname: m["nickname"].(string),
      Population: m["population"].(int),
    }
    rowList = append(rowList, curRow)
    m = map[string]interface{}{}
  }

  json.NewEncoder(w).Encode(rowList)
  fmt.Println(rowList)

  //json.NewEncoder(w).Encode(tables.MonkeyResponse{Species: "aa", Nickname: "aaa", Population: 1000})
}