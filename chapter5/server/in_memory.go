package main

import (
	"fmt"
	"time"

	pb "github.com/PacktPublishing/gRPC-Go-for-Professionals/proto/todo/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// inMemoryDb is a fake database.
// Its purpose is to let us focus on gRPC
// and to not add any database dependencies such as ORM, ...
type inMemoryDb struct {
	tasks []*pb.Task
}

// New creates a new instance of inMemoryDb
func New() db {
	return &inMemoryDb{}
}

// addTask appends a Task, generated from description and dueDate, to the underlying array.
// It never returns an error but a real database might!
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

// getTasks applies the function f to all the Tasks in the underlying array.
// If any error happens, it will short circuit and return the error.
func (d *inMemoryDb) getTasks(f func(interface{}) error) error {
	for _, task := range d.tasks {
		if err := f(task); err != nil {
			return err
		}
	}
	return nil
}

// updateTask applies an update of task to the first task that has the task.Id == id.
// If a task is not found with the given id, it returns an error.
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

// deleteTask deletes a task with task.Id == id in the underlying array.
// If a task is not found with the given id, it returns an error.
func (d *inMemoryDb) deleteTask(id uint64) error {
	for i, task := range d.tasks {
		if task.Id == id {
			d.tasks = append(d.tasks[:i], d.tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}
