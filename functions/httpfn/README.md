# HTTP Cloud Function (gen 2)

This project is a getting starter guide for working with HTTP Cloud Functions.

## Download template locally with gonew

* Install [gonew](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) if you have
  not already.

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

* Download this template locally:

```bash
gonew github.com/GoogleCloudPlatform/go-templates/functions/httpfn your.domain/httpfn
```

## Deploy

```bash
gcloud functions deploy hello-http \
--gen2 \
--runtime=go120 \
--region=us-central1 \
--source=. \
--entry-point=HelloHTTP \
--trigger-http \
--allow-unauthenticated
```

*Note*, in projection environments consider removing the
`--allow-unauthenticated` flag.

## Resources

* [Create and deploy an HTTP Cloud Function with Go](https://cloud.google.com/functions/docs/create-deploy-http-go)
* [Cloud Function Go Runtime](https://cloud.google.com/functions/docs/concepts/go-runtime)
* [gcloud functions deploy reference docs](https://cloud.google.com/sdk/gcloud/reference/functions/deploy)
