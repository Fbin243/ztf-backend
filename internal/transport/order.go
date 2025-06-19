package transport

import (
	"fmt"
	"strconv"

	biz "ztf-backend/internal/business"
	"ztf-backend/internal/entity"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderBusiness *biz.OrderBusiness
}

func NewOrderHandler(orderBusiness *biz.OrderBusiness) *OrderHandler {
	return &OrderHandler{
		orderBusiness: orderBusiness,
	}
}

func (hdl *OrderHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := hdl.orderBusiness.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve orders"})
		return
	}
	ctx.JSON(200, orders)
}

func (hdl *OrderHandler) GetOrderById(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := hdl.orderBusiness.FindById(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Order not found"})
		return
	}
	ctx.JSON(200, order)
}

func (hdl *OrderHandler) CreateOrder(ctx *gin.Context) {
	var order entity.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id, err := hdl.orderBusiness.InsertOne(&order)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}
	ctx.JSON(201, gin.H{"id": id})
}

func (hdl *OrderHandler) UpdateOrder(ctx *gin.Context) {
	var order entity.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id, err := hdl.orderBusiness.UpdateOne(&order)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (hdl *OrderHandler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	val, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	deletedId, err := hdl.orderBusiness.DeleteOne(uint(val))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}
	ctx.JSON(200, gin.H{"deleted_id": deletedId})
}
