package scrlog

import (
	"testing"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func TestLogDefalut(t *testing.T) {
	sss := log.New()
	a, err := NewScreenLogPlugin(nil)
	if err != nil {
		t.Error(err)
		return
	}
	sss.AddHook(a)
	ssr := sss.WithFields(log.Fields{
		"prog": "cmd",
	})
	ssr.Infof("love t")
	ssr.Errorf("love WW")
}

func TestLogLevel(t *testing.T) {
	sss := log.New()
	a, err := NewScreenLogPlugin([]OptionItem{
		OptionItem{K: "Level", V: []logrus.Level{log.WarnLevel, log.InfoLevel}},
	})
	if err != nil {
		t.Error(err)
		return
	}
	sss.AddHook(a)
	ssr := sss.WithFields(log.Fields{
		"prog": "cmd",
	})
	ssr.Infof("love t")
	ssr.Errorf("love WW")
}
