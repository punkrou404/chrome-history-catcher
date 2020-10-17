# Chrome-history-catcher

ただchrome履歴を出力するためだけのコマンドです。

- title
- url
- timestamp

がtsv形式で出力されます

# environment

- OS
    - linux
- dependency
    - chrome
    - golang
    - bash

# Setup

```
git clone https://github.com/punkrou404/chrome-history-catcher.git
cd chrome-history-catcher
sh ./init.sh 
# 以降はウィザードに従ってセットアップしてください。
# /usr/local/bin/chc に構築されます
```

# Usecase

```
# "www.sejuku.net"以外の履歴を見たい時
chc | grep -v www.sejuku.net > without.sehuku.history.list
```