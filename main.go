package main

import (
    "log"
    "io/ioutil"
	"os"
    "time"
    "github.com/faiface/beep/wav"
    "github.com/faiface/beep/speaker"
)

const PATH = "./audio-assets"
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
        file := audioFiles[RandomIntBetween(0, len(audioFiles))]
        log.Println(Stringify(file))
		playFile(file)
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
    // looping
    select {}
}



// package main
// import (
//     "os"
//     "time"
//     "github.com/faiface/beep/wav"
//     "github.com/faiface/beep/speaker"
// )
// func main() {
//     // open file
//     // underscore mean we are ignoring error
//     f, _ := os.Open("./audio-assets/kick.wav")
//     // decoding mp3 file
//     // 3 outputs
//     // stream , format and error
//     streamer, format, _ := wav.Decode(f)
//     // activate speakers
//     speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
//     // play
//     speaker.Play(streamer)
//     // looping
//     select {}
// }
