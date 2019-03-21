package warewulf

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func Nodes() ([]Node, error) {
	cmd := exec.Command("wwsh", "object", "jsondump", "-t", "node")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return []Node{}, err
	}

	var nodes []Node
	if err := json.Unmarshal(out.Bytes(), &nodes); err != nil {
		return []Node{}, err
	}

	return nodes, nil
}
