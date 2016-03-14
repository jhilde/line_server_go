if [ "$1" == "" ]; then
    echo "usage: run.sh file"
else
    go run line_server.go $1
fi
