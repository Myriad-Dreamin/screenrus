package scrlog

import (
	"io"
	"os"

	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
)

type ScreenLogPlugin struct {
	Formatter log.Formatter
	Out       io.Writer

	sllv []log.Level
}

func NewScreenLogPlugin(kvopt []OptionItem) (slp *ScreenLogPlugin, err error) {
	slp = new(ScreenLogPlugin)

	opt, llv, err := FillOption(kvopt)
	if err != nil {
		return nil, err
	}
	slp.Formatter = &log.TextFormatter{
		ForceColors:     opt.ForceColors,
		FullTimestamp:   opt.FullTimestamp,
		TimestampFormat: opt.TimestampFormat,
	}
	slp.sllv = llv
	slp.Out = ansicolor.NewAnsiColorWriter(os.Stdout)
	return
}

func (slp *ScreenLogPlugin) Fire(entry *log.Entry) error {
	if entry.Level > entry.Logger.Level {
		return nil
	}

	dataBytes, err := slp.Formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = slp.Out.Write(dataBytes)

	return err
}

func (slp *ScreenLogPlugin) Levels() []log.Level {
	return slp.sllv
}
