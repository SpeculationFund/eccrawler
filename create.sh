#!/bin/bash

Usage() {
    echo "Usage: $1 {exchange center name}" && exit
}


[ "$1" == "" ] && Usage $0

MARKET=$1
FILE=${MARKET,,}
echo "copying example"
cp ./crawler.goexample ./$FILE.go
echo "replace by exchange center name"
sed -i "s/_CRAWLER_/$MARKET/g" ./$FILE.go
echo "SUCCESS, Please follow TODO to modify your code"
