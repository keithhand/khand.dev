ARG GO_VERSION
ARG ALPINE_VERSION
ARG TEMPL_VERSION
ARG AIR_VERSION

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as template
WORKDIR /template
ARG TEMPL_VERSION
ENV TEMPL_VERSION=${TEMPL_VERSION}
RUN go install github.com/a-h/templ/cmd/templ@${TEMPL_VERSION}
COPY ./handlers ./handlers
RUN templ generate

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
COPY --from=template ./template .
RUN go build -v -o ./tmp/main ./cmd/server/server.go

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as develop
WORKDIR /develop
ARG AIR_VERSION
ENV AIR_VERSION=${AIR_VERSION}
RUN go install github.com/air-verse/air@${AIR_VERSION}
COPY --from=build ./build .
COPY --from=template  /go/bin/templ /go/bin/templ
ENTRYPOINT ["air"]

FROM alpine:${ALPINE_VERSION} as deploy
USER guest
COPY --from=build /build/tmp/main /srv/app/server
ENTRYPOINT ["/srv/app/server"]
