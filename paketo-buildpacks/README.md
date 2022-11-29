https://paketo.io/docs/


cd samples/java/maven

pack build paketo-demo-app --builder paketobuildpacks/builder:base


docker run -d -p 8080:8080 -e PORT=8080 paketo-demo-app

pack build samples/java --path java/maven --volume $HOME/.m2:/home/cnb/.m2:rw

curl -s "https://get.sdkman.io" | bash

source "$HOME/.sdkman/bin/sdkman-init.sh"
 
pack build samples/java --path target/demo-0.0.1-SNAPSHOT.jar --volume ${HOME}/.m2:/home/cnb/.m2:rw

pack build samples/java --env BP_JVM_JLINK_ENABLED=true --env BP_JVM_JLINK_ARGS="--no-header-files --compress=1 --add-modules java.base,java.se"

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



