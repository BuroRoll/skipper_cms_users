#!/usr/bin/env bash

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./pkg ./pkg
RUN chmod a+x /app/pkg/config/auth_model.conf
RUN chmod a+x /app/pkg/config/policy.csv
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app pkg/main.go
EXPOSE 8080
ENTRYPOINT ["./main"]

FROM scratch

WORKDIR /app

COPY --from=build /app/pkg/config/auth_model.conf /app/pkg/config/auth_model.conf
COPY --from=build /app/pkg/config/policy.csv /app/pkg/config/policy.csv
COPY --from=build /app/main ./main
EXPOSE 8080
ENTRYPOINT ["./main"]