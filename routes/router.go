package routes

import (
	v1 "GinBlog/api/v1"
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		//用户路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		//分类模块
		routerV1.POST("category/add", v1.AddCategory)
		routerV1.GET("categories", v1.GetCate)
		routerV1.PUT("category/:id", v1.EditCate)
		routerV1.DELETE("category/:id", v1.DeleteCate)
		//文章
		routerV1.POST("article/add", v1.AddArticle)
		routerV1.GET("articles", v1.GetArt)
		routerV1.GET("article/info/:id", v1.GetArtInfo)
		routerV1.GET("article/list/:id", v1.GetCateArt)
		routerV1.PUT("article/:id", v1.EditArt)
		routerV1.DELETE("article/:id", v1.DeleteArt)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
