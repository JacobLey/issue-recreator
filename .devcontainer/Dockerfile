# https://github.com/devcontainers/images/blob/main/src/base-debian/history
FROM mcr.microsoft.com/devcontainers/base:1.0.19-debian12

ARG NODE_VERSION=22.13.1
ARG PNPM_VERSION=9.15.2
ARG DAGGER_VERSION=0.15.3
ARG GO_VERSION=1.23.2

RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install vim -y

ARG USER=issue-recreator
RUN adduser $USER && \
    echo $USER:$USER | chpasswd && \
    adduser $USER sudo
ENV HOME="/home/$USER"

# Install Go
ENV PATH=$PATH:$HOME/go/bin
RUN wget https://dl.google.com/go/go${GO_VERSION}.linux-arm64.tar.gz && \
    tar -C $HOME -xzf go${GO_VERSION}.linux-arm64.tar.gz && \
    go install golang.org/x/tools/gopls@latest

# Install dagger
RUN curl -fsSL https://dl.dagger.io/dagger/install.sh | \
    DAGGER_VERSION=${DAGGER_VERSION} BIN_DIR=$HOME/.local/bin sh

# Grant permissions to non-root
RUN mkdir -p $HOME/.config/vscode-dev-containers && \
    chown -R $USER:$USER $HOME