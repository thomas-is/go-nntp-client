package nntp

import (
  "strings"
  "strconv"
)


type Article struct {
  Status     Status
  Number     int
  MessageId  string
  Path       []string
  References []string
  Header     map[string]string
  Body       []string
}


func (n Client) Article(id string) Article {

  var article Article

  n = n.Command("ARTICLE "+id)

  article.Status = n.Status

/* 220 n <a> article retrieved - head and body follow
           (n = article number, <a> = message-id)
   221 n <a> article retrieved - head follows
   222 n <a> article retrieved - body follows
   223 n <a> article retrieved - request text separately
   412 no newsgroup has been selected
   420 no current article has been selected
   423 no such article number in this group
   430 no such article found                              */

  if n.Status.Code < 220 || n.Status.Code > 223 {
    return article
  }

  info := strings.Split(n.Status.Message, " ")
  article.Number, _ = strconv.Atoi(info[0])
  article.MessageId = info[1]

  var field string
  emptyLine := 0
  header := make(map[string]string)

  for i := 0; i < len(n.Answer); i++ {
    if n.Answer[i] == "" {
      /* empty line, assuming end of header */
      emptyLine = i
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

  article.Header = header

  var value string
  var found bool

  if value, found = header["References"]; found {
    article.References = strings.Split(value, "\t")
  }

  if value, found = header["Path"]; found {
    article.Path = strings.Split(value, "!")
  }

  if emptyLine+1 < len(n.Answer)-1 {
    article.Body = n.Answer[emptyLine+1:len(n.Answer)-1]
  }

  return article

}
