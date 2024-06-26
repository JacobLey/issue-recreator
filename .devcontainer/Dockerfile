FROM mcr.microsoft.com/devcontainers/base:debian

ARG NODE_VERSION=20.11.1
ARG PNPM_VERSION=9.0.2

RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install vim -y

ARG USER=issue-recreator
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
    nvm use $NODE_VERSION && \
    # Remove default npm
    npm uninstall -g npm

# Install pnpm
ENV PNPM_HOME="${HOME}/.local/share/pnpm"
ENV PATH="${PNPM_HOME}:${PATH}"
RUN \. "$NVM_DIR/nvm.sh" && \
    corepack enable && \
    corepack prepare "pnpm@${PNPM_VERSION}" --activate

# Install packages
RUN \. "$NVM_DIR/nvm.sh"