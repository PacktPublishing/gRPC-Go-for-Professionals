# This script is for maintenance of the projects.
# It cleans the different folders by removing
# generated code and bazel builds.

source ./scripts/constants.sh

pwd=$(pwd)

for nb in ${CHAPTERS_NB[@]}
do
	if [ -d "$pwd/chapter$nb" ]
	then
		echo "enter $pwd/chapter$nb"
		cd $pwd/chapter$nb
		find proto -name "*.pb.go" -type f -exec rm -rf '{}' \;
		bazel clean --expunge
	fi
done