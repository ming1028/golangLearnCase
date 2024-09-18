package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"time"

	nls "github.com/aliyun/alibabacloud-nls-go-sdk"
)

const ()

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

func onResultTranslated(text string, param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}
	logger.Println("onCustomHandler:", text)
}

func onClose(param interface{}) {
	logger, ok := param.(*nls.NlsLogger)
	if !ok {
		log.Default().Fatal("invalid logger")
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

var lk sync.Mutex
var fail = 0
var reqNum = 0

func testMultiInstance(num int) {
	pcm, err := os.Open("./tingwu/test1.pcm")
	if err != nil {
		log.Default().Fatalln(err)
	}

	buffers := nls.LoadPcmInChunk(pcm, 320)
	param := nls.DefaultSpeechTranscriptionParam()
	config, _ := nls.NewConnectionConfigWithAKInfoDefault("wss://tingwu-realtime-cn-shanghai.aliyuncs.com/api/ws/v1?mc=BmqoBXSX-8gBGoc8eK1r4JcySVeP1RoV69O_TGEAFN51dyQvl0PiXDYt2NXs8OcPRMQ7nZVhg-aNPFhNEVNbSmQ9r3bYzdkNLEMBho7XYq3gWaYk2IfgFxg2uMosOhdV",
		APPKEY, AKID, AKKEY)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			strId := fmt.Sprintf("ID%d   ", id)
			logger := nls.NewNlsLogger(os.Stderr, strId, log.LstdFlags|log.Lmicroseconds)
			logger.SetLogSil(false)
			logger.SetDebug(true)
			logger.Printf("Test Normal Case for SpeechRecognition:%s", strId)
			st, err := nls.NewSpeechTranscription(config, logger,
				onTaskFailed, onStarted,
				onSentenceBegin, onSentenceEnd, onResultChanged,
				onCompleted, onClose, logger)
			if err != nil {
				logger.Fatalln(err)
				return
			}
			st.SetCustomHandler("ResultTranslated", onResultTranslated)
			test_ex := make(map[string]interface{})
			test_ex["test"] = "hello"

			for {
				lk.Lock()
				reqNum++
				lk.Unlock()
				logger.Println("ST start")
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

				for _, data := range buffers.Data {
					if data != nil {
						fmt.Println(st.SendAudioData(data.Data))
						time.Sleep(10 * time.Millisecond)
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
		}(i)
	}

	wg.Wait()
}

func main() {
	go func() {
		log.Default().Println(http.ListenAndServe(":6060", nil))
	}()
	coroutineId := flag.Int("num", 1, "coroutine number")
	flag.Parse()
	log.Default().Printf("start %d coroutines", *coroutineId)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			lk.Lock()
			log.Printf(">>>>>>>>REQ NUM: %d>>>>>>>>>FAIL: %d", reqNum, fail)
			lk.Unlock()
			os.Exit(0)
		}
	}()
	testMultiInstance(*coroutineId)
}
