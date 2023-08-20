#!/bin/bash
topic=$1
if [ -z $topic ];then
	echo "please provide topic name"
	echo "$0 topic"
	exit 1
fi
docker exec -it broker kafka-topics --create --topic $topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
