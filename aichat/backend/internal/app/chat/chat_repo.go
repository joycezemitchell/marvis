package chat

import "aichat/internal/app/dto"

type Repository interface {
	GetChat(*dto.ChatRequest) error
}

type repository struct {
	// Database connection goes here
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetChat(*dto.ChatRequest) error {
	// Database access code goes here

	return nil
}
