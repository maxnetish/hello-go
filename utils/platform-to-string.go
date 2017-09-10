package utils

import (
	"runtime"
	"fmt"
	"os"
	"strings"
)

type OsAndArch struct {
	Title string
	Arch  string
}

type Formattable interface {
	Format() string
}

func formatOsAndArch(inp *OsAndArch) string {
	return inp.Title + " " + inp.Arch
}

func (self *OsAndArch) Create() *OsAndArch {
	self.Title = runtime.GOOS
	self.Arch = runtime.GOARCH
	return self
}

func (self *OsAndArch) Format() (result string) {
	result = formatOsAndArch(self)
	return
}

func ReadEnvironmentVariables() (result map[string]string) {
	var oneEnv []string
	envSlice := os.Environ()
	result = make(map[string]string)

	for _, env := range envSlice {
		oneEnv = strings.Split(env, "=")
		result[oneEnv[0]] = oneEnv[1]
	}
	return
}

func AnotherPlatformInfo() (result string) {
	var localInfo OsAndArch
	var localFormattable Formattable

	localInfo.Create()
	localFormattable = &localInfo
	result = localFormattable.Format()
	return;
}

func PlatformInfo() (result string) {
	var osTitle string
	var parts [2]string
	var osArch string

	switch os := runtime.GOOS; os {
	case "darwin":
		osTitle = "OS X"
	case "linux":
		osTitle = "Linux"
	default:
		osTitle = fmt.Sprint(os)
	}

	osArch = fmt.Sprint(runtime.GOARCH)

	parts[0] = osTitle
	parts[1] = osArch

	for ind, part := range parts {
		result = result + part
		if ind < len(parts)-1 {
			result = result + " "
		}
	}

	return
}
