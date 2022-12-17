# Node.js Sample app using Yarn and a React framework

## Building

```shell
pack build react-sample --buildpack paketo-buildpacks/nodejs
```

## Running

```shell
docker run --rm --interactive --tty --init --env PORT=8080 --publish 8080:8080 react-sample
```

## Viewing

```shell
curl http://localhost:8080
```
