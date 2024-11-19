package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bytedance/sonic"
)

var links map[string]string = nil
var linksUpdateTime time.Time
var linksUpdateLock sync.Mutex
var linksDecodeErr error = nil

func updateLinks() error {
	const name = "links.json"

	stat, err := os.Stat(name)
	if err != nil {
		return err
	}
	modTime := stat.ModTime()
	if modTime.Equal(linksUpdateTime) {
		return linksDecodeErr
	}

	linksUpdateLock.Lock()
	defer linksUpdateLock.Unlock()

	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	stat, err = f.Stat()
	if err != nil {
		return err
	}
	modTime = stat.ModTime()
	if modTime.Equal(linksUpdateTime) {
		return linksDecodeErr
	}

	log.Println("updateLinks")

	links = map[string]string{}
	err = sonic.ConfigDefault.NewDecoder(f).Decode(&links)
	if err != nil {
		err = fmt.Errorf("updateLinks error: %s", err)
	}
	linksDecodeErr = err
	linksUpdateTime = modTime
	return err
}
