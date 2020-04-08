// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"log"

	"github.com/gofiber/fiber"
)

type baseServant struct {
	// just empty
}

func (s baseServant) success(c *fiber.Ctx, data interface{}) {
	err := c.JSON(fiber.Map{
		"status": "success",
		"todos":  data,
	})
	if err != nil {
		log.Print(err)
	}
}

func (s baseServant) failure(c *fiber.Ctx, code int, err error) {
	log.Printf("failure: %s\n", err)
	resErr := c.Status(code).JSON(fiber.Map{
		"status": "failure",
		"error":  err.Error(),
	})
	if resErr != nil {
		log.Print(err)
	}
}
