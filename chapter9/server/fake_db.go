package main

import (
	"fmt"
	"time"
)

type FakeDb struct {
	d    *inMemoryDb
	opts testOptions
}

func NewFakeDb(opt ...TestOption) *FakeDb {
	opts := defaultTestOptions
	for _, o := range opt {
		o.apply(&opts)
	}

	return &FakeDb{
		d:    &inMemoryDb{},
		opts: opts,
	}
}

func (db *FakeDb) Reset() {
	db.opts = defaultTestOptions
	db.d = &inMemoryDb{}
}

func (db *FakeDb) addTask(description string, dueDate time.Time) (uint64, error) {
	if !db.opts.isAvailable {
		return 0, fmt.Errorf(
			"couldn't access the database",
		)
	}
	return db.d.addTask(description, dueDate)
}

func (db *FakeDb) getTasks(f func(interface{}) error) error {
	if !db.opts.isAvailable {
		return fmt.Errorf(
			// the error message is different only because we
			// do not handle the error, we directly return it
			// in ListTasks. For other endpoints, we generally
			// prepend with "unexpected error: ".
			"unexpected error: couldn't access the database",
		)
	}
	return db.d.getTasks(f)
}

func (db *FakeDb) updateTask(id uint64, description string, dueDate time.Time, done bool) error {
	if !db.opts.isAvailable {
		return fmt.Errorf(
			"couldn't access the database",
		)
	}
	return db.d.updateTask(id, description, dueDate, done)
}

func (db *FakeDb) deleteTask(id uint64) error {
	if !db.opts.isAvailable {
		return fmt.Errorf(
			"couldn't access the database",
		)
	}
	return db.d.deleteTask(id)
}
