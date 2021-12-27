package main

import "fmt"

type Response struct {
	Name string
	Code int
}

func main() {
	response := &Response{
		Name: "mingyuan",
		Code: 204,
	}

	fmt.Println(Marshal(response))

	fmt.Println(MarshalToString(response))


	request := "{\"Name\":\"mingyuan\",\"Code\":203334}"

	UnmarshalFromString(request, response)

	fmt.Println(MarshalIndent(response, "", " "))

}
