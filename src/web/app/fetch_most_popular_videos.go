package api // サブパッケージの概念

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		key := "AIzaSyBiR_3VkJE9a6kmEwUjZX53-k_TvWs0Hps"

		ctx := context.Background()

		yts, err := youtube.NewService(ctx, option.WithAPIKey(key)) // APIを実行するためのサービスを生成する。

		if err != nil {
			logrus.Fatalf("Error creating new Youtube service: %v", err)
		}

		// var score [2]string = [2]string{"id", "snippet"}
		// 上記がだめなのはなぜ?
		queryList := []string{"id", "snippet"}

		call := yts.Videos.List(queryList).Chart("mostPopular").MaxResults(3)

		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
