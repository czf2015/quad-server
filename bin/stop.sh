#!/bin/bash

ENTRY="/home/czf/workspace/develop/revstream"

# 获取进程 pid
PIDS=`ps -ef | grep "$ENTRY" | grep -v "grep" | sort -k3 -nr | awk '{print $2}'`

# 杀死所有相关进程
echo "Stopping..."
for PID in $PIDS; do
    kill -9 $PID
done

# PROCESSES=`ps aux -P | grep nuxt | cut -d" " -f 4`
# for i in $PROCESSES;do
#     kill $i
# done