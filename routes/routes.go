package routes

import (
	v1 "Mygin/api/v1"
	"Mygin/middleware"
	"Mygin/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArtInfo)
		auth.GET("admin/article", v1.GetArt)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)
	}

	routerV1 := r.Group("api/v1")
	{
		//用户模块的路由接口routerV1.POST("user/add", v1.AddUser)
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		//分类模块的路由接口
		routerV1.GET("category", v1.GetCate)
		//文章模块的路由接口
		routerV1.GET("article", v1.GetArt)
		routerV1.GET("article/list/:id", v1.GetCateArt)
		routerV1.GET("article/info/:id", v1.GetArtInfo)
		//登陆模块的接口
		routerV1.POST("login", v1.Login)

	}

	_ = r.Run(utils.HttpPort)
}
