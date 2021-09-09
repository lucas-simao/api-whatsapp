package mensagem

import (
	"os/exec"
	"time"
)

func PostReceber(url, numero, texto string) {

	varEnviaMensagem := []string{
		"shell",
		"am",
		"start",
		"-a",
		"android.intent.action.VIEW",
		"'" + url + "'",
	}

	cmdEnviaMensagem := exec.Command("adb", varEnviaMensagem...)

	cmdEnviaMensagem.Run()

	time.Sleep(2 * time.Second)

	varPressEnter := []string{
		"shell",
		"input",
		"keyevent",
		"66",
	}

	cmdPressEnter := exec.Command("adb", varPressEnter...)

	cmdPressEnter.Run()

	PostStatus(numero, texto)
}
