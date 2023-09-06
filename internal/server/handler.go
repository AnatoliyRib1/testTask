package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"testTask/contracts"
)

var (
	ErrBlocked = errors.New("blocked")
	maxItems   = 100
	interval   = 2 * time.Second
)

type Handler struct {
	// Здесь можно добавить поля для хранения зависимостей или состояния обработчика.
}

func NewHandler() *Handler {
	return &Handler{}
}

// GetLimits обрабатывает запрос GET /api/items/limits.
func (h *Handler) GetLimits(c echo.Context) error {
	limits := contracts.Limits{
		MaxItems: maxItems,
		Interval: interval,
	}

	return c.JSON(http.StatusOK, limits)
}

// Process обрабатывает запрос POST /api/items/process.
func (h *Handler) Process(c echo.Context) error {
	var request contracts.ProcessRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, contracts.HTTPError{
			Message: "Invalid request body",
		})
	}
	if len(request.Batch.Items) > maxItems {
		return c.JSON(http.StatusTooManyRequests, ErrBlocked)
	}

	// Здесь должна быть обработка запроса и ошибок

	return c.NoContent(http.StatusOK)
}
