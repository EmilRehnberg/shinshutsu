package alarmclock

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (
	printf  = fmt.Printf
	sprintf = fmt.Sprintf
)

// tolls the system bell on linux and runs a audio file in the repo on mac.
// not implemented for other OSes.
func Toll() {
	host := &environment{}
	bell := &hostBell{}
	tollBell(host, bell)
}

func tollBell(host operatingSystem, bell systemBell) {
	switch os := host.Os(); os {
	case "darwin":
		bell.macToll()
	case "linux":
		bell.linuxToll()
	default:
		bell.abortWithNotImpletedError(os)
	}
}

type operatingSystem interface {
	Os() string
}

type environment struct{}

func (e *environment) Os() string {
	return systemOs()
}

var systemOs = func() string {
	return runtime.GOOS
}

type Instruction struct{}

func (e *Instruction) Command(executable, soundFilePath string) CmdRunner {
	return &Executioner{
		instruction: buildCommand(executable, soundFilePath),
	}
}

var buildCommand = func(executable, soundFilePath string) CmdRunner {
	return exec.Command(executable, soundFilePath)
}

type systemBell interface {
	toll(Commander)
	macToll()
	linuxToll()
	abortWithNotImpletedError(string)
}

type hostBell struct {
	executable, soundFilePath string
}

func (bell *hostBell) toll(instruction Commander) {
	instruction.Command(bell.executable, bell.soundFilePath).Run()
}

type CmdRunner interface {
	Run() error
}

type Executioner struct {
	instruction CmdRunner
}

func (e *Executioner) Run() error {
	return e.instruction.Run()
}

type Commander interface {
	Command(string, string) CmdRunner
}

func (bell *hostBell) macToll() {
	bell.executable = "afplay"
	bell.soundFilePath = macAlarmPath()
	bell.toll(&Instruction{})
}

func (bell *hostBell) linuxToll() {
	bell.executable = "paplay"
	bell.soundFilePath = "/usr/share/sounds/freedesktop/stereo/alarm-clock-elapsed.oga"
	bell.toll(&Instruction{})
}

func macAlarmPath() string {
	return sprintf("%s/src/github.com/emilrehnberg/shinshutsu/media/alarm-clock.aiff", os.Getenv("GOPATH"))
}

func (b *hostBell) abortWithNotImpletedError(osName string) {
	printf("Bell is not implemeted for %v", osName)
	dirtyExit()
}

var dirtyExit = func() {
	os.Exit(1)
}
