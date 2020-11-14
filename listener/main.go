package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        "io/ioutil"
        "io"
)

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide port number")
                return
        }

        data, err := ioutil.ReadFile("banner.txt")
        fmt.Println(string(data))

        PORT := ":" + arguments[1]
        l, err := net.Listen("tcp", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()

        c, err := l.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }

        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")

                text, _ := reader.ReadString('\n')

                c.Write([]byte(text))



                netData, err := bufio.NewReader(c).ReadString('\n')
                netData = strings.Replace(netData, " ", "\n", -1)
                if err != nil {
                        if err == io.EOF {
                          fmt.Println("Exiting TCP server!")
                          return
                        }
                        fmt.Println(err)
                        return
                }
                if strings.TrimSpace(string([]byte(text))) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }

                fmt.Print("-> ", string(netData))

        }
}
