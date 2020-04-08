// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Todos))
}

// Site site interface info
type Todos struct {
	AllTodos         Get    `mir:"/api/todos"`
	TodoByID         Get    `mir:"/api/todos/:id"`
	CreateTodo       Post   `mir:"/api/todos"`
	ToggleTodoStatus Patch  `mir:"/api/todos/:id"`
	DeleteTodo       Delete `mir:"/api/todos/:id"`
}
