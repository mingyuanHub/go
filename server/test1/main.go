package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"bytes"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":10888", nil)
}

func hello(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "[hello]")
	body, _ := getRequestBody(request)
	fmt.Println("[hello]", string(body))
	io.WriteString(response, string(body))
}

func getRequestBody(request *http.Request) ([]byte, error) {
	if request.Body != nil {
		byts, err := ioutil.ReadAll(request.Body) // io.ReadAll as Go 1.16, below please use ioutil.ReadAll
		if err != nil {
			return nil, err
		}
		request.Body = ioutil.NopCloser(bytes.NewReader(byts))
		return byts, nil
	}
	return make([]byte, 0), nil
}