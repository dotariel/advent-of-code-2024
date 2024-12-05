#!/bin/bash

DAY=$1

if [ -z "${DAY}" ]; then
  echo "usage: new.sh <DD>"
  exit
fi

DIR="day-${DAY}"

mkdir -p ${DIR}
cd ${DIR}

go mod init dotariel.org/aoc2024

touch input.txt

# main.go
touch main.go
echo "package main"  >> main.go
echo ""              >> main.go
echo "func main() {" >> main.go
echo "}"             >> main.go

# main_test.go
touch main_test.go
echo "package main"  >> main_test.go

cd -