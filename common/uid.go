package common

import (
	"encoding/base64"

	"github.com/google/uuid"
)

func EncodeUID(id uuid.UUID) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(id.String()))

	return encoded
}

func DecodeUID(encoded string) (uuid.UUID, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		return uuid.Nil, err
	}

	if id, err := uuid.Parse(string(decoded)); err != nil {
		return uuid.Nil, err
	} else {
		return id, nil
	}
}
