package reqvalidator

import (
	"github.com/jumayevgadam/music-app/pkg/errlst"
	"github.com/labstack/echo/v4"
	"log"
)

// ReadRequest body and validate
func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(&request); err != nil {
		log.Print("err: [reqvalidator][ReadRequest]")
		return errlst.ParseErrors(err)
	}

	return validate.StructCtx(ctx.Request().Context(), request)
}
