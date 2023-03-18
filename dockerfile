FROM golang:1.19.6-alpine3.17

RUN apk update \
# Install git
    && apk add --no-cache --upgrade git \
# Install bash
    && apk add --no-cache --upgrade bash \
# Install make
    && apk add --no-cache make

# Get keeper src
WORKDIR /keeper
COPY . .
RUN ls -ltr
RUN \
    # Build keeper
    make build-server && chmod +x ./keeper-server


CMD ["./bash]"]
