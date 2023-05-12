"""Defines the dependencies' versions.

This allow us to keep the book generic and update cheaply the versions.
"""

# Go
# The last version can be found here: https://go.dev/dl/
# Note: You do not need to download the last version,
#       Just update this variable. Bazel will download it for you.
GO_VERSION = "1.20.2"

# Rules_go
# The last version and SHA256 can be found here: https://github.com/bazelbuild/rules_go/releases
RULES_GO_VERSION = "v0.39.0"
RULES_GO_SHA256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f"

# Gazelle
# The last version and SHA256 can be found here: https://github.com/bazelbuild/bazel-gazelle/blob/master/README.rst
GAZELLE_VERSION = "v0.29.0"
GAZELLE_SHA256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d"

# Protobuf
# The last version can be found here: https://github.com/protocolbuffers/protobuf/releases
# Note: You do not need to download the last version,
#       Just update this variable. Bazel will download it for you.
PROTO_VERSION = "v3.21.12"