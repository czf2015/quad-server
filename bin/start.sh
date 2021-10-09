#!/bin/bash

source "$0/../env.sh"

go build -o $PROJECT_NAME

PROCESSES=`ps aux -P | grep "$PROJECT_NAME" | sed -e 's/ubuntu *//' | cut -d" " -f 1`
for i in $PROCESSES;do
    kill $i
done

cd $PROJECT_PATH
./$PROJECT_NAME &