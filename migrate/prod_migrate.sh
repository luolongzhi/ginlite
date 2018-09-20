#!/bin/sh 

action=$1
step=$2

echo "action: $action, step: $step"

migrate -path ./migrations -database "mysql://root:123456@tcp(127.0.0.1:3306)/ginlite" $action $step 
