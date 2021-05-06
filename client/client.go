package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jaakali/ipfilter"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	t0 := time.Now()
	// ctx0, cancel0 := context.WithTimeout(context.Background(), time.Second)
	conn, err := grpc.Dial(":41563", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println(time.Since(t0))
	defer conn.Close()
	c := pb.NewIpFilterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	for {
		r, err := c.Rewrite(ctx, &pb.IpReq{Ip4: "17.0.2.1.3"})
		if err != nil {
			log.Printf("could not greet: %v", err)
		} else {
			log.Printf("Greeting: %t", r.GetRet())
		}
		time.Sleep(2*time.Second)
	}
}