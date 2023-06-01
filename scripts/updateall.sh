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
		buf generate proto

		for dir in server proto client
		do
			echo "enter $dir"
			cd $pwd/chapter$nb/$dir
			go get -u ./...
			go mod tidy
		done
	fi
done

cd $pwd/chapter9 && bazel run //:gazelle-update-repos