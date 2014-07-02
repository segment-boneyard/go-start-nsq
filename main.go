package main

import . "github.com/visionmedia/go-gracefully"
import "github.com/visionmedia/go-flags"
import "strconv"
import "os/exec"
import "log"
import "os"

type Options struct {
	Count int `short:"n" default:"1" description:"number of nsqd nodes"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func run(quit chan bool, cmd string, args ...string) {
	log.Printf("exec %s %v", cmd, args)
	proc := exec.Command(cmd, args...)
	proc.Stderr = os.Stderr
	proc.Stdout = os.Stdout

	err := proc.Start()
	check(err)

	<-quit

	log.Printf("kill %s", cmd)
	err = proc.Process.Kill()
	check(err)
}

func main() {
	opts := &Options{}
	quit := make(chan bool)
	http := 5000
	tcp := 5001

	_, err := flags.Parse(opts)
	check(err)

	check(os.MkdirAll("/tmp/nsqd", 0755))

	for i := 0; i < opts.Count; i++ {
		httpArg := ":" + strconv.Itoa(http)
		tcpArg := ":" + strconv.Itoa(tcp)
		go run(quit, "nsqd", "--http-address", httpArg, "--tcp-address", tcpArg, "--lookupd-tcp-address", ":4160", "--data-path", "/tmp/nsqd")
		http += 2
		tcp += 2
	}

	go run(quit, "nsqadmin", "--lookupd-http-address", ":4161")
	go run(quit, "nsqlookupd")

	Shutdown()
	close(quit)
}
