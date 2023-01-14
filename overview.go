package nntp

import (
//  "net"
//  "fmt"
//  "os"
  "strings"
//  "strconv"
)

type Overview struct {
  Data []string
}


func (n *Client) Xover( arg string ) []Overview {

  n.Command("XOVER "+arg)

  overview := make([]Overview,len(n.Answer))

  for i := 0; i < len(n.Answer); i++ {
    overview[i].Data = strings.Split(n.Answer[i], "\t")
  }
  return overview
}


