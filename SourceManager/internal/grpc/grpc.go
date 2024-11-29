package grpc

import (
	pb "SourceManager/internal/proto"
	"fmt"
)

type serverAPI struct {
	pb.UnimplementedDataProcessorServer
}

func ProcessData()

func grpc() {
	fmt.Println("h")
}
