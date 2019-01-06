package main

import (
	"github.com/golang/glog"
	"github.com/minio/minci/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}
