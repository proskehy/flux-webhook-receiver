A tool to receive webhooks regarding Git/Docker image changes and send a notification to [Weave flux](https://github.com/weaveworks/flux).

# Supported services  

`flux-webhook-receiver` lets you setup Git repository webhooks from
 * `GitHub`,
 * `GitLab`,
 * `Bitbucket` and
 * `Bitbucket Server`.  
 
Docker image webhooks can be configured for `DockerHub` and `Nexus`.

For the services which allow it, you can configure a secret to verify the incoming payload.  
You can also set the Git branch that you want to receive events from.  

# Setup

It is intended to run this tool as a sidecar to Flux (see [Get started with Flux](https://docs.fluxcd.io/en/latest/tutorials/get-started.html)).

Example Flux deployment with the sidecar configured can be found in `examples/flux-deployment.yaml`. 
Then, you will want to expose the deployment with Ingress, example of that is in `examples/flux-service.yaml` and `examples/flux-ingress.yaml`.  
Sensitive values for the webhook secrets should be kept in Kubernetes Secrets. You can either create those directly or use something like [SealedSecrets](https://github.com/bitnami-labs/sealed-secrets) or [HashiCorp's Vault](https://www.vaultproject.io/).  
Apply the other manifests in the `examples` folder (namespace, serviceaccount, secret) as well if those resources are not present in your cluster yet.

The service runs on port `3033` and individual webhook handlers are exposed on paths `/gitSync` and `/imageSync`.

Environment variables to configure the deployment:

  * `GIT_ENABLED`: Enable the endpoint for Git repo webhooks (default `true`)
  * `DOCKER_ENABLED`: Enable the endpoit for Docker registry webhooks (default `true`)
  * `GIT_WEBHOOK_SECRET`: Secret to verify the Git repo webhook payload with (not set by default)
  * `GIT_BRANCHES`: Git branches to receive webhooks from specified as space separated values e.g. `master develop` (default `master`)
  * `GIT_HOST`: Git repository host that the webhooks will be coming from (default `github`) 
  * `FLUX_DOCKER_HOST`: Docker registry host that the webhooks will be coming from (default `dockerhub`) 
  * `DOCKER_WEBHOOK_SECRET`: Secret to verify the Docker registry webhook payload with (not set by default)
