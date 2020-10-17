FROM golang:1.14.10-stretch AS base
ENV LINT_VERSION=v1.27.0
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s ${GOLANG_CI_LINT_VERSION}
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
