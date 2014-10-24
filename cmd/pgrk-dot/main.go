package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func randomString() string {
	h := md5.New()
	io.WriteString(h, strconv.Itoa(int(time.Now().UnixNano())))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	fmt.Printf("digraph { \"%s\"\n", randomString())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 1 {
			continue
		}
		for _, f := range fields[1:] {
			fmt.Printf("    %s -> %s;\n", fields[0], f)
		}
	}
	fmt.Println("}")
}
