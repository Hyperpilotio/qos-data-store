FROM alpine:3.4

COPY qos-data-store /opt/qos-data-store/qos-data-store
COPY documents/dev.config /opt/qos-data-store/dev.config

CMD ["/opt/qos-data-store/qos-data-store", "--config", "/opt/qos-data-store/dev.config"]