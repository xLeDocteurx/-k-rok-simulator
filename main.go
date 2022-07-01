package main

import (
    "log"
    "io/ioutil"
	"os"
    "time"
    // "image"
    // "image/color"
    "github.com/faiface/beep"
    "github.com/faiface/beep/wav"
    "github.com/faiface/beep/speaker"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/data/binding"
)

const AUDIO_ASSETS_PATH = "./audio-assets"
const WINDOW_WIDTH = 200
const WINDOW_HEIGHT = 200
const BUTTONS_HEIGHT = 30

// FUN CONST TO PLAY WITH
const MIN_WAIT_IN_SECONDS = 2
const MAX_WAIT_IN_SECONDS = 15



func play(wantToPlay *bool) {
    (*wantToPlay) = true
}

func stop(wantToPlay *bool) {
    (*wantToPlay) = false
}

func main() {
    canPlay := true
    wantToPlay := false
    audioFiles := []AudioFile{}
    fyneApp := app.New()
    logoIcon, err := fyne.LoadResourceFromPath("./assets/logo.jpeg")
    if err != nil {
        log.Fatal(err)
    }
    fyneApp.SetIcon(logoIcon)
    mainWindow := fyneApp.NewWindow("PokRok")



    log.Println("App is starting")
    files, err := ioutil.ReadDir(AUDIO_ASSETS_PATH)
    if err != nil {
        log.Fatal(err)
    }
    for _, file := range files {
		audioFiles = append(audioFiles, NewAudioFile(AUDIO_ASSETS_PATH, file.Name()))
    }
    speaker.Init(44100, 4410)



    logoImage := canvas.NewImageFromFile("./assets/logo.jpeg")
	logoImage.FillMode = canvas.ImageFillStretch
    logoImage.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT)) 
    logoImage.Move(fyne.NewPos(0, 0)) 
    playButton := widget.NewButton("Play", func(){play(&wantToPlay)})
    playButton.Resize(fyne.NewSize(WINDOW_WIDTH / 2, BUTTONS_HEIGHT)) 
    playButton.Move(fyne.NewPos(0, WINDOW_WIDTH + 2))
    stopButton := widget.NewButton("Stop", func(){stop(&wantToPlay)})
    stopButton.Resize(fyne.NewSize(WINDOW_WIDTH / 2, BUTTONS_HEIGHT)) 
    stopButton.Move(fyne.NewPos(WINDOW_WIDTH / 2, WINDOW_WIDTH + 2))
	mainContainer := container.NewWithoutLayout(logoImage, playButton, stopButton)
    
    mainWindow.SetContent(mainContainer)
    mainWindow.SetMaster()
    mainWindow.Resize(fyne.NewSize(WINDOW_WIDTH + 8, WINDOW_HEIGHT + BUTTONS_HEIGHT + 8))
    mainWindow.CenterOnScreen()

    // playSomethingFromFilesList()
    go func() {
        for {
            if (canPlay && wantToPlay) {
                // canPlay = false
                audioFile := audioFiles[RandomIntBetween(0, len(audioFiles) - 1)]
                f, err := os.Open(audioFile.Fullpath)
                if err != nil {
                    log.Fatal(err)
                }
                // streamer, format, _ := wav.Decode(f)
                streamer, _, _ := wav.Decode(f)
                if err != nil {
                    log.Fatal(err)
                } 
                log.Println(audioFile.name)
                done := make(chan bool)
                speaker.Play(streamer, beep.Callback(func () {
                    done <- true
                    // streamer.Close()
                    // defer streamer.Close()
                }))
                <-done
                // defer speaker.Close()
                time.Sleep(time.Second * time.Duration(RandomIntBetween(MIN_WAIT_IN_SECONDS, MAX_WAIT_IN_SECONDS)))
                // canPlay = true
                // playSomethingFromFilesList(audioFiles, canPlay)
            }
            //  else {
            //     return
            // }

        }
    }()

	mainWindow.ShowAndRun()
    log.Println("App is exiting... Very much")
}

// func playSomethingFromFilesList(audioFiles []AudioFile, canPlay *bool) {
//     if (*canPlay) {
//         audioFile := audioFiles[RandomIntBetween(0, len(audioFiles) - 1)]
//         f, err := os.Open(audioFile.Fullpath)
//         if err != nil {
//             log.Fatal(err)
//         }
//         // streamer, format, _ := wav.Decode(f)
//         streamer, _, _ := wav.Decode(f)
//         if err != nil {
//             log.Fatal(err)
//         } 
//         log.Println(audioFile.name)
//         done := make(chan bool)
//         speaker.Play(streamer, beep.Callback(func () {
//             done <- true
//             // streamer.Close()
//             // defer streamer.Close()
//         }))
//         <-done
//         // defer speaker.Close()
//         time.Sleep(time.Second * time.Duration(RandomIntBetween(MIN_WAIT_IN_SECONDS, MAX_WAIT_IN_SECONDS)))
//         // playSomethingFromFilesList(audioFiles, canPlay)
//     }
//     //  else {
//     //     return
//     // }
// }
