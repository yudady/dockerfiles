# Node.js Sample app using Yarn and a React framework

## Building
pack build react-webpack --builder gcr.io/buildpacks/builder:v1


```shell
pack build react-sample --buildpack paketo-buildpacks/nodejs
```

## Running

```shell
docker run --rm --interactive --tty --init --env PORT=3000 --publish 3000:3000 react-webpack -y
```

## Viewing

```shell
curl http://localhost:3000
```
