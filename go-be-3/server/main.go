// Package main implements a server for Beef service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

	pb "be3/beef"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement beef.BeefServer.
type server struct {
	pb.UnimplementedBeefServer
}

func GetJson(content string) string {
	beefMap := make(map[string]int)

	// split to words
	zp := regexp.MustCompile(`[,\.\s ]+`)
	beefs := zp.Split(content, -1)

	// after splitting, Empty string ("") always exists at the last position of variable 'beefs'
	beefsLen := len(beefs)
	if beefs[beefsLen-1] == "" {
		beefsLen--
	}
	log.Printf("beefs.len:%d\n", beefsLen)

	// summary
	for i := 0; i < beefsLen; i++ {
		beefMap[strings.ToLower(beefs[i])]++
	}

	// format json string
	json := ""
	for k, v := range beefMap {
		json += fmt.Sprintf("\"%s\":%d,", k, v)
	}

	// cut last ','
	json = json[:len(json)-1]

	return "{\"beef\":{" + json + "}}"
}

// Summary implements beef.BeefServer
func (s *server) Summary(_ context.Context, in *pb.SummaryRequest) (*pb.SummaryReply, error) {
	log.Println("Client Called")

	json := GetJson(in.GetContent())

	return &pb.SummaryReply{Json: json}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBeefServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
