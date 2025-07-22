package router

import (
	"github.com/gin-gonic/gin"
)

// RestHandler is a collection of *gin.HandlerFunc functions that implement CRUD for a resource.
type RestHandler interface {
	// Group returns the Gin RouterGroup that handler functions should be registered under.
	Group(r *gin.RouterGroup) *gin.RouterGroup

	// Options implements the HTTP OPTIONS verb for this handler.
	Options(ctx *gin.Context)

	// Head implements the HTTP HEAD verb for this handler.
	Head(ctx *gin.Context)

	// Create implements resource creation. This will be registered with the Router using the POST verb.
	Create(ctx *gin.Context)

	// List implements fetching multiple resources. This will be registered with the Router using the GET verb.
	List(ctx *gin.Context)

	// Fetch implements fetching a single resource. This will be registered with the Router using the GET verb.
	Fetch(ctx *gin.Context)

	// Replace implements idempotent updating of a resource.
	// This will be registered with the Router using the PUT verb.
	Replace(ctx *gin.Context)

	// Update implements partial updating of a resource.
	// This will be registered with the Router using the PATCH verb.
	Update(ctx *gin.Context)

	// Delete implements destruction of a resource. This will be registered with the Router using the DELETE verb.
	Delete(ctx *gin.Context)
}
