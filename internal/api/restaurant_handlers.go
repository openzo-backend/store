

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/service"
)

type RestaurantHandler struct {
	restaurantService service.RestaurantService
}

func NewRestaurantHandler(restaurantService *service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{restaurantService: *restaurantService}
}

func (h *RestaurantHandler) CreateRestaurantDetails(ctx *gin.Context) {
	
	var restaurant models.RestaurantDetails

	ctx.ShouldBindJSON(&restaurant)

	createdRestaurant, err := h.restaurantService.CreateRestaurantDetails(ctx, restaurant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdRestaurant)

}

func (h *RestaurantHandler) UpdateRestaurantDetails(ctx *gin.Context) {
	var restaurant models.RestaurantDetails

	ctx.ShouldBindJSON(&restaurant)

	updatedRestaurant, err := h.restaurantService.UpdateRestaurantDetails(ctx, restaurant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedRestaurant)
}

func (h *RestaurantHandler) GetRestaurantByID(ctx *gin.Context) {
	id := ctx.Param("id")

	restaurant, err := h.restaurantService.GetRestaurantByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, restaurant)
}

func (h *RestaurantHandler) GetRestaurantDetailsByStoreID(ctx *gin.Context) {
	id := ctx.Param("store_id")

	restaurantDetails, err := h.restaurantService.GetRestaurantDetailsByStoreID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, restaurantDetails)
}

func (h *RestaurantHandler) GetRestaurantsByUserID(ctx *gin.Context) {

	id := ctx.Param("user_id")

	restaurants, err := h.restaurantService.GetRestaurantsByUserID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, restaurants)
}

func (h *RestaurantHandler) DeleteRestaurant(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.restaurantService.DeleteRestaurant(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurant deleted successfully"})
}


