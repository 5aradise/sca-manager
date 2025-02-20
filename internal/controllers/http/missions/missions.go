package missions

import (
	"github.com/5aradise/sca-manager/internal/models"
	"github.com/gofiber/fiber/v3"
)

func (h *handler) CreateMission(c fiber.Ctx) error {
	var req models.Mission
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := h.s.CreateMission(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *handler) GetMission(c fiber.Ctx) error {
	id, err := getMissionIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := h.s.GetMission(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *handler) ListMissions(c fiber.Ctx) error {
	res, err := h.s.ListMissions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if res == nil {
		res = []models.Mission{}
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

type updateMissionRequest struct {
	IsCompleted bool   `json:"is_completed"`
	CatID       *int32 `json:"cat_id"`
}

func (h *handler) UpdateMission(c fiber.Ctx) error {
	id, err := getMissionIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var req updateMissionRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var res models.Mission
	if req.CatID != nil {
		res, err = h.s.AssignCat(c.Context(), id, *req.CatID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	if req.IsCompleted {
		res, err = h.s.MarkMissionAsCompleted(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	if res.ID == 0 {
		res, err = h.s.GetMission(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *handler) DeleteMission(c fiber.Ctx) error {
	id, err := getMissionIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = h.s.DeleteMission(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
