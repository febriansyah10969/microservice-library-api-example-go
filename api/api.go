package api

import (
	"gitlab.com/p9359/backend-prob/febry-go/internal/app"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, app app.BookApp) {
	// keyWebPosOnline := os.Getenv("JWT_SECRET_ONLINE")
	// keyWebDashboard := os.Getenv("JWT_SECRET")
	// jwt_key_web_online := []byte(keyWebPosOnline)
	// jwt_key_dashboard_online := []byte(keyWebDashboard)
	// authService := middleware.NewJWTService(jwt_key_web_online, jwt_key_dashboard_online)

	// r.Use(middleware.CORSMiddleware())
	api := r.Group("api/v1")
	{
		api.Use()
		{
			// api.GET("/:id", app.GetUnreviewedList)
			// api.GET("/reviewed/:id", app.GetReviewedProductList)
			// api.GET("/dash-reviewed", app.GetDashboardReviewedProductList)
			// api.POST("/california", app.ReviewConcurrently)
			// api.POST("/californian", app.ReviewSequentially)
		}
	}
}
