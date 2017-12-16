/*
* @Author: jiangzhuolin
* @Date:   2017-12-16 18:56:41
* @Last Modified by:   jiangzhuolin
* @Last Modified time: 2017-12-17 01:01:53
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
    io.WriteString(w, "<p>Trying to deploy the webserver,please wait a minute...<p><br>")
    reLaunch()
    io.WriteString(w, "<p>Finished deploying the webserver<p>")

}

func main() {
    http.HandleFunc("/", firstPage)
    http.ListenAndServe(":5000", nil)
}