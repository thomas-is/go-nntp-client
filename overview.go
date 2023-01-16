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


func (n *Client) Xover( arg string ) []map[string]string {

  n.OverviewFmt().Command("XOVER "+arg)

  var value []string
  overview := make([]map[string]string,len(n.Answer)-1)

  for i:=0;i<len(n.Answer)-1;i++ {
    overview[i] = make(map[string]string)
    value = strings.Split(n.Answer[i], "\t")
    for j:=0;j<len(value);j++ {
      overview[i][n.OverviewFormat[j]] = value[j]
    }
  }

  return overview

}


