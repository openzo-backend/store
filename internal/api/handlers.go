package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/middlewares"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/service"
)

type Handler struct {
	storeService service.StoreService
}

func NewHandler(storeService *service.StoreService) *Handler {
	return &Handler{storeService: *storeService}
}

func (h *Handler) CreateStore(ctx *gin.Context) {
	var store models.Store

	store.Name = ctx.PostForm("name")
	store.Pincode = ctx.PostForm("pincode")
	store.UserEmail = ctx.PostForm("user_email")
	store.UserID = ctx.PostForm("user_id")
	store.Address = ctx.PostForm("address")
	store.Phone = ctx.PostForm("phone")
	store.Location = ctx.PostForm("location")
	store.OpeningTime = ctx.PostForm("opening_time")
	store.ClosingTime = ctx.PostForm("closing_time")
	store.OnlineDiscovery = ctx.PostForm("online_discovery") == "true"
	store.SelfDeliveryService = ctx.PostForm("self_delivery_service") == "true"
	store.StoreType = models.StoreType(ctx.PostForm("store_type"))
	store.Description = ctx.PostForm("description")
	store.FCMToken = ctx.PostForm("fcm_token")

	log.Println("Phone no", store.Phone)

	createdStore, err := h.storeService.CreateStore(ctx, store)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdStore)

}

func (h *Handler) GetStoreByID(ctx *gin.Context) {
	id := ctx.Param("id")

	store, err := h.storeService.GetStoreByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := ctx.MustGet("user").(middlewares.User)
	log.Printf("User : %+v", user)

	if store.UserID != user.ID {
		store.StorePrivate = models.StorePrivate{}
	}

	ctx.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoreByPhoneNo(ctx *gin.Context) {
	phoneNo := ctx.Param("phone_no")

	store, err := h.storeService.GetStoreByPhoneNo(ctx, phoneNo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoreByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	user := ctx.MustGet("user").(middlewares.User)
	if user.ID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	store, err := h.storeService.GetStoreByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoresByPincode(ctx *gin.Context) {
	pincode := ctx.Param("pincode")

	stores, err := h.storeService.GetStoresByPincode(ctx, pincode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stores)
}

func (h *Handler) GetStoresByPincodeAndCategory(ctx *gin.Context) {
	pincode := ctx.Param("pincode")
	category := ctx.Param("category")

	stores, err := h.storeService.GetStoresByPincodeAndCategory(ctx, pincode, category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stores)
}

func (h *Handler) GetCategories(ctx *gin.Context) {
	categories, err := h.storeService.GetCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (h *Handler) UpdateStore(ctx *gin.Context) {
	var store models.Store
	store.ID = ctx.Param("id")
	store.Name = ctx.PostForm("name")
	store.Pincode = ctx.PostForm("pincode")
	store.UserEmail = ctx.PostForm("user_email")
	store.Address = ctx.PostForm("address")
	store.Phone = ctx.PostForm("phone")
	store.Location = ctx.PostForm("location")
	store.OpeningTime = ctx.PostForm("opening_time")
	store.ClosingTime = ctx.PostForm("closing_time")
	store.OnlineDiscovery = ctx.PostForm("online_discovery") == "true"
	store.SelfDeliveryService = ctx.PostForm("self_delivery_service") == "true"
	store.StoreType = models.StoreType(ctx.PostForm("store_type"))
	store.Description = ctx.PostForm("description")
	store.FCMToken = ctx.PostForm("fcm_token")

	updatedStore, err := h.storeService.UpdateStore(ctx, store)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedStore)
}
