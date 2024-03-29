FROM mcr.microsoft.com/devcontainers/base:debian

ARG NX_VERSION=18.1.3
ARG NODE_VERSION=20.11.1

RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install vim -y

ARG USER=jacobley
RUN adduser $USER
USER $USER
ENV HOME="/home/${USER}"

# Install node
ENV NVM_DIR="$HOME/.nvm"
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/master/install.sh | bash && \
    # Setup nvm
    \. "$NVM_DIR/nvm.sh" && \
    # Install node
    nvm install $NODE_VERSION && \
    nvm use $NODE_VERSION

# Install packages
RUN \. "$NVM_DIR/nvm.sh" && \
    npm i -g "nx@${NX_VERSION}"