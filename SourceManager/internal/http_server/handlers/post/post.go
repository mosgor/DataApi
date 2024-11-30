package post

import (
	"SourceManager/internal/config"
	"SourceManager/internal/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct{} // заменить на данные с дб

var data []Data

func ServerPost() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var newData Data
		cfg := config.MustLoad()
		log := logger.SetupLogger(cfg.Env)
		err := json.NewDecoder(r.Body).Decode(&newData)
		if err != nil {
			log.Error(err.Error())
			return
		}

		data = append(data, newData)
		log.Info(fmt.Sprintf("New data read successfully:%v", newData))

	})
}
