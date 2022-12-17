# Node.js Sample App using NPM

## Building

```shell
pack build npm-sample --buildpack paketo-buildpacks/nodejs
```

## Running

```shell
docker run --interactive --tty --init --publish 8080:8080 npm-sample
```

## Viewing

```shell
curl http://localhost:8080
```


