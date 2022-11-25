(curl -sSL "https://github.com/buildpacks/pack/releases/download/v0.27.0/pack-v0.27.0-linux.tgz" | sudo tar -C /usr/local/bin/ --no-same-owner -xzv pack)

vscode ➜ /workspace $ pack builder suggest
Suggested builders:
        Google:                gcr.io/buildpacks/builder:v1      Ubuntu 18 base image with buildpacks for .NET, Go, Java, Node.js, and Python                                                      
        Heroku:                heroku/buildpacks:20              Base builder for Heroku-20 stack, based on ubuntu:20.04 base image                                                                
        Paketo Buildpacks:     paketobuildpacks/builder:base     Ubuntu bionic base image with buildpacks for Java, .NET Core, NodeJS, Go, Python, Ruby, Apache HTTPD, NGINX and Procfile          
        Paketo Buildpacks:     paketobuildpacks/builder:full     Ubuntu bionic base image with buildpacks for Java, .NET Core, NodeJS, Go, Python, PHP, Ruby, Apache HTTPD, NGINX and Procfile     
        Paketo Buildpacks:     paketobuildpacks/builder:tiny     Tiny base image (bionic build image, distroless-like run image) with buildpacks for Java, Java Native Image and Go                

Tip: Learn more about a specific builder with:
        pack builder inspect <builder-image>



pack build sample-app --path apps/java-maven --builder heroku/buildpacks:20


pack build sample-app --path apps/java-maven --builder cnbs/sample-builder:bionic




