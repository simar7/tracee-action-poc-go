package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/sethvargo/go-githubactions"
)

//const (
//	TraceeBPFExe = "/tracee/tracee-ebpf"
//)

func startTracee() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "aquasecurity/tracee:latest", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "tracee",
		Cmd:   []string{"help"},
		Tty:   false,
	}, &container.HostConfig{Privileged: true}, nil, nil, "tracee")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

func main() {
	args := githubactions.GetInput("args")
	profileMode := githubactions.GetInput("profile")
	failOnDiff := githubactions.GetInput("fail-on-diff")
	//createPR := githubactions.GetInput("create-pr")

	fmt.Println("profileMode: ", profileMode, "failOnDiff: ", failOnDiff)
	fmt.Println("args: ", args)

	// run tracee in a docker container
	startTracee()

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
