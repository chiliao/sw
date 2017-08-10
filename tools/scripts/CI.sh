#!/bin/sh

warn () {
  echo "$@" >&2
}

die () {
  status="$1"
  shift
  warn "$@"
  exit "$status"
}

# verify vagrant cluster is in sane state
if ! tools/scripts/verify_cluster.sh; then
  die "failed to verify vagrant cluster"
fi

# build, test and verify coverage
if ! vagrant ssh -c "cd /import/src/github.com/pensando/sw; make" node1; then
    die 1 "Make failed"
fi

# create a pensando cluster
if ! vagrant ssh -c "cd /import/src/github.com/pensando/sw; make cluster" node1; then
    die 1 "Make cluster failed"
fi

# wait for few seconds to confirm cluster formed correctly
sleep 5

exit 0
