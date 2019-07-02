package scrlog

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Option struct {
	ForceColors     bool
	FullTimestamp   bool
	TimestampFormat string
}

type OptionItem struct {
	K string
	V interface{}
}

func FillOption(kvopt []OptionItem) (opt *Option, llv []log.Level, err error) {
	opt = &Option{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	llv = log.AllLevels

	for _, kv := range kvopt {
		switch kv.K {
		case "ForceColors":
			if v, ok := kv.V.(bool); ok {
				opt.ForceColors = v
			} else {
				return nil, nil, fmt.Errorf("reflect seg 'ForceColors' error: not bool value")
			}
		case "FullTimestamp":
			if v, ok := kv.V.(bool); ok {
				opt.FullTimestamp = v
			} else {
				return nil, nil, fmt.Errorf("reflect seg 'FullTimestamp' error: not bool value")
			}
		case "TimestampFormat":
			if v, ok := kv.V.(string); ok {
				opt.TimestampFormat = v
			} else {
				return nil, nil, fmt.Errorf("reflect seg 'TimestampFormat' error: not string value")
			}
		case "Level":
			if v, ok := kv.V.([]log.Level); ok {
				llv = v
			} else {
				return nil, nil, fmt.Errorf("reflect seg 'Level' error: not []logrus.Level value")
			}
		}

	}
	return
}
