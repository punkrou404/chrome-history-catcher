# Chrome-history-catcher

ただchrome履歴を出力するためだけのコマンドです。

- ページタイトル
- URL
- 訪問日時

がtsv形式で出力されます

# Prerequisites

下記以上のバージョンであれば基本問題ないはず

```
# Google Chrome
81.0.4044.138 (Official Build) （64 ビット）

# Golang
$ go version
go version go1.14.2 linux/amd64

# Bash
$ sh --version
GNU bash, バージョン 5.0.16(1)-release (x86_64-pc-linux-gnu)
```

# Getting Started

```
git clone https://github.com/punkrou404/chrome-history-catcher.git
cd chrome-history-catcher
sh ./init.sh 
# 以降はウィザードに従ってセットアップしてください。
# /usr/local/bin/chc に構築されます
```

# Chrome-history-catcher

ただchrome履歴を出力するためだけのコマンドです。

- ページタイトル
- URL
- 訪問日時

がtsv形式で出力されます

# Prerequisites

下記以上のバージョンであれば基本問題ないはず

```
# Google Chrome
81.0.4044.138 (Official Build) （64 ビット）

# Golang
$ go version
go version go1.14.2 linux/amd64

# Bash
$ sh --version
GNU bash, バージョン 5.0.16(1)-release (x86_64-pc-linux-gnu)
```

# Getting Started

```
git clone https://github.com/punkrou404/chrome-history-catcher.git
cd chrome-history-catcher
sh ./init.sh 
# 以降はウィザードに従ってセットアップしてください。
# /usr/local/bin/chc に構築されます
```

# Usecase

あとは好きにつかってください

```
# "www.sejuku.net"以外の履歴を見たい時
chc | grep -v www.sejuku.net | less

# 最新10件だけ見たい時
chc | head -n 10

# タイトルとURLだけみたいとき(重複なし)
chc | awk -F'[\t]' '{print $1 $2}' | uniq 
```
