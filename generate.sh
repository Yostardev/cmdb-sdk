#!/usr/bin/env bash
read_dir() {
  dir=$1
  files=$(\ls -a $dir)
  for file in $files; do
    file_path=$dir/$file
    if [ -d $file_path ]; then
      if [[ $file != '.' && $file != '..' ]]; then
        read_dir $file_path
      fi
    else
      if [ "${file##*.}"x = "proto"x ]; then
        protoc -I ../ -I ./ --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $file_path && protoc-go-inject-tag -input=$dir/*.pb.go
      fi
    fi
  done
}

for file in ./*; do
  if [ -d $file ]; then
    read_dir $file
  fi
done
