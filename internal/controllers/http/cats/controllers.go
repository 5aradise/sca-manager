package cats

import (
	"github.com/5aradise/sca-manager/internal/models"
	"github.com/gofiber/fiber/v3"
)

func (h *handler) CreateCat(c fiber.Ctx) error {
	var req models.Cat
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := h.s.CreateCat(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *handler) ListCats(c fiber.Ctx) error {
	res, err := h.s.ListCats(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if res == nil {
		res = []models.Cat{}
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *handler) GetCat(c fiber.Ctx) error {
	id, err := getIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := h.s.GetCat(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

type updateCatSalaryRequest struct {
	Salary *models.Money
}

func (h *handler) UpdateCatSalary(c fiber.Ctx) error {
	id, err := getIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var req updateCatSalaryRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var res models.Cat
	if req.Salary != nil {
		res, err = h.s.UpdateCatSalary(c.Context(), id, *req.Salary)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	if res.ID == 0 {
		res, err = h.s.GetCat(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *handler) DeleteCat(c fiber.Ctx) error {
	id, err := getIdFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = h.s.DeleteCat(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
