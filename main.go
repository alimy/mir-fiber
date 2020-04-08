// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/alimy/mir-fiber/internal/servants"
	"github.com/alimy/mir-fiber/mirc/gen/api"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectionString := os.Getenv("MONGODB_URI")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	err := mgm.SetDefaultConfig(nil, "todos", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	// register servants to fiber
	ts := servants.NewTodoServant()
	api.RegisterTodosServant(app, ts)

	// start servant service
	if err := app.Listen(3000); err != nil {
		log.Fatal(err)
	}
}
