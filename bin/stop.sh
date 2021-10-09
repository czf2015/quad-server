#!/bin/bash

source "$0/../env.sh"

# 获取进程 pid
PIDS=`ps -ef | grep "$PROJECT_NAME" | grep -v "grep" | sort -k3 -nr | awk '{print $2}'`

# 杀死所有相关进程
echo "Stopping..."
for PID in $PIDS; do
    kill -9 $PID
done