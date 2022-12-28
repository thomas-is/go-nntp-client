package nntp

import (
  "net"
  "fmt"
  "os"
  "strings"
  "strconv"
)

type Status struct {
  Code    int
  Message string
}

type Client struct {
  Socket  net.Conn
  Status  Status
  Answer  []string
}

const NNTP_BUFFER_SIZE = 100000
const NNTP_EOL         = "\r\n"


func Dial( host string, port int) Client {

  target := fmt.Sprintf("%s:%d", host, port)

  var n Client
  var err error

  n.Socket, err = net.Dial("tcp", target)
  if err != nil {
    fmt.Println("Dial failed:", err.Error())
    os.Exit(1)
  }

  return n.Read()

}


func (n Client) Read() Client {

  buffer := make([]byte, NNTP_BUFFER_SIZE)

  _, err := n.Socket.Read(buffer)
  if err != nil {
    println("Read data failed: ", err.Error())
    os.Exit(1)
  }

  lines := strings.Split(string(buffer), NNTP_EOL)
  info := strings.Split(lines[0], " ")

  var status Status
  status.Code, _ = strconv.Atoi(info[0])
  status.Message = strings.Join(info[1:]," ")

  n.Status = status
  // remove first line (status) and last "line" (zero filled)
  n.Answer = lines[1:len(lines)-1]

  return n

}


func (n Client) Command(message string) Client {

  message = message +"\n"

  _, err := n.Socket.Write([]byte(message))
  if err != nil {
	  println("Write data failed:", err.Error())
	  os.Exit(1)
  }

  return n.Read()

}


func (n Client) Quit() Client {

  n.Command("QUIT").Socket.Close()

  return n

}
