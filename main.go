package main

import (
	"fmt"
	"time"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli"
)

func main() {
	Do()

	fmt.Printf("Wait forever.\n")
	for {
		time.Sleep(1 * time.Second)
	}
}

func Do() {
	fmt.Printf("Entered...\n")
	c := cli.New()

	a := &action.Configuration{
		Capabilities: chartutil.DefaultCapabilities,
	}
	debug := func(msg string, args ...interface{}) {
		fmt.Printf(fmt.Sprintf("** %s\n", msg), args...)
	}
	a.Init(c.RESTClientGetter(), "default", "", debug)
	fmt.Printf("Initialized action.\n")

	act := action.NewInstall(a)
	act.Atomic = true
	act.Namespace = "default"
	act.ReleaseName = "foo"
	fmt.Printf("Created install action.\n")

	ch, err := loader.Load("/opt/charts/foo")
	if err != nil {
		fmt.Printf("failed to load chart from:%s\n", err.Error())
		return
	}
	fmt.Printf("Loaded chart\n")

	if _, err := act.Run(ch, nil); err != nil {
		fmt.Printf("helm install failed: %s\n", err.Error())
		return
	}
	fmt.Printf("Installed chart\n")
}
