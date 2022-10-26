FROM shoothzj/base:go AS build
COPY . /opt/compile
WORKDIR /opt/compile/pkg
RUN go build -o kafka_mate .


FROM ttbb/kafka:nake

COPY docker-build /opt/kafka/mate

COPY --from=build /opt/compile/pkg/kafka_mate /opt/kafka/mate/kafka_mate

COPY config/kraft_server_original.properties /opt/kafka/config/kraft/server_original.properties
COPY config/server_original.properties /opt/kafka/config/server_original.properties

WORKDIR /opt/kafka

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/kafka/mate/scripts/start.sh"]
