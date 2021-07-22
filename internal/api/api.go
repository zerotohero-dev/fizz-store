/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package api

import (
	"context"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-store/internal/endpoint"
	"github.com/zerotohero-dev/fizz-store/internal/service"
	"github.com/zerotohero-dev/fizz-store/internal/transport"
)

const prefix = "/store"

func InitializeEndpoints(e env.FizzEnv, router *mux.Router) {
	svc := service.New(e, context.Background())

	// Stripe callback.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeSubscribeEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeSubscribeRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/subscribe",
	)
}
