package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kelseyhightower/envconfig"
	"github.com/ymohl-cl/fizzbuzz/api"
	"github.com/ymohl-cl/gopkg/server"
)

const (
	// tableableFormat constant to use to display usage in a tabular format
	tableFormat = `
			KEY	TYPE	DEFAULT	REQUIRED	DESCRIPTION
			{{range .}}{{usage_key .}}	{{usage_type .}}	{{usage_default .}}	{{usage_required .}}	{{usage_description .}}
			{{end}}`
)

// usage create a tabwriter instance to support table output like ask by envconfig
// Example: https://github.com/kelseyhightower/envconfig/blob/master/usage.go
func usage() error {
	var err error
	tabs := tabwriter.NewWriter(os.Stdout, 1, 0, 4, ' ', 0)

	// parameters usage
	fmt.Printf("Usage ./%s\n", *appName)
	flag.PrintDefaults()
	fmt.Println()

	// environment usage
	fmt.Println("This application is configured via the environment")
	fmt.Println()
	fmt.Println("To configure the server the following environment variables can be used:")
	if err = envconfig.Usagef(*appName, &server.Config{}, tabs, tableFormat); err != nil {
		return err
	}
	tabs.Flush()
	fmt.Println()
	fmt.Println("To configure postgres the following environment variables can be used:")
	if err = envconfig.Usagef(*appName, &api.Config{}, tabs, tableFormat); err != nil {
		return err
	}
	tabs.Flush()

	// happy close
	fmt.Println()
	fmt.Println("Enjoy !")
	return nil
}
