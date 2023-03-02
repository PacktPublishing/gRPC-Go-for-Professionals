package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter3/proto"
	"google.golang.org/protobuf/proto"
)

type AccountJSON struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
}

type AccountsJSON struct {
	Values []AccountJSON
}

func readFromJSON() time.Duration {
	data, err := os.ReadFile("accounts.json")
	if err != nil {
		log.Fatalln(err)
	}

	var accounts AccountsJSON

	start := time.Now()
	if err := json.Unmarshal(data, &accounts.Values); err != nil {
		log.Fatalf("failed to unmarshal: %v\n", err)
	}
	return time.Since(start)
}

func readFromPB() time.Duration {
	in, err := os.ReadFile("accounts.bin")
	if err != nil {
		log.Fatalln(err)
	}

	var data pb.Accounts

	start := time.Now()
	if err = proto.Unmarshal(in, &data); err != nil {
		log.Fatalln("failed to unmarshal: %v\n", err)
	}
	return time.Since(start)
}

const sizeData int = 100

func mean(data [sizeData]time.Duration) float64 {
	var sum float64

	for _, d := range data {
		sum += float64(d.Milliseconds())
	}

	return sum / float64(len(data))
}

func main() {
	var jsonTimes [sizeData]time.Duration
	var pbTimes [sizeData]time.Duration

	for i := 0; i < sizeData; i++ {
		elapsed := readFromJSON()
		jsonTimes[i] = elapsed

		elapsed = readFromPB()
		pbTimes[i] = elapsed
	}

	log.Printf("JSON: %fms", mean(jsonTimes))
	log.Printf("PB: %fms", mean(pbTimes))
}
