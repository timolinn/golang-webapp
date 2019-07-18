package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	shopController shop
)

// starts up the application
func StartUp(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	homeController.loginTemplate = templates["login.html"]
	homeController.standLocatorTemplate = templates["stand_locator.html"]

	// register routes
	homeController.registerRoutes()
	shopController.registerRoutes()

	// static routing
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
