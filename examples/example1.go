package main

/* (c) 2023 e1z0 <e1z0@vintage2000.org>
 * Simple example how to send files over IrDA using obexftp library golang bindings
 */

import (
    "obexftp"
    "os"
    "fmt"
    "path/filepath"
    "errors"
)


func main() {
    if len(os.Args) < 2 {
        fmt.Println("Tiny OBEX (written in go)")
        fmt.Println("Push tool for IrDA")
        fmt.Printf("Usage: %s <filename>\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
        fmt.Printf("File %s does not seem to exist\n",filename)
        os.Exit(1)
    }
    filebase := filepath.Base(filename)
    cli, err := obexftp.Open()
    if err != nil {
        fmt.Printf("Error: %s\n",err)
        os.Exit(1)
    }
    err = obexftp.Connect(cli)
    if err != nil {
        fmt.Printf("Error: %s\n",err)
        os.Exit(1)
    }
    err = obexftp.Push(cli,filename,filebase)
    if err != nil {
        fmt.Printf("Error: %s\n",err)
        os.Exit(1)
    }
    err = obexftp.Disconnect(cli)
    if err != nil {
        fmt.Printf("Error: %s\n",err)
        os.Exit(1)
    }
    obexftp.Close(cli)
    fmt.Printf("Everything went ok\n")
}

