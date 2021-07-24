package main

import (
	"encoding/json"
	"log"
	"strings"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

//import "fmt"

func ReadMap(_env *LnsEnv, inStream Types_InStream) LnsAny {
	txt, _ := inStream.ReadAll(_env)

	if txt != nil {
		return Txt2Map(_env, txt.(string))
	}
	return nil
}

func Txt2Map(_env *LnsEnv, txt string) LnsAny {
	jsonMap := map[string]LnsAny{}
	decoder := json.NewDecoder(strings.NewReader(txt))
	if err := decoder.Decode(&jsonMap); err != nil {
		log.Print(err)
		return nil
	}
	ret := Lns_mapFromGo(jsonMap)
	return ret
}

func WriteObj(_env *LnsEnv, outStream Types_OutStream, jsonMap *LnsMap) {
	outStream.Write(_env, Obj2Txt(_env, jsonMap))
}

func Obj2Txt(_env *LnsEnv, jsonMap *LnsMap) string {
	goMap := Lns_valToGo(jsonMap, true)

	bytes, _ := json.Marshal(goMap)
	return string(bytes)
}
