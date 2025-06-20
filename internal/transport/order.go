package transport

import (
	biz "ztf-backend/internal/business"
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"

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
	stringId := ctx.Param("id")
	uintId, err := utils.ConvertStringToUInt(stringId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	order, err := hdl.orderBusiness.FindById(uintId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Order not found"})
		return
	}
	ctx.JSON(200, order)
}

func (hdl *OrderHandler) CreateOrder(ctx *gin.Context) {
	var order entity.CreateOrderInput
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
	stringId := ctx.Param("id")
	uintId, err := utils.ConvertStringToUInt(stringId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var order entity.UpdateOrderInput
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id, err := hdl.orderBusiness.UpdateOne(uintId, &order)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (hdl *OrderHandler) DeleteOrder(ctx *gin.Context) {
	stringId := ctx.Param("id")
	uintId, err := utils.ConvertStringToUInt(stringId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	id, err := hdl.orderBusiness.DeleteOne(uintId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}
	ctx.JSON(200, gin.H{"deleted_id": id})
}
