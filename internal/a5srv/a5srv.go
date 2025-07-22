package a5srv

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/derhabicht/capa5/internal/handlers/healthcheck"
	"github.com/derhabicht/capa5/internal/router"
)

func Setup() (*router.Router, error) {
	r := router.NewRouter("/api/v1")

	hc := healthcheck.NewHandler()
	r.RegisterHandler(http.MethodGet, "/healthcheck", hc.Check)

	return r, nil
}

func Run(r *router.Router) error {
	err := r.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
