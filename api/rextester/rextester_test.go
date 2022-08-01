package rextester

import (
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"sync"
	"telcompiler/global"
	"testing"
	"time"
)

func init() {
	global.Client = &fasthttp.Client{
		Name:                     "telCompiler_test",
		NoDefaultUserAgentHeader: true,
		TLSConfig:                &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost:          5000,
		MaxIdleConnDuration:      5 * time.Second,
	}
}

func TestAPIRequest(t *testing.T) {
	tests := []struct {
		LanguageCode int
		Code         string
	}{
		{
			LanguageCode: 6,
			Code: `
						#include  <stdio.h>
						int main(void)
						{
							printf("Hello, world!\n");
							return 0;
						}
					`,
		},
		{
			LanguageCode: 20,
			Code: `
						package main  
						import "fmt"
						func main() { 
							for {
								fmt.Printf("hello, world\n") 
								}
						}
					`,
		},
		{
			LanguageCode: 7,
			Code: `
						#include <iostream>
						int main()
						{
							std::cout << "Hello, world!\n";
						}
					`,
		},
		{
			LanguageCode: 24,
			Code:         "print('hello world')",
		},
		{
			LanguageCode: 46,
			Code: `
						fn main() {
							println!("Hello, world!");
						}
					`,
		},
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(tests))
	for _, tt := range tests {
		go func(languageCode int, code string, wg *sync.WaitGroup) {
			result, err := RequestCompileCode(languageCode, code)
			if err != nil {
				t.Error(err)
			}
			t.Log(result)
			wg.Done()
		}(tt.LanguageCode, tt.Code, wg)
	}
	wg.Wait()
}

func TestNormalizeStats(t *testing.T) {
	test := &rexTesterRequest{
		LanguageChoiceWrapper: 20,
		Program: `
						package main  
						import "fmt"
						func main() { 
							fmt.Printf("hello, world\n") 
						}
					`,
		CompilerArgs: goArgs,
	}
	resp, err := apiRequest(test)
	if err != nil {
		t.Error(err)
	}
	res := &Result{}
	res.NormalizeStats(resp.Stats)
	t.Log(res.Stats)
}
