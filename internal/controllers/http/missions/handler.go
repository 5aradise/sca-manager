package missions

import (
	"context"

	"github.com/5aradise/sca-manager/internal/models"
	"github.com/gofiber/fiber/v3"
)

type (
	missionService interface {
		AddTarget(ctx context.Context, missionID int32, target models.Target) (models.Mission, error)
		AssignCat(ctx context.Context, missionID int32, catID int32) (models.Mission, error)
		CreateMission(ctx context.Context, mission models.Mission) (models.Mission, error)
		DeleteMission(ctx context.Context, id int32) error
		DeleteTarget(ctx context.Context, missionID int32, targetNum int) (models.Mission, error)
		GetMission(ctx context.Context, id int32) (models.Mission, error)
		ListMissions(ctx context.Context) ([]models.Mission, error)
		MarkMissionAsCompleted(ctx context.Context, id int32) (models.Mission, error)
		MarkTargetAsCompleted(ctx context.Context, missionID int32, targetNum int) (models.Mission, error)
		UpdateTargetNotes(ctx context.Context, missionID int32, targetNum int, notes string) (models.Mission, error)
	}

	handler struct {
		s missionService
	}
)

func New(s missionService) *handler {
	return &handler{s}
}

func (h *handler) Init(r fiber.Router) {
	r.Post("", h.CreateMission)
	r.Get("", h.ListMissions)
	r.Get("/:id", h.GetMission)
	r.Patch("/:id", h.UpdateMission)
	r.Delete("/:id", h.DeleteMission)

	targets := r.Group("/:id/targets")
	targets.Post("", h.AddTarget)
	targets.Patch("/:tidx", h.UpdateTarget)
	targets.Delete("/:tidx", h.DeleteTarget)
}
