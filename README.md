# Screen-rus     

easiest usage:

```go
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


>---------------------------------------------------
INFO[2019-07-03 01:02:15] love t                                        prog=cmd
time="2019-07-03T01:02:15+08:00" level=info msg="love t" prog=cmd
ERRO[2019-07-03 01:02:15] love WW                                       prog=cmd
time="2019-07-03T01:02:15+08:00" level=error msg="love WW" prog=cmd
```

set level:

```go
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


>---------------------------------------------------
INFO[2019-07-03 01:02:15] love t                                        prog=cmd
time="2019-07-03T01:02:15+08:00" level=info msg="love t" prog=cmd
time="2019-07-03T01:02:15+08:00" level=error msg="love WW" prog=cmd
```

following is all the options you can set:

| Key Name        | Value Type     | Defalut               |
| --------------- | -------------- | --------------------- |
| ForceColors     | bool           | true                  |
| FullTimestamp   | bool           | true                  |
| TimestampFormat | string         | "2006-01-02 15:04:05" |
| Level           | []logrus.Level | log.AllLevels         |

