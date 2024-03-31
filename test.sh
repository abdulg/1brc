#!/usr/bin/env bash

for f in samples/*.txt
do
  testFile=$(basename -s .txt "$f")
  echo "${testFile}"
  go run main.go -file samples/"${testFile}".txt > test.out && diff test.out samples/"${testFile}".out
done
