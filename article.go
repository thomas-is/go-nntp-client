package nntp

import (
  "strings"
  "strconv"
)


type Article struct {
  Number  int
  Header  map[string]string
  Body    []string
}

func (n *Client) Article(id string) Article {

  var article Article
  header := make(map[string]string)
  var body []string

  article.Number = 0
  article.Header = header
  article.Body   = body

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
  if n.Status.Code < 220 || n.Status.Code > 223 {
    return article
  }

  info := strings.Split(n.Status.Message, " ")
  article.Number, _ = strconv.Atoi(info[0])
//  message.MessageId = info[1]

  var field string
  var bodyLine int

  for i := 0; i < len(n.Answer); i++ {

    bodyLine = i

    if n.Answer[i] == "" {
      /* empty line, assuming end of header */
      break
    }

    atom := strings.Split(n.Answer[i], ": ")
    if len(atom) == 1 {
      /* append single atom to last known field */
      header[field] = header[field] + atom[0]
      continue
    }

    field = atom[0]
    header[field] = strings.Join(atom[1:], ": ")

  }

  bodyLine += 1

  article.Header = header

  if bodyLine < len(n.Answer)-1 {
    article.Body = n.Answer[bodyLine:len(n.Answer)-1]
  }

  return article

}

func (a Article) References() []string {
  value, found := a.Header["References"]
  if ! found {
    return make([]string, 0)
  }
  f := func(c rune) bool {
      return c == ' ' || c == '\t'
  }
  return strings.FieldsFunc(value, f)
}

func (a Article) Path() []string {
  value, found := a.Header["Path"]
  if ! found {
    return make([]string, 0)
  }
  return strings.Split(value, "!")
}

