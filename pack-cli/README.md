# http

```shell
pack config default-builder paketobuildpacks/builder:full
```


## golang

```shell
# build image
pack build yudady/mygo:v3 --buildpack paketo-buildpacks/go --builder paketobuildpacks/builder:tiny 

# run
docker run -it --rm -p 8080:8080 yudady/mygo:v3

```


## java-gradle

```shell
# build image
./gradlew bootBuildImage --imageName=yudady/java-gradle:v1

# run
docker run -it --rm -p 8080:8080 yudady/java-gradle:v1

```


## java-maven

```shell

./mvnw spring-boot:build-image 


# run
docker run -it --rm -p 8080:8080 yudady/java-maven:v1

```

docker system prune