package task

import (
	"context"
	"fmt"
	cl "github.com/Suhach/test_task-service/internal/client"
)

type Service struct {
	repo       *Repository
	userClient *cl.UserClient
}

func NewService(repo *Repository, userClient *cl.UserClient) *Service {
	return &Service{repo: repo, userClient: userClient}
}

func (s *Service) Create(task string, isDone bool, userID uint) (*Task, error) {
	// Проверяем существование пользователя через gRPC
	_, err := s.userClient.GetUser(context.Background(), uint32(userID))
	if err != nil {
		return nil, fmt.Errorf("user with ID %d not found: %w", userID, err)
	}
	t := &Task{Task: task, IsDone: isDone, UserID: userID}
	if err := s.repo.Create(t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) GetByUserID(userID uint) ([]Task, error) {
	// Проверяем существование пользователя
	_, err := s.userClient.GetUser(context.Background(), uint32(userID))
	if err != nil {
		return nil, fmt.Errorf("user with ID %d not found: %w", userID, err)
	}
	return s.repo.GetByUserID(userID)
}

func (s *Service) GetAll() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *Service) Get(id uint) (*Task, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(id uint, task string, isDone bool, userID uint) (*Task, error) {
	// Проверяем существование пользователя
	_, err := s.userClient.GetUser(context.Background(), uint32(userID))
	if err != nil {
		return nil, fmt.Errorf("user with ID %d not found: %w", userID, err)
	}
	t, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	t.Task = task
	t.IsDone = isDone
	t.UserID = userID
	if err := s.repo.Update(t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
