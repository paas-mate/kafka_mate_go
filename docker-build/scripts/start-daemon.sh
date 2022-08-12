#!/bin/bash

mkdir -p $KAFKA_HOME/logs
nohup $KAFKA_HOME/mate/kafka_mate >>$KAFKA_HOME/logs/kafka_mate.stdout.log 2>>$KAFKA_HOME/logs/kafka_mate.stderr.log &

