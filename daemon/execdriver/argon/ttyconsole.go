// +build windows

package argon

import (
	//"os/exec"

	//log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/daemon/execdriver"
	//	"github.com/docker/docker/pkg/hcsshim"
	//"gopkg.in/natefinch/npipe.v2"
)

type TtyConsole struct {
	//	MasterPty *os.File
}

func NewTtyConsole(processConfig *execdriver.ProcessConfig, pipes *execdriver.Pipes) (*TtyConsole, error) {
	//ptyMaster, ptySlave, err := pty.Open()
	//if err != nil {
	//	return nil, err
	//}

	tty := &TtyConsole{
	//		MasterPty: ptyMaster,
	}

	//if err := tty.AttachPipes(&processConfig.Cmd, pipes); err != nil {
	//	tty.Close()
	//	return nil, err
	//}

	//processConfig.Console = tty.MasterPty.Name()

	return tty, nil
}

func (t *TtyConsole) Resize(h, w int) error {
	//return term.SetWinsize(t.MasterPty.Fd(), &term.Winsize{Height: uint16(h), Width: uint16(w)})
	return nil
}

func (t *TtyConsole) Close() error {
	//	t.SlavePty.Close()
	//	return t.MasterPty.Close()
	return nil
}
