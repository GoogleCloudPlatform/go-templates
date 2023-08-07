# Cloud Run Template Microservice

A template repository for a Cloud Run microservice, written in Go.

## Download template locally with gonew

* Install [gonew](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) if you have
  not already.

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

* Download this template locally:

```bash
gonew github.com/GoogleCloudPlatform/go-templates/run/microservice your.domain/microservice
```

## Prerequisite

* Enable the Cloud Run API via the CLI:

```bash
gcloud services enable run.googleapis.com
gcloud services enable artifactregistry.googleapis.com
```

## Features

* **gorilla/mux**: A request router and dispatcher
* **Dockerfile**: Container build instructions
* **SIGTERM handler**: Catch termination signal for cleanup before Cloud Run stops the container
* **Structured logging w/ Log Correlation** JSON formatted logger, parsable by Cloud Logging, with [automatic correlation of container logs to a request log](https://cloud.google.com/run/docs/logging#correlate-logs).
* **Unit tests** Basic unit tests setup for the microservice

## Local development

```bash
# Set your project ID
export GOOGLE_CLOUD_PROJECT=<GCP_PROJECT_ID>

#Build and Start the server
go build -o server && ./server
```

## Deploying to Cloud Run

### Setup

```bash
# Set your project ID
export GOOGLE_CLOUD_PROJECT=<GCP_PROJECT_ID>
# Name of your Artifact Registry repository 
export ARTIFACT_REPOSITORY="run-images"
# Name of the container image
export IMAGE_NAME="micro"
# The Google Cloud region your code will run in
export REGION=us-central1

# Enable Artifact Registry
gcloud services enable artifactregistry.googleapis.com

# Create Artifact Registry repository
gcloud artifacts repositories create $ARTIFACT_REPOSITORY --location $REGION --repository-format "docker"

# Use the gcloud credential helper to authorize Docker to push to your Artifact Registry
gcloud auth configure-docker
```

### Deploy

```bash
# Build container image in Cloud Build and store it in Artifact Registry
gcloud builds submit --pack image=$REGION-docker.pkg.dev/$GOOGLE_CLOUD_PROJECT/$ARTIFACT_REPOSITORY/$IMAGE_NAME:manual
# Deploy image to Cloud Run
gcloud run deploy microservice-template \
--image $REGION-docker.pkg.dev/$GOOGLE_CLOUD_PROJECT/$ARTIFACT_REPOSITORY/$IMAGE_NAME:manual
```

## Resources

* [Deploy a Go service to Cloud Run](https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service)
* [Overview of Cloud Build](https://cloud.google.com/build/docs/overview)
* [Artifact Registry: Connect to Cloud Build](https://cloud.google.com/artifact-registry/docs/configure-cloud-build)
