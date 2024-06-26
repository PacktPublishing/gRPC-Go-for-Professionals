#!/bin/sh

# This script is for maintenance of the projects.
# It updates all the dependencies for the different
# go projects.

. ./scripts/constants.sh

pwd=$(pwd)

for nb in $CHAPTERS_NB
do
	if [ -d "$pwd/chapter$nb" ]
	then
		echo "enter $pwd/chapter$nb"
		cd $pwd/chapter$nb
		protoc -Iproto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative --validate_out="lang=go,paths=source_relative:proto" $(find . -type f -name "*.proto")

		for dir in server proto client
		do
			echo "enter $dir"
			cd $pwd/chapter$nb/$dir
			go get -u -d ./...
			go mod tidy
		done

		go work sync
		bazel run //:gazelle-update-repos
	fi
done

echo "enter $pwd/helpers"
cd $pwd/helpers
go get -u -d ./...
go mod tidy

echo "enter $pwd/proto"
cd $pwd/proto
go get -u -d ./...
go mod tidy
