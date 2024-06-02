package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/repository"
)

type RestaurantService interface {
	CreateRestaurantDetails(ctx *gin.Context, req models.RestaurantDetails) (models.RestaurantDetails, error)
	UpdateRestaurantDetails(ctx *gin.Context, req models.RestaurantDetails) (models.RestaurantDetails, error)
	GetRestaurantByID(ctx *gin.Context, id string) (models.Restaurant, error)
	GetRestaurantDetailsByStoreID(ctx *gin.Context, id string) (models.RestaurantDetails, error)
	GetRestaurantsByUserID(ctx *gin.Context, id string) ([]models.Restaurant, error)
	DeleteRestaurant(ctx *gin.Context, id string) error
}

type restaurantService struct {
	restaurantRepository repository.RestaurantRepository
}

func NewRestaurantService(restaurantRepository repository.RestaurantRepository) RestaurantService {
	return &restaurantService{
		restaurantRepository: restaurantRepository,
	}
}

func (s *restaurantService) CreateRestaurantDetails(ctx *gin.Context, req models.RestaurantDetails) (models.RestaurantDetails, error) {
	createdRestaurant, err := s.restaurantRepository.CreateRestaurantDetails(req)
	if err != nil {
		return models.RestaurantDetails{}, err
	}
	return createdRestaurant, nil
}

func (s *restaurantService) UpdateRestaurantDetails(ctx *gin.Context, req models.RestaurantDetails) (models.RestaurantDetails, error) {
	updatedRestaurant, err := s.restaurantRepository.UpdateRestaurantDetails(req)
	if err != nil {
		return models.RestaurantDetails{}, err
	}
	return updatedRestaurant, nil
}

func (s *restaurantService) GetRestaurantByID(ctx *gin.Context, id string) (models.Restaurant, error) {
	restaurant, err := s.restaurantRepository.GetRestaurantByID(id)
	if err != nil {
		return models.Restaurant{}, err
	}
	return restaurant, nil
}

func (s *restaurantService) GetRestaurantDetailsByStoreID(ctx *gin.Context, id string) (models.RestaurantDetails, error) {
	restaurantDetails, err := s.restaurantRepository.GetRestaurantDetailsByStoreID(id)
	if err != nil {
		return models.RestaurantDetails{}, err
	}
	return restaurantDetails, nil
}

func (s *restaurantService) GetRestaurantsByUserID(ctx *gin.Context, id string) ([]models.Restaurant, error) {
	restaurants, err := s.restaurantRepository.GetRestaurantsByUserID(id)
	if err != nil {
		return []models.Restaurant{}, err
	}
	return restaurants, nil
}

func (s *restaurantService) DeleteRestaurant(ctx *gin.Context, id string) error {
	err := s.restaurantRepository.DeleteRestaurant(id)
	if err != nil {
		return err
	}
	return nil
}
