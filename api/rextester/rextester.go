package rextester

import (
	"encoding/json"
	"errors"
	"github.com/ja7ad/telcompiler/api/client"
	"strings"
)

const rexTesterAddress = "https://rextester.com/rundotnet/Run"

const (
	cArgs   = "-Wall -std=gnu99 -O2 -o a.out source_file.c"
	cppArgs = "-Wall -std=c++14 -O2 -o a.out source_file.cpp"
	goArgs  = "-o a.out source_file.go"
)

var statsList = strings.NewReplacer(
	"Absolute running time", "> مدت زمان اجرا",
	"cpu time", "> مدت پردازش",
	"memory peak", "> مقدار رم مصرف شده",
	"absolute service time", "> مدت استفاده از سرویس",
	"Compilation time", "> مدت زمان کامپایل",
	"Execution time", "> مدت زمان اجرا",
	"rows selected", "> ردیف های انتخاب شده",
	"rows affected", "> ردیف های تحت تاثیر",
	"sec", "ثانیه",
	"absolute running time", "> مدت زمان در حال اجرا",
	"Mb", "مگابایت",
	"Kb", "کیلوبایت",
	"Gb", "گیگابایت",
	", ", "\n",
)

type rexTesterRequest struct {
	LanguageChoiceWrapper int    `json:"LanguageChoiceWrapper"`
	Program               string `json:"Program"`
	CompilerArgs          string `json:"CompilerArgs"`
}

type rexTesterResponse struct {
	Errors       any `json:"Errors"`
	Result       any `json:"Result"`
	Stats        any `json:"Stats"`
	Warnings     any `json:"Warnings"`
	languageCode int `json:"languageCode"`
}

type Result struct {
	Language string `json:"language"`
	Errors   string `json:"errors"`
	Result   string `json:"result"`
	Stats    string `json:"stats"`
	Warnings string `json:"warnings"`
}

func RequestCompileCode(languageCode int, code string) (*Result, error) {
	if len(code) == 0 {
		return nil, errors.New("error code is empty")
	}
	reqBody := &rexTesterRequest{}
	switch languageCode {
	case 6:
		reqBody.LanguageChoiceWrapper = languageCode
		reqBody.Program = code
		reqBody.CompilerArgs = cArgs
	case 7:
		reqBody.LanguageChoiceWrapper = languageCode
		reqBody.Program = code
		reqBody.CompilerArgs = cppArgs
	case 20:
		reqBody.LanguageChoiceWrapper = languageCode
		reqBody.Program = code
		reqBody.CompilerArgs = goArgs
	default:
		reqBody.LanguageChoiceWrapper = languageCode
		reqBody.Program = code
	}

	resp, err := apiRequest(reqBody)
	if err != nil {
		return nil, err
	}
	result := &Result{
		Language: resp.Language(),
	}
	if resp.Errors != nil {
		result.Errors = resp.Errors.(string)
	}
	if resp.Result != nil {
		result.Result = resp.Result.(string)
	}
	if resp.Stats != nil {
		result.Stats = result.NormalizeStats(resp.Stats)
	}
	if resp.Warnings != nil {
		result.Warnings = resp.Warnings.(string)
	}
	return result, nil
}

func (r *Result) NormalizeStats(stats any) string {
	return statsList.Replace(stats.(string))
}

func (r *rexTesterResponse) Language() string {
	switch r.languageCode {
	case 20:
		return "Go"
	case 6:
		return "C"
	case 7:
		return "Cpp"
	case 24:
		return "Python"
	case 46:
		return "Rust"
	default:
		return "unknown"
	}
}

func apiRequest(body *rexTesterRequest) (*rexTesterResponse, error) {
	response := &rexTesterResponse{}
	resp, err := client.APIRequest(rexTesterAddress, body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resp, response); err != nil {
		return nil, err
	}
	response.languageCode = body.LanguageChoiceWrapper
	return response, nil
}
