package chat

import (
	"aichat/internal/app/dto"
	"context"
	openai "github.com/sashabaranov/go-openai"
	"os"
)

type Service interface {
	ProcessChat(*dto.ChatRequest) (*dto.ChatResponse, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ProcessChatxx(req *dto.ChatRequest) (*dto.ChatResponse, error) {
	key := os.Getenv("OPENAPI_KEY")
	client := openai.NewClient(key)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: req.Message,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	// fmt.Print(resp.Choices[0].Message.Content)
	return &dto.ChatResponse{
		Message: resp.Choices[0].Message.Content,
	}, nil
	// return s.repo.GetChat()
}

func (s *service) ProcessChat(req *dto.ChatRequest) (*dto.ChatResponse, error) {
	key := os.Getenv("OPENAPI_KEY")
	client := openai.NewClient(key)
	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: req.Message,
	})

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		return nil, err
	}

	return &dto.ChatResponse{
		Message: resp.Choices[0].Message.Content,
	}, nil
}
