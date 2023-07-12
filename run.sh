#!/bin/sh

CMD='go run github.com/zacharysyoung/SnippetBox'
PROC_NAME='SnippetBox' # (final part of) module name
HOST_PORT=':4000' # see hostPort var in main.go

startServer() {
    $CMD &
    sleep 1 && echo  # force terminal to refresh and get back to prompt
}

stopServer() {
    pkill $PROC_NAME
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
    lsof -n -i "$HOST_PORT"
    ;;
    *)
    echo "run start|stop|restart|status" && exit 1
esac