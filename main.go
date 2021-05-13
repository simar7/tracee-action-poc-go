package main

import (
	"fmt"

	"github.com/sethvargo/go-githubactions"
)

//const (
//	TraceeBPFExe = "/tracee/tracee-ebpf"
//)

func main() {
	args := githubactions.GetInput("args")
	profileMode := githubactions.GetInput("profile")
	failOnDiff := githubactions.GetInput("fail-on-diff")
	//createPR := githubactions.GetInput("create-pr")

	fmt.Println("profileMode: ", profileMode, "failOnDiff: ", failOnDiff)
	fmt.Println("args: ", args)

	// run tracee and capture exit code
	//var args []string
	//if profileMode != "" {
	//	args = append(args, "--capture=exec", "--capture=profile")
	//}

	//if err := exec.Command(TraceeBPFExe, args...).Run(); err != nil {
	//	log.Fatal(err)
	//}

	// fail if failondiff mode set

	// create pr if diff foundy
}
