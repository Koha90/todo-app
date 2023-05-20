// Package handler has methods for api.
package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/koha90/todo-app/pkg/service"
)

// Handler has fields for handler api.
type Handler struct {
	services *service.Service
}

// NewHandler constructor for handler
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes initilazing routes api.
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListByID)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemByID)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
