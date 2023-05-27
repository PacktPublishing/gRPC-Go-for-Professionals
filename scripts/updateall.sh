# This script is for maintenance of the projects.
# It updates all the dependencies for the different
# go projects.

source ./scripts/constants.sh

pwd=$(pwd)

for nb in ${CHAPTERS_NB[@]}
do
	if [ -d "$pwd/chapter$nb" ]
	then
		echo "enter $pwd/chapter$nb"
		cd $pwd/chapter$nb
		find proto -type f -name "*.proto" -exec protoc -Iproto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {} ";"

		for dir in server proto client
		do
			echo "enter $dir"
			cd $pwd/chapter$nb/$dir
			go get -u ./...
			go mod tidy
			bazel run //:gazelle-update-repos
		done
	fi
done