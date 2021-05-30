package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ecnepsnai/logtic"
	"github.com/ecnepsnai/startpage"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		startpage.Stop()
		os.Exit(1)
	}()
	processArgs()
	fmt.Printf("Startpage %s (Runtime %s). Copyright Ian Spence 2021 - All rights reserved.\n", startpage.StartpageVersion, runtime.Version())
	startpage.Start()
}

func printHelpAndExit() {
	fmt.Printf("Usage %s [options]\n\n", os.Args[0])
	fmt.Printf("Options:\n")
	fmt.Printf("-b --bind-addr <socket>     Specify the listen address for the web server\n")
	fmt.Printf("-s --static-dir <dir>       Specify the directory where static assets are located\n")
	fmt.Printf("-v --verbose                Set the log level to debug\n")
	os.Exit(1)
}

func processArgs() {
	if len(os.Args) == 1 {
		return
	}
	args := os.Args[1:]

	i := 0
	count := len(args)
	for i < count {
		arg := args[i]
		if arg == "-v" || arg == "--verbose" {
			logtic.Log.Level = logtic.LevelDebug
		} else if arg == "-b" || arg == "--bind-addr" {
			if i == count-1 {
				fmt.Fprintf(os.Stderr, "%s requires exactly 1 parameter\n", arg)
				printHelpAndExit()
			}

			value := args[i+1]
			startpage.HTTPBindAddress = value
			i++
		} else if arg == "-s" || arg == "--static-dir" {
			if i == count-1 {
				fmt.Fprintf(os.Stderr, "%s requires exactly 1 parameter\n", arg)
				printHelpAndExit()
			}

			value := args[i+1]
			startpage.StaticPath = value
			i++
		} else {
			printHelpAndExit()
		}
		i++
	}
}
