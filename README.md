# Get Secrets

This is a tool for reading various config files from S3 and exposing them to a sub-process as Env-Vars

## Raison d'Etre

The main idea of this is to avoid shipping config files that contain secrets with the
app, or for copying them into place as part of the deployment process; they can be
"deployed" at the last moment before the app is started.

In most cases, other systems can manage this for you (e.g. Puppet or Kubernetes
Secrets) but in cases such as running non-k8s Docker containers, this becomes more
complicated partly because the idealised smaller image size means configuration
management tooling is not really appropriate.

This is a middle-road option that avoids needing to add config management code,
mount special file-systems, or include secrets-delivery client code (e.g. Hashicorp
Vault).

## How it works

Rather than running your app / binary directly, you would instead run `get-secrets`
and pass the path/file of your app plus any command-line arguments _your_ app needs (as
additional arguments for `get-secrets`).

It will then use certain configurations (see below) to find the secrets files, parse the contents
of those into new env vars, and then start your app (with the additional args) with the new
env vars (from the parsed secrets files) passed-on to your app.

Whilst primarily designed for Docker containers, `get-secrets` is agnostic to the environment it
runs in.

The only potential gotcha is how you pass AWS authentication info for use:  EC2 and ECS expose
this access via an http://169... address, but you can also use the `$AWS_...` env vars;  in a
Docker container, you won't be able to use `~/.aws/credentials` (etc) without specifically
creating or mounting that dir into the container.

## Configuration

`get-secrets` uses the [Viper](https://github.com/spf13/viper) library to configure itself.

This is set-up to look for a `.secrets.toml` file in three possible places (in order of
preference):
- in the `$HOME` dir
- in the directory that `$SECRETS_BASE` points to
- in the directory that the `get-secrets` binary runs from

Within that file, it can contain something like the following:

```yaml
## This sets debug-mode for `get-secrets` (only):
## Can also use "$SECRETS_DEBUG"
debug = true
## The runtime and local directory:
## Can also use "$SECRETS_BASE"
base = "RUNTIME_DIR"

[application]
## Really only used for logging. Allows you to give a "friendly" name in logging for _your_ app:
## Can also use "$APPLICATION_NAME"
name = "APP-NAME"
## Really only used for logging. Allows you to give a "friendly" name in logging for the type
## of environment your app will be running in -- e.g. "production" or "development":
## Can also use "$ENVIRONMENT"
environment = "ENV-NAME"

[dotenv]
## Whether `get-secrets` should skip straight to running the new app, rather than downloading
## and processing the secrets files from S3:
## Can also use "$SKIP_SECRETS"
skip = false

[s3]
## Base path on S3 that contains the secrets files:
## Can also use "$SECRETS_S3_DOTENV_PATH"
path = "s3://BUCKET/DIR-PATH"

[logging]
## This is the logging output-format.  Either "text" or "json":
## Can also use "$SECRETS_LOGGING_FORMAT"
format = "text"

[logging.sentry]
## This allows sending errors or failures within `get-secrets` to https://sentry.io
## Can also use "$SECRETS_LOGGING_SENTRY_DSN"
dsn = "https://ID:TOKEN@sentry.io/PROJECT"
```

Note that these all map to optional env vars, this means you can override the above settings
by passing in an env var to `get-secrets` as mentioned in the code comments, above.

## Alternatives

- Kubernetes Secrets
- Include code that can read directly from KMS, Hashicorp Vault (et al)
- Inject a file into system at run-time
- Mount a read-only volume containing secrets into your system
- Mount a named-pipe from host into the container that delivers one-shot-access
  secret data (containerised only)
