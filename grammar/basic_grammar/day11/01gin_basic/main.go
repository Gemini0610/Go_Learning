package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// gin demo
// gin.Context是用来处理HTTP请求和响应的上下文对象，封装了HTTP请求的信息
// 处理get请求的/index路径并在indexHandler函数中进行相关的处理和响应
func indexHandler(c *gin.Context) {
	//具体的处理请求的业务逻辑
	c.JSON(http.StatusOK, gin.H{
		"msg": "helloworld!",
	})
}

// 处理登录的函数，但他过于简单，一般要连接数据库进行校验
// func loginHandler(c *gin.Context) {
// 	//获取当前HTTP请求的方法
// 	if c.Request.Method == "POST" {
// 		//处理用户提交过来的请求数据,这里的“username”就是input的name
// 		username := c.PostForm("username")
// 		password := c.PostForm("password")
// 		c.JSON(http.StatusOK, gin.H{
// 			//用json来返回值
// 			"username": username,
// 			"password": password,
// 		})
// 	} else {
// 		//返回一个登录的界面
// 		c.HTML(http.StatusOK, "login.html", nil)
// 	}
// }

// 下面是数据库校验的login函数
func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		//处理用户提交过来的请求数据
		//声明userInfo变量,将用户在网页提交的数据传入到变量中
		var u UserInfo
		//将用户提交的数据与结构体进行绑定bind
		//func (c *Context) ShouldBind(obj any) error 返回的时err
		err := c.ShouldBind(&u)
		if err != nil {
			//在给客户端返回错误响应时，使用fmt.Println输出错误信息并不会将其作为有效的JSON响应返回给客户端。c.JSON方法用于将数据序列化为JSON格式并发送给客户端。
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		}
		//去数据库中校验数据
		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	} else {
		//返回一个登录的界面
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

// 获取path参数
func postHandler(c *gin.Context) {
	//取到path参数,param用于获取url的参数
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")
	c.JSON(http.StatusOK, gin.H{
		"yaer":  year,
		"month": month,
		"day":   day,
	})
}

// querystring
func searchHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}
func main() {
	router := gin.Default() //得到一个默认的处理引擎

	//加载HTML文件
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", indexHandler)
	// 所有请求的URL以v1开头都交给下面的v1Group====>路由
	v1Group := router.Group("/v1")
	{
		v1Group.GET("/index", indexHandler)
	}
	//登录
	router.Any("/login", loginHandler)

	//querystring 用searchHandler函数
	router.GET("/search", searchHandler)
	//获取path参数,用自建的postHandler函数获取
	router.GET("/posts/:year/:month/:day", postHandler)
	router.Run(":9090")
}
