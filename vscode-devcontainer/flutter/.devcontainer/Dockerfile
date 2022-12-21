# Note: You can use any Debian/Ubuntu based image you want. 
FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN set -ex;
RUN apt-get update;
RUN apt-get upgrade;
RUN apt-get install -y  --no-install-recommends \
    bash-completion \
    build-essential \
    curl \
    file \
    zip \
    unzip \
    htop \
    jq \
    less \
    locales \
    man-db \
    software-properties-common \
    sudo \
    time \
    vim \
    multitail \
    lsof \
    openssh-client \
    iproute2 \
    lsb-release \
    xz-utils \
    libglu1-mesa \
    snapd \
    gzip \
    cmake \
    ninja-build \
    clang \
    pkg-config \
    libgtk-3-dev \
    && rm -rf /var/lib/apt/lists/*;

RUN wget https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_3.3.10-stable.tar.xz
RUN mkdir /flutter && tar -xvf ./flutter_linux_3.3.10-stable.tar.xz -C /
RUN echo export PATH="$PATH:/flutter/bin" >> /root/.bashrc
RUN git config --global --add safe.directory /flutter
