#!/usr/bin/env bash
mkdir ./day$1
cp ./template/day.go ./day$1/day$1.go 
cp ./template/day_test.go ./day$1/day$1_test.go
touch ./data/day$1.txt