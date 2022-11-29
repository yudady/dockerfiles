# Note: You can use any Debian/Ubuntu based image you want. 
FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN wget https://github.com/buildpacks/pack/releases/download/v0.27.0/pack-v0.27.0-linux.tgz

RUN tar zxvf ./pack-v0.27.0-linux.tgz -C /usr/local/bin/

RUN rm -f ./pack-v0.27.0-linux.tgz

#RUN (curl -sSL \"https://github.com/buildpacks/pack/releases/download/v0.27.0/pack-v0.27.0-linux.tgz\" | sudo tar -C /usr/local/bin/ --no-same-owner -xzv pack)
