package main

import (
	_ "encoding/json"
	_ "encoding/xml"
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	_ "io/ioutil"
	_ "net/http"
	_ "net/http/cookiejar"
	_ "net/url"
	"os"
	_ "strings"
)

// オプション情報
type Option struct {
	IsVersion      *bool
	IsAnsi         *bool   // ログ出力をAnsiカラーにするか
	LogLevel       *string // ログレベル
	LogDestination *string // ログ出力場所
}

func main() {
	o := optionParser()

	if *o.IsVersion {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	initLogger(o)
	defer log.Flush()

	log.Info(getVersion())

}

func optionParser() Option {
	o := Option{}
	o.IsVersion = flag.Bool("v", false, "Show version")
	o.IsAnsi = flag.Bool("ansi", true, "Enable Ansi color")
	o.LogLevel = flag.String("l", "debug", "Log level")
	o.LogDestination = flag.String("logdest", "./var/log/nicony.log", "Log destination path")

	flag.Parse()

	return o
}
