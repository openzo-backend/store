package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/service"
)

type ReviewHandler struct {
	reviewService service.ReviewService
}

func NewReviewHandler(reviewService *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService: *reviewService}
}

func (h *ReviewHandler) CreateReview(ctx *gin.Context) {

	var review models.Review

	ctx.ShouldBindJSON(&review)

	createdReview, err := h.reviewService.CreateReview(ctx, review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdReview)

}

func (h *ReviewHandler) GetReviewByID(ctx *gin.Context) {
	id := ctx.Param("id")

	review, err := h.reviewService.GetReviewByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) GetReviewsByStoreID(ctx *gin.Context) {
	id := ctx.Param("store_id")

	reviews, err := h.reviewService.GetReviewsByStoreID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) GetReviewsByUserID(ctx *gin.Context) {
	id := ctx.Param("user_id")

	reviews, err := h.reviewService.GetReviewsByUserID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) UpdateReview(ctx *gin.Context) {
	var review models.Review

	ctx.ShouldBindJSON(&review)

	updatedReview, err := h.reviewService.UpdateReview(ctx, review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedReview)
}

func (h *ReviewHandler) DeleteReview(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.reviewService.DeleteReview(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
