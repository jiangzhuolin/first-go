/*
* @Author: jiangzhuolin
* @Date:   2017-12-16 18:56:41
* @Last Modified by:   jiangzhuolin
* @Last Modified time: 2017-12-17 00:28:38
*/
package main

import (
    "io"
    "net/http"
    "os/exec"
    "log"
)

func reLaunch() {
    cmd := exec.Command("sh", "./deploy.sh")
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }
    err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1> Hello, this is my deploy server!<h1>")
    reLaunch()
}

func main() {
    http.HandleFunc("/", firstPage)
    http.ListenAndServe(":5000", nil)
}