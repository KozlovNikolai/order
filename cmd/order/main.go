package main

import (
	"time"

	"github.com/KozlovNikolai/order/configs"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	sugar = zap.NewExample().Sugar()
	defer sugar.Sync()

}

func main() {
	sugar.Info("Приложение запущено")

	err := configs.ReadConfig()
	if err != nil {
		sugar.Panicf("error Read Config:%w", err)
	}
	url := "yandex.ru"
	sugar.Errorw("eRROR MESSAGE",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Warnw("eRROR MESSAGE",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
}
