package missions

import (
	"github.com/5aradise/sca-manager/internal/models"
	"github.com/gofiber/fiber/v3"
)

func (h *handler) AddTarget(c fiber.Ctx) error {
	id, err := getMissionIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var req models.Target
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := h.s.AddTarget(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

type updateTargetRequest struct {
	IsCompleted bool    `json:"is_completed"`
	Notes       *string `json:"notes"`
}

func (h *handler) UpdateTarget(c fiber.Ctx) error {
	missionID, err := getMissionIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	targetIDX, err := getTargetIdxFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var req updateTargetRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var res models.Mission
	if req.Notes != nil {
		res, err = h.s.UpdateTargetNotes(c.Context(), missionID, targetIDX, *req.Notes)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	if req.IsCompleted {
		res, err = h.s.MarkTargetAsCompleted(c.Context(), missionID, targetIDX)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	if res.ID == 0 {
		res, err = h.s.GetMission(c.Context(), missionID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *handler) DeleteTarget(c fiber.Ctx) error {
	missionID, err := getMissionIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	targetIDX, err := getTargetIdxFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := h.s.DeleteTarget(c.Context(), missionID, targetIDX)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
