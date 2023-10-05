package nntp

import (
  "strings"
)

type Overview struct {
  data []string
}


func (n *Client) Xover( arg string ) []map[string]string {

  n.OverviewFmt().Command("XOVER "+arg)

  var value []string
  overview := make([]map[string]string,len(n.answer)-1)

  for i:=0;i<len(n.answer)-1;i++ {
    overview[i] = make(map[string]string)
    value = strings.Split(n.answer[i], "\t")
    for j:=0;j<len(value);j++ {
      overview[i][n.overviewFormat[j]] = value[j]
    }
  }

  return overview

}
