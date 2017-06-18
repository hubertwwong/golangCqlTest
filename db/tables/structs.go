package tables

import "github.com/gocql/gocql"

type HeartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}

type MonkeyResponse struct {
  Identifier gocql.UUID `json:identifier`
  Species string `json:"species"`
  Nickname string `json:"nickname"`
  Population int `json:"population"`
}