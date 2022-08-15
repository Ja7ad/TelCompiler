package client

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/ja7ad/telcompiler/global"
	"github.com/valyala/fasthttp"
	"log"
)

func APIRequest(url string, body any) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	b, err := json.Marshal(body)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("error on marshal body api request %v", err)
		return nil, err
	}
	req.SetBody(b)
	return apiResponse(req)
}

func apiResponse(request *fasthttp.Request) ([]byte, error) {
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)
	defer fasthttp.ReleaseRequest(request)
	if err := global.Client.Do(request, response); err != nil {
		sentry.CaptureException(err)
		log.Printf("error on api request %v", err)
		return nil, err
	}
	return response.Body(), nil
}
