package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/service"
)

type TableHandler struct {
	tableService service.TableService
}

func NewTableHandler(tableService *service.TableService) *TableHandler {
	return &TableHandler{tableService: *tableService}
}

func (h *TableHandler) CreateTable(ctx *gin.Context) {

	var table models.ResTable

	ctx.ShouldBindJSON(&table)

	createdTable, err := h.tableService.CreateTable(ctx, table)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdTable)

}

func (h *TableHandler) GetTablesByStoreID(ctx *gin.Context) {
	id := ctx.Param("store_id")

	tables, err := h.tableService.GetTablesByStoreID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tables)
}

func (h *TableHandler) UpdateTable(ctx *gin.Context) {
	var table models.ResTable

	ctx.ShouldBindJSON(&table)

	updatedTable, err := h.tableService.UpdateTable(ctx, table)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedTable)
}

func (h *TableHandler) DeleteTable(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.tableService.DeleteTable(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Table deleted successfully"})
}
