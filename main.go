package main

import (
	"fmt"
	"os"

	"github.com/mackerelio/mackerel-client-go"
)

func printfStderr(format string, v ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func confirm() bool {
	var response string
	_, _ = fmt.Scanln(&response)

	if response == "y" {
		return true
	}

	return false
}

func main() {

	envName := "MACKEREL_APIKEY"
	apikey := os.Getenv(envName)
	if apikey == "" {
		printfStderr("ERROR: %s is not set\n", envName)
		os.Exit(1)
	}

	serviceNames := os.Args[1:]
	if len(serviceNames) == 0 {
		printfStderr("ERROR: Args (Mackerel Service Names) is not set\n")
		os.Exit(1)
	}

	mkr := mackerel.NewClient(apikey)

	org, err := mkr.GetOrg()
	if err != nil {
		printfStderr("ERROR: GetOrg: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Organization: %s\n", org.Name)
	fmt.Print("ok? [y/N]: ")
	if !confirm() {
		printfStderr("user canceled\n")
		os.Exit(1)
	}

	for _, serviceName := range serviceNames {
		fmt.Printf("%s\n", serviceName)
		_, err := mkr.DeleteService(serviceName)
		if err != nil {
			printfStderr("ERROR: DeleteService: %v\n", err)
			os.Exit(1)
		}
	}
}

