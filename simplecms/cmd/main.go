package main

import (
  "github.com/ardeshir/simplecms"
  "net/http"
)

func main() {

   http.HandleFunc("/",      simplecms.ServeIndex)
   http.HandleFunc("/new",   simplecms.HandleNew)
   http.HandleFunc("/page/", simplecms.ServePage )
   http.HandleFunc("/post/", simplecms.ServePost )
   http.ListenAndServe(":3000", nil)

}
