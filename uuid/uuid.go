package uuid

import (
	"github.com/google/uuid"
	"github.com/stillwondering/minar"
)

func GenerateID() minar.MinutesID {
	uuid := uuid.New()

	return minar.MinutesID(uuid.String())
}
