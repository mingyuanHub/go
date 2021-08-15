package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"test1-client/services"
)

func main()  {
	creds, err := credentials.NewClientTLSFromFile("keys/server_no_passwd.crt", "mingyuan.com")
	if err != nil {
		fmt.Println(err)
	}
	con, err := grpc.Dial(":8084", grpc.WithTransportCredentials(creds))

	//con, err := grpc.Dial(":8084", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer  con.Close()
	//
	//prodClient := services.NewProdServiceClient(con)
	//prodRes, err := prodClient.GetProdStock(context.Background(),
	//	&services.ProdRequest{ProdId: 12})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(prodRes.ProdStock)

	maClient := services.NewMaServiceClient(con)
	maRes, err := maClient.GetName(context.Background(), &services.MaRequest{Age: 12})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(maRes.Name)




}
