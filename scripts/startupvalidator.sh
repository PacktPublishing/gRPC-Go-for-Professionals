#!/bin/sh

# This script is used for testing that the application
# can be run with `go run` and with `bazel run`.
# Note: assert_contains was slightly modified but the original
# code comes from: https://github.com/torokmark/assert.sh

source ./scripts/constants.sh

__COLOR=false

while getopts ':c' 'OPTKEY'; do
	case ${OPTKEY} in
		'c')
			__COLOR=true
			;;
		'?')
			echo "INVALID OPTION -- ${OPTARG}" >&2
			exit 1
			;;
		':')
			echo "MISSING ARGUMENT for option -- ${OPTARG}" >&2
			exit 1
			;;
		*)
			echo "UNIMPLEMENTED OPTION -- ${OPTKEY}" >&2
			exit 1
			;;
	esac
done

if ${__COLOR} && command -v tput &>/dev/null && tty -s ; then
  RED=$(tput setaf 1)
  GREEN=$(tput setaf 2)
  NORMAL=$(tput sgr0)
elif ${__COLOR}; then
  RED=$(echo -en "\e[31m")
  GREEN=$(echo -en "\e[32m")
  NORMAL=$(echo -en "\e[00m")
else
  RED=""
  GREEN=""
  NORMAL=""
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

pwd=$(pwd)
for nb in ${CHAPTERS_NB[@]}
do
  if [ -d "$pwd/chapter$nb" ]
	then
    echo "enter $pwd/chapter$nb"
    cd $pwd/chapter$nb
    buf generate proto

    stderr=$(go run ./server 2>&1)
    assert_contain "$stderr" "usage: server" "check go run ./server"
    stderr=$(bazel run --ui_event_filters=-info,-stdout,-stderr --noshow_progress //server:server 2>&1)
    assert_contain "$stderr" "usage: server" "check bazel run //server:server"

    stderr=$(go run ./client 2>&1)
    assert_contain "$stderr" "usage: client" "check go run ./client"
    stderr=$(bazel run --ui_event_filters=-info,-stdout,-stderr --noshow_progress //client:client 2>&1)
    assert_contain "$stderr" "usage: client" "check bazel run //client:client"
  fi
done