package handler

import (
	"assignment2-golang-hacktiv8/dto"
	"assignment2-golang-hacktiv8/pkg/errs"
	"assignment2-golang-hacktiv8/service/order_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
    OrderService order_service.Service
}

func NewOrderHandler(orderService order_service.Service) orderHandler {
    return orderHandler{
        OrderService: orderService,
    }
}

// CreateOrder godoc
// @Tags orders
// @Description Create Order Data
// @ID create-new-order
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewOrderRequestDto true "request body json"
// @Success 201 {object} dto.NewOrderResponseDto
// @Router /orders [post]
func (oh *orderHandler) CreateOrder(ctx *gin.Context) {
    var newOrderRequest dto.NewOrderRequestDto

    if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
        errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
        ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
        return
    }

    response, err := oh.OrderService.CreateOrderWithItems(newOrderRequest)

    if err != nil {
        ctx.AbortWithStatusJSON(err.Status(), err)
        return
    }

    ctx.JSON(response.StatusCode, response)
}

// @Tags orders
// @Description Get Order with Item Data
// @ID get-orders-with-items
// @Produce json
// @Success 200 {object} GetOrdersResponse
// @Router /orders [get]
func (oh *orderHandler) GetOrders(ctx *gin.Context) {
    response, err := oh.OrderService.GetOrders()

    if err != nil {
        errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
        ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
        return
    }

    ctx.JSON(response.BaseResponse.StatusCode, response)
}

// @Tags orders
// @Description Update Order Data By Id
// @ID update-order
// @Accept json
// @Produce json
// @Param orderId path int true "order's id"
// @Param RequestBody body dto.NewOrderRequestDto true "request body json"
// @Success 200 {object} dto.NewOrderResponseDto
// @Router /orders/{orderId} [put]
func (oh *orderHandler) UpdateOrder(ctx *gin.Context) {
    var newOrderRequest dto.NewOrderRequestDto

    orderId, errParam := strconv.Atoi(ctx.Param("orderId"))

    if errParam != nil {
        errParsingParam := errs.NewBadRequest("orderId has to be a valid number value")
        ctx.AbortWithStatusJSON(errParsingParam.Status(), errParsingParam)
        return
    }

    if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
        errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
        ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
        return
    }

    response, err := oh.OrderService.UpdateOrder(orderId, newOrderRequest)

    if err != nil {
        ctx.AbortWithStatusJSON(err.Status(), err)
        return
    }

    ctx.JSON(response.StatusCode, response)
}

// @Tags orders
// @Description Delete Order Data By Id
// @ID delete-order
// @Param orderId path int true "order's id"
// @Success 200 {string} string "Order deleted successfully"
// @Router /orders/{orderId} [delete]
func (oh *orderHandler) DeleteOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
			errParsingParam := errs.NewBadRequest("orderId has to be a valid number value")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errParsingParam)
			return
	}

	// Panggil service untuk menghapus pesanan berdasarkan ID
	err = oh.OrderService.DeleteOrder(orderId)
	if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
