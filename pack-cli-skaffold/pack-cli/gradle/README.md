# Java Gradle Sample Application

## Building

```bash
pack build yudady/gradle:v1 --buildpack paketo-buildpacks/java --builder paketobuildpacks/builder:tiny 
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 yudady/gradle:v1
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
