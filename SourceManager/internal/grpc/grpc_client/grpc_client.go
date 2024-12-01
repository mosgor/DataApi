package grpc_client

import (
	"SourceManager/internal/config"
	"SourceManager/internal/http_client/handlers/get"
	pb "SourceManager/internal/proto"
	"SourceManager/internal/repositories"
	"context"
	"log/slog"
	"net/http"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func GrpcClientConnection(pb_client pb.DataProcessorClient, IDs [][]int32, ctx context.Context, client http.Client, log *slog.Logger, cfg config.Config, repo repositories.SourceRepository) {
	stream, err := pb_client.ProcessData(ctx)
	var bsons string
	if err != nil {
		log.Error(err.Error())
	}
	for _, id := range IDs {
		bsons += get.GetMultipleData(client, id, cfg, repo)
		if err := stream.Send(&pb.Data{
			SourceId:    id,
			DataJson:    bsons,
			ArrivalTime: timestamppb.Now(),
		}); err != nil {
			log.Error(err.Error())
		}
	}
	status, err := stream.CloseAndRecv()
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("Grpc working:", "status:", status)

}
