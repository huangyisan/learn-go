package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

type Processes map[string]string

func parse(path string) (Processes, error) {
	fhandler, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer fhandler.Close()
	pses := make(Processes)
	reader := bufio.NewReader(fhandler)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineText := strings.TrimSpace(string(line))
		if strings.HasPrefix(lineText, "#") {
			continue
		}
		nodes := strings.SplitN(lineText, "=", 2)
		fmt.Println(nodes)
		if len(nodes) != 2 {
			continue
		}
		pses[nodes[0]] = nodes[1]

	}
	return pses, nil
}

func run(name, path string, stoped chan<- string) {
	log.Printf("run process %s[%s]", name, path)
	cmd := exec.Command(path)
	cmd.Start()
	cmd.Wait()
	log.Printf("process %s is stoped", name)
	// 结束后，则将name放入stopped管道中
	stoped <- name
}

func main() {
	pses, err := parse("process.conf")
	if err != nil {
		log.Fatal(err)
	}

	stoped := make(chan string)
	interrupt := make(chan os.Signal)
	// 监听SIGTERM和SIGINT信号，如果监听到，则放入interrupt管道中。
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	// 启动进程
	for name, path := range pses {
		go run(name, path, stoped)
	}

	END:
		for {
			select {
			// 阻塞，当interrupt管道有数据的时候，break END
			case <-interrupt:
				break END
			// 阻塞，当stoped管道有数据的时候，重启进程
			case name := <-stoped:
				go run(name, pses[name], stoped) // 重启进程
			}
		}
	}
