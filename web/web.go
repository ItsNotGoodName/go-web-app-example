package web

import (
	"context"
	"net/url"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

//go:generate pnpm install
//go:generate pnpm run build

// ---------- Meta

type Meta struct {
	URL      *url.URL
	TimeZone *time.Location
}

type contextKey string

var metaContextKey contextKey = "meta"

func SetMeta(ctx context.Context, m Meta) context.Context {
	return context.WithValue(ctx, metaContextKey, m)
}

func GetMeta(ctx context.Context) Meta {
	meta, ok := ctx.Value(metaContextKey).(Meta)
	if !ok {
		panic("meta not set")
	}
	return meta
}

func MetaMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Time zone
			timeZone := time.UTC
			tz, err := c.Cookie("tz")
			if err == nil {
				loc, err := time.LoadLocation(tz.Value)
				if err == nil {
					timeZone = loc
				}
			}

			c.SetRequest(c.Request().WithContext(SetMeta(c.Request().Context(), Meta{
				URL:      c.Request().URL,
				TimeZone: timeZone,
			})))
			return next(c)
		}
	}
}

// ---------- Render

var head templ.Component

func init() {
	head = templ.Raw(HeadHTML())
}

func Render(c echo.Context, body templ.Component) error {
	return Base(head, body).Render(c.Request().Context(), c.Response())
}
