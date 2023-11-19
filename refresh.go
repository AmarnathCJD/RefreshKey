package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gen2brain/beeep"
)

var wd, _ = os.Getwd()

var AVAILABLE_REFRESH_RATES = []string{"60", "144"} // Change this to your available refresh rates

func main() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		var ref string = getCurrentRefreshRate()
		var toRef string = AVAILABLE_REFRESH_RATES[0]
		for _, v := range AVAILABLE_REFRESH_RATES {
			if !strings.Contains(ref, v) {
				toRef = v
				break
			}
		}

		csrPath := filepath.Join(wd, "bin", "csr.exe")

		cmd = exec.Command("cmd", "/C", csrPath, fmt.Sprintf("/f=%s", toRef), "/d=0")
		if err := cmd.Run(); err != nil {
			fmt.Println("Error changing refresh rate: " + err.Error())
			return
		}
		ref = getCurrentRefreshRate()
		sendNotification(ref)
	}
}

func getCurrentRefreshRate() string {
	cmd := exec.Command("cmd", "/C", "wmic", "PATH", "Win32_videocontroller", "get", "currentrefreshrate", "/value")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "unknown"
	}

	outputString := string(out)
	lines := strings.Split(outputString, "\n")
	return strings.TrimSpace(strings.SplitN(lines[5], "=", 2)[1]) + "Hz"
}

func sendNotification(ref string) {
	err := beeep.Alert("Refresh Rate Changed", "Refresh Rate: "+ref, filepath.Join(wd, "assets", "refresh.png"))
	if err != nil {
		panic(err)
	}
}
