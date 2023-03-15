# This script is used for testing that the application
# can be run with `go run` and with `bazel run`.
# Note: assert_contains was slightly modified but the original
# code comes from: https://github.com/torokmark/assert.sh

if command -v tput &>/dev/null && tty -s; then
  RED=$(tput setaf 1)
  GREEN=$(tput setaf 2)
  NORMAL=$(tput sgr0)
else
  RED=$(echo -en "\e[31m")
  GREEN=$(echo -en "\e[32m")
  NORMAL=$(echo -en "\e[00m")
fi

log_success() {
  printf "${GREEN}✔ %s${NORMAL}\n" "$@" >&2
}

log_failure() {
  printf "${RED}✖ %s${NORMAL}\n" "$@" >&2
}

assert_contain() {
  local haystack="$1"
  local needle="${2-}"
  local msg="${3-}"

  if [ -z "${needle:+x}" ]; then
    return 0;
  fi

  if [ -z "${haystack##*$needle*}" ]; then
    [ "${#msg}" -gt 0 ] && log_success "$haystack contains $needle :: $msg" || true
    return 0
  else
    [ "${#msg}" -gt 0 ] && log_failure "$haystack doesn't contain $needle :: $msg" || true
    return 1
  fi
}

stderr=$(go run server/main.go 2>&1)
assert_contain "$stderr" "usage: server" "check go run server/main.go"
stderr=$(bazel run --ui_event_filters=-info,-stdout,-stderr --noshow_progress //server:server 2>&1)
assert_contain "$stderr" "usage: server" "check bazel run //server:server"

stderr=$(go run client/main.go 2>&1)
assert_contain "$stderr" "usage: client" "check go run client/main.go"
stderr=$(bazel run --ui_event_filters=-info,-stdout,-stderr --noshow_progress //client:client 2>&1)
assert_contain "$stderr" "usage: client" "check bazel run //client:client"