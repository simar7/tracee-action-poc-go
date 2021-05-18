#!/bin/bash
set -x

TRACEE_EBPF_EXE=${TRACEE_EBPF_EXE:="/tracee/tracee-ebpf"}
TRACEE_RULES_EXE=${TRACEE_RULES_EXE:="/tracee/tracee-rules"}

while getopts "a:" o; do
   case "${o}" in
       a)
         export profile=${OPTARG}
       ;;
   esac
done

profile=$(echo $profile | tr -d '\r')
echo "$profile"

if [ "$profile" = "start" ]; then
  $TRACEE_EBPF_EXE --output out-file:/tmp/tracee/tracee.stdout.log -capture exec -capture profile
elif [ "$profile" = "stop" ]; then
  docker ps -a
  out=$(docker kill --signal=SIGINT tracee-profiler)
  echo "killed $out"
  docker ps -a
#  trcid=$(docker inspect --format="{{.Id}}" tracee-profiler)
#  echo -e "POST /containers/$trcid/kill?signal=SIGINT HTTP/1.0\r\n" | nc -U /var/run/docker.sock
else
  $TRACEE_EBPF_EXE --output out-file:/tmp/tracee/tracee.stdout.log -capture exec -capture profile
fi