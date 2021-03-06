package main

import (
	"flag"
	"github.com/iwittenberg/paneless/arrangements"
	"github.com/iwittenberg/paneless/handler"
	"github.com/iwittenberg/paneless/tray"
	"log"
)

func main() {
	p := flag.String("preferences", "preferences.json", "Preferences file path")
	s := flag.String("snapshot", "snapshot.json", "Snapshot file path")
	flag.Parse()

	as, err := arrangements.NewFromFile(*p)
	if err != nil {
		log.Fatal("Couldnt read from file", err)
	}

	handler.RegisterHotkeysAndListen(as)
	t := tray.GetInstance(as, p, s)
	t.Init()
}
