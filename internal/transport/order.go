package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	biz "ztf-backend/internal/business"
	"ztf-backend/internal/entity"
	"ztf-backend/internal/utils"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (hdl *OrderHandler) GetOrderById(ctx *gin.Context) {
	stringId := ctx.Param("id")
	uintId, err := utils.ConvertStringToUInt(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	order, err := hdl.orderBusiness.FindById(uintId)
	if err == utils.ErrorNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (hdl *OrderHandler) CreateOrder(ctx *gin.Context) {
	var order entity.CreateOrderInput
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := GetValidator().Struct(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.orderBusiness.InsertOne(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (hdl *OrderHandler) UpdateOrder(ctx *gin.Context) {
	stringId := ctx.Param("id")
	uintId, err := utils.ConvertStringToUInt(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var order entity.UpdateOrderInput
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := hdl.orderBusiness.UpdateOne(uintId, &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *OrderHandler) DeleteOrder(ctx *gin.Context) {
	stringId := ctx.Param("id")
	uintId, err := utils.ConvertStringToUInt(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	id, err := hdl.orderBusiness.DeleteOne(uintId)
	if err == utils.ErrorNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
