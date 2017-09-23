package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/profile"

	"github.com/malfanmh/go-shop/internal/model"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

var (
	name    = "go-shop"
	version = "unversioned"
	token   = name + "/" + version
	path_config = os.Getenv("GOSHOP_CONFIG")

	fversion = flag.Bool("version", false, "print the version.")
	fconfig  = flag.String("config", path_config, "set the config file path.")
	fprofile = flag.String("profile", "", "enable profiler, value either one of [cpu, mem, block].")

)

func init() {
	flag.Parse()

	if *fversion {
		printVersion()
	}
}

func main() {
	switch *fprofile {
	case "cpu":
		defer profile.Start(profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile).Stop()
	}

	if err := ReadConfig(*fconfig, &config); err != nil {
		log.Fatal(err)
	}

	//init database Connection
	if err := model.ConnectDB(config.Database.Endpoint); err != nil {
		log.Fatal(err)
	}

	//#### setup Route and middleware #######
	app := newApp()
	app.Use(recover.New())

	log.Printf("Server is listening on %q\n", config.Server.Host)

	err := app.Run(iris.Addr(config.Server.Host))
	if err != nil {
		panic(err)
	}
}

func printVersion() {
	fmt.Println(token)
	os.Exit(0)
}