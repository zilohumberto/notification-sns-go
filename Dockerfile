# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-notification-sns-go"
LABEL REPO="https://github.com/zilohumberto/notification-sns-go"

ENV PROJPATH=/go/src/github.com/zilohumberto/notification-sns-go

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/zilohumberto/notification-sns-go
WORKDIR /go/src/github.com/zilohumberto/notification-sns-go

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/zilohumberto/notification-sns-go"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/notification-sns-go/bin

WORKDIR /opt/notification-sns-go/bin

COPY --from=build-stage /go/src/github.com/zilohumberto/notification-sns-go/bin/notification-sns-go /opt/notification-sns-go/bin/
RUN chmod +x /opt/notification-sns-go/bin/notification-sns-go

# Create appuser
RUN adduser -D -g '' notification-sns-go
USER notification-sns-go

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/notification-sns-go/bin/notification-sns-go"]
