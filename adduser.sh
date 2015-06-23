#!/bin/bash

name=${1}
if [ "$name" = "" ]; then
    echo -n "username: "
    read name
fi
key=`date +%s`_$name
id=`echo "$key" | md5sum | awk '{print $1}'`
echo "INSERT INTO users (id, name, num) VALUES ('$id', '$name', 0);" | sqlite3 main.db
id=`echo "SELECT id FROM users WHERE name = '$name';" | sqlite3 main.db`

url="http://shootora.chimay.blue/vote?id=$id"
echo "$name様 ハイカットおじさん第二回引退界隈焼肉会参加可否シークレット投票先です．投票よろしくお願いいたします．$url"
