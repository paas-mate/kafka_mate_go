#!/bin/bash

nohup $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties >>$KAFKA_HOME/logs/kafka.stdout.log 2>>$KAFKA_HOME/logs/kafka.stderr.log &