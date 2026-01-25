# Fast, secure and self-hostable github webhook server

I built this small utility mainly to receive **github webhooks** on my infrastructure and use them to trigger deployment jobs on the server.

## Pre-build configuration

After you cloned this repository, you'll need to edit `.env.example` and `config/services.json.example` with your own configuration.

**`.env`**
 - `WEBHOOK_SECRET`: base64 encoded string, can be generate by running (for example) the following command `head -c 30 /dev/random | base64`
 - `WEBHOOK_PORT`: the port used to listen for github webhooks
 - `DEPLOY_TOKEN_<NAME>`: token deployment that will be included a specifique `X-Deploy-Token`

**`config/services.json`**
```json
{
    "github-repository-name": {
        "url": "<public-or-local-url>",
        "token_env": "<name-of-env-variable-containing-the-secret>"
    }
}
```

> [!TIP]
> Don't forget to rename files to `.env` and `config/services.json` after editing the .example versions

## Build

This service is clearly intended to run inside a container (podman preferably). To build the container, execute the following command:
```sh
podman build -t webgohook .
```

## Run

Indise the root folder, execute the run command:

```sh
podman run -d --name webgohook --env-file .env -p 0.0.0.0:<desired-port>:8081 webgohook:latest
```
> [!NOTE]
> The webgohook server is by default listening on port 8081 inside the container if you didn't set the `WEBHOOK_PORT` variable in the `.env` file
