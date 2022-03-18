package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/middleware"
	"go-gin-example/pkg/setting"
	v1 "go-gin-example/routers/api/v1"
)

func InitRouter()*gin.Engine{
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	engine.GET("/auth",v1.GetAuth)
	v1api := engine.Group("/api/v1")
	v1api.Use(middleware.JWT())
	{
		//获取标签列表
		v1api.GET("/tags", v1.GetTags)
		//新建标签
		v1api.POST("/tags", v1.AddTag)
		//更新指定标签
		v1api.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		v1api.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		v1api.GET("/articles", v1.GetArticles)
		//获取指定文章
		v1api.GET("/articles/:id", v1.GetArticle)
		//新建文章
		v1api.POST("/articles", v1.AddArticle)
		//更新指定文章
		v1api.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		v1api.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return engine
}
