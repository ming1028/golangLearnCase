package main

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/aliyun/alibabacloud-nls-go-sdk"
)

func onTaskFailed(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}
	logger.Println("TaskFailed:", text)
}

func onStarted(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}

	logger.Println("onStarted:", text)
}

func onSentenceBegin(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}

	logger.Println("onSentenceBegin:", text)
}

func onSentenceEnd(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}

	logger.Println("onSentenceEnd:", text)
}

func onResultChanged(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}

	logger.Println("onResultChanged:", text)
}

func onCompleted(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}

	logger.Println("onCompleted:", text)
}

func onClose(params interface{}) {
	logger, ok := params.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}
	var err error
	logger.Println("onClosed:")
	st, err = nls.NewSpeechTranscription(config, logger,
		onTaskFailed, onStarted,
		onSentenceBegin, onSentenceEnd, onResultChanged,
		onCompleted, onClose, logger)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	test_ex := make(map[string]interface{})
	test_ex["test"] = "hello"

	ready, err := st.Start(param, test_ex)
	if err != nil {
		lk.Lock()
		fail++
		lk.Unlock()
		st.Shutdown()
		return
	}
	err = waitReady(ready, logger)
	if err != nil {
		lk.Lock()
		fail++
		lk.Unlock()
		st.Shutdown()
		return
	}
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

var lk sync.Mutex
var fail = 0
var reqNum = 0
var st *nls.SpeechTranscription
var config *nls.ConnectionConfig
var param nls.SpeechTranscriptionStartParam
var logger *nls.NlsLogger

func testMultiInstance() {
	pcm, err := os.Open("./tingwu/test1.pcm")
	if err != nil {
		log.Default().Fatalln(err)
	}

	buffers := nls.LoadPcmInChunk(pcm, 320)
	param = nls.DefaultSpeechTranscriptionParam()
	// param.Format = "mp3"
	//config := nls.NewConnectionConfigWithToken(PRE_URL_WSS,
	//        APPKEY, TOKEN)
	config, _ = nls.NewConnectionConfigWithAKInfoDefault(WSSURL, APPKEY, AKID, AKKEY)
	logger = nls.NewNlsLogger(os.Stderr, "1", log.LstdFlags|log.Lmicroseconds)
	logger.SetLogSil(false)
	logger.SetDebug(true)
	st, err = nls.NewSpeechTranscription(config, logger,
		onTaskFailed, onStarted,
		onSentenceBegin, onSentenceEnd, onResultChanged,
		onCompleted, onClose, logger)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	test_ex := make(map[string]interface{})
	test_ex["test"] = "hello"

	ready, err := st.Start(param, test_ex)
	if err != nil {
		lk.Lock()
		fail++
		lk.Unlock()
		st.Shutdown()
		return
	}

	err = waitReady(ready, logger)
	if err != nil {
		lk.Lock()
		fail++
		lk.Unlock()
		st.Shutdown()
		return
	}

	i := 0
	for _, data := range buffers.Data {
		if data != nil {
			st.SendAudioData(data.Data)
			if i == 5 {
				time.Sleep(20 * time.Second)
			} else {
				time.Sleep(time.Millisecond * 10)
			}
			i++
		}
	}

	logger.Println("send audio done")
	ready, err = st.Stop()
	if err != nil {
		lk.Lock()
		fail++
		lk.Unlock()
		st.Shutdown()
		return
	}

	err = waitReady(ready, logger)
	if err != nil {
		lk.Lock()
		fail++
		lk.Unlock()
		st.Shutdown()
		return
	}

	logger.Println("Sr done")
	st.Shutdown()
}

func main() {

	testMultiInstance()
}
