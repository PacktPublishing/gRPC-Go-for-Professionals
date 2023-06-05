"""Defines the dependencies' versions.

This allow us to keep the book generic and update cheaply the versions.
"""

# Go
# The last version can be found here: https://go.dev/dl/
# Note: You do not need to download the last version,
#       Just update this variable. Bazel will download it for you.
GO_VERSION = "1.20.4"

# Rules_go
# The last version and SHA256 can be found here: https://github.com/bazelbuild/rules_go/releases
RULES_GO_VERSION = "v0.39.1"
RULES_GO_SHA256 = "6dc2da7ab4cf5d7bfc7c949776b1b7c733f05e56edc4bcd9022bb249d2e2a996"

# Gazelle
# The last version and SHA256 can be found here: https://github.com/bazelbuild/bazel-gazelle/blob/master/README.rst
GAZELLE_VERSION = "v0.30.0"
GAZELLE_SHA256 = "727f3e4edd96ea20c29e8c2ca9e8d2af724d8c7778e7923a854b2c80952bc405"

# Protobuf
# The last version can be found here: https://github.com/protocolbuffers/protobuf/releases
# Note: You do not need to download the last version,
#       Just update this variable. Bazel will download it for you.
PROTO_VERSION = "v23.2"

# protoc-gen-validate
# The last version can be found here: https://github.com/bufbuild/protoc-gen-validate/releases
PROTOC_GEN_VALIDATE_VERSION = "v1.0.1"
