package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/toolkits/pkg/runner"
	"initCake/conf"
	"initCake/pkg/ctx"
	"initCake/pkg/httpx"
	"initCake/pkg/logx"
	"initCake/pkg/osx"
	"initCake/pkg/version"
	"initCake/router"
	"initCake/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	showVersion = flag.Bool("version", false, "show version.")
	configDir   = flag.String("configs", osx.GetEnv("DAGGER_CONFIGS", "etc"), "Specify configuration directory.(env:DAGGER_CONFIGS)")
	cryptoKey   = flag.String("crypto-key", "", "Specify the secret key for configuration file field encryption.")
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(version.Version)
	}
	printEnv()
	cleanFunc, err := initialize(*configDir, *cryptoKey)
	if err != nil {
		log.Fatalln("failed to initialize:", err)
	}
	code := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

EXIT:
	for {
		sig := <-sc
		fmt.Println("received signal:", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			code = 0
			break EXIT
		case syscall.SIGHUP:
			// reload configuration?
		default:
			break EXIT
		}
	}

	cleanFunc()
	fmt.Println("process exited")
	os.Exit(code)
}

func printEnv() {
	runner.Init()
	fmt.Println("runner.cwd:", runner.Cwd)
	fmt.Println("runner.hostname:", runner.Hostname)
	fmt.Println("runner.fd_limits:", runner.FdLimits())
	fmt.Println("runner.vm_limits:", runner.VMLimits())
}

func initialize(configDir string, cryptoKey string) (func(), error) {
	config, err := conf.InitConfig(configDir, cryptoKey)
	logxClean, err := logx.Init(config.Log)
	if err != nil {
		return nil, err
	}
	// todo
	db, err := storage.New(config.DB)
	//db, err := storage.New(config.DB)
	//
	//if err != nil {
	//	return nil, err
	//}
	ctx := ctx.NewContext(context.Background(), db)
	rt := router.New(config.HTTP, ctx)
	r := httpx.GinEngine(config.Global.RunMode, config.HTTP)
	rt.Config(r)
	httpClean := httpx.Init(config.HTTP, r)
	return func() {
		logxClean()
		httpClean()
	}, nil
}
