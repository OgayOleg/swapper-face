package dto

import "encoding/json"

type Input struct {
	Target               string `json:"target"`
	Source               string `json:"source"`
	DisableSafetyChecker bool   `json:"disable_safety_checker"`
}

type FaceSwapReq struct {
	Version string `json:"version"`
	Input   Input  `json:"input"`
}

type ResponseID struct {
	Code   int `json:"code"`
	Result struct {
		TaskID string `json:"task_id"`
	}
}

type ResponseURL struct {
	Result struct {
		Error  json.RawMessage `json:"error"`
		Output json.RawMessage `json:"output"`
		Status string          `json:"status"`
	} `json:"result"`
}

func NewFaceSwapReq(target, source, version string) *FaceSwapReq {
	return &FaceSwapReq{
		Version: version,
		Input: Input{
			Target:               target,
			Source:               source,
			DisableSafetyChecker: true,
		},
	}
}
