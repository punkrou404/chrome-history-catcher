#!/bin/bash
cd $(dirname $0)

echo "----------------------------"
echo "chcセットアップウィザードです"
echo "以下に従い、chromeの履歴ファイルの場所を特定してください"
echo "1. chromeを開いてください"
echo "2. アドレスバーに\"chrome://version/\"を入力してください"
echo "2. 「プロフィール パス」が表示されます。コピーして以下に入力してください。"
echo "----------------------------"
echo ""
read -p "ここにペーストしてください: " input

if [ -z "$input" ] ; then
    # undefined error
    echo "[ERROR]未入力です。やり直してください。"

elif [ -e "$input/History" ] ; then
    # succeed
    if [ ! -d "$HOME/.chc" ] ; then
        mkdir $HOME/.chc
        echo "セットアップ開始..."
    else
        rm -rf $HOME/.chc
        mkdir $HOME/.chc
        echo "再セットアップ開始..."
    fi
    echo $input/History | sudo tee $HOME/.chc/chc.config
    echo "" | sudo tee -a $HOME/.chc/chc.config
    echo "[Enter]押したらセットアップ開始"
    read Wait

    # setup
    make
    
else 
    # validation error
    echo "[ERROR]Chrome指定のパスではない可能性があります。やり直してください。"
    echo $input

fi