// This code is transcompiled by LuneScript.
package main
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Util bool
var Util__mod__ string
// for 11
func Util_convExp0_76(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 10: decl @Util.readFile
func Util_readFile(_env *LnsEnv, path string) LnsAny {
    var fileObj Lns_luaStream
    
    {
        _fileObj := Util_convExp0_76(Lns_2DDD(Lns_io_open(path, nil)))
        if _fileObj == nil{
            return nil
        } else {
            fileObj = _fileObj.(Lns_luaStream)
        }
    }
    return fileObj.Read(_env, "*a")
}


func Lns_Util_init(_env *LnsEnv) {
    if init_Util { return }
    init_Util = true
    Util__mod__ = "@Util"
    Lns_InitMod()
    Lns_Types_init(_env)
}
func init() {
    init_Util = false
}
