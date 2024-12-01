package get

import (
	"SourceManager/internal/config"
	"SourceManager/internal/logger"
	"SourceManager/internal/repositories"
	"context"
	"net/http"
	"strconv"
)

func GetData(client http.Client, source_id int, cfg config.Config, repository repositories.SourceRepository) string {

	log := logger.SetupLogger(cfg.Env)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.Timeout)
	defer cancel()
	read, err := repository.ReadOne(ctx, source_id)
	if err != nil {
		log.Error("FROM http_client/handlers/get", "Error:", err.Error())
	}

	addr := read.URL
	response, err := client.Get(addr)
	if err != nil {
		log.Error(err.Error())
	}
	body := make([]byte, response.ContentLength)
	response.Body.Read(body)
	return string(body)
}

func GetMultipleData(client http.Client, source_ids []int32, cfg config.Config, repository repositories.SourceRepository) string {
	log := logger.SetupLogger(cfg.Env)
	stringStart := "{"
	for _, id := range source_ids {
		stringStart += "\"" + strconv.Itoa(int(id)) + "\":" + GetData(client, int(id), cfg, repository)
	}
	stringStart += "}"
	log.Info("Data merged successfully", "data:", stringStart)
	return stringStart
}
