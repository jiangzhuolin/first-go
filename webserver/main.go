/*
* @Author: jiangzhuolin
* @Date:   2017-12-16 13:34:37
* @Last Modified by:   jiangzhuolin
* @Last Modified time: 2017-12-17 00:48:09
*/

package main

import (
    "io"
    "net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1> Hello, this is a test message for DevOps!<h1>")
}

func main() {
    http.HandleFunc("/", firstPage)
    http.ListenAndServe(":8000", nil)
}