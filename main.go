package main

import (
    "log"
    "io/ioutil"
	"os"
    "time"
    "github.com/faiface/beep"
    "github.com/faiface/beep/wav"
    "github.com/faiface/beep/speaker"
)

const PATH = "./audio-assets"
// const MIN_WAIT_IN_SECONDS = 5
const MIN_WAIT_IN_SECONDS = 1
// const MAX_WAIT_IN_SECONDS = 30
const MAX_WAIT_IN_SECONDS = 3

var audioFiles = []AudioFile{}

func main() {
    files, err := ioutil.ReadDir(PATH)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
		audioFiles = append(audioFiles, NewAudioFile(PATH, file.Name()))
    }

    log.Println("START")
    speaker.Init(44100, 4410)

    playSomethingFromFilesList()
    log.Println("END")
}

func playSomethingFromFilesList() {
    log.Println("long before the storm")
    audioFile := audioFiles[RandomIntBetween(0, len(audioFiles) - 1)]
    f, err := os.Open(audioFile.Fullpath)
    if err != nil {
        log.Fatal(err)
    }
    // log.Println("2")
    streamer, format, _ := wav.Decode(f)
    if err != nil {
        log.Fatal(err)
    } 
    // log.Println("3")
    log.Println(format.SampleRate, format.SampleRate.N(time.Second/10))
    log.Println("------")
    done := make(chan bool)
    log.Println("before the storm")
    speaker.Play(streamer, beep.Callback(func () {
        done <- true
    }))
    <-done
    streamer.Close()
    log.Println("TICK")
    time.Sleep(time.Second * time.Duration(RandomIntBetween(MIN_WAIT_IN_SECONDS, MAX_WAIT_IN_SECONDS)))
    log.Println("TOCK")
    playSomethingFromFilesList()
	defer speaker.Close()
}
