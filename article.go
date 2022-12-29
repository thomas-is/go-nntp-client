package nntp

import (
  "strings"
  "strconv"
)


type Article struct {
  Number    int
  Header    map[string]string
  Body      []string
  Client    Client            `json:"-"`
}


func (n Client) Article(id string) Article {

  var article Article
  header := make(map[string]string)
  var body []string

  article.Number = 0
  article.Header = header
  article.Body   = body
  article.Client = n.Command("ARTICLE "+id)

/* 220 n <a> article retrieved - head and body follow
           (n = article number, <a> = message-id)
   221 n <a> article retrieved - head follows
   222 n <a> article retrieved - body follows
   223 n <a> article retrieved - request text separately
   412 no newsgroup has been selected
   420 no current article has been selected
   423 no such article number in this group
   430 no such article found                              */

  if article.Client.Status.Code < 220 || article.Client.Status.Code > 223 {
    return article
  }

  info := strings.Split(article.Client.Status.Message, " ")
  article.Number, _ = strconv.Atoi(info[0])
//  message.MessageId = info[1]

  var field string
  emptyLine := 0

  for i := 0; i < len(article.Client.Answer); i++ {
    if article.Client.Answer[i] == "" {
      /* empty line, assuming end of header */
      emptyLine = i
      break
    }

    atom := strings.Split(article.Client.Answer[i], ": ")
    if len(atom) == 1 {
      /* append single atom to last known field */
      header[field] = header[field] + atom[0]
      continue
    }

    field = atom[0]
    header[field] = strings.Join(atom[1:], ": ")

  }

  article.Header = header

//  var value string
//  var found bool
//
//  if value, found = header["References"]; found {
//    article.References = strings.Split(value, "\t")
//  }
//
//  if value, found = header["Path"]; found {
//    article.Path = strings.Split(value, "!")
//  }

  if emptyLine+1 < len(article.Client.Answer)-1 {
    article.Body = article.Client.Answer[emptyLine+1:len(article.Client.Answer)-1]
  }

  return article

}
