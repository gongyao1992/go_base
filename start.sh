#!/bin/sh
echo "开始执行"

runName="xx_draw_webserver"
tport=1235

# 获取进程id
pid=$(ps aux | grep "./$runName" | grep -v "grep" | awk '{print $2}')

# 启动 关闭
if [ $# -eq 0 ]
then
    zhiling="start"
else
    zhiling=$1
fi

# switch语句
case $zhiling in
"start") echo "开始(start|restart)中..."
  # shellcheck disable=SC2071
  if [ "$pid"x != ""x ]
    then
      kill -9 $pid
      echo "kill success $pid"
  fi
  # 执行
  cd ./server/ && go build -o $runName
  nohup ./$runName --port $tport &

  echo "loading success"
  ;;
"stop") echo "关闭(stop)中..."
  if [ "$pid"x != ""x ]
    then
      kill -9 $pid
      echo "kill success $pid"
  else
      echo "没有可关闭的进程"
  fi
  ;;
*) echo "参数错误[start:开始|stop:结束]"
esac