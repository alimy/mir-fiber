// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package models

import (
	"github.com/Kamva/mgm/v2"
)

// Todo is the model that defines a todo entry
type Todo struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Description      string `json:"description" bson:"description"`
	Done             bool   `json:"done" bson:"done"`
}

// NewTodo is a wrapper that creates a new todo entry
func NewTodo(title, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Done:        false,
	}
}

// NilTodo return an empty todo entry
func NilTodo() *Todo {
	return &Todo{}
}
