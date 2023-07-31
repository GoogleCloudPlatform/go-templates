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
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	*http.Server
	projectID string
	log       *logging.Logger
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

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	app, err := newApp(ctx, port, projectID)
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

func newApp(ctx context.Context, port, projectID string) (*App, error) {
	if projectID == "" {
		projID, err := metadata.ProjectID()
		if err != nil {
			return nil, fmt.Errorf("unable to detect Project ID from GOOGLE_CLOUD_PROJECT or metadata server: %w", err)
		}
		projectID = projID
	}

	app := &App{
		Server: &http.Server{
			Addr: ":" + port,
			// Add some defaults, should be changed to suit your use case.
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		projectID: projectID,
	}

	client, err := logging.NewClient(ctx, fmt.Sprintf("projects/%s", app.projectID),
		// We don't need to make any requests when logging to stderr.
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		))
	if err != nil {
		return nil, fmt.Errorf("unable to initialize logging client: %v", err)
	}
	app.log = client.Logger("test-log", logging.RedirectAsJSON(os.Stderr))

	// Setup request router.
	r := mux.NewRouter()
	r.HandleFunc("/", app.Handler).Methods("GET")
	app.Server.Handler = r
	return app, nil
}
