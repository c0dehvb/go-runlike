package utils

import (
	"encoding/json"
	"fmt"
	"github.com/c0dehvb/go-runlike/runlike/utils/exec"
	"strings"
)

//func ContainerInspect(containerName string) (types.ContainerJSON, error) {
//	result := types.ContainerJSON{}
//	if strings.TrimSpace(containerName) == "" {
//		return result, fmt.Errorf("containerName不能为空")
//	}
//	cmd := exec.New().Command("docker", "inspect", containerName)
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		return result, fmt.Errorf("%v, %s", err, strings.TrimSpace(string(output)))
//	}
//	var containers []types.ContainerJSON
//	err = json.Unmarshal(output, &containers)
//	if err != nil || len(containers) == 0 {
//		return result, err
//	}
//	return containers[0], nil
//}

func ContainerInspectMap(containerName string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	if strings.TrimSpace(containerName) == "" {
		return result, fmt.Errorf("containerName不能为空")
	}
	cmd := exec.New().Command("docker", "inspect", containerName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.HasPrefix(string(output), "[]\n") {
			output = []byte(string(output)[3:])
		}
		return result, fmt.Errorf("%s", strings.TrimSpace(string(output)))
	}
	var containers []map[string]interface{}
	err = json.Unmarshal(output, &containers)
	if err != nil || len(containers) == 0 {
		return result, err
	}
	return containers[0], nil
}
