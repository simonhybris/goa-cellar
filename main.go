// +build !appengine

package main

import (
	"github.com/raphael/goa"
	"github.com/raphael/goa-cellar/app"
	"github.com/raphael/goa-cellar/controllers"
	"github.com/raphael/goa-cellar/js"
	"github.com/raphael/goa-cellar/schema"
	"github.com/raphael/goa-cellar/swagger"
	"github.com/raphael/goa-middleware/middleware"
)

func main() {
	// Create goa service
	service := goa.New("cellar")

	// Setup basic middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest())
	service.Use(middleware.Recover())

	// Mount account controller onto service
	ac := controllers.NewAccount(service)
	app.MountAccountController(service, ac)

	// Mount bottle controller onto service
	bc := controllers.NewBottle(service)
	app.MountBottleController(service, bc)

	// Mount Swagger Spec controller onto service
	swagger.MountController(service)

	// Mount JSON Schema controller onto service
	schema.MountController(service)

	// Mount JavaScript example
	js.MountController(service)

	// Run service
	service.ListenAndServe(":8080")
}
