# https://github.com/devcontainers/images/blob/main/src/go/README.md
FROM mcr.microsoft.com/devcontainers/go:1.19

#######
RUN curl -fsSL "https://get.sdkman.io" | bash; \
    bash -c ". /root/.sdkman/bin/sdkman-init.sh \
    && sdk install java 22.3.r17-grl \
    && sdk install maven \
    && sdk install gradle \
    && sdk flush archives \
    && sdk flush temp \
    && mkdir /root/.m2 \
    && mkdir /root/.gradle \
    && mkdir /root/.bashrc.d \
    && echo 'export SDKMAN_DIR=\"/root/.sdkman\"' >> /root/.bashrc.d/99-java \
    && echo '[[ -s \"/root/.sdkman/bin/sdkman-init.sh\" ]] && source \"/root/.sdkman/bin/sdkman-init.sh\"' >> /root/.bashrc.d/99-java"; \
    # above, we are adding the sdkman init to .bashrc (executing sdkman-init.sh does that), because one is executed on interactive shells, the other for non-interactive shells (e.g. plugin-host)
    # Clean up
    apt-get autoremove -y; \
    apt-get clean -y; \
    rm -rf /var/lib/apt/lists/* 


RUN wget https://github.com/buildpacks/pack/releases/download/v0.27.0/pack-v0.27.0-linux.tgz

RUN tar zxvf ./pack-v0.27.0-linux.tgz -C /usr/local/bin/

RUN rm -f ./pack-v0.27.0-linux.tgz