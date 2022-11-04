package get

import "errors"

type DTO struct {
	Code string `json:"code"`
}

func (dto *DTO) Validate() error {
	if dto.Code == "" {
		return errors.New("code is empty")
	}

	if len(dto.Code) != 10 {
		return errors.New("code must be no less or more than 10 chars")
	}

	return nil
}
