package add

import "errors"

type DTO struct {
	Url string `json:"url"`
}

func (dto *DTO) Validate() error {
	if dto.Url == "" {
		return errors.New("url is empty")
	}
	return nil
}
