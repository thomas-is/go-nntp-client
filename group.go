package nntp

import (
  "strings"
  "strconv"
)

type Group struct {
  name    string
  total   int
  first   int
  last    int
}

func (n *Client) Group(name string) Group {

  var g Group
  n = n.Command("GROUP "+name)

  if n.status.code != 211 {
    return g
  }

  info := strings.Split(n.status.message, " ")

  g.total, _ = strconv.Atoi(info[0])
  g.first, _ = strconv.Atoi(info[1])
  g.last,  _ = strconv.Atoi(info[2])
  g.name     = info[3]

  return g

}

