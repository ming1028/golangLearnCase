package main

import (
	"errors"
	"log"
	"os"
	"time"

	nls "github.com/aliyun/alibabacloud-nls-go-sdk"
)

const (
	TOKEN  = "default"
	APPKEY = "default"
)

func onTaskFailed(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		return
	}

	logger.Println("TaskFailed:", text)
}

func onStarted(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		return
	}

	logger.Println("onStarted:", text)
}

func onSentenceBegin(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		return
	}

	logger.Println("onSentenceBegin:", text)
}

func onSentenceEnd(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		return
	}

	logger.Println("onSentenceEnd:", text)
}

func onResultChanged(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		return
	}

	logger.Println("onResultChanged:", text)
}

func onCompleted(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		return
	}

	logger.Println("onCompleted:", text)
}

func onClose(param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		return
	}

	logger.Println("onClosed:")
}

func waitReady(ch chan bool, logger *nls.NlsLogger) error {
	select {
	case done := <-ch:
		{
			if !done {
				logger.Println("Wait failed")
				return errors.New("wait failed")
			}
			logger.Println("Wait done")
		}
	case <-time.After(20 * time.Second):
		{
			logger.Println("Wait timeout")
			return errors.New("wait timeout")
		}
	}
	return nil
}

func run_push_audio_stream(url string) {
	// 使用本地pcm或者opus格式文件模拟真实场景下音频流实时采集
	pcm, err := os.Open("tingwu-sample-16k.pcm")
	if err != nil {
		log.Default().Fatalln(err)
	}

	buffers := nls.LoadPcmInChunk(pcm, 320)
	param := nls.DefaultSpeechTranscriptionParam()
	config := nls.NewConnectionConfigWithToken(url, APPKEY, TOKEN)
	logger := nls.NewNlsLogger(os.Stderr, "1", log.LstdFlags|log.Lmicroseconds)
	logger.SetLogSil(false)
	logger.SetDebug(true)
	st, err := nls.NewSpeechTranscription(config, logger,
		onTaskFailed, onStarted,
		onSentenceBegin, onSentenceEnd, onResultChanged,
		onCompleted, onClose, logger)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	logger.Println("Start pushing audio stream")
	ready, err := st.Start(param, nil)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	err = waitReady(ready, logger)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	for _, data := range buffers.Data {
		if data != nil {
			st.SendAudioData(data.Data)
			time.Sleep(10 * time.Millisecond)
		}
	}

	ready, err = st.Stop()
	if err != nil {
		logger.Fatalln(err)
		return
	}

	err = waitReady(ready, logger)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	st.Shutdown()
	logger.Println("Push audio stream done")
}

func main() {
	// 此处url来自于用户通过OpenAPI创建记录时返回的推流url
	run_push_audio_stream("wss://tingwu-realtime-cn-hangzhou-pre.aliyuncs.com/api/ws/v1?mc=*********h-moSNWGZO5mq-uZzu1EQbVBABVn9y8VGzWmVcAEiLNE1idoml7JU_wr17G4dDdxwQ6jiMg8OCQCrptlCnSk4hJ9K_fVfP8ngWaYk2If*********")
}
