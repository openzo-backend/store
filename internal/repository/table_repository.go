package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"gorm.io/gorm"
)

type TableRepository interface {
	CreateTable(Table models.ResTable) (models.ResTable, error)
	GetTablesByStoreID(id string) ([]models.ResTable, error)
	UpdateTable(Table models.ResTable) (models.ResTable, error)
	DeleteTable(id string) error
}

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {

	return &tableRepository{db: db}
}

func (r *tableRepository) CreateTable(Table models.ResTable) (models.ResTable, error) {
	log.Printf("CreateTable called with Table: %+v", Table)
	Table.ID = uuid.New().String()
	tx := r.db.Create(&Table)

	if tx.Error != nil {
		return models.ResTable{}, tx.Error
	}

	return Table, nil
}

func (r *tableRepository) GetTablesByStoreID(id string) ([]models.ResTable, error) {
	log.Println("GetTablesByStoreID called with id: ", id)
	var resTables []models.ResTable
	tx := r.db.Where("store_id = ?", id).Find(&resTables)
	// tx := r.db.Find(&resTables)

	if tx.Error != nil {
		return []models.ResTable{}, tx.Error
	}
	return resTables, nil
}

func (r *tableRepository) UpdateTable(Table models.ResTable) (models.ResTable, error) {
	tx := r.db.Save(&Table)

	if tx.Error != nil {
		return models.ResTable{}, tx.Error
	}

	return Table, nil
}

func (r *tableRepository) DeleteTable(id string) error {
	tx := r.db.Where("id = ?", id).Delete(&models.ResTable{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
