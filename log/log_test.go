package log

import "testing"

func TestLog(t *testing.T) {
	log := Logger()
	log.Debugf(`I'm Debug%s`, "f")
	log.Info(`I'm Info`)
	Logger().Notice(`I'm Notice`)
	Logger().Warning(`I'm Warning`)
	L().Error(`I'm Error`)
	L().Critical(`I'm Critical`)
	//log.Fatal(`I'm Fatal`) /// os.Exit(1)
	//log.Panic(`I'm Panic`) /// panic() inside
}
