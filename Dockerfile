# To build the njuno image, just run:
# > docker build -t njuno .
#
# In order to work properly, this Docker container needs to have a volume that:
# - as source points to a directory which contains a config.yaml and firebase-config.yaml files
# - as destination it points to the /home folder
#
# Simple usage with a mounted data directory (considering ~/.njuno/config as the configuration folder):
# > docker run -it -v ~/.njuno/config:/home njuno njuno parse config.yaml firebase-config.json
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -v ~/.njuno/config:/home --name njuno njuno
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it njuno bash
#
# To exit the bash, just execute
# > exit
FROM golang:bullseye AS build-env

# Set working directory for the build
WORKDIR /go/src/github.com/forbole/njuno

# Add source files
COPY . .

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apt-get update
RUN apt-get install curl make git build-essential libssl-dev pkg-config clang -y && \
    make install

# Final image
FROM rust:bullseye as build-nomic

# Install ca-certificates
WORKDIR /home

# Install bash
RUN set -ex \
    && sed -i -- 's/# deb-src/deb-src/g' /etc/apt/sources.list \
    && apt-get update \
    && apt-get install git build-essential libssl-dev pkg-config clang -y
RUN git clone https://github.com/MonikaCat/nomic.git /home/nomic
WORKDIR /home/nomic
RUN cargo build --locked

# Copy over binaries from the build-env
FROM debian:bullseye-slim
#WORKDIR /njuno
ENV HOME=/home/forbole
COPY --from=build-env /go/bin/njuno /usr/bin/njuno
COPY --from=build-env /go/src/github.com/forbole/njuno/modules/staking/utils/validators_query.sh $HOME/.njuno/validators_query.sh
COPY --from=build-nomic /home/nomic/target/debug/nomic /usr/bin/nomic