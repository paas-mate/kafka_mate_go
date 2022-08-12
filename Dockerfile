FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o kafka_mate .


FROM ttbb/kafka:nake

COPY docker-build /opt/sh/kafka/mate

COPY --from=build /opt/sh/compile/pkg/kafka_mate /opt/sh/kafka/mate/kafka_mate

COPY config/kraft_server_original.properties /opt/sh/kafka/config/kraft/server_original.properties
COPY config/server_original.properties /opt/sh/kafka/config/server_original.properties

WORKDIR /opt/sh/kafka

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/kafka/mate/scripts/start.sh"]
