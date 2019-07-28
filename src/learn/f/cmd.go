package f

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func runCmd(cmd string) (out string, err error) {
	//command := strings.Fields(cmd)
	//fmt.Println(cmd)

	c := exec.Command("sh", "-c", cmd)
	//c := exec.Command("cat", "/sys/class/net/*/address")
	stderr := &bytes.Buffer{}
	c.Stderr = stderr

	output, err := c.Output()
	if err != nil {
		fmt.Printf("ERROR: %s %s", err, stderr)
		return
	}
	out = string(output)
	out = strings.Trim(out, " ")
	return
}

func tet() {
	cmd := "cat /sys/class/net/*/address"
	out, _ := runCmd(cmd)
	fmt.Println(out)
}
