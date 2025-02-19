package cats

import (
	"context"

	"github.com/5aradise/sca-manager/internal/models"
	"github.com/gofiber/fiber/v3"
)

type (
	catService interface {
		CreateCat(context.Context, models.Cat) (models.Cat, error)
		GetCat(context.Context, int32) (models.Cat, error)
		ListCats(context.Context) ([]models.Cat, error)
		UpdateCatSalary(context.Context, int32, models.Money) (models.Cat, error)
		DeleteCat(context.Context, int32) error
	}

	handler struct {
		s catService
	}
)

func New(s catService) *handler {
	return &handler{s}
}

func (h *handler) Init(r fiber.Router) {
	r.Post("", h.CreateCat)
	r.Get("", h.ListCats)
	r.Get("/:id", h.GetCat)
	r.Patch("/:id", h.UpdateCatSalary)
	r.Delete("/:id", h.DeleteCat)
}
