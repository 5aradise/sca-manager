package cats

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

var ErrEmptyString = errors.New("empty string")

func getIdFromParams(c fiber.Ctx) (int32, error) {
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
