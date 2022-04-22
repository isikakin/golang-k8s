package category

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-k8s/internal/domain/enitity"
	"golang-k8s/internal/service"
	"net/http"
)

// GetAllCategories godoc
// @Summary Get All Categories
// @Description Get all category list
// @Produce  json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Tags Category
// @Router /all [get]
func GetAllCategories(e *echo.Echo, categoryService service.CategoryService) {
	e.GET("/all", func(c echo.Context) error {

		ctx := context.Background()

		var (
			categories []enitity.Category
			err        error
		)

		if categories, err = categoryService.GetAll(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Println("lencat", len(categories))
		return c.JSON(http.StatusOK, categories)

	})
}

// GetById godoc
// @Summary Get Category By Id
// @Description Get category by id
// @Produce  json
// @Param id path string true "categoryId"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Tags Category
// @Router /{id} [get]
func GetById(e *echo.Echo, categoryService service.CategoryService) {
	e.GET("/:id", func(c echo.Context) error {
		var (
			err      error
			category *enitity.Category
		)

		ctx := context.Background()

		categoryId := c.Param("id")

		if category, err = categoryService.GetById(ctx, categoryId); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, category)
	})
}

// InsertCategory godoc
// @Summary Insert category
// @Description Insert new category
// @Produce  json
// @Param request body InsertCategory true "request"
// @Success 201 {object} string
// @Failure 500 {object} string
// @Tags Category
// @Router /insert [post]
func InsertCategory(e *echo.Echo, categoryService service.CategoryService) {
	e.POST("/insert", func(c echo.Context) error {

		var (
			err error
		)

		ctx := context.Background()

		request := new(InsertCategoryRequest)

		if err = c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("request deserialize error: %s", err))
		}

		category := enitity.NewCategory(request.Name)

		if err = categoryService.Insert(ctx, category); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, nil)

	})
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update category fields
// @Produce  json
// @Param request body UpdateCategoryRequest true "request"
// @Success 201 {object} string
// @Failure 500 {object} string
// @Tags Category
// @Router /update/{id} [put]
func UpdateCategory(e *echo.Echo, categoryService service.CategoryService) {
	e.PUT("/update/:id", func(c echo.Context) error {

		var (
			err error
		)

		ctx := context.Background()

		request := new(UpdateCategoryRequest)
		categoryId := c.Param("id")

		if err = c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("request deserialize error: %s", err))
		}

		category := enitity.NewCategory(request.Name)

		if err = categoryService.UpdateById(ctx, categoryId, category); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)

	})
}

// DeleteById godoc
// @Summary Delete Category By Id
// @Description Delete category by id
// @Produce  json
// @Param id path string true "categoryId"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Tags Category
// @Router /delete/{id} [delete]
func DeleteById(e *echo.Echo, categoryService service.CategoryService) {
	e.DELETE("delete/:id", func(c echo.Context) error {
		var (
			err error
		)

		ctx := context.Background()

		categoryId := c.Param("id")

		if err = categoryService.DeleteById(ctx, categoryId); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, nil)
	})
}
