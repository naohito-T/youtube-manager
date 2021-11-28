package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/naohito-T/youtube-manager-api/src/web/api"
)

func Init(e *echo.Echo) {
	// ルートをグループ化していて、このブロックの中で定義するエンドポイントはすべて接頭辞として/apiが付与される。
	g := e.Group("/api")
	{
		// バリデーションは何が有名？
		// エンドポイントは/api/popularになる。
		g.GET("/popular", api.FetchMostPopularVideos()) // アクセス時に実行するmethod
	}

}
