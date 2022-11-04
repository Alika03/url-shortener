package get

import (
	"context"
	"errors"
	"net/http"
	"shortener/domain/shortener/command"
	"shortener/domain/shortener/command/get"
	reddis "shortener/infrastructure/cache/redis"
	"shortener/infrastructure/db/postgres"
	"shortener/server/http/helper"
	"time"
)

// Handle @Summary Get
// @Tags get
// @Description get url by code
// @ID get-url
// @Accept json
// @Produce json
// @Param code path string true "url info"
// @Success 200
// @Failure 400
// @Failure 409
// @Router /{code} [get]
func Handle(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{Code: helper.GetRequestParam(request)}

	if err := dto.Validate(); err != nil {
		helper.ErrorResponse(err, response, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	model, err := get.NewHandler(
		command.NewUrlRepository(postgres.GetDbConnection()),
		command.NewCacheRepository(reddis.GetDbConnection()),
	).Handle(ctx, dto.Code)
	if err != nil {
		helper.ErrorResponse(err, response, http.StatusConflict)
		return
	}

	if model.FullUrl == "" {
		helper.ErrorResponse(errors.New("url not found"), response, http.StatusNotFound)
		return
	}

	http.Redirect(response, request, model.FullUrl, http.StatusMovedPermanently)
}
