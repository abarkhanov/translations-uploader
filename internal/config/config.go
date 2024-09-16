package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func PromtParameter(paramName string, required bool) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(fmt.Sprintf("Enter %s: ", paramName))
	v, err := reader.ReadString('\n')
	v = strings.TrimSpace(strings.TrimSuffix(v, "\n"))
	if err != nil {
		return "", err
	}

	if required == true && len(v) == 0 {
		return "", errors.New(fmt.Sprintf("%s is required parameter", paramName))
	}

	if required == false && len(v) == 0 {
		fmt.Println(fmt.Sprintf("%s is not set", paramName))
		return "", nil
	}

	return v, nil
}
