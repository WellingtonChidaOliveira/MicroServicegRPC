package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"gitub.com/wellingtonchida/micro/client"
)

func main() {
	client := client.New("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "BT")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", price)

	return
	listenAddr := flag.String("listen-addr", ":3000", "server listen address")
	svc := NewLoggingService(NewMetricsService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
