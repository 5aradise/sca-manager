package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Money int32

func NewMoney(f float64) Money {
	return Money((f * 100) + 0.5)
}

func NewMoneyFromCents(c int32) Money {
	return Money(c)
}

func (m Money) Cents() int32 {
	return int32(m)
}

func (m Money) Float64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

func (m Money) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f", x)
}

func (m Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Float64())
}

var ErrTooManyDecPlcs = errors.New("invalid Money string: too many decimal places")

func (m *Money) UnmarshalJSON(data []byte) error {
	dpi := bytes.LastIndexByte(data, '.')
	if dpi == -1 {
		i, err := strconv.ParseInt(string(data), 10, 32)
		if err != nil {
			return err
		}

		*m = Money(i * 100)
		return nil
	}

	if len(data)-dpi > 3 {
		return ErrTooManyDecPlcs
	}

	f, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}

	*m = NewMoney(f)
	return nil
}
