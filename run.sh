#!/bin/sh

HOST_PORT=':4000'
PROC_NAME='SnippetBox'


startServer() {
    echo "starting $PROC_NAME"
    go run ./cmd/web/$PROC_NAME -addr="$HOST_PORT" &
    sleep 1
}

stopServer() {
    pkill $PROC_NAME
    if [ $? = 1 ]; then
        echo "error: could not kill proc named $PROC_NAME; get more info with status command"
        exit 1
    else
        sleep 1
        echo "stopped $PROC_NAME"
    fi
}

getStatus() {
    # output will include header, which starts with 'COMMAND   PID...'
    output=$(lsof -n -i "$HOST_PORT" | grep 'COMMAND\|LISTEN')
    
    if [ $? = 1 ]; then
        echo "could not find anything listing on $HOST_PORT"
        exit 1
    fi

    lineCount=$(echo "$output" | wc -l | tr -d ' ')
    if [ "$lineCount" -gt 2 ]; then
        echo "warning: expected only one open file listening on $HOST_PORT"
        echo "-----"
        echo "$output"
        exit 1
    fi

    output=$(echo "$output" | grep -v COMMAND)
    cmd=$( echo "$output" | cut -w -f1)
    pid=$( echo "$output" | cut -w -f2)
    user=$(echo "$output" | cut -w -f3)
    name=$(echo "$output" | cut -w -f9-10)
    
    echo "COMMAND: $cmd"
    echo "PID:     $pid"
    echo "USER:    $user"
    echo "NAME:    $name"
}

case $1 in
    start)
    startServer
    ;;
    stop)
    stopServer
    ;;
    restart)
    stopServer
    startServer
    ;;
    status)
    getStatus
    ;;
    *)
    echo "run start|stop|restart|status" && exit 1
esac