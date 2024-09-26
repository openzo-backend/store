package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/middlewares"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/pb"
	"github.com/tanush-128/openzo_backend/store/internal/service"
	"github.com/tanush-128/openzo_backend/store/internal/utils"
)

type Handler struct {
	storeService service.StoreService
}

func NewHandler(storeService *service.StoreService) *Handler {
	return &Handler{storeService: *storeService}
}

// ParseFormFields extracts and validates store fields from the context
func parseFormFields(ctx *gin.Context) (*models.Store, error) {
	log.Println(ctx.PostForm("opening_time"))
	store := models.Store{

		StorePublic: models.StorePublic{
			Name:    ctx.PostForm("name"),
			Pincode: ctx.PostForm("pincode"),
			Address: ctx.PostForm("address"),

			Phone:       ctx.PostForm("phone"),
			Location:    ctx.PostForm("location"),
			OpeningTime: ctx.PostForm("opening_time"),
			ClosingTime: ctx.PostForm("closing_time"),
			// OpeningTime2:        ctx.PostForm("opening_time2"),
			// ClosingTime2:       ctx.PostForm("closing_time2"),
			SelfDeliveryService: ctx.PostForm("self_delivery_service") == "true",

			StoreType:   ctx.PostForm("store_type"),
			Category:    ctx.PostForm("category"),
			SubCategory: ctx.PostForm("sub_category"),
			Description: ctx.PostForm("description"),

			MetaDescription: ctx.PostForm("meta_description"),
			MetaTags:        ctx.PostForm("meta_tags"),
			Busy:            ctx.PostForm("busy") == "true",
			DeliveryCharge:  utils.StringToInt(ctx.PostForm("delivery_charge")),
			PackagingCharge: utils.StringToInt(ctx.PostForm("packaging_charge")),
			Ranking:         utils.StringToInt(ctx.PostForm("ranking")),
			RestaurantDetails: models.RestaurantDetails{
				PrimaryCuisine:     ctx.PostForm("primary_cuisine"),
				SecondaryCuisine:   ctx.PostForm("secondary_cuisine"),
				AvgPricePerPerson:  utils.StringToInt(ctx.PostForm("avg_price_per_person")),
				PureVeg:            ctx.PostForm("pure_veg") == "true",
				Alcohol:            ctx.PostForm("alcohol") == "true",
				TableCount:         utils.StringToInt(ctx.PostForm("table_count")),
				SeatingCapacity:    utils.StringToInt(ctx.PostForm("seating_capacity")),
				ReserveTableOnline: ctx.PostForm("reserve_table_online") == "true",
			},
		},
		StorePrivate: models.StorePrivate{
			FCMToken:        ctx.PostForm("fcm_token"),
			UserID:          ctx.PostForm("user_id"),
			UserEmail:       ctx.PostForm("user_email"),
			OnlineDiscovery: ctx.PostForm("online_discovery") == "true",
			DetailsComplete: ctx.PostForm("details_complete") == "true",
		},
	}
	// log.Println("Opening Time:", store.OpeningTime2)
	// Validate required fields
	if store.Name == "" || store.UserID == "" || store.Pincode == "" {
		return nil, errors.New("missing required fields: name, user_id, or pincode")
	}

	return &store, nil
}

func (h *Handler) CreateStore(ctx *gin.Context) {
	store, err := parseFormFields(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Phone no:", store.Phone)
	log.Println("User ID:", store.UserID)

	createdStore, err := h.storeService.CreateStore(ctx, *store)
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
	log.Printf("User: %+v", user)

	// Remove private store information if the user does not own the store
	if store.UserID != user.ID && user.Role != pb.Role_ADMIN {
		log.Println("Removing private store information")
		store.StorePrivate = models.StorePrivate{}
	}

	ctx.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoreBasicDetailsByID(ctx *gin.Context) {
	id := ctx.Param("id")
	store, err := h.storeService.GetStoreBasicDetailsByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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

	// Ensure the authenticated user is the same as the store owner
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

func (h *Handler) GetStoresByPincodeAndSubCategory(ctx *gin.Context) {
	pincode := ctx.Param("pincode")
	subCategory := ctx.Param("sub_category")

	stores, err := h.storeService.GetStoresByPincodeAndSubCategory(ctx, pincode, subCategory)
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
	store, err := parseFormFields(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store.ID = ctx.Param("id")

	updatedStore, err := h.storeService.UpdateStore(ctx, *store)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedStore)
}
