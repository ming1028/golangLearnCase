package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	engine := gin.Default()
	engine.Use(statCost())
	// 自定义模板函数(坑点：自定义函数必须在文件引入之前)
	engine.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	engine.LoadHTMLGlob("gin/templates/**/*") // 只解析到第二层级下所有文件 /* 这一层所有的
	// engine.LoadHTMLFiles("gin/templates/index.tmpl")
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

	engine.GET("/auth", authHandler)
	engine.GET("/home", JWTAuthMiddleware(), homeHandler)
	// 启动http服务 默认8080端口
	//engine.Run()

	srv := http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 监听信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	for sign := range quit {
		switch sign {
		case syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown: ", err)
			}
			log.Println("Server exiting")
		default:
			log.Println("step 5: unknown signal", sign)
		}
	}
}

// 定义中间件
func statCost() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		// 调用该请求的剩余处理程序
		context.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println("耗费时间", cost.String(), cost.Nanoseconds())
	}
}

const TokenExpireDuration = time.Hour * 2

var tokenSecret = []byte("jwtSecret")

type MyClaims struct {
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

func genToken(username string) (string, error) {
	c := MyClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "gin-demo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(tokenSecret)
}

func parseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, err
	}
	return nil, errors.New("invalid token")
}

func authHandler(ctx *gin.Context) {
	username := "zhangsan"
	tokenString, err := genToken(username)
	fmt.Println(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"token": tokenString,
		},
	})
	return
}

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			ctx.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		fmt.Println(parts)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			ctx.Abort()
			return
		}
		mc, _ := parseToken(parts[1])
		fmt.Println(mc)
		ctx.Set("username", mc.UserName)
		ctx.Next()
	}
}

func homeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}
