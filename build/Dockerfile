# -----
# BUILD
# -----

FROM golang:1.24.4-alpine3.22 AS build

# set up workdir
WORKDIR /go/src

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# compile app
COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal
RUN go build -o ./app ./cmd/app/main.go

# ---
# RUN
# ---

FROM alpine:3.22 AS run

WORKDIR /app

ENV SHELL=/bin/sh

# copy compiled app and swagger files
COPY --from=build /go/src/app .

# run app
CMD ["/bin/sh", "-c", "/app/app"]
