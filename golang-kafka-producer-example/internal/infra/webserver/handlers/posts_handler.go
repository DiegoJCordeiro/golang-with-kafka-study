package handlers

import (
	"encoding/json"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/dto"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/usecase"
	"io"
	"net/http"
	"strconv"
)

type PostsHandler struct {
	createPostsUseCase   usecase.CreatePostsUseCase
	queryAllPostsUseCase usecase.QueryAllPostsUseCase
}

func NewPostsHandler(createPostsUseCase usecase.CreatePostsUseCase, queryAllPostsUseCase usecase.QueryAllPostsUseCase) *PostsHandler {
	return &PostsHandler{
		createPostsUseCase:   createPostsUseCase,
		queryAllPostsUseCase: queryAllPostsUseCase,
	}
}

// CreateHandler Create a posts godoc
//
// @Summary     Create a posts
// @Description This endpoint is used to create some posts.
// @Tags        Posts
// @Accept      json
// @Produces    json
// @Param       request      body   	dto.CreatePostsDTO      true      "CreatePostsDTO Request"
// @Success     200 {object} dto.CreatePostsDTO
// @Failure     500         {object}      dto.ErrorDTO
// @Router      /v1/posts  [post]
func (handler *PostsHandler) CreateHandler(response http.ResponseWriter, request *http.Request) {

	var createPostsDTO dto.CreatePostsDTO

	requestBody, err := io.ReadAll(request.Body)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorDto := dto.ErrorDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		_ = json.NewEncoder(response).Encode(errorDto)
		return
	}

	err = json.Unmarshal(requestBody, &createPostsDTO)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errorDto := dto.ErrorDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		_ = json.NewEncoder(response).Encode(errorDto)
		return
	}

	err = handler.createPostsUseCase.Execute(createPostsDTO.Messages)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorDto := dto.ErrorDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		_ = json.NewEncoder(response).Encode(errorDto)
		return
	}

	response.WriteHeader(http.StatusOK)
}

// QueryAllHandler Query all posts godoc
//
// @Summary     Query all posts
// @Description This endpoint is used to query all posts.
// @Tags        Posts
// @Accept      json
// @Produces    json
// @Param       limit    query     int64  false  "limit of data"
// @Param       offset   query     int64  false  "offset of data"
// @Success     200 {object} dto.QueryPostsDTO
// @Failure     500         {object}      dto.ErrorDTO
// @Router      /v1/posts  [get]
func (handler *PostsHandler) QueryAllHandler(response http.ResponseWriter, request *http.Request) {

	limit := request.URL.Query().Get("limit")
	offset := request.URL.Query().Get("offset")
	limitConv, _ := strconv.ParseInt(limit, 10, 64)
	offsetConv, _ := strconv.ParseInt(offset, 10, 64)

	posts, err := handler.queryAllPostsUseCase.Execute(limitConv, offsetConv)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorDto := dto.ErrorDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		_ = json.NewEncoder(response).Encode(errorDto)
		return
	}

	err = json.NewEncoder(response).Encode(posts)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorDto := dto.ErrorDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		_ = json.NewEncoder(response).Encode(errorDto)
		return
	}

	response.WriteHeader(http.StatusOK)
}
