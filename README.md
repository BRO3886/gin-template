# Learning Gin, following clean arch

## What is Gin

Gin is a framework written in Go. To be honest, it is not that popular and I had a hard time reading and understanding it's structure and code.

## What is clean architecture

Refer these:

* https://github.com/BRO3886/clean-go-notes
* https://github.com/supercmmetry/bandersnatch
* https://medium.com/gdg-vit/clean-architecture-the-right-way-d83b81ecac6


## Why use Gin?

0. **It uses context since its based on [httprouter](https://github.com/julienschmidt/httprouter)**

1. **Route grouping and great support for middleware**

```go
[main.go]

v1 := r.Group("api/v1")
{
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	usrGroup := v1.Group("/user")
	{
		usrGroup.POST("/register", handlers.RegisterUser(userSvc))
		usrGroup.POST("/login", handlers.LoginUser(userSvc))
		usrGroup.Use(middleware.BasicJWTAuth(userSvc))
		{
			usrGroup.GET("/getdetails", handlers.GetUserDetails(userSvc))
		}
	}
	// v1.Group("article")
}
```

2. **Logger support**

```bash
[GIN-debug] GET    /api/v1/ping              --> main.main.func1 (3 handlers)
[GIN-debug] POST   /api/v1/user/register     --> github.com/BRO3886/gin-learn/api/handlers.RegisterUser.func1 (3 handlers)
[GIN-debug] POST   /api/v1/user/login        --> github.com/BRO3886/gin-learn/api/handlers.LoginUser.func1 (3 handlers)
[GIN-debug] GET    /api/v1/user/getdetails   --> github.com/BRO3886/gin-learn/api/handlers.GetUserDetails.func1 (4 handlers)
2020/08/15 18:33:03 runnning on port 5432
[GIN-debug] Listening and serving HTTP on :5432
[GIN] 2020/08/15 - 18:34:14 | 200 |       199.2µs |             ::1 | GET      "/api/v1/ping"
[GIN] 2020/08/15 - 18:34:14 | 404 |         4.1µs |             ::1 | GET      "/favicon.ico"
```

3. **Reduces a lot of boilerplate code attached with net/http and is a lot faster than net/http**

```go
[api/handlers/user.go, Login function]

user, err := svc.Login(user.Email, user.Password)
if err != nil {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	ctx.Abort()
	return
}
```

4. **Middleware support**

```go
[main.go]

articleGroup := v1.Group("/articles")
articleGroup.Use(middleware.BasicJWTAuth(userSvc))
{
    articleGroup.POST("/create", handlers.CreateNewArticle(articleSvc))
    articleGroup.GET("/myarticles", handlers.GetArticlesByUser(articleSvc))
    articleGroup.GET("", handlers.GetAllArticles(articleSvc))
}
```

> If you liked the repo please conisder giving it a ⭐


