package api

import (
	"os"

	"gitlab.com/p9359/backend-prob/febry-go/internal/app"
	"gitlab.com/p9359/backend-prob/febry-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, app app.BookApp) {
	keyWebPosOnline := os.Getenv("JWT_SECRET")
	jwt_key_web_online := []byte(keyWebPosOnline)
	authService := middleware.NewJWTService(jwt_key_web_online)

	r.Use(middleware.CORSMiddleware())
	api := r.Group("api/v1")
	{
		api.Use(middleware.AuthHandler(authService))
		{
			api.GET("/category/detail/:id", app.GetCategoryDetail)

			api.GET("/book", app.GetListBook)
			api.GET("/book/detail/:uuid", app.GetBookDetail)
			api.POST("/book", app.CreateBook)
			api.PUT("/book/:uuid", app.UpdateBook)
			api.DELETE("/book/:uuid", app.DeleteBook)

			api.GET("book/stock/history/:uuid", app.GetBookHistory)
			api.PUT("book/stock/increase/:uuid", app.IncreaseBookStock)
			api.PUT("book/stock/decrease/:uuid", app.DecreaseBookStock)

			api.GET("book/transactions", app.GetTransactions)
			api.POST("book/transaction/add-to-cart", app.AddToCart)
			api.POST("book/transaction/on-borrow", app.OnBorrow)
			api.POST("book/transaction/finish", app.Finish)
			api.POST("book/transaction/cancel", app.Cancel)
		}
	}
}
