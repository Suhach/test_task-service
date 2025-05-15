package grpc

import (
	"context"

	pb "github.com/Suhach/test_protoc-cont/proto/task"
	"github.com/Suhach/test_task-service/internal/task"
)

type Handler struct {
	pb.UnimplementedTaskServiceServer
	svc *task.Service
}

func NewHandler(svc *task.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task, err := h.svc.Create(req.Task, req.IsDone, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return &pb.CreateTaskResponse{Id: uint32(task.ID), Task: task.Task, IsDone: task.IsDone, UserId: uint32(task.UserID)}, nil
}

func (h *Handler) GetTasksByUser(ctx context.Context, req *pb.GetTasksByUserRequest) (*pb.GetTasksByUserResponse, error) {
	tasks, err := h.svc.GetByUserID(uint(req.UserId))
	if err != nil {
		return nil, err
	}
	var pbTasks []*pb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &pb.Task{Id: uint32(t.ID), Task: t.Task, IsDone: t.IsDone, UserId: uint32(t.UserID)})
	}
	return &pb.GetTasksByUserResponse{Tasks: pbTasks}, nil
}

func (h *Handler) GetAllTasks(ctx context.Context, req *pb.GetAllTasksRequest) (*pb.GetAllTasksResponse, error) {
	tasks, err := h.svc.GetAll()
	if err != nil {
		return nil, err
	}
	var pbTasks []*pb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &pb.Task{Id: uint32(t.ID), Task: t.Task, IsDone: t.IsDone, UserId: uint32(t.UserID)})
	}
	return &pb.GetAllTasksResponse{Tasks: pbTasks}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	task, err := h.svc.Get(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetTaskResponse{Id: uint32(task.ID), Task: task.Task, IsDone: task.IsDone, UserId: uint32(task.UserID)}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	task, err := h.svc.Update(uint(req.Id), req.Task, req.IsDone, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTaskResponse{Id: uint32(task.ID), Task: task.Task, IsDone: task.IsDone, UserId: uint32(task.UserID)}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	if err := h.svc.Delete(uint(req.Id)); err != nil {
		return nil, err
	}
	return &pb.DeleteTaskResponse{}, nil
}
