# This script is for maintenance of the projects.
# It cleans the different folders by removing
# generated code and bazel builds.
pwd=$(pwd)
for nb in 4 5 6
do
	if [ -d "$pwd/chapter$nb" ]
	then
		echo "enter $pwd/chapter$nb"
		cd $pwd/chapter$nb
		find proto -name "*.pb.go" -type f -exec rm -rf '{}' \;
		bazel clean --expunge
	fi
done