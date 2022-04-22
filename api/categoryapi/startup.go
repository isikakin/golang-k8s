package categoryapi

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang-k8s/internal/category"
	"golang-k8s/internal/category/docs"
	"golang-k8s/internal/domain/repository"
	"golang-k8s/internal/service"
	"golang-k8s/pkg/config"
	"golang-k8s/pkg/mongodb"
	"net/http"
	"time"
)

func Init(cmd *cobra.Command, args []string) error {

	docs.Init()

	var configuration config.Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic("configuration is invalid!")
	}

	fmt.Println("mongo", configuration.Database.Address)

	var mongoClient, mongoErr = mongodb.NewClient(configuration.Database.Address, configuration.Database.Replicaset, 10*time.Second)
	if mongoErr != nil {
		fmt.Println("mongo err", err.Error())
		panic(err)
	}

	var categoryRepository = repository.NewCategoryRepository(mongoClient, configuration.Database.DatabaseName)
	var categoryService = service.NewCategoryService(categoryRepository)

	e := echo.New()

	e.GET("/healthcheck", func(c echo.Context) error {
		fmt.Println("categoryapitest")
		return c.String(http.StatusOK, configuration.AppSettings.HealthCheck)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)


	category.GetAllCategories(e, categoryService)
	category.GetById(e, categoryService)
	category.InsertCategory(e, categoryService)
	category.UpdateCategory(e, categoryService)
	category.DeleteById(e, categoryService)

	if err = e.Start(":5000"); err != nil {
		panic(err)
	}

	return nil
}
