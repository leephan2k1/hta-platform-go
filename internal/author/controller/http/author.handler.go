package http

import (
	"hta-platform/internal/author/application/service"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorService service.AuthorService
}

func NewAuthorHandler(authorService service.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}

func (a *AuthorHandler) CreateAuthor(c *gin.Context) (res interface{}, err error) {
	return nil, nil
}

func (a *AuthorHandler) GetAuthorByUrl(c *gin.Context) (res interface{}, err error) {
	return nil, nil
}
