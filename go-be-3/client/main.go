// Package main implements a client for Beef service.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	pb "be3/beef"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func getContent(fileName string) *string {

	fmt.Printf("File Name:%s\n", fileName)

	path := getExePath()
	file, err := os.ReadFile(path + "/../files/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	content := flag.String("content"+fileName, string(file[:]), "Content to count")
	flag.Parse()

	return content
}

func GetConnection() *grpc.ClientConn {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func GetJson(c pb.BeefClient, fileName string) string {

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Summary(ctx, &pb.SummaryRequest{Content: *getContent(fileName)})
	if err != nil {
		log.Fatalf("could not summary: %v", err)
	}

	return r.GetJson()
}

func getExePath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("unable to get the current filename")
		return ""
	}
	dirname := filepath.Dir(filename)
	return dirname
}

func main() {
	// Set up a connection to the server.
	conn := GetConnection()
	defer conn.Close()
	c := pb.NewBeefClient(conn)

	/* First File */
	fileName := "short.txt"
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(GetJson(c, fileName)), "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	log.Printf("Json:\n%s\n", prettyJSON.String())

	/* Second File */
	fileName = "beef.txt"
	var prettyJSON2 bytes.Buffer
	error = json.Indent(&prettyJSON2, []byte(GetJson(c, fileName)), "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	log.Printf("Json:\n%s\n", prettyJSON.String())

}
