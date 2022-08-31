#!/bin/bash

declare	-a arrFound=("ian", "Ian", "iuiygaygn", "I d skd a efju N")
declare -a arrNotFound=("ihhhhhn", "ina", "xian")


echo "FOUND"
for i in "${arrFound[@]}"
do
  echo $i
  echo $i | go run findian.go
done


echo "NOT FOUND"
for i in "${arrNotFound[@]}"
do
  echo $i
  echo $i | go run findian.go
done

