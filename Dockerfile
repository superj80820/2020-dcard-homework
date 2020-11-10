# build stage
FROM golang:1.14.6-alpine AS build-env

RUN apk add --no-cache git
RUN git config --global url."https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.com/".insteadOf https://gitlab.com/
RUN go env -w GOPRIVATE=gitlab.com/${CI_PROJECT_NAMESPACE}

ADD . /src
RUN cd /src/cmd && go build -o app

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/cmd/app /app/
ENTRYPOINT ./app