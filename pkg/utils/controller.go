package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"gorm.io/gorm"
)

type handler struct {
	Db *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	r := &handler{
		Db: db,
	}

	h := func(c *fiber.Ctx) error {
		sleepTime, _ := time.ParseDuration(c.Params("sleepTime") + "ms")
		if err := sleepWithContext(c.UserContext(), sleepTime); err != nil {
			return fmt.Errorf("%w: execution error", err)
		}
		return nil
	}

	routes := app.Group("/api/v1/utils")
	app.Static("/", "./public")
	routes.Get("/migrate", r.Migrate, timeout.NewWithContext(h, 10*time.Second))
	routes.Get("/customers", r.GetCustomerList)
	routes.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	routes.Get("/:table", r.GetCountIdTable)
	routes.Get("/:mes/:ano", r.GetCaixaEnd, timeout.NewWithContext(h, 10*time.Second))
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)

	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
		return context.DeadlineExceeded
	case <-timer.C:
	}
	return nil
}
