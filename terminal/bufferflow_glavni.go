package main

import (
	"log"
	//"net/http"
	//"encoding/hex"
	//"testing"
	//"regexp"
	//"strconv"
	//"time"
)



type BufferflowDefault struct {
	Name string
	Port string
}


var ()



func (b *BufferflowDefault) Init() {
	log.Println("Initting default buffer flow (which means no buffering)")
}

func (b *BufferflowDefault) RewriteSerialData(cmd string, id string) string {
	return ""
}

func (b *BufferflowDefault) BlockUntilReady(cmd string, id string) (bool, bool, string) {
	//log.Printf("BlockUntilReady() start\n")
	return true, false, ""
}

func (b *BufferflowDefault) OnIncomingData(data string) {
//log.Printf("OnIncomingData() start. data:%v\n", data)
	id := data

	var prvi_znak, zadnji_znak string = string([]rune(id)[0]), string([]rune(id)[13])

	log.Println("Deluje")
	log.Println("Id: ",id)
	if id == "\u0002070092BB6846\u0003"{	//prvi znak = 0002  Zadnji znak = 0003
		log.Println("Pozdravljen admin")
	}else{
		log.Println("Pozdravljen uporabnik")
	}

/*
	prvi_znak == string([]rune(id)[0])
	drugi_znak == string([]rune(id)[13])
*/


	log.Println("Prvi znak ", prvi_znak, ", zadnji znak ", zadnji_znak)

	
	
}

// Clean out b.sem so it can truly block
func (b *BufferflowDefault) ClearOutSemaphore() {
}

func (b *BufferflowDefault) BreakApartCommands(cmd string) []string {
	return []string{cmd}
}

func (b *BufferflowDefault) Pause() {
	return
}

func (b *BufferflowDefault) Unpause() {
	return
}

func (b *BufferflowDefault) SeeIfSpecificCommandsShouldSkipBuffer(cmd string) bool {
	return false
}

func (b *BufferflowDefault) SeeIfSpecificCommandsShouldPauseBuffer(cmd string) bool {
	return false
}

func (b *BufferflowDefault) SeeIfSpecificCommandsShouldUnpauseBuffer(cmd string) bool {
	return false
}

func (b *BufferflowDefault) SeeIfSpecificCommandsShouldWipeBuffer(cmd string) bool {
	return false
}

func (b *BufferflowDefault) SeeIfSpecificCommandsReturnNoResponse(cmd string) bool {
	return false
}

func (b *BufferflowDefault) ReleaseLock() {
}

func (b *BufferflowDefault) IsBufferGloballySendingBackIncomingData() bool {
	return false
}

func (b *BufferflowDefault) Close() {
}

func (b *BufferflowDefault) GetManualPaused() bool {
	return false
}

func (b *BufferflowDefault) SetManualPaused(isPaused bool) {
}
