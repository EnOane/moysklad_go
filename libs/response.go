package libs

import (
	"moysklad/libs/btypes"
)

type MoyskladResponse[T any] struct {
	Meta btypes.Meta `json:"meta"`
	Rows []T         `json:"rows"`
}
