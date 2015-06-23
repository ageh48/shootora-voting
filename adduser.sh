#!/bin/bash

name=${1}
if [ "$name" = "" ]; then
    exit 1
fi
key=`date +%s`_$name
id=`echo "$key" | md5sum | awk '{print $1}'`
echo "INSERT INTO users (id, name, num) VALUES ('$id', '$name', 0);" | sqlite3 main.db > /dev/null 2>&1
id=`echo "SELECT id FROM users WHERE name = '$name';" | sqlite3 main.db`

url="http://shootora.chimay.blue/vote?id=$id"
echo "$name様 ハイカットおじさん第二回引退界隈焼肉会参加可否シークレット投票先です．投票内容は公開されません．素直に投票してください．では，投票よろしくお願いいたします． $url"
