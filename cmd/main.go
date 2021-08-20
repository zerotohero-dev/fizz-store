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

package main

import (
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-store/internal/api"
)

func main() {
	e := *env.New()

	appEnv := e.Store

	svcName := appEnv.ServiceName

	app.Configure(e, svcName, appEnv.HoneybadgerApiKey, appEnv.Sanitize)

	r := mux.NewRouter()
	api.InitializeEndpoints(e, r)
	app.RouteHealthEndpoints(e.Store.PathPrefix, r)

	app.ListenAndServe(e, svcName, appEnv.Port, r)
}
