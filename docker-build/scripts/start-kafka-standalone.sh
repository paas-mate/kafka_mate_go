#!/bin/bash

if [ ${ZK_ADDR} ];then
    sed -i s/localhost:2181/${ZK_ADDR}/g $KAFKA_HOME/config/server.properties
else
    nohup $KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties >>$KAFKA_HOME/logs/zookeeper.stdout.log 2>>$KAFKA_HOME/logs/zookeeper.stderr.log &
fi
nohup $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties >>$KAFKA_HOME/logs/kafka.stdout.log 2>>$KAFKA_HOME/logs/kafka.stderr.log &
