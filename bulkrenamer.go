package main

import (
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	"io/ioutil"
	"os"
	"path"
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

	targetDir := flag.Arg(0)

	log.Debug("Target directory is: " + targetDir)
	if targetDir == "" {
		log.Error("Target directory is REQUIRE")
	}

	execute(targetDir)

}

func execute(basepath string) {
	fileInfos, err := ioutil.ReadDir(basepath)
	if err != nil {
		log.Error(err)
	}
	i := 1
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			execute(basepath + "/" + fileInfo.Name())
		} else {
			// ignore
			if fileInfo.Name() == ".DS_Store" {
				log.Info(".DS_Store Skip")
				continue
			}
			if fileInfo.Name() == "Thumbs.db" {
				log.Info("Thumbs.db Skip")
				continue
			}

			ext := path.Ext(fileInfo.Name())
			index := fmt.Sprintf("%02d", i)
			basename := path.Base(basepath) + "_" + index + ext

			oldname := path.Join(basepath, fileInfo.Name())
			newname := path.Join(basepath, basename)
			log.Info("Rename: " + oldname + " -> " + newname)

			os.Rename(oldname, newname)
			i += 1
		}
	}
}

func optionParser() Option {
	o := Option{}
	o.IsVersion = flag.Bool("v", false, "Show version")
	o.IsAnsi = flag.Bool("ansi", true, "Enable Ansi color")
	o.LogLevel = flag.String("l", "debug", "Log level")
	o.LogDestination = flag.String("logdest", "./var/log/bulkrenamer.log", "Log destination path")

	flag.Parse()

	return o
}
