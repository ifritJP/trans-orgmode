package main

import (
	"io"
	"net/http"
	"strings"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

func Req(_env *LnsEnv, url, method string, headerMap *LnsMap, body string) (LnsAny, LnsAny) {
	client := http.Client{}
	inBody := strings.NewReader(body)
	request, err := http.NewRequest(method, url, inBody)
	if err != nil {
		return err.Error(), nil
	}
	for key, val := range headerMap.Items {
		request.Header.Add(key.(string), val.(string))
	}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	outBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error(), nil
	}

	respHeaderMap := map[LnsAny]LnsAny{}
	for key, val := range resp.Header {
		respHeaderMap[key] = val[0]
	}

	respObj := NewHttpIF_Response(
		_env, resp.StatusCode, NewLnsMap(respHeaderMap), string(outBody))

	return nil, respObj
}
