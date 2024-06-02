package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/repository"
)

type TableService interface {
	CreateTable(ctx *gin.Context, req models.ResTable) (models.ResTable, error)
	GetTablesByStoreID(ctx *gin.Context, id string) ([]models.ResTable, error)
	UpdateTable(ctx *gin.Context, req models.ResTable) (models.ResTable, error)
	DeleteTable(ctx *gin.Context, id string) error
}

type tableService struct {
	tableRepository repository.TableRepository
}

func NewTableService(tableRepository repository.TableRepository) TableService {
	return &tableService{
		tableRepository: tableRepository,
	}
}

func (s *tableService) CreateTable(ctx *gin.Context, req models.ResTable) (models.ResTable, error) {
	createdTable, err := s.tableRepository.CreateTable(req)
	if err != nil {
		return models.ResTable{}, err
	}
	return createdTable, nil
}

func (s *tableService) GetTablesByStoreID(ctx *gin.Context, id string) ([]models.ResTable, error) {
	tables, err := s.tableRepository.GetTablesByStoreID(id)
	if err != nil {
		return []models.ResTable{}, err
	}
	return tables, nil
}

func (s *tableService) UpdateTable(ctx *gin.Context, req models.ResTable) (models.ResTable, error) {
	updatedTable, err := s.tableRepository.UpdateTable(req)
	if err != nil {
		return models.ResTable{}, err
	}
	return updatedTable, nil
}

func (s *tableService) DeleteTable(ctx *gin.Context, id string) error {
	err := s.tableRepository.DeleteTable(id)
	if err != nil {
		return err
	}
	return nil
}
