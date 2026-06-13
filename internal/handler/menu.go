package handler

import (
	"coffee-project/internal/models"
	"coffee-project/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	menuUsecase usecase.MenuUsecase
}

func NewMenuHandler(u usecase.MenuUsecase) *MenuHandler {
	return &MenuHandler{menuUsecase: u}
}

// GetMenu godoc
// @Summary
// @Description  Возвращает полный список доступных напитков из БД
// @Tags         menu
// @Produce      json
// @Success      200  {array}   models.MenuItem
// @Failure      500  {object}  models.ErrorResponse
// @Router       /menu [get]
func (h *MenuHandler) GetMenu(c *gin.Context) {
	items, err := h.menuUsecase.FetchMenu()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// CreateItem godoc
// @Summary
// @Description
// @Tags         menu
// @Accept       json
// @Produce      json
// @Param        input body      models.CreateMenuItemInput true "New menu item data"
// @Success      201   {object}  models.MenuItem
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /menu [post]
func (h *MenuHandler) CreateItem(c *gin.Context) {
	var input models.CreateMenuItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	newItem, err := h.menuUsecase.AddMenuItem(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

// ErrorResponse используется для возврата ошибок в формате JSON
type ErrorResponse struct {
	Error string `json:"error"`
}
