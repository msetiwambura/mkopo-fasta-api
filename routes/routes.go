package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"loanapi/controllers"
	"loanapi/middlewares"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"ChannelID", "IPAddress", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))
	r.POST("/login", controllers.Login)
	r.GET("/customers/api/customers/fetch", controllers.GetCustomers)
	r.GET("/loans/api/loans/fetch", controllers.GetLoans)
	r.POST("/customers/api/customers/create", controllers.CreateCustomer)
	r.POST("/payments/api/payments/create", controllers.CreatePayment)
	// Protected routes
	authorized := r.Group("/api")
	authorized.Use(middlewares.AuthMiddleware())
	{

		authorized.GET("/protected", func(c *gin.Context) {
			ChannelID, _ := c.Get("ChannelID")
			IPAddress, _ := c.Get("IPAddress")
			c.JSON(http.StatusOK, gin.H{"message": "Hello, " + ChannelID.(string) + " from " + IPAddress.(string)})
		})
		//authorized.GET("/customers", controllers.GetCustomers)
		authorized.GET("/customers/:id", controllers.GetCustomer)
		//authorized.POST("/customers", controllers.CreateCustomer)
		authorized.PUT("/customers/:id", controllers.UpdateCustomer)
		authorized.DELETE("/customers/:id", controllers.DeleteCustomer)
		authorized.POST("/loans", controllers.CreateLoan)
		authorized.POST("/collaterals", controllers.CreateCollateral)
		authorized.POST("/guarantors", controllers.CreateGuarantor)
		//authorized.POST("/payments", controllers.CreatePayment)
		authorized.GET("/loans/customer/:customerID", controllers.GetLoansByCustomerID)
		authorized.GET("/loans/summary", controllers.SummarizeLoans)
		//authorized.GET("/loans", controllers.GetLoans)
		authorized.PUT("/loans", controllers.UpdateLoans)
		authorized.GET("/loans/:id", controllers.GetLoanById)
	}

	return r
}
