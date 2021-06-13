package adapter

import (
	"github.com/cadasmeq/golang-ddd/internal/app/adapter/repository"
	"github.com/cadasmeq/golang-ddd/internal/app/application/usecase"

	"github.com/gin-gonic/gin"
)

// Controller
type Controller struct{}

var (
	orderRepository = repository.Order{}
)

// Router
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}

	// Routes
	r.GET("/order", ctrl.order)
	return r
}

func (ctrl Controller) order(c *gin.Context) {
	order := usecase.AddNewTransaction(orderRepository)
	c.JSON(200, order)
}
