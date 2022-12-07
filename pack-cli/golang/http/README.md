# http

```shell
# build image
pack build yudady/mygo:v3 --buildpack paketo-buildpacks/go --builder paketobuildpacks/builder:tiny 

# run
docker run -it --rm -p 8080:8080 yudady/mygo:v3

```

