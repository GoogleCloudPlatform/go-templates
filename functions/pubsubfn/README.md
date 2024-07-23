# PubSub Event (Gen 2)

This project is a getting starter guide for working with a Pub/Sub event
handling Cloud Functions.

## Download template locally with gonew

* Install [gonew](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) if you have
  not already.

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

* Download this template locally:

```bash
gonew github.com/GoogleCloudPlatform/go-templates/functions/pubsubfn your.domain/pubsubfn
```

## Deploy

```bash
# The name of the Pub/Sub topic that the Cloud Function will listen on
EXPORT TOPIC_NAME=func-test

# Create the Pub/Sub Topic
gcloud pubsub topics create $TOPIC_NAME

# Deploy the Cloud Function
gcloud functions deploy hello-pubsub \
--gen2 \
--runtime=go122 \
--region=us-central1 \
--source=. \
--entry-point=HelloPubSub \
--trigger-topic=func-test

# Send a test message to your topic
gcloud pubsub topics publish $$TOPIC_NAME --message="Test message from gcloud"
```

## Resources

* [Cloud Functions Pub/Sub tutorial](https://cloud.google.com/functions/docs/tutorials/pubsub#functions-prepare-environment-go)
* [Cloud Function Go Runtime](https://cloud.google.com/functions/docs/concepts/go-runtime)
* [gcloud functions deploy reference docs](https://cloud.google.com/sdk/gcloud/reference/functions/deploy)
* [What is Pub/Sub?](https://cloud.google.com/pubsub/docs/overview)
* [Eventarc triggers](https://cloud.google.com/functions/docs/calling/eventarc)
