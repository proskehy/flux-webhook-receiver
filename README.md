A tool to receive webhooks regarding Git/image changes and send a notification to [Weave flux](https://github.com/weaveworks/flux).

It lets you setup webhooks from GitHub, GitLab, Bitbucket and Bitbucket Server.  
For the services which allow it, you can configure a secret to verify the incoming payload.  
You can also set the branch that you want to receive events from.

# Setup

It is intended to run this tool as a sidecar to Flux (see [Get started with Flux](https://docs.fluxcd.io/en/latest/tutorials/get-started.html)).

Example Flux deployment with the sidecar configured can be found in `examples/flux-deployment.yaml`. Then, you will want to expose the deployment with Ingress, example of that is in `examples/flux-service.yaml` and `examples/flux-ingress.yaml`.

Environment variables to configure the deployment:

  * `GIT_WEBHOOK_SECRET`: secret to verify the payload with (not set by default) 
  * `GIT_BRANCH`: branch to receive webhooks from (default `master`)
  * `GIT_HOST`: repository host that the webhooks will be coming from (default `github`) 
  
 There are no releases of this image yet, so you will need to build it and host it somewhere yourself for now.