/*
* @Author: jiangzhuolin
* @Date:   2017-12-16 13:34:37
* @Last Modified by:   jiangzhuolin
* @Last Modified time: 2017-12-17 01:35:46
*/

package main

import (
    "io"
    "net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1> Hello, I'm Charlie Jiang!<h1>")
}

func main() {
    http.HandleFunc("/", firstPage)
    http.ListenAndServe(":8000", nil)
}