package main

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"log"
)

type TranscodeingParam struct {
	TargetAudioFormat     string `json:"TargetAudioFormat,omitempty"`
	TargetVideoFormat     string `json:"TargetVideoFormat,omitempty"`
	VideoThumbnailEnabled bool   `json:"VideoThumbnailEnabled,omitempty"`
	SpectrumEnabled       bool   `json:"SpectrumEnabled,omitempty"`
}

type DiarizationParam struct {
	SpeakerCount int `json:"SpeakerCount,omitempty"`
}

type TranscriptionParam struct {
	AudioEventDetectionEnabled bool              `json:"AudioEventDetectionEnabled,omitempty"`
	DiarizationEnabled         bool              `json:"DiarizationEnabled,omitempty"`
	Diarization                *DiarizationParam `json:"Diarization,omitempty"`
	OutputLevel                int               `json:"OutputLevel,omitempty"`
}

type TranslationParam struct {
	TargetLanguages []string `json:"TargetLanguages,omitempty"`
}

type SummarizationParam struct {
	Types []string `json:"Types,omitempty"`
}

type ExtraParamerters struct {
	Transcoding              *TranscodeingParam  `json:"Transcoding,omitempty"`
	Transcription            *TranscriptionParam `json:"Transcription,omitempty"`
	TranslationEnabled       bool                `json:"TranslationEnabled,omitempty"`
	Translation              *TranslationParam   `json:"Translation,omitempty"`
	AutoChaptersEnabled      bool                `json:"AutoChaptersEnabled,omitempty"`
	MeetingAssistanceEnabled bool                `json:"MeetingAssistanceEnabled,omitempty"`
	SummarizationEnabled     bool                `json:"SummarizationEnabled,omitempty"`
	Summarization            *SummarizationParam `json:"Summarization,omitempty"`
	TextPolishEnabled        bool                `json:"TextPolishEnabled,omitempty"`
}

type InputParam struct {
	SourceLanguage string `json:"SourceLanguage"`
	FileUrl        string `json:"FileUrl,omitempty"`
	TaskKey        string `json:"TaskKey,omitempty"`
	Format         string `json:"Format,omitempty"`
	SampleRate     int    `json:"SampleRate,omitempty"`
}

type TaskBodyParam struct {
	Appkey      string            `json:"AppKey"`
	Input       InputParam        `json:"Input"`
	Paramerters *ExtraParamerters `json:"Parameters,omitempty"`
}

type CreateTaskResponse struct {
	RequestId string `json:"RequestId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	Data      struct {
		TaskId         string `json:"TaskId"`
		TaskKey        string `json:"TaskKey"`
		MeetingJoinUrl string `json:"MeetingJoinUrl,omitempty"`
	} `json:"Data"`
}

func init_request_param() *ExtraParamerters {
	param := new(ExtraParamerters)
	param.Transcoding = new(TranscodeingParam)

	transcription := new(TranscriptionParam)
	transcription.DiarizationEnabled = true
	transcription.Diarization = &DiarizationParam{
		SpeakerCount: 0,
	}
	transcription.OutputLevel = 2
	param.Transcription = transcription

	return param
}

func test_submit_realtime_meeting_task() string {
	client, err := sdk.NewClientWithAccessKey("cn-shanghai", akkey, aksecret)
	if err != nil {
		log.Default().Fatalln(err)
		return ""
	}

	request := requests.NewCommonRequest()
	request.Method = "PUT"
	request.Domain = "tingwu.cn-shanghai.aliyuncs.com"
	request.Version = "2023-09-30"
	request.SetContentType("application/json")
	request.PathPattern = "/openapi/tingwu/v2/tasks"
	request.QueryParams["type"] = "realtime"

	param := new(TaskBodyParam)
	param.Appkey = "9US2295hIkPhuVfJ"
	param.Input.SourceLanguage = "cn"
	param.Input.Format = "pcm"
	param.Input.SampleRate = 16000
	param.Paramerters = init_request_param()

	b, _ := json.Marshal(param)
	log.Default().Print("request body:\n", string(b))
	request.SetContent(b)
	request.SetScheme("https")

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		log.Default().Fatalln(err)
		return ""
	}

	log.Default().Print("response body:\n", string(response.GetHttpContentBytes()))

	var resp CreateTaskResponse
	err = json.Unmarshal(response.GetHttpContentBytes(), &resp)
	if err != nil {
		log.Default().Fatalln(err)
		return ""
	}

	log.Default().Println("TaskId:", resp.Data.TaskId)
	log.Default().Println("MeetingJoinUrl:", resp.Data.MeetingJoinUrl)

	return resp.Data.MeetingJoinUrl
}

func main() {
	test_submit_realtime_meeting_task()
}
