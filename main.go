package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/linkchain/cmd"
	"github.com/linkchain/common/util/log"
)

func main() {
	var (
		loglevel = flag.Int("loglevel", 3, "log level")
	)
	flag.Parse()
	//init log
	log.Root().SetHandler(
		log.LvlFilterHandler(log.Lvl(*loglevel),
			log.StreamHandler(os.Stdout, log.TerminalFormat(true))))

	start := strings.Fields("start")

	cmd.RootCmd.SetArgs(start)
	cmd.RootCmd.Execute()

	/*time.Sleep(time.Duration(1) * time.Second)

	cmd.RootCmd.SetArgs(send)
	cmd.RootCmd.Execute()

	time.Sleep(time.Duration(1) * time.Second)

	mine := strings.Fields("mine")

	cmd.RootCmd.SetArgs(mine)
	cmd.RootCmd.Execute()

	time.Sleep(time.Duration(1) * time.Second)

	cmd.RootCmd.SetArgs(mine)
	cmd.RootCmd.Execute()*/
	/*mine := strings.Fields("mine")

	cmd.RootCmd.SetArgs(mine)
	cmd.RootCmd.Execute()*/

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			words := strings.Fields(text)

			cmd.RootCmd.SetArgs(words)
			cmd.RootCmd.Execute()
		}

	}
}
