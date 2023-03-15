"""Defines the dependencies' versions.

This allow us to keep the book generic and update cheaply the versions.
"""

# Go
# The last version can be found here: https://go.dev/dl/
# Note: You do not need to download the last version,
#       Just update this variable. Bazel will download it for you.
GO_VERSION = "1.20.1"

# Rules_go
# The last version and SHA256 can be found here: https://github.com/bazelbuild/rules_go/releases
RULES_GO_VERSION = "v0.38.1"
RULES_GO_SHA256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d"

# Gazelle
# The last version and SHA256 can be found here: https://github.com/bazelbuild/bazel-gazelle/blob/master/README.rst
GAZELLE_VERSION = "v0.29.0"
GAZELLE_SHA256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d"

# Protobuf
# The last version can be found here: https://github.com/protocolbuffers/protobuf/releases
# Note: You do not need to download the last version,
#       Just update this variable. Bazel will download it for you.
PROTO_VERSION = "v21.12"