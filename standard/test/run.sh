#!/bin/bash

echo "print"

read -p "请输入一个数字" num
if [[ "$num" =~ ^|0-9|+$ ]]; then
  echo $funtype
else
  echo "error"
fi

function list() {
    echo 1.执行一个基本的测试
}