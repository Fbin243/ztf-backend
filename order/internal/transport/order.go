package transport

import (
	"net/http"

	biz "ztf-backend/order/internal/business"
	"ztf-backend/order/internal/entity"
	dto2 "ztf-backend/order/internal/transport/dto"
	errs "ztf-backend/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type OrderHandler struct {
	orderBusiness    *biz.OrderBusiness
	merchantBusiness *biz.MerchantBusiness
	userBusiness     *biz.UserBusiness
}

func NewOrderHandler(
	orderBusiness *biz.OrderBusiness,
	merchantBusiness *biz.MerchantBusiness,
	userBusiness *biz.UserBusiness,
) *OrderHandler {
	return &OrderHandler{
		orderBusiness:    orderBusiness,
		merchantBusiness: merchantBusiness,
		userBusiness:     userBusiness,
	}
}

func (hdl *OrderHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := hdl.orderBusiness.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Find merchant and user info of each order
	userIDs := []string{}
	merchantIDs := []string{}
	for _, order := range orders {
		if order.UserId != nil {
			userIDs = append(userIDs, *order.UserId)
		}
		merchantIDs = append(merchantIDs, order.MerchantId)
	}

	merchants, err := hdl.merchantBusiness.FindByIds(merchantIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	merchantMap := map[string]entity.Merchant{}
	for _, merchant := range merchants {
		merchantMap[merchant.Id] = merchant
	}

	users, err := hdl.userBusiness.FindByIds(userIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userMap := map[string]entity.User{}
	for _, user := range users {
		userMap[user.Id] = user
	}

	orderDtos := lo.Map(orders, func(order entity.Order, _ int) dto2.Order {
		merchant := merchantMap[order.MerchantId]
		var user *dto2.User
		if order.UserId != nil {
			user = &dto2.User{
				Id:       *order.UserId,
				Username: userMap[*order.UserId].Username,
				Email:    userMap[*order.UserId].Email,
			}
		}

		return dto2.Order{
			Id:        order.Id,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			PayAmount: order.PayAmount,
			Info:      order.Info,
			Merchant: dto2.Merchant{
				Id:       order.MerchantId,
				Username: merchant.Username,
				Email:    merchant.Email,
			},
			User: user,
		}
	})

	ctx.JSON(http.StatusOK, orderDtos)
}

func (hdl *OrderHandler) GetOrderById(ctx *gin.Context) {
	stringId := ctx.Param("id")
	order, err := hdl.orderBusiness.FindByIdWithMerchantAndUser(stringId)
	if err == errs.ErrorNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderDto := &dto2.Order{}
	err = copier.Copy(orderDto, order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderDto)
}

func (hdl *OrderHandler) CreateOrder(ctx *gin.Context) {
	var order entity.CreateOrderInput
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := GetValidator().Struct(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.orderBusiness.InsertOne(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (hdl *OrderHandler) UpdateOrder(ctx *gin.Context) {
	stringId := ctx.Param("id")
	var order entity.UpdateOrderInput
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := GetValidator().Struct(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.orderBusiness.UpdateOne(stringId, &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *OrderHandler) DeleteOrder(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := hdl.orderBusiness.DeleteOne(stringId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *OrderHandler) PayForOrder(ctx *gin.Context) {
	stringId := ctx.Param("id")
	var input entity.PayOrderInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := GetValidator().Struct(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.orderBusiness.PayForOrder(stringId, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
