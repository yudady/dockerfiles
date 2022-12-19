# [pack](https://paketo.io/docs/)

- [paketo-buildpacks](https://github.com/paketo-buildpacks)
- [buildpacks](https://buildpacks.io/docs/app-developer-guide/build-an-app/)


## set default builder

```shell
pack config default-builder gcr.io/buildpacks/builder:v1

pack config default-builder paketobuildpacks/builder:full
pack config default-builder paketobuildpacks/builder:base
pack config default-builder paketobuildpacks/builder:tiny
```

## clean docker image

```shell
docker system prune
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


```shell
pack build samples/java-native \
  --env BP_NATIVE_IMAGE=true \
  --env BP_LOG_LEVEL=DEBUG \
  --path java/native-image/java-native-image-sample \
  --volume ${HOME}/.m2:/home/cnb/.m2:rw


pack build samples/java-native \
  --env BP_NATIVE_IMAGE=true \
  --path java/native-image/java-native-image-sample \
  --volume ${HOME}/.m2:/home/cnb/.m2:rw \
  --buildpack paketo-buildpacks/graalvm \
  --buildpack paketo-buildpacks/java-native-image \
  --env BP_LOG_LEVEL=DEBUG
```

