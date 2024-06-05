package id

import (
	"fmt"

	"github.com/google/uuid"
)

type Newsletter uuid.UUID

func (u *Newsletter) FromString(s string) error {
	id, err := uuid.Parse(s)
	if err != nil {
		return err
	}

	*u = Newsletter(id)
	return nil
}

func (u Newsletter) String() string {
	return uuid.UUID(u).String()
}

func (u *Newsletter) Scan(data any) error {
	return scanUUID((*uuid.UUID)(u), "Newsletter", data)
}

func (u Newsletter) MarshalText() ([]byte, error) {
	return []byte(uuid.UUID(u).String()), nil
}

func (u *Newsletter) UnmarshalText(data []byte) error {
	return unmarshalUUID((*uuid.UUID)(u), "Newsletter", data)
}

func scanUUID(u *uuid.UUID, idTypeName string, data any) error {
	if err := u.Scan(data); err != nil {
		return fmt.Errorf("scanning %q id value: %w", idTypeName, err)
	}
	return nil
}

func unmarshalUUID(u *uuid.UUID, idTypeName string, data []byte) error {
	if err := u.UnmarshalText(data); err != nil {
		return fmt.Errorf("parsing %q id value: %w", idTypeName, err)
	}
	return nil
}

func (u Newsletter) IsEmpty() bool {
	return uuid.UUID(u) == uuid.Nil
}
