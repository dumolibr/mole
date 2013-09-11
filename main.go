package main

import (
	"crypto/tls"
	"errors"
	"log"
	"os"
	"os/user"
	"path"
	"runtime"

	"github.com/jessevdk/go-flags"
)

var errParams = errors.New("incorrect command line parameters")

var (
	buildVersion string
	buildDate    string
	buildUser    string
	homeDir      string
)

var globalOpts struct {
	Debug bool `short:"d" long:"debug" description:"Show debug output"`
}

var globalParser = flags.NewParser(&globalOpts, flags.Default)

func main() {
	globalParser.ApplicationName = "mole"
	if _, e := globalParser.Parse(); e != nil {
		os.Exit(1)
	}
}

func setup() {
	globalOpts.Debug = globalOpts.Debug || isDebug

	if globalOpts.Debug {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		log.Println("Debug enabled")
	} else {
		log.SetFlags(0)
	}

	if globalOpts.Debug {
		printVersion()
	}

	u, e := user.Current()
	if e != nil {
		log.Fatal(e)
	}
	homeDir = path.Join(u.HomeDir, ".mole")
	if globalOpts.Debug {
		log.Println("homeDir", homeDir)
	}
}

func printVersion() {
	log.Printf("mole (%s-%s)", runtime.GOOS, runtime.GOARCH)
	if buildVersion != "" {
		log.Printf("  %s (%s)", buildVersion, buildKind)
	}
	if buildDate != "" {
		log.Printf("  %s by %s", buildDate, buildUser)
	}
}

func certificate() tls.Certificate {
	cert, err := tls.LoadX509KeyPair(path.Join(homeDir, "mole.crt"), path.Join(homeDir, "mole.key"))
	if err != nil {
		log.Fatal(err)
	}
	return cert
}