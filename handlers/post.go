package handlers

import (
	"context"
	"fmt"
	"net/http"

	"gofiber-blog/database"
	"gofiber-blog/models"

	"github.com/danielgtaylor/huma/v2"
)

// --- I/O types ---

type PostBody struct {
	Title  string `json:"title" doc:"Post title" minLength:"3"`
	Body   string `json:"body" doc:"Post body" minLength:"1"`
	Author string `json:"author" doc:"Post author" minLength:"1"`
}

type UpdateBody struct {
	Title  string `json:"title,omitempty" doc:"Post title"`
	Body   string `json:"body,omitempty" doc:"Post body"`
	Author string `json:"author,omitempty" doc:"Post author"`
}

type PostIDInput struct {
	ID uint `path:"id" doc:"Post ID"`
}

type CreatePostInput struct {
	Body PostBody
}

type UpdatePostInput struct {
	ID   uint `path:"id" doc:"Post ID"`
	Body UpdateBody
}

type PostOutput struct {
	Body models.Post
}

type PostListOutput struct {
	Body struct {
		Data  []models.Post `json:"data"`
		Count int           `json:"count"`
	}
}

type MessageOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

// --- Handlers ---

func GetPosts(_ context.Context, _ *struct{}) (*PostListOutput, error) {
	var posts []models.Post
	database.DB.Order("created_at desc").Find(&posts)

	out := &PostListOutput{}
	out.Body.Data = posts
	out.Body.Count = len(posts)
	return out, nil
}

func GetPost(_ context.Context, input *PostIDInput) (*PostOutput, error) {
	var post models.Post
	if result := database.DB.First(&post, input.ID); result.Error != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("post %d not found", input.ID))
	}
	return &PostOutput{Body: post}, nil
}

func CreatePost(_ context.Context, input *CreatePostInput) (*PostOutput, error) {
	post := models.Post{
		Title:  input.Body.Title,
		Body:   input.Body.Body,
		Author: input.Body.Author,
	}
	if result := database.DB.Create(&post); result.Error != nil {
		return nil, huma.Error500InternalServerError("failed to create post")
	}
	return &PostOutput{Body: post}, nil
}

func UpdatePost(_ context.Context, input *UpdatePostInput) (*PostOutput, error) {
	var post models.Post
	if result := database.DB.First(&post, input.ID); result.Error != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("post %d not found", input.ID))
	}

	updates := map[string]interface{}{}
	if input.Body.Title != "" {
		updates["title"] = input.Body.Title
	}
	if input.Body.Body != "" {
		updates["body"] = input.Body.Body
	}
	if input.Body.Author != "" {
		updates["author"] = input.Body.Author
	}
	database.DB.Model(&post).Updates(updates)

	return &PostOutput{Body: post}, nil
}

func DeletePost(_ context.Context, input *PostIDInput) (*MessageOutput, error) {
	var post models.Post
	if result := database.DB.First(&post, input.ID); result.Error != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("post %d not found", input.ID))
	}
	database.DB.Delete(&post)

	out := &MessageOutput{}
	out.Body.Message = fmt.Sprintf("post %d deleted", input.ID)
	return out, nil
}

// --- Route registration ---

func RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "list-posts",
		Method:      http.MethodGet,
		Path:        "/posts",
		Summary:     "List all posts",
		Tags:        []string{"posts"},
	}, GetPosts)

	huma.Register(api, huma.Operation{
		OperationID: "get-post",
		Method:      http.MethodGet,
		Path:        "/posts/{id}",
		Summary:     "Get a post by ID",
		Tags:        []string{"posts"},
	}, GetPost)

	huma.Register(api, huma.Operation{
		OperationID:  "create-post",
		Method:       http.MethodPost,
		Path:         "/posts",
		Summary:      "Create a new post",
		Tags:         []string{"posts"},
		DefaultStatus: http.StatusCreated,
	}, CreatePost)

	huma.Register(api, huma.Operation{
		OperationID: "update-post",
		Method:      http.MethodPut,
		Path:        "/posts/{id}",
		Summary:     "Update a post",
		Tags:        []string{"posts"},
	}, UpdatePost)

	huma.Register(api, huma.Operation{
		OperationID: "delete-post",
		Method:      http.MethodDelete,
		Path:        "/posts/{id}",
		Summary:     "Delete a post",
		Tags:        []string{"posts"},
	}, DeletePost)
}
