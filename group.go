package nntp

import (
  "strings"
  "strconv"
)

type Group struct {
  Status  Status
  Name    string
  Total   int
  First   int
  Last    int
}

func (n Client) Group(name string) Group {

  n = n.Command("GROUP "+name)

  var g Group

  g.Status = n.Status
  if g.Status.Code != 211 {
    return g
  }

  info := strings.Split(n.Status.Message, " ")

  g.Total, _ = strconv.Atoi(info[0])
  g.First, _ = strconv.Atoi(info[1])
  g.Last,  _ = strconv.Atoi(info[2])
  g.Name     = info[3]

  return g

}

