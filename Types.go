// This code is transcompiled by LuneScript.
package main
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Types bool
var Types__mod__ string
// for 29
func Types_convExp0_135(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 12: decl @Types.LuaInStream.readStream
func (self *Types_LuaInStream) readStream(_env *LnsEnv, mode LnsAny)(LnsAny, string) {
    {
        _bin := self.stream.Read(_env, mode)
        if !Lns_IsNil( _bin ) {
            bin := _bin.(string)
            return bin, ""
        }
    }
    return nil, "err"
}
// 18: decl @Types.LuaInStream.read
func (self *Types_LuaInStream) Read(_env *LnsEnv, size LnsInt)(LnsAny, string) {
    return self.FP.readStream(_env, size)
}
// 21: decl @Types.LuaInStream.readAll
func (self *Types_LuaInStream) ReadAll(_env *LnsEnv)(LnsAny, string) {
    return self.FP.readStream(_env, "*a")
}
// 28: decl @Types.LuaOutStream.write
func (self *Types_LuaOutStream) Write(_env *LnsEnv, bin string) string {
    var err LnsAny
    _,err = self.stream.Write(_env, bin)
    if err != nil{
        err_45 := err.(string)
        return err_45
    }
    return ""
}
type Types_InStream interface {
        Read(_env *LnsEnv, arg1 LnsInt)(LnsAny, string)
        ReadAll(_env *LnsEnv)(LnsAny, string)
}
func Lns_cast2Types_InStream( obj LnsAny ) LnsAny {
    if _, ok := obj.(Types_InStream); ok { 
        return obj
    }
    return nil
}

type Types_OutStream interface {
        Write(_env *LnsEnv, arg1 string) string
}
func Lns_cast2Types_OutStream( obj LnsAny ) LnsAny {
    if _, ok := obj.(Types_OutStream); ok { 
        return obj
    }
    return nil
}

// declaration Class -- LuaInStream
type Types_LuaInStreamMtd interface {
    Read(_env *LnsEnv, arg1 LnsInt)(LnsAny, string)
    ReadAll(_env *LnsEnv)(LnsAny, string)
    readStream(_env *LnsEnv, arg1 LnsAny)(LnsAny, string)
}
type Types_LuaInStream struct {
    stream Lns_iStream
    FP Types_LuaInStreamMtd
}
func Types_LuaInStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Types_LuaInStream).FP
}
type Types_LuaInStreamDownCast interface {
    ToTypes_LuaInStream() *Types_LuaInStream
}
func Types_LuaInStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_LuaInStreamDownCast)
    if ok { return work.ToTypes_LuaInStream() }
    return nil
}
func (obj *Types_LuaInStream) ToTypes_LuaInStream() *Types_LuaInStream {
    return obj
}
func NewTypes_LuaInStream(_env *LnsEnv, arg1 Lns_iStream) *Types_LuaInStream {
    obj := &Types_LuaInStream{}
    obj.FP = obj
    obj.InitTypes_LuaInStream(_env, arg1)
    return obj
}
func (self *Types_LuaInStream) InitTypes_LuaInStream(_env *LnsEnv, arg1 Lns_iStream) {
    self.stream = arg1
}

// declaration Class -- LuaOutStream
type Types_LuaOutStreamMtd interface {
    Write(_env *LnsEnv, arg1 string) string
}
type Types_LuaOutStream struct {
    stream Lns_oStream
    FP Types_LuaOutStreamMtd
}
func Types_LuaOutStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Types_LuaOutStream).FP
}
type Types_LuaOutStreamDownCast interface {
    ToTypes_LuaOutStream() *Types_LuaOutStream
}
func Types_LuaOutStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_LuaOutStreamDownCast)
    if ok { return work.ToTypes_LuaOutStream() }
    return nil
}
func (obj *Types_LuaOutStream) ToTypes_LuaOutStream() *Types_LuaOutStream {
    return obj
}
func NewTypes_LuaOutStream(_env *LnsEnv, arg1 Lns_oStream) *Types_LuaOutStream {
    obj := &Types_LuaOutStream{}
    obj.FP = obj
    obj.InitTypes_LuaOutStream(_env, arg1)
    return obj
}
func (self *Types_LuaOutStream) InitTypes_LuaOutStream(_env *LnsEnv, arg1 Lns_oStream) {
    self.stream = arg1
}

func Lns_Types_init(_env *LnsEnv) {
    if init_Types { return }
    init_Types = true
    Types__mod__ = "@Types"
    Lns_InitMod()
}
func init() {
    init_Types = false
}
