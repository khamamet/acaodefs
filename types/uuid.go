package types

import (
	"fmt"
	"io"

	"github.com/moroz/uuidv7-go"
)

type UUID struct {
	uuidv7.UUID
}

func NewUUID() UUID {
	return UUID{uuidv7.Generate()}
}

func ParseOrZeroUUID(s string) UUID {
	parsed, err := uuidv7.Parse(s)
	if err != nil {
		return ZeroUUID()
	}
	return UUID{parsed}
}

func ParseUUID(s string) (UUID, error) {
	parsed, err := uuidv7.Parse(s)
	if err != nil {
		return UUID{}, fmt.Errorf("invalid UUID format: %w", err)
	}
	return UUID{parsed}, nil
}
func ZeroUUID() UUID {
	uuid, err := uuidv7.Parse("00000000-0000-0000-0000-000000000000")
	if err == nil {
		return UUID{uuid}
	}
	return NewUUID()
}

func (u UUID) String() string {
	return u.UUID.String()
}

func (u *UUID) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UUID must be a string")
	}
	parsed, err := uuidv7.Parse(str)
	if err != nil {
		return fmt.Errorf("invalid UUID format: %w", err)
	}
	u.UUID = parsed
	return nil
}

func (u UUID) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%s"`, u.UUID.String())
}
