package main

import (
	"fmt"
	"time"

	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter5/proto/todo/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type inMemoryDb struct {
	tasks []*pb.Task
}

func New() db {
	return &inMemoryDb{}
}

func (d *inMemoryDb) addTask(description string, dueDate time.Time) (uint64, error) {
	nextId := uint64(len(d.tasks) + 1)
	task := &pb.Task{
		Id:          nextId,
		Description: description,
		DueDate:     timestamppb.New(dueDate),
	}

	d.tasks = append(d.tasks, task)
	return nextId, nil
}

func (d *inMemoryDb) getTasks(f func(interface{}) error) error {
	for _, task := range d.tasks {
		err := f(task)

		if err != nil {
			return err
		}
	}
	return nil
}

func (d *inMemoryDb) updateTask(id uint64, description string, dueDate time.Time, done bool) error {
	for i, task := range d.tasks {
		if task.Id == id {
			t := d.tasks[i]
			t.Description = description
			t.DueDate = timestamppb.New(dueDate)
			t.Done = done
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}

func (d *inMemoryDb) deleteTask(id uint64) error {
	for i, task := range d.tasks {
		if task.Id == id {
			d.tasks = append(d.tasks[:i], d.tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}
