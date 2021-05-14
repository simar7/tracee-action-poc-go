#!/bin/bash
set -e

while getopts "a:" o; do
   case "${o}" in
       a)
         export profile=${OPTARG}
       ;;
   esac
done

profile=$(echo $profile | tr -d '\r')
echo "$profile"

TRACEE_EBPF_EXE=${TRACEE_EBPF_EXE:="/tracee/tracee-ebpf"}
TRACEE_RULES_EXE=${TRACEE_RULES_EXE:="/tracee/tracee-rules"}

$TRACEE_EBPF_EXE trace -capture exec -capture profile