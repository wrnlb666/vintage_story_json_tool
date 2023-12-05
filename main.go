package main

import (
    "os"
    "io"
    "log"
    "os/user"
    "strings"
)


func rootDir() string {
    user, err := user.Current()
    if err != nil {
        log.Fatal(err.Error())
    }
    dir := user.HomeDir
    return dir
}


func targetFile(rd string) string {
    return rd + "\\AppData\\Roaming\\Tecent\\test.json"
}


func main() {
    fileName := targetFile(rootDir())

    read, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }

    data, err := io.ReadAll(read)
    if err != nil {
        log.Fatal(err)
    }

    read.Close()

    str := string(data)

    lines := strings.Split(str, "\n")
    for i, v := range lines {
        if strings.Contains(v, "language") {
            lines[i] = "\"language\": \"zh-cn\""
        }
    }

    str = strings.Join(lines, "\n")
    write, err := os.OpenFile(fileName, os.O_WRONLY, 0600);
    defer write.Close()
    io.WriteString(write, str)
}
