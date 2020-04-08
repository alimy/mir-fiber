// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"errors"
	"net/http"

	"github.com/Kamva/mgm/v2"
	"github.com/alimy/mir-fiber/internal/models"
	"github.com/alimy/mir-fiber/mirc/gen/api"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	errNotFound    = errors.New("not found")
	errTitleOrDesc = errors.New("title or description not specified")
)

type simpleTodos struct {
	baseServant
}

// AllTodos - GET /api/todos
func (s *simpleTodos) AllTodos(c *fiber.Ctx) {
	todo := models.NilTodo()
	collection := mgm.Coll(todo)
	var todos []models.Todo

	if err := collection.SimpleFind(&todos, bson.D{}); err == nil {
		s.success(c, todos)
	} else {
		s.failure(c, http.StatusInternalServerError, err)
	}
}

// TodoByID - GET /api/todos/:id
func (s *simpleTodos) TodoByID(c *fiber.Ctx) {
	id := c.Params("id")

	todo := models.NilTodo()
	collection := mgm.Coll(todo)

	if err := collection.FindByID(id, todo); err == nil {
		s.success(c, todo)
	} else {
		s.failure(c, http.StatusNotFound, errNotFound)
	}
}

// CreateTodo - POST /api/todos
func (s *simpleTodos) CreateTodo(c *fiber.Ctx) {
	params := new(struct {
		Title       string
		Description string
	})

	if err := c.BodyParser(&params); err != nil {
		s.failure(c, http.StatusInternalServerError, err)
		return
	}

	if len(params.Title) == 0 || len(params.Description) == 0 {
		s.failure(c, http.StatusBadRequest, errTitleOrDesc)
		return
	}

	todo := models.NewTodo(params.Title, params.Description)
	if err := mgm.Coll(todo).Create(todo); err == nil {
		s.success(c, todo)
	} else {
		s.failure(c, http.StatusInternalServerError, err)
	}
}

// ToggleTodoStatus - PATCH /api/todos/:id
func (s *simpleTodos) ToggleTodoStatus(c *fiber.Ctx) {
	id := c.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		s.failure(c, http.StatusNotFound, errNotFound)
		return
	}

	todo.Done = !todo.Done
	err = collection.Update(todo)
	if err = collection.Update(todo); err == nil {
		s.success(c, todo)
	} else {
		s.failure(c, http.StatusInternalServerError, err)
	}
}

// DeleteTodo - DELETE /api/todos/:id
func (s *simpleTodos) DeleteTodo(c *fiber.Ctx) {
	id := c.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		s.failure(c, http.StatusNotFound, errNotFound)
		return
	}

	if err = collection.Delete(todo); err == nil {
		s.success(c, todo)
	} else {
		s.failure(c, http.StatusInternalServerError, err)
	}
}

// NewTodoServant return a simple todo service
func NewTodoServant() api.Todos {
	return &simpleTodos{
		baseServant: baseServant{},
	}
}
