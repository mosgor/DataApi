package main

import (
	"SourceManager/internal/config"
	"SourceManager/internal/db/sources"
	"SourceManager/internal/grpc/grpc_client"
	"SourceManager/internal/http_client"
	"SourceManager/internal/http_server"
	logs "SourceManager/internal/logger"
	pb "SourceManager/internal/proto"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	log := logs.SetupLogger(cfg.Env)

	log.Info("Logger is up")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.Timeout)
	defer cancel()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://admin:%v@postgres:5432/DataApi", cfg.DatabasePass))
	if err != nil {
		log.Error("unable to connect to postgres")
		return
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Error("unable to ping postgres")
		return
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://mongo:27017").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Error("unable to connect to mongo in main")
		return
	}

	repo := sources.NewRepository(pool, client, log)

	server := http_server.CreateServer(repo)

	addr := fmt.Sprintf("data_processor" + cfg.GRPC.Port)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("FROM grpc_client", "Error:", err)
	}
	defer conn.Close()

	grpcClient := pb.NewDataProcessorClient(conn)
	ids := http_client.UpdateData(*cfg)
	httpClient := http.Client{
		Timeout: cfg.HTTP.Timeout,
	}
	go grpc_client.GrpcClientConnection(grpcClient, ids, ctx, httpClient, log, *cfg, repo)

	log.Info("Starting listening address")
	er := server.ListenAndServe()
	if er != nil {
		log.Error("Some error with listening")
		log.Error(er.Error())
	}
}
