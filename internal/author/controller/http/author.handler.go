package http

import (
	"hta-platform/internal/author/application/service"
	"hta-platform/internal/author/controller/dto"
	"hta-platform/pkg/response"
	"hta-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthorHandler struct {
	authorService service.AuthorService
}

func NewAuthorHandler(authorService service.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}

func (a *AuthorHandler) CreateAuthor(c *gin.Context) (res interface{}, err error) {
	var author dto.AuthorReq

	err = c.ShouldBindJSON(&author)
	if err != nil {
		return nil, err
	}
	author.Normalize()

	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&author, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	newAuthor, err := a.authorService.CreateAuthor(c, &author)
	if err != nil {
		return nil, err
	}

	return newAuthor, nil
}

func (a *AuthorHandler) GetAuthorByUrl(c *gin.Context) (res interface{}, err error) {
	url := c.Param("url")
	author, err := a.authorService.GetAuthorByUrl(c, url)

	if err != nil {
		return nil, err
	}

	return author, nil
}
