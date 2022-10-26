#!/bin/bash

uuid=`bash $KAFKA_HOME/bin/kafka-storage.sh random-uuid`
bash $KAFKA_HOME/bin/kafka-storage.sh format -t $uuid -c $KAFKA_HOME/config/kraft/server.properties

nohup $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/kraft/server.properties >>$KAFKA_HOME/logs/kafka.stdout.log 2>>$KAFKA_HOME/logs/kafka.stderr.log &
