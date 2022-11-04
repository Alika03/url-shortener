package command

import (
	"shortener/domain/shortener/model"
	"shortener/utils"
	"time"
)

type Link struct {
	Link *model.Link
}

func NewLink(url string) Link {
	t := time.Now().Add(2 * time.Hour).UTC()
	return Link{
		Link: &model.Link{
			Id:        utils.GenerateUuid(),
			Code:      utils.GetUniqueCode(),
			FullUrl:   url,
			ExpiredAt: t,
		},
	}
}
