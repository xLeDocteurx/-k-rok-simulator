package main

import (
    "log"
    "io/ioutil"
	"os"
    "time"
    "github.com/faiface/beep/wav"
    "github.com/faiface/beep/speaker"
)

const PATH = "./assets"
const MIN_WAIT_IN_SECONDS = 5
const MAX_WAIT_IN_SECONDS = 30

var audioFiles = []AudioFile{}

func main() {
    files, err := ioutil.ReadDir(PATH)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
		audioFiles = append(audioFiles, NewAudioFile(PATH, file.Name()))
    }

    if err != nil {
        log.Fatal(err)
    }

	for {
		playFile(audioFiles[RandomIntBetween(0, len(audioFiles))])
		time.Sleep(time.Second * time.Duration(RandomIntBetween(MIN_WAIT_IN_SECONDS, MAX_WAIT_IN_SECONDS)))
	}

}

func playFile(audioFile AudioFile) {
    f, err := os.Open(audioFile.Fullpath)
    if err != nil {
        log.Fatal(err)
    }
    streamer, format, _ := wav.Decode(f)
    if err != nil {
        log.Fatal(err)
    } 
	defer streamer.Close()
    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
    speaker.Play(streamer)
    // // looping
    // select {}
}
