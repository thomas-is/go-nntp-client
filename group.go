package nntp

import (
  "strings"
  "strconv"
)

type Group struct {
  Name    string
  Total   int
  First   int
  Last    int
  Client  Client
}

func (n Client) Group(name string) Group {

  var g Group
  g.Client = n.Command("GROUP "+name)

  if g.Client.Status.Code != 211 {
    return g
  }

  info := strings.Split(g.Client.Status.Message, " ")

  g.Total, _ = strconv.Atoi(info[0])
  g.First, _ = strconv.Atoi(info[1])
  g.Last,  _ = strconv.Atoi(info[2])
  g.Name     = info[3]

  return g

}

