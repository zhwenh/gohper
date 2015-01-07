package test

import (
	"github.com/cosiner/golib/types"
	"github.com/cosiner/gomodule/config"
	"testing"
)

func TestConf(t *testing.T) {
	cfg := config.NewConfig()
	cfg.ParseFile("app.conf")
	cfg.SetCurrSec("log")

	bufsize, _ := cfg.Val("bufsize")
	t.Log(types.Str2Bytes(bufsize))
	t.Log(cfg.Val("maxsize"))
	t.Log(cfg.Val("logdir"))
	t.Log(cfg.Val("level"))
	t.Log(cfg.Val("flushinterval"))
}