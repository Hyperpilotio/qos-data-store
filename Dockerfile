FROM lacion/docker-alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/hyperpilotio/qos-data-store"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/qos-data-store/bin

WORKDIR /opt/qos-data-store/bin

COPY documents/dev.config /opt/qos-data-store/bin/dev.config
COPY bin/qos-data-store /opt/qos-data-store/bin/
RUN chmod +x /opt/qos-data-store/bin/qos-data-store
RUN chmod 755 /opt/qos-data-store/bin/dev.config

CMD ["/opt/qos-data-store/bin/qos-data-store", "--config", "/opt/qos-data-store/bin/dev.config"]

