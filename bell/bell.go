package bell

import (
	"os/exec"
)

func Toll() {
	cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/alarm-clock-elapsed.oga")
	cmd.Run()
}
