package add

import (
	"context"
	"net/http"
	"shortener/config"
	"shortener/domain/shortener/command"
	"shortener/domain/shortener/command/add"
	reddis "shortener/infrastructure/cache/redis"
	"shortener/infrastructure/db/postgres"
	"shortener/server/http/helper"
	"time"
)

// Handle @Summary Add
// @Tags add
// @Description generate url
// @ID add-url
// @Accept json
// @Produce json
// @Param input body DTO true "url info"
// @Success 200
// @Failure 400
// @Failure 409
// @Router /add [post]
func Handle(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{}

	if err := helper.BindRequest(request, dto); err != nil {
		helper.ErrorResponse(err, response, http.StatusUnprocessableEntity)
		return
	}

	if err := dto.Validate(); err != nil {
		helper.ErrorResponse(err, response, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	code, err := add.NewHandler(
		command.NewUrlRepository(postgres.GetDbConnection()),
		command.NewCacheRepository(reddis.GetDbConnection()),
	).Handle(ctx, dto.Url)
	if err != nil {
		helper.ErrorResponse(err, response, http.StatusConflict)
		return
	}

	resp := map[string]string{
		"short_url": config.GetConfig().ServerParam.Host + ":" + config.GetConfig().ServerParam.Port + "/" + code,
	}
	helper.JsonResponse(response, resp)
}
