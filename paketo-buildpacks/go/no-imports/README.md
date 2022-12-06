# Go Sample App using no imports

## Building

`pack build go-sample --buildpack paketo-buildpacks/go`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 go-sample`

## Viewing

`curl http://localhost:8080`


pack build yudady/go-app:v3 --buildpack paketo-buildpacks/go --builder paketobuildpacks/builder:tiny 

docker run --interactive --tty --env PORT=8080 --publish 8080:8080 yudady/go-app:v3


docker run -it -p 8080:8080 --env PORT=8080 yudady/go-app:v2
