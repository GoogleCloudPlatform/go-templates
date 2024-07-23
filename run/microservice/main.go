// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	*http.Server
	logger *slog.Logger
}

func main() {
	ctx := context.Background()
	// Listen for SIGINT to gracefully shutdown.
	nctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on port %s", port)

	app, err := newApp(port)
	if err != nil {
		log.Fatalf("unable to initialize application: %v", err)
	}
	log.Println("starting HTTP server")
	go func() {
		if err := app.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server closed: %v", err)
		}
	}()

	<-nctx.Done()
	log.Println("shutdown initiated")
	// Cloud Run gives apps 10 seconds to shutdown. See
	// https://cloud.google.com/blog/topics/developers-practitioners/graceful-shutdowns-cloud-run-deep-dive
	// for more details.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	app.Shutdown(ctx)
	log.Println("shutdown")
}

func newApp(port string) (*App, error) {
	app := &App{
		Server: &http.Server{
			Addr: ":" + port,
			// Add some defaults, should be changed to suit your use case.
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		logger: slog.New(slog.NewJSONHandler(os.Stderr, nil)),
	}

	// Setup request router.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Handler)
	app.Server.Handler = mux
	return app, nil
}
