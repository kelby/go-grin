package api

import "log"
import "net/http"

func Get(url string) (resp *Response, err error) {
  resp, err := http.Get(url)

  if err != nil {
    log.Println("Get error ", err.Error())
  }
}

func Post(url string, contentType string, body io.Reader) (resp *Response, err error) {
  resp, err := http.Post(url, contentType, body)

  if err != nil {
    log.Println("Post error ", err.Error())
  }
}
