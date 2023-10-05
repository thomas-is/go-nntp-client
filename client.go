package nntp

import (
  "net"
  "fmt"
  "os"
  "strings"
  "strconv"
)

const (
  BUFFER_SIZE = 100000
  EOL         = "\r\n"
)

type Status struct {
  code    int
  message string
}

type Client struct {
  socket          net.Conn  `json:"-"`
  status          Status
  overviewFormat  []string
  answer          []string
}

/* client init */
func Dial( host string, port int) *Client {
  var client  Client
  var err     error
  client.socket, err = net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
  if err != nil {
    fmt.Println("Dial failed:", err.Error())
    os.Exit(1)
  }
  client.Read()
  return &client
}

/* LIST OVERVIEW.FMT */
func (n *Client) OverviewFmt() *Client {
  n.Command("LIST OVERVIEW.FMT")
  if n.status.code != 215 {
    return n
  }
  n.overviewFormat = make([]string, len(n.answer))
  n.overviewFormat[0] = "Number"
  for i:=1;i<len(n.answer);i++ {
    n.overviewFormat[i] = strings.Split(n.answer[i-1], ":")[0]
  }
  return n
}

/* read response */
func (n *Client) Read() *Client {
  buffer := make([]byte, BUFFER_SIZE)
  _, err := n.socket.Read(buffer)
  if err != nil {
    println("Read data failed: ", err.Error())
    os.Exit(1)
  }
  lines := strings.Split(string(buffer), EOL)
  info  := strings.Split(lines[0], " ")
  n.status.code, _ = strconv.Atoi(info[0])
  n.status.message = strings.Join(info[1:]," ")
  /* remove first line (status)
     and last "line" (zero filled) */
  n.answer = lines[1:len(lines)-1]
  return n
}

/* send command */
func (n *Client) Command(message string) *Client {
  _, err := n.socket.Write([]byte(message+"\n"))
  if err != nil {
	  println("Write data failed:", err.Error())
	  os.Exit(1)
  }
  return n.Read()
}

/* QUIT */
func (n *Client) Quit() *Client {
  n.Command("QUIT").socket.Close()
  return n
}

