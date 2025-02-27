#!/bin/bash

PID=$(pgrep -f "demand-paging\.py")

if [ -z "${PID}" ]; then
    echo "demand-paging.pyプロセスが存在しないので$0 より先に起動。" >&2
    exit 1
fi

while true; do
    DATE=$(date | tr -d '\n')
    INFO=$(ps -h -o vsz,rss,maj_flt,min_flt -p ${PID})
    if [ $? -ne 0 ]; then
        echo "$DATE: demand-paging.pyプロセスは終了" >&2
        exit 1
    fi
    echo "${DATE}: ${INFO}"
    sleep 1
done