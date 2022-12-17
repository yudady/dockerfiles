# Java Maven Sample Application

## Building

```bash
pack build yudady/maven --buildpack paketo-buildpacks/java --builder paketobuildpacks/builder:tiny 
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 yudady/maven
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
