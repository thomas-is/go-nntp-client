package nntp

import (
  "strings"
  "strconv"
)


type Article struct {
  number  int
  header  map[string]string
  body    []string
}

func (n *Client) Article( id string ) Article {

  var article Article
  header := make(map[string]string)
  var body []string

  article.number = 0
  article.header = header
  article.body   = body

/* 220 n <a> article retrieved - head and body follow
           (n = article number, <a> = message-id)
   221 n <a> article retrieved - head follows
   222 n <a> article retrieved - body follows
   223 n <a> article retrieved - request text separately
   412 no newsgroup has been selected
   420 no current article has been selected
   423 no such article number in this group
   430 no such article found                              */

  n = n.Command("ARTICLE "+id)
  if n.status.code < 220 || n.status.code > 223 {
    return article
  }

  info := strings.Split(n.status.message, " ")
  article.number, _ = strconv.Atoi(info[0])

  var field string
  var bodyLine int

  for i := 0; i < len(n.answer); i++ {

    bodyLine = i

    if n.answer[i] == "" {
      /* empty line, assuming end of header */
      break
    }

    atom := strings.Split(n.answer[i], ": ")
    if len(atom) == 1 {
      /* append single atom to last known field */
      header[field] = header[field] + atom[0]
      continue
    }

    field = atom[0]
    header[field] = strings.Join(atom[1:], ": ")

  }

  bodyLine += 1

  article.header = header

  if bodyLine < len(n.answer)-1 {
    article.body = n.answer[bodyLine:len(n.answer)-1]
  }

  return article

}

func (a Article) References() []string {
  value, found := a.header["References"]
  if ! found {
    return make([]string, 0)
  }
  f := func(c rune) bool {
      return c == ' ' || c == '\t'
  }
  return strings.FieldsFunc(value, f)
}

func (a Article) Path() []string {
  value, found := a.header["Path"]
  if ! found {
    return make([]string, 0)
  }
  return strings.Split(value, "!")
}

