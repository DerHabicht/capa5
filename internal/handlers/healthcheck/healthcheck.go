package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/derhabicht/capa5/internal/config"
)

type ReturnView struct {
	Version string `json:"version"`
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (hc *Handler) Check(ctx *gin.Context) {
	rv := ReturnView{
		Version: config.GetString(config.Version),
	}

	ctx.JSON(http.StatusOK, rv)
}
