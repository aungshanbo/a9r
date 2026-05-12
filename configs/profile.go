package configs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Loadconfig() [] string{
	var profiles [] string
	homeDir ,err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home, err")
	}

	configFile := filepath.Join(homeDir,".aws","config")

	file , err := os.Open(configFile)
	if err != nil {
		fmt.Println("Error getting config file",err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "[") && strings.HasSuffix( line, "]") {
			line = strings.Trim(line,"[]")
			line = strings.TrimPrefix(line,"profile ")
			profiles = append(profiles,line)

		}
	}
	return profiles
}