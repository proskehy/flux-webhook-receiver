A tool to receive webhooks regarding Git/Docker image changes and send a notification to [Weave flux](https://github.com/weaveworks/flux).

# Supported services  

`flux-webhook-receiver` lets you setup Git repository webhooks from
 * GitHub,
 * GitLab,
 * Bitbucket,
 * Bitbucket Server  
and Docker image webhooks from
 * DockerHub and
 * Nexus.

For the services which allow it, you can configure a secret to verify the incoming payload.  
You can also set the Git branch that you want to receive events from.  

# Setup

It is intended to run this tool as a sidecar to Flux (see [Get started with Flux](https://docs.fluxcd.io/en/latest/tutorials/get-started.html)).

Example Flux deployment with the sidecar configured can be found in `examples/flux-deployment.yaml`. 
Then, you will want to expose the deployment with Ingress, example of that is in `examples/flux-service.yaml` and `examples/flux-ingress.yaml`.  

The service runs on port `3033` and individual webhook handlers are exposed on paths `/gitSync` and `/imageSync`.

Environment variables to configure the deployment:

  * `GIT_WEBHOOK_SECRET`: secret to verify the git repo webhook payload with (not set by default) 
  * `GIT_BRANCH`: git branch to receive webhooks from (default `master`)
  * `GIT_HOST`: git repository host that the webhooks will be coming from (default `github`) 
  * `DOCKER_HOST`: Docker registry host that the webhooks will be coming from (default `dockerhub`) 
  * `DOCKER_WEBHOOK_SECRET`: secret to verify the Docker registry webhook payload with (not set by default)
