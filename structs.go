package main

import (

)

func NewAudioFile(path string, name string) AudioFile {
	return AudioFile{path+"/"+name}
}
type AudioFile struct {
	Fullpath string
}

func (this *AudioFile) stop() {

}


type AnyNumber interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint | uintptr | float32 | float64
}

type AnyT interface {
    bool | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint | uintptr | float32 | float64 | complex64 | complex128 | string | map[string]func() | interface{}
}
