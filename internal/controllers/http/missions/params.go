package missions

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

var ErrEmptyString = errors.New("empty string")

func getMissionIdFromParams(c fiber.Ctx) (int32, error) {
	strID := c.Params("id")
	if strID == "" {
		return 0, fmt.Errorf("invalid id param: %w", ErrEmptyString)
	}
	id, err := strconv.ParseInt(strID, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid id param: %w", err)
	}
	return int32(id), nil
}

func getTargetIdxFromParams(c fiber.Ctx) (int, error) {
	strIDX := c.Params("tidx")
	if strIDX == "" {
		return 0, fmt.Errorf("invalid target idx param: %w", ErrEmptyString)
	}
	idx, err := strconv.Atoi(strIDX)
	if err != nil {
		return 0, fmt.Errorf("invalid target idx param: %w", err)
	}
	return idx - 1, nil
}
