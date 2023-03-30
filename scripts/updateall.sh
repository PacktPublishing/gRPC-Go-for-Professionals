# This script is for maintenance of the projects.
# It updates all the dependencies for the different
# go projects.
pwd=$(pwd)
for nb in 4 5 6
do
	echo "enter $pwd/chapter$nb"
	cd $pwd/chapter$nb
	find proto -type f -name "*.proto" -exec protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {} ";"

	for dir in server proto client
	do
		echo "enter $dir"
		cd $pwd/chapter$nb/$dir
		go get -u ./...
		go mod tidy
		bazel run //:gazelle-update-repos
	done
done