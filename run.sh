#!bin/sh
export GIN_MODE=release
echo "start run.sh"
cd /root/go/src/course
./main
if [ $? -ne 0 ]; then
    echo "failed"
else
    echo "succeed"
fi