package gRPC

import (
	pb "ModelOrchestrator/pkg/internal/proto"
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/structs"
	"context"
	"google.golang.org/grpc"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type rpcServer struct {
	pb.UnimplementedModelOrchestratorServer
	logsRepo  repositories.LogsRepository
	modelRepo repositories.ModelRepository
	netClient *http.Client
	timeout   time.Duration
	log       *slog.Logger
}

func (r *rpcServer) SendData(stream pb.ModelOrchestrator_SendDataServer) error {
	r.log.Info("Starting gRPC connection")
	for {
		processed, err := stream.Recv()
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, r.timeout)
		if err == io.EOF {
			r.log.Info("Closing gRPC connection")
			cancel()
			return stream.SendAndClose(&pb.Status{
				Message: "OK",
			})
		}
		if err != nil {
			cancel()
			return err
		}
		modelId := processed.ModelId
		one, err := r.modelRepo.ReadOne(ctx, int(modelId))
		if err != nil {
			cancel()
			return err
		}
		url := one.Url
		lm := structs.LogsModel{
			SourceId: processed.SourceId,
			ModelId:  int(processed.ModelId),
			Time:     time.Since(processed.ArrivalTime.AsTime()),
			Status:   "OK",
		}

		// TODO: add test ml (server)
		_, err = r.netClient.Get(url)
		if err != nil {
			lm.Status = "ERROR"
			r.log.Warn("No response from data!")
			//cancel()
			//return err
		}

		lm.TimeWithResponse = time.Since(processed.ArrivalTime.AsTime())
		err = r.logsRepo.Create(ctx, &lm)
		if err != nil {
			cancel()
			return err
		}
	}
}

func NewRpcServer(l repositories.LogsRepository, m repositories.ModelRepository,
	n *http.Client, t time.Duration, log *slog.Logger) *grpc.Server {
	serv := &rpcServer{
		logsRepo:  l,
		modelRepo: m,
		netClient: n,
		timeout:   t,
		log:       log,
	}
	s := grpc.NewServer()
	pb.RegisterModelOrchestratorServer(s, serv)
	return s
}
