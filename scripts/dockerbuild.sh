# This is supposed to be run inside a chapter folder

# If the buildx node is not created, uncomment the line below
# sudo docker buildx create  --name mybuild --driver=docker-container
sudo docker buildx build \
	--tag clementjean/packt-book:server \
	--file server/Dockerfile \
	--platform linux/arm64,linux/amd64 \
	--builder mybuild \
	--push .

sudo docker buildx build \
	--tag clementjean/packt-book:client \
	--file client/Dockerfile \
	--build-arg SERVER_ADDR="dns:///todo-server.default.svc.cluster.local:50051" \
	--platform linux/arm64,linux/amd64 \
	--builder mybuild \
	--push .