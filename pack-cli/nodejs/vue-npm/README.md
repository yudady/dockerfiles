# Node.js Sample app using Npm and a Vue framework

## Building

```shell
pack build vue-sample --buildpack paketo-buildpacks/nodejs --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"
```

## Running

```shell
docker run --interactive --tty --init --publish 8080:8080 --rm vue-sample
```

## Viewing

```shell
curl http://localhost:8080
```

### Note

We need the additional flag `--env "NODE_ENV=development"` when running `pack build` since we need the `vue-cli-service` provided in the devDependencies.
