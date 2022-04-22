package productapi

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang-k8s/internal/domain/repository"
	"golang-k8s/internal/product"
	"golang-k8s/internal/service"
	"golang-k8s/pkg/config"
	"golang-k8s/pkg/httpclient/category"
	"golang-k8s/pkg/mongodb"
	"net/http"
	"time"
)

func Init(cmd *cobra.Command, args []string) error {

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

	var productRepository = repository.NewProductRepository(mongoClient, configuration.Database.DatabaseName)
	var categoryClient = category.NewClient(configuration.CategoryApi.Url, time.Minute)
	var productService = service.NewProductService(productRepository, categoryClient)

	e := echo.New()

	e.GET("/healthcheck", func(c echo.Context) error {
		fmt.Println("productapitest")
		return c.String(http.StatusOK, configuration.AppSettings.HealthCheck)
	})

	product.GetProductDetailById(e, productService)
	product.InsertProduct(e, productService)
	product.ReadinessProbe(e,productService)

	if err = e.Start(":5001"); err != nil {
		panic(err)
	}

	return nil
}
