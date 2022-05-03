package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	engine := gin.Default()
	// 自定义模板函数(坑点：自定义函数必须在文件引入之前)
	engine.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	engine.LoadHTMLGlob("gin/templates/**/*")
	engine.LoadHTMLFiles("gin/templates/index.tmpl")
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})

	// restful api
	// 查询
	engine.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "get",
		})
	})

	// 保存
	engine.POST("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "post",
		})
	})

	// 更新
	engine.PUT("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "put",
		})
	})

	// 删除
	engine.DELETE("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "delete",
		})
	})

	engine.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	engine.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	engine.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", "<a href='https://www.baidu.com'>百度</a>")
	})

	// 获取querystring参数
	engine.GET("/user/search", func(ctx *gin.Context) {
		username := ctx.DefaultQuery("username", "张三")
		address := ctx.Query("address")
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "success",
			"username": username,
			"address":  address,
		})
	})
	// form
	engine.POST("/user/search", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		address := ctx.PostForm("address")
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "success",
			"username": username,
			"address":  address,
		})
	})

	// json
	engine.POST("json", func(context *gin.Context) {
		jsonBytes, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(jsonBytes, &m)
		context.JSON(http.StatusOK, m)
	})
	// 启动http服务 默认8080端口
	engine.Run()
}
