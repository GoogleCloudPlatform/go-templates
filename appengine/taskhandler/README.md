# App Engine Task Handler

This project is a getting starter guide for working with App Engine task
handlers in Go.

## Download template locally with gonew

* Install [gonew](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) if you have
  not already.

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

* Download this template locally:

```bash
gonew github.com/GoogleCloudPlatform/go-templates/appengine/taskhandler your.domain/taskhandler
```

## Deploy

```bash
# The name of the Cloud Task queue the handler will receive message from
EXPORT AE_QUEUE=ae-queue

# Deploy the app to App Engine
gcloud app deploy

# Create the Cloud Task queue
gcloud tasks queues create $AE_QUEUE

# Connect the Cloud Task queue to the App Engine handler endpoint
gcloud tasks create-app-engine-task hello --queue=$AE_QUEUE --relative-uri=/task_handler --body-content="a friendly hello"
```

## Resources

* [Create App Engine task handlers](https://cloud.google.com/tasks/docs/creating-appengine-handlers#go)
* [Cloud Task queues with App Engine](https://cloud.google.com/tasks/docs/dual-overview#appe)
* [gcloud tasks create-app-engine-task](https://cloud.google.com/sdk/gcloud/reference/tasks/create-app-engine-task)
* [Create task with Client library](https://cloud.google.com/tasks/docs/samples/cloud-tasks-appengine-create-task?hl=en#cloud_tasks_appengine_create_task-go)
