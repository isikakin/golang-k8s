package product

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-k8s/internal/domain/enitity"
	"golang-k8s/internal/service"
	"golang-k8s/internal/service/models"
	"net/http"
)

func GetProductDetailById(e *echo.Echo, productService service.ProductService) {
	e.GET("/:id", func(c echo.Context) error {
		var (
			err           error
			productDetail *models.ProductDetail
		)

		ctx := context.Background()

		productId := c.Param("id")

		fmt.Println("productId", productId)

		if productDetail, err = productService.GetProductDetailById(ctx, productId); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, productDetail)
	})
}

func InsertProduct(e *echo.Echo, productService service.ProductService) {
	e.POST("/insert", func(c echo.Context) error {

		var (
			err error
		)

		ctx := context.Background()

		request := new(InsertProductRequest)

		if err = c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("request deserialize error: %s", err))
		}

		product := enitity.NewProduct(
			request.CategoryId,
			request.Name,
			request.Price,
		)

		if err = productService.Insert(ctx, product); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, nil)

	})
}

func ReadinessProbe(e *echo.Echo, productService service.ProductService) {
	e.GET("/readiness", func(c echo.Context) error {

		var (
			ctx = context.Background()
			err error
		)

		if err = productService.Ping(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "pong")
	})
}
