package main

import "fmt"

type IPlayer interface {
	Play()
}

type TapePlayer struct {
	Song      string
	Batteries string
}

func (t TapePlayer) Play() {
	fmt.Println("Playing", t.Song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

type TapeRecorder struct {
	Song        string
	Microphones int
}

func (t TapeRecorder) Play() {
	fmt.Println("Playing", t.Song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func play(p IPlayer) {
	p.Play()
}

type MyError string

func (m MyError) Error() string {
	return string(m)
}

func main() {
	var player IPlayer

	player = TapePlayer{Song: "tpSong"}
	player.Play()
	play(player)

	player = TapeRecorder{Song: "trSong"}
	player.Play()
	play(player)

	tapeR, ok := player.(TapeRecorder)

	fmt.Println(ok)
	tapeR.Stop()

	fmt.Println(MyError("My error"))
}
