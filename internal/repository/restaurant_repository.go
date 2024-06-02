                                                                                                                                                                                                                                                                    
package repository

import (
	"github.com/google/uuid"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"gorm.io/gorm"
)

type RestaurantRepository interface {
	CreateRestaurantDetails(Restaurant models.RestaurantDetails) (models.RestaurantDetails, error)
	UpdateRestaurantDetails(Restaurant models.RestaurantDetails) (models.RestaurantDetails, error)

	GetRestaurantByID(id string) (models.Restaurant, error)
	GetRestaurantDetailsByStoreID(id string) (models.RestaurantDetails, error)

	GetRestaurantsByUserID(id string) ([]models.Restaurant, error)
	DeleteRestaurant(id string) error

	// Add more methods for other Stor e operations (GetStoreByEmail, UpdateStore, etc.)

}

type restaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {

	return &restaurantRepository{db: db}
}

func (r *restaurantRepository) CreateRestaurantDetails(Restaurant models.RestaurantDetails) (models.RestaurantDetails, error) {
	Restaurant.ID = uuid.New().String()
	tx := r.db.Create(&Restaurant)

	if tx.Error != nil {
		return models.RestaurantDetails{}, tx.Error
	}

	return Restaurant, nil
}

func (r *restaurantRepository) UpdateRestaurantDetails(Restaurant models.RestaurantDetails) (models.RestaurantDetails, error) {
	tx := r.db.Save(&Restaurant)

	if tx.Error != nil {
		return models.RestaurantDetails{}, tx.Error
	}

	return Restaurant, nil
}

func (r *restaurantRepository) GetRestaurantByID(id string) (models.Restaurant, error) {
	var Restaurant models.Restaurant
	var RestaurantDetails models.RestaurantDetails
	var Store models.Store

	tx := r.db.Where("id = ?", id).First(&Store)
	if tx.Error != nil {
		return models.Restaurant{}, tx.Error
	}
	tx = r.db.Where("store_id = ?", id).First(&RestaurantDetails)
	if tx.Error != nil {
		return models.Restaurant{}, tx.Error
	}
	Restaurant.Store = Store
	Restaurant.RestaurantDetails = RestaurantDetails
	Restaurant.ID = Store.ID
	Restaurant.Store.StorePrivate = models.StorePrivate{}

	return Restaurant, nil
}

func (r *restaurantRepository) GetRestaurantDetailsByStoreID(id string) (models.RestaurantDetails, error) {
	var RestaurantDetails models.RestaurantDetails
	tx := r.db.Where("store_id = ?", id).First(&RestaurantDetails)
	if tx.Error != nil {
		return models.RestaurantDetails{}, tx.Error
	}

	return RestaurantDetails, nil
}

func (r *restaurantRepository) GetRestaurantsByUserID(id string) ([]models.Restaurant, error) {
	var Restaurants []models.Restaurant
	tx := r.db.Where("user_id = ?", id).Find(&Restaurants)
	if tx.Error != nil {
		return []models.Restaurant{}, tx.Error
	}
	return Restaurants, nil
}

func (r *restaurantRepository) DeleteRestaurant(id string) error {
	tx := r.db.Delete(&models.Restaurant{}, id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}


