package handler

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.GET("/", h.getAllUsers)
			user.GET("/:id", h.getUserById)
			user.PUT("/:id", h.updateUser)
			user.DELETE("/:id", h.deleteUser)

			cart := user.Group("/cart")
			{
				cart.POST(":id/", h.addBeatToCart)
				cart.GET(":id/", h.getAllBeatsFromCart)
				cart.GET(":id/:id", h.getBeatByIdFromCart)
				cart.DELETE(":id/", h.deleteAllBeatsInCart)
				cart.DELETE(":id/:id", h.deleteBeatInCart)
			}

			beats := user.Group("/beats")
			{
				beats.POST("/", h.addBeat)
				beats.PUT("/:id", h.updateBeat)
				beats.DELETE("/:id", h.deleteBeat)
			}
		}

		beats := router.Group("/beats")
		{
			beats.GET("/", h.getAllBeats)
			beats.GET("/:id", h.getBeatById)
			beats.PUT("/:id", h.updateBeat)
			beats.DELETE("/:id", h.deleteBeat)
		}
	}

	return router
}
