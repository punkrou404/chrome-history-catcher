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
    echo "なにも入力していないじゃないか？"

elif [ -e "$input/History" ] ; then
    # succeed
    if [ ! -d "$HOME/.chc" ] ; then
        mkdir $HOME/.chc
        echo "Start the initial setup."
    else
        rm -rf $HOME/.chc
        mkdir $HOME/.chc
        echo "Start the initial re-setup."
    fi
    echo $input/History | sudo tee $HOME/.chc/chc.config
    echo "" | sudo tee -a $HOME/.chc/chc.config
    echo "いいんじゃないか？"
    echo "Enter押したらSetupするぞ"
    read Wait

    # setup
    make
    
else 
    # validation error
    echo "ほんとにお前が入力したものはあってるか？"
    echo "一度見直して出直してこい"
    echo $input

fi