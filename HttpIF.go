// This code is transcompiled by LuneScript.
package main
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_HttpIF bool
var HttpIF__mod__ string
// declaration Class -- Response
type HttpIF_ResponseMtd interface {
    Get_body(_env *LnsEnv) LnsAny
    Get_header(_env *LnsEnv) *LnsMap
    Get_httpStatus(_env *LnsEnv) LnsInt
}
type HttpIF_Response struct {
    httpStatus LnsInt
    header *LnsMap
    body LnsAny
    FP HttpIF_ResponseMtd
}
func HttpIF_Response2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*HttpIF_Response).FP
}
type HttpIF_ResponseDownCast interface {
    ToHttpIF_Response() *HttpIF_Response
}
func HttpIF_ResponseDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(HttpIF_ResponseDownCast)
    if ok { return work.ToHttpIF_Response() }
    return nil
}
func (obj *HttpIF_Response) ToHttpIF_Response() *HttpIF_Response {
    return obj
}
func NewHttpIF_Response(_env *LnsEnv, arg1 LnsInt, arg2 *LnsMap, arg3 LnsAny) *HttpIF_Response {
    obj := &HttpIF_Response{}
    obj.FP = obj
    obj.InitHttpIF_Response(_env, arg1, arg2, arg3)
    return obj
}
func (self *HttpIF_Response) InitHttpIF_Response(_env *LnsEnv, arg1 LnsInt, arg2 *LnsMap, arg3 LnsAny) {
    self.httpStatus = arg1
    self.header = arg2
    self.body = arg3
}
func (self *HttpIF_Response) Get_httpStatus(_env *LnsEnv) LnsInt{ return self.httpStatus }
func (self *HttpIF_Response) Get_header(_env *LnsEnv) *LnsMap{ return self.header }
func (self *HttpIF_Response) Get_body(_env *LnsEnv) LnsAny{ return self.body }


func Lns_HttpIF_init(_env *LnsEnv) {
    if init_HttpIF { return }
    init_HttpIF = true
    HttpIF__mod__ = "@HttpIF"
    Lns_InitMod()
}
func init() {
    init_HttpIF = false
}
