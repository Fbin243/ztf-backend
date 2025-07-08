package rest

import (
	"net/http"
	"ztf-backend/services/order/internal/auth"
	"ztf-backend/services/order/internal/entity"
	"ztf-backend/services/order/internal/transport/dto"
	"ztf-backend/services/order/internal/transport/validation"
	"ztf-backend/services/order/pkg/convert"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"

	biz "ztf-backend/services/order/internal/business"

	errs "ztf-backend/services/order/internal/errors"
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
	orders, err := hdl.orderBusiness.GetOrderList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Find merchant and user info of each order
	userIDs := []int64{}
	merchantIDs := []int64{}
	for _, order := range orders {
		if order.UserId != nil {
			userIDs = append(userIDs, *order.UserId)
		}
		merchantIDs = append(merchantIDs, order.MerchantId)
	}

	merchants, err := hdl.merchantBusiness.FindByIds(ctx, merchantIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	merchantMap := map[int64]entity.Merchant{}
	for _, merchant := range merchants {
		merchantMap[merchant.Id] = merchant
	}

	users, err := hdl.userBusiness.FindByIds(ctx, userIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userMap := map[int64]entity.User{}
	for _, user := range users {
		userMap[user.Id] = user
	}

	orderDtos := lo.Map(orders, func(order entity.Order, _ int) dto.Order {
		merchant := merchantMap[order.MerchantId]
		var user *dto.User
		if order.UserId != nil {
			user = &dto.User{
				Id:       *order.UserId,
				Username: userMap[*order.UserId].Username,
				Email:    userMap[*order.UserId].Email,
			}
		}

		return dto.Order{
			Id:        order.Id,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			PayAmount: order.PayAmount,
			Info:      order.Info,
			Merchant: dto.Merchant{
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
	id := convert.MustConvStrToInt(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
		return
	}

	order, err := hdl.orderBusiness.GetOrderWithMerchantAndUser(ctx, id)
	if err == errs.ErrorNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderDto := &dto.Order{}
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

	err := validation.GetValidator().Struct(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.orderBusiness.CreateOrder(ctx, &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (hdl *OrderHandler) UpdateOrder(ctx *gin.Context) {
	orderId := convert.MustConvStrToInt(ctx.Param("id"))
	if orderId == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
		return
	}

	var order entity.UpdateOrderInput
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validation.GetValidator().Struct(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.orderBusiness.UpdateOrder(ctx, orderId, &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *OrderHandler) DeleteOrder(ctx *gin.Context) {
	orderId := convert.MustConvStrToInt(ctx.Param("id"))
	if orderId == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
		return
	}

	id, err := hdl.orderBusiness.DeleteOrder(ctx, orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (hdl *OrderHandler) PayForOrder(ctx *gin.Context) {
	userId := convert.MustConvStrToInt(ctx.GetHeader("X-User-Id"))
	orderId := convert.MustConvStrToInt(ctx.Param("id"))
	if userId == 0 || orderId == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	reqCtx := auth.SetAuthKey(ctx.Request.Context(), userId)

	var input entity.PayOrderInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validation.GetValidator().Struct(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := hdl.orderBusiness.PayForOrder(reqCtx, orderId, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
