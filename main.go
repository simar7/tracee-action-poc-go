package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/sethvargo/go-githubactions"

	"github.com/docker/docker/api/types/mount"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func startTracee(args []string) {
	log.Println("Starting Tracee...")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := "simar7/trcghaction:latest"
	_, err = cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	//io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   args,
	}, &container.HostConfig{
		Privileged: true,
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/proc",
				Target: "/proc",
			},
			{
				Type:   mount.TypeBind,
				Source: "/boot",
				Target: "/boot",
			},
			{
				Type:   mount.TypeBind,
				Source: "/lib/modules/",
				Target: "/lib/modules/",
			},
			{
				Type:   mount.TypeBind,
				Source: "/usr/src",
				Target: "/usr/src",
			},
			{
				Type:   mount.TypeBind,
				Source: "/tmp/tracee",
				Target: "/tmp/tracee",
			},
		},
	}, nil, nil, "tracee")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	log.Println("Tracee Started with container ID: ", resp.ID)
}
func stopTracee(failOnDiff string) {
	fmt.Println("Stopping Tracee...")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	var traceeCID string
	for _, c := range containers {
		if findStringinSlice("/tracee", c.Names) {
			traceeCID = c.ID
			fmt.Println("Stopping Tracee container: ", traceeCID)
			break
		}
	}

	if err := cli.ContainerKill(ctx, traceeCID, "SIGINT"); err != nil {
		panic(err)
	}

	n, _ := cli.ContainerWait(ctx, traceeCID, container.WaitConditionNotRunning)
	traceeExitCode := (<-n).StatusCode
	if strings.ToLower(failOnDiff) == "true" && traceeExitCode != 0 {
		b, _ := ioutil.ReadFile("/tmp/tracee/tracee.stdout")
		fmt.Println(string(b))
		fmt.Println("Tracee stopped with exit code: ", traceeExitCode)
		os.Exit(int(traceeExitCode))
	}
}

func findStringinSlice(needle string, haystack []string) bool {
	for _, n := range haystack {
		if n == needle {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(os.MkdirAll("/tmp/tracee", 0755))
	//profile := "start"
	//profile := "stop"
	profile := githubactions.GetInput("profile")
	switch strings.ToLower(profile) {
	case "start":
		startTracee([]string{"trace", "--output", "out-file:/tmp/tracee/tracee.stdout", "--capture", "exec", "--capture", "profile"})
	case "stop":
		//failOnDiff := githubactions.GetInput("fail-on-diff")
		failOnDiff := "true"
		stopTracee(failOnDiff)
	default:
		log.Fatalf("invalid option specified: %s, (valid options: start, stop)", profile)
	}
}
