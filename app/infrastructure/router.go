package infrastructure

import (
	"app/app/utils"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *Router {
	e := echo.New()

	e.Use(
		CustomContextMiddleware(),
		middleware.Logger(),
	)

	e.HTTPErrorHandler = httpErrorHandler

	zaimController := controllers.NewZaimController()

	return &Router{Echo: e, closables: []closable{}}
}

type Router struct {
	Echo      *echo.Echo
	closables []closable
}

type closable interface {
	Close() error
}

func (r *Router) Start() {
	go func() {
		if err := r.Echo.Start(":8080"); err != nil {
			r.Echo.Logger.Fatal(err)
		}
	}()
}

func (r *Router) Shutdown(ctx context.Context) {
	if err := r.Echo.Shutdown(ctx); err != nil {
		r.Echo.Logger.Error(err)
	}

	r.closeAll(ctx)
}

func (r *Router) closeAll(ctx context.Context) {
	var wg utils.WaitGroup

	for _, c := range r.closables {
		c := c
		wg.Go(func() {
			select {
			case <-ctx.Done():
				r.Echo.Logger.Error(ctx.Err())
			default:
				if err := c.Close(); err != nil {
					r.Echo.Logger.Error(err)
				}
			}
		})
	}

	wg.Wait()
}

type ControllerFunc func(controllers.Context) error

func newHandlerFunc(f ControllerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return f(c.(*CustomContext))
	}
}
