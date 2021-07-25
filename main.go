// This code is transcompiled by LuneScript.
package main
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_main bool
var main__mod__ string
// decl enum -- Mode 
type main_Mode = string
const main_Mode__Github = "github"
const main_Mode__MkReq = "mkreq"
const main_Mode__Org = "org"
const main_Mode__Trans = "trans"
var main_ModeList_ = NewLnsList( []LnsAny {
  main_Mode__Org,
  main_Mode__MkReq,
  main_Mode__Trans,
  main_Mode__Github,
})
func main_Mode_get__allList_2_(_env *LnsEnv) *LnsList{
    return main_ModeList_
}
var main_ModeMap_ = map[string]string {
  main_Mode__Github: "Mode.Github",
  main_Mode__MkReq: "Mode.MkReq",
  main_Mode__Org: "Mode.Org",
  main_Mode__Trans: "Mode.Trans",
}
func main_Mode__from_1_(_env *LnsEnv, arg1 string) LnsAny{
    if _, ok := main_ModeMap_[arg1]; ok { return arg1 }
    return nil
}

func main_Mode_getTxt(arg1 string) string {
    return main_ModeMap_[arg1];
}
var main_LimitSize LnsInt
var main_LimitTxtNum LnsInt
// for 195
func main_convExp0_1266(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 308
func main_convExp0_1906(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 223
func main_convExp0_1406(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 224
func main_convExp0_1438(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 243
func main_convExp0_1592(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 244
func main_convExp0_1608(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 74
func main_convExp0_711(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 175
func main_convExp0_1145(arg1 []LnsAny) (*LnsList, *OrgDoc_TransCtrl) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsList), Lns_getFromMulti( arg1, 1 ).(*OrgDoc_TransCtrl)
}
// for 184
func main_convExp0_1193(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 217
func main_convExp0_1365(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 233
func main_convExp0_1495(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 236
func main_convExp0_1531(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 242
func main_convExp0_1577(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 339
func main_convExp0_2036(arg1 []LnsAny) (*LnsList, *OrgDoc_TransCtrl) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsList), Lns_getFromMulti( arg1, 1 ).(*OrgDoc_TransCtrl)
}
// 9: decl @main.doc2org
func main_doc2org_0_(_env *LnsEnv, doc OrgDoc_Document,transCtrl *OrgDoc_TransCtrl) {
    for _, _item := range( doc.GetItemList(_env).Items ) {
        item := _item.(OrgDoc_Item)
        var id LnsInt
        id = item.Get_Id(_env)
        if _switch0 := item.Get_Kind(_env); _switch0 == OrgDoc_ItemKind__Text {
            Lns_print([]LnsAny{OrgDoc_getTxt(_env, doc, id, transCtrl)})
        } else if _switch0 == OrgDoc_ItemKind__Headline {
            var headline OrgDoc_Headline
            headline = doc.GetHeadlineList(_env).GetAt(id).(OrgDoc_Headline)
            Lns_print([]LnsAny{_env.GetVM().String_format("\n\n%s %s", []LnsAny{_env.GetVM().String_rep("*", headline.Get_Level(_env)), OrgDoc_getTxt(_env, doc, headline.Get_TxtId(_env), transCtrl)})})
            {
                _customId := headline.Get_CustomId(_env)
                if !Lns_IsNil( _customId ) {
                    customId := _customId.(string)
                    Lns_print([]LnsAny{_env.GetVM().String_format(":PROPERTIES:\n:CUSTOM_ID: %s\n:END:\n", []LnsAny{customId})})
                }
            }
        } else if _switch0 == OrgDoc_ItemKind__Table {
            var table OrgDoc_Table
            table = doc.GetTableList(_env).GetAt(id).(OrgDoc_Table)
            var itemNum LnsInt
            itemNum = table.Get_RowList(_env).GetAt(1).(OrgDoc_TableRow).Get_TxtIdList(_env).Len()
            Lns_print([]LnsAny{_env.GetVM().String_format("|%s-|", []LnsAny{_env.GetVM().String_rep("-|", itemNum - 1)})})
            for _, _tableRow := range( table.Get_RowList(_env).Items ) {
                tableRow := _tableRow.(OrgDoc_TableRow)
                if tableRow.Get_TxtIdList(_env).Len() == 0{
                    Lns_print([]LnsAny{_env.GetVM().String_format("|%s-|", []LnsAny{_env.GetVM().String_rep("-+", itemNum - 1)})})
                } else { 
                    Lns_io_stdout.Write(_env, "| ")
                    for _, _textId := range( tableRow.Get_TxtIdList(_env).Items ) {
                        textId := _textId.(LnsInt)
                        Lns_io_stdout.Write(_env, OrgDoc_getTxt(_env, doc, textId, transCtrl))
                        Lns_io_stdout.Write(_env, " | ")
                    }
                    Lns_io_stdout.Write(_env, "\n")
                }
            }
        } else if _switch0 == OrgDoc_ItemKind__Verb {
            var verb OrgDoc_Verb
            verb = doc.GetVerbList(_env).GetAt(id).(OrgDoc_Verb)
            Lns_print([]LnsAny{OrgDoc_getTxt(_env, doc, verb.Get_TxtId(_env), transCtrl)})
        } else if _switch0 == OrgDoc_ItemKind__Exp {
        } else if _switch0 == OrgDoc_ItemKind__Block || _switch0 == OrgDoc_ItemKind__Name {
            var block OrgDoc_Block
            block = doc.GetBlockList(_env).GetAt(id).(OrgDoc_Block)
            Lns_print([]LnsAny{OrgDoc_getTxt(_env, doc, block.Get_TxtId(_env), transCtrl)})
        } else if _switch0 == OrgDoc_ItemKind__List {
            var list OrgDoc_ListItem
            list = doc.GetListList(_env).GetAt(id).(OrgDoc_ListItem)
            Lns_print([]LnsAny{_env.GetVM().String_format("%s- %s", []LnsAny{_env.GetVM().String_rep(" ", (list.Get_Level(_env) - 1) * 2), OrgDoc_getTxt(_env, doc, list.Get_TxtId(_env), transCtrl)})})
        } else if _switch0 == OrgDoc_ItemKind__Descript {
            var list OrgDoc_Descript
            list = doc.GetDescriptList(_env).GetAt(id).(OrgDoc_Descript)
            Lns_print([]LnsAny{_env.GetVM().String_format("%s- %s :: %s", []LnsAny{_env.GetVM().String_rep(" ", (list.Get_Level(_env) - 1) * 2), OrgDoc_getTxt(_env, doc, list.Get_TermId(_env), transCtrl), OrgDoc_getTxt(_env, doc, list.Get_TxtId(_env), transCtrl)})})
        } else if _switch0 == OrgDoc_ItemKind__Rule {
            Lns_print([]LnsAny{"\n-----"})
        } else if _switch0 == OrgDoc_ItemKind__Blank {
            Lns_print([]LnsAny{""})
        } else if _switch0 == OrgDoc_ItemKind__Emphasis {
            var emphasis OrgDoc_Exp
            emphasis = doc.GetExpList(_env).GetAt(id).(OrgDoc_Exp)
            var txt string
            txt = main_convExp0_711(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(OrgDoc_getTxt(_env, doc, emphasis.Get_TxtId(_env), transCtrl),"^%s+", "")).(string),"%s+$", "")))
            Lns_print([]LnsAny{_env.GetVM().String_format("%s%s%s", []LnsAny{emphasis.Get_Delimit(_env), txt, emphasis.Get_Delimit(_env)})})
        } else if _switch0 == OrgDoc_ItemKind__Keyword {
            var keyword OrgDoc_Keyword
            keyword = doc.GetKeywordList(_env).GetAt(id).(OrgDoc_Keyword)
            Lns_print([]LnsAny{_env.GetVM().String_format("#+%s: %s", []LnsAny{keyword.Get_Keyword(_env), OrgDoc_getTxt(_env, doc, keyword.Get_TxtId(_env), transCtrl)})})
        } else if _switch0 == OrgDoc_ItemKind__Comment {
            var comment OrgDoc_Comment
            comment = doc.GetCommentList(_env).GetAt(id).(OrgDoc_Comment)
            Lns_print([]LnsAny{_env.GetVM().String_format("# %s", []LnsAny{OrgDoc_getTxt(_env, doc, comment.Get_TxtId(_env), transCtrl)})})
        }
    }
}

// 91: decl @main.isOnlyAscii
func main_isOnlyAscii_1_(_env *LnsEnv, txt string) bool {
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(txt,"[^%g%s]+", nil, nil))){
        return false
    }
    return true
}

// 128: decl @main.createTransReq
func main_createTransReq_5_(_env *LnsEnv, doc OrgDoc_Document)(*LnsList, *OrgDoc_TransCtrl) {
    var ignoreTxtIdSet *LnsSet
    ignoreTxtIdSet = NewLnsSet([]LnsAny{})
    for _, _block := range( doc.GetBlockList(_env).Items ) {
        block := _block.(OrgDoc_Block)
        ignoreTxtIdSet.Add(block.Get_TxtId(_env))
    }
    for _, _verb := range( doc.GetVerbList(_env).Items ) {
        verb := _verb.(OrgDoc_Verb)
        ignoreTxtIdSet.Add(verb.Get_TxtId(_env))
    }
    var request *main_TransRequest
    request = Newmain_TransRequest(_env)
    var src2convId *LnsMap
    src2convId = NewLnsMap( map[LnsAny]LnsAny{})
    for _index, _txt := range( doc.GetTextList(_env).Items ) {
        index := _index + 1
        txt := _txt.(string)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(ignoreTxtIdSet.Has(index))) &&
            _env.SetStackVal( Lns_op_not(main_isOnlyAscii_1_(_env, txt))) ).(bool)){
            src2convId.Set(index,request.FP.AddTxt(_env, Lns_car(_env.GetVM().String_gsub(txt,"\\n", " ")).(string)))
        }
    }
    var reqList *LnsList
    reqList = NewLnsList([]LnsAny{})
    var workReq *main_TransRequest
    workReq = Newmain_TransRequest(_env)
    var size LnsInt
    size = 0
    var count LnsInt
    count = 0
    for _, _txt := range( request.FP.Get_q(_env).Items ) {
        txt := _txt.(string)
        workReq.FP.AddTxt(_env, txt)
        count = count + 1
        size = size + len(txt)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( size > main_LimitSize) ||
            _env.SetStackVal( count >= main_LimitTxtNum) ).(bool){
            reqList.Insert(Obj2Txt(_env, workReq.FP.ToMap()))
            size = 0
            count = 0
            workReq = Newmain_TransRequest(_env)
        }
    }
    if workReq.FP.Get_q(_env).Len() > 0{
        reqList.Insert(Obj2Txt(_env, workReq.FP.ToMap()))
    }
    return reqList, NewOrgDoc_TransCtrl(_env, doc.GetTextList(_env), src2convId)
}

// 174: decl @main.translate
func main_translate_7_(_env *LnsEnv, doc OrgDoc_Document,conf *main_Conf) LnsAny {
    var requestBodyList *LnsList
    var transCtrl *OrgDoc_TransCtrl
    requestBodyList,transCtrl = main_createTransReq_5_(_env, doc)
    var header *LnsMap
    header = NewLnsMap( map[LnsAny]LnsAny{"Authorization":"Bearer " + conf.FP.Get_token(_env),"Content-Type":"application/json; charset=utf-8",})
    var transList *LnsList
    transList = NewLnsList([]LnsAny{})
    for _, _requestBody := range( requestBodyList.Items ) {
        requestBody := _requestBody.(string)
        var err LnsAny
        var resp LnsAny
        err,resp = Req(_env, "https://translation.googleapis.com/language/translate/v2", "POST", header, requestBody)
        if err != nil{
            err_151 := err.(string)
            Lns_print([]LnsAny{err_151})
            return nil
        }
        if resp != nil{
            resp_153 := resp.(*HttpIF_Response)
            if resp_153.FP.Get_httpStatus(_env) == 200{
                var _map LnsAny
                _map = Txt2Map(_env, Lns_unwrapDefault( resp_153.FP.Get_body(_env), "").(string))
                {
                    _translation := main_convExp0_1266(Lns_2DDD(main_TranslationData__fromStem_4_(_env, _env.NilAccFin( _env.NilAccPush(
                    _map) && 
                    _env.NilAccPush( _env.NilAccPop().(*LnsMap).Get("data"))),nil)))
                    if !Lns_IsNil( _translation ) {
                        translation := _translation.(*main_TranslationData)
                        for _, _trans := range( translation.FP.Get_translations(_env).Items ) {
                            trans := _trans.(main_TranslationDownCast).Tomain_Translation()
                            transList.Insert(trans.FP.Get_translatedText(_env))
                        }
                    }
                }
            } else { 
                Lns_print([]LnsAny{_env.GetVM().String_format("httpStatus = %d", []LnsAny{resp_153.FP.Get_httpStatus(_env)})})
                Lns_print([]LnsAny{_env.GetVM().String_format("body = %s", []LnsAny{resp_153.FP.Get_body(_env)})})
                return nil
            }
        } else {
            Lns_print([]LnsAny{"resp is nil"})
            return nil
        }
    }
    transCtrl.FP.Set_transTxtList(_env, transList)
    return transCtrl
}

// 216: decl @main.org2github
func main_org2github_8_(_env *LnsEnv, doc OrgDoc_Document,path string) LnsAny {
    var fileObj Lns_luaStream
    
    {
        _fileObj := main_convExp0_1365(Lns_2DDD(Lns_io_open(path, nil)))
        if _fileObj == nil{
            return _env.GetVM().String_format("failed to open -- %s", []LnsAny{path})
        } else {
            fileObj = _fileObj.(Lns_luaStream)
        }
    }
    var customId2headline *LnsMap
    customId2headline = NewLnsMap( map[LnsAny]LnsAny{})
    for _, _headline := range( doc.GetHeadlineList(_env).Items ) {
        headline := _headline.(OrgDoc_Headline)
        var txt string
        txt = doc.GetTextList(_env).GetAt(headline.Get_TxtId(_env)).(string)
        txt = main_convExp0_1406(Lns_2DDD(_env.GetVM().String_gsub(txt,"^%*(.*)%*$", "%1")))
        txt = main_convExp0_1438(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(txt,"^%s+(.*)%s+$", "%1")).(string),"[^%w%s]", "")).(string),"%s", "-")))
        customId2headline.Set(headline.Get_CustomId(_env),_env.GetVM().String_lower(txt))
    }
    for  {
        var line string
        
        {
            _line := fileObj.Read(_env, "*l")
            if _line == nil{
                break
            } else {
                line = _line.(string)
            }
        }
        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"[[#", 1, true))){
            var customId string
            customId = main_convExp0_1495(Lns_2DDD(_env.GetVM().String_gsub(line,".*%[%[#(%g+)%]%].*", "%1")))
            {
                _headline := customId2headline.Get(customId)
                if !Lns_IsNil( _headline ) {
                    headline := _headline.(string)
                    var pattern string
                    pattern = _env.GetVM().String_format("%%[%%[#%s%%]%%]", []LnsAny{customId})
                    var link string
                    link = main_convExp0_1531(Lns_2DDD(_env.GetVM().String_gsub(line,pattern, _env.GetVM().String_format("[[#%s]]", []LnsAny{headline}))))
                    Lns_print([]LnsAny{link})
                } else {
                    Lns_print([]LnsAny{line})
                }
            }
        } else if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"^[%*]+ ", nil, nil))){
            var txt string
            txt = main_convExp0_1577(Lns_2DDD(_env.GetVM().String_gsub(line,"^[%*]+ ", "")))
            txt = main_convExp0_1592(Lns_2DDD(_env.GetVM().String_gsub(txt,"^%*(.*)%*$", "%1")))
            txt = main_convExp0_1608(Lns_2DDD(_env.GetVM().String_gsub(txt,"^%s+(.*)%s+$", "%1")))
            Lns_print([]LnsAny{Lns_car(_env.GetVM().String_gsub(line,"^([%*]+ ).*", "%1")).(string) + txt})
        } else { 
            Lns_print([]LnsAny{line})
        }
    }
    return nil
}

// 263: decl @main.__main
func Main___main(_env *LnsEnv, argList *LnsList) LnsInt {
    Lns_main_init( _env )
    var printUsage func(_env *LnsEnv, code LnsInt)
    printUsage = func(_env *LnsEnv, code LnsInt) {
        Lns_print([]LnsAny{_env.GetVM().String_format("usage: %s orgfile conffile [-v]", []LnsAny{argList.GetAt(1).(string)})})
        _env.GetVM().OS_exit(code)
    }
    if argList.Len() == 1{
        Lns_print([]LnsAny{_env.GetVM().String_format("usage: %s orgfile conffile [-v]", []LnsAny{argList.GetAt(1).(string)})})
        return 1
    }
    var argIndex LnsInt
    argIndex = 1
    var getNextOp func(_env *LnsEnv) LnsAny
    getNextOp = func(_env *LnsEnv) LnsAny {
        if argList.Len() <= argIndex{
            return nil
        }
        argIndex = argIndex + 1
        return argList.GetAt(argIndex).(string)
    }
    var getNextOpNonNil func(_env *LnsEnv) string
    getNextOpNonNil = func(_env *LnsEnv) string {
        {
            _nextOp := getNextOp(_env)
            if !Lns_IsNil( _nextOp ) {
                nextOp := _nextOp.(string)
                return nextOp
            }
        }
        printUsage(_env, 1)
    // insert a dummy
        return ""
    }
    var path string
    path = ""
    var conf LnsAny
    conf = nil
    var enableLog bool
    enableLog = false
    var mode string
    mode = main_Mode__Org
    for  {
        var arg string
        
        {
            _arg := getNextOp(_env)
            if _arg == nil{
                break
            } else {
                arg = _arg.(string)
            }
        }
        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(arg,"^-", nil, nil))){
            if _switch0 := arg; _switch0 == "-c" {
                var confPath string
                confPath = getNextOpNonNil(_env)
                var confTxt string
                
                {
                    _confTxt := Util_readFile(_env, confPath)
                    if _confTxt == nil{
                        Lns_print([]LnsAny{"failed to open -- ", confPath})
                        printUsage(_env, 1)
                    } else {
                        confTxt = _confTxt.(string)
                    }
                }
                {
                    __exp := main_convExp0_1906(Lns_2DDD(main_Conf__fromStem_4_(_env, Txt2Map(_env, confTxt),nil)))
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*main_Conf)
                        conf = _exp
                    } else {
                        Lns_print([]LnsAny{"load error"})
                        printUsage(_env, 1)
                    }
                }
            } else if _switch0 == "-v" {
                enableLog = true
            } else if _switch0 == "-m" {
                var nextOp string
                nextOp = getNextOpNonNil(_env)
                {
                    __exp := main_Mode__from_1_(_env, nextOp)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(string)
                        mode = _exp
                    } else {
                        Lns_print([]LnsAny{_env.GetVM().String_format("illegal mode -- %s", []LnsAny{nextOp})})
                        printUsage(_env, 1)
                    }
                }
            }
        } else { 
            path = arg
        }
    }
    var doc OrgDoc_Document
    doc = LoadOrg(_env, path, enableLog)
    if _switch1 := mode; _switch1 == main_Mode__Org {
        main_doc2org_0_(_env, doc, NewOrgDoc_TransCtrl(_env, doc.GetTextList(_env), NewLnsMap( map[LnsAny]LnsAny{})))
    } else if _switch1 == main_Mode__MkReq {
        var requestBodyList *LnsList
        requestBodyList,_ = main_createTransReq_5_(_env, doc)
        for _, _requestBody := range( requestBodyList.Items ) {
            requestBody := _requestBody.(string)
            Lns_print([]LnsAny{requestBody})
        }
    } else if _switch1 == main_Mode__Trans {
        if conf != nil{
            conf_233 := conf.(*main_Conf)
            {
                _transCtrl := main_translate_7_(_env, doc, conf_233)
                if !Lns_IsNil( _transCtrl ) {
                    transCtrl := _transCtrl.(*OrgDoc_TransCtrl)
                    main_doc2org_0_(_env, doc, transCtrl)
                } else {
                    printUsage(_env, 1)
                }
            }
        } else {
            Lns_print([]LnsAny{"no conf"})
            printUsage(_env, 1)
        }
    } else if _switch1 == main_Mode__Github {
        main_org2github_8_(_env, doc, path)
    }
    return 0
}




// 113: decl @main.TransRequest.addTxt
func (self *main_TransRequest) AddTxt(_env *LnsEnv, txt string) LnsInt {
    self.q.Insert(txt)
    return self.q.Len()
}
// declaration Class -- TransRequest
type main_TransRequestMtd interface {
    ToMap() *LnsMap
    AddTxt(_env *LnsEnv, arg1 string) LnsInt
    Get_format(_env *LnsEnv) string
    Get_model(_env *LnsEnv) string
    Get_q(_env *LnsEnv) *LnsList
    Get_source(_env *LnsEnv) string
    Get_target(_env *LnsEnv) string
}
type main_TransRequest struct {
    model string
    q *LnsList
    format string
    target string
    source string
    FP main_TransRequestMtd
}
func main_TransRequest2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*main_TransRequest).FP
}
type main_TransRequestDownCast interface {
    Tomain_TransRequest() *main_TransRequest
}
func main_TransRequestDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(main_TransRequestDownCast)
    if ok { return work.Tomain_TransRequest() }
    return nil
}
func (obj *main_TransRequest) Tomain_TransRequest() *main_TransRequest {
    return obj
}
func Newmain_TransRequest(_env *LnsEnv) *main_TransRequest {
    obj := &main_TransRequest{}
    obj.FP = obj
    obj.Initmain_TransRequest(_env)
    return obj
}
func (self *main_TransRequest) Get_model(_env *LnsEnv) string{ return self.model }
func (self *main_TransRequest) Get_q(_env *LnsEnv) *LnsList{ return self.q }
func (self *main_TransRequest) Get_format(_env *LnsEnv) string{ return self.format }
func (self *main_TransRequest) Get_target(_env *LnsEnv) string{ return self.target }
func (self *main_TransRequest) Get_source(_env *LnsEnv) string{ return self.source }
func (self *main_TransRequest) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["model"] = Lns_ToCollection( self.model )
    obj.Items["q"] = Lns_ToCollection( self.q )
    obj.Items["format"] = Lns_ToCollection( self.format )
    obj.Items["target"] = Lns_ToCollection( self.target )
    obj.Items["source"] = Lns_ToCollection( self.source )
    return obj
}
func (self *main_TransRequest) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func main_TransRequest__fromMap_8_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return main_TransRequest_FromMap( arg1, paramList )
}
func main_TransRequest__fromStem_9_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return main_TransRequest_FromMap( arg1, paramList )
}
func main_TransRequest_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := main_TransRequest_FromMapSub(obj,false, paramList);
    return conv,mess
}
func main_TransRequest_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &main_TransRequest{}
    newObj.FP = newObj
    return main_TransRequest_FromMapMain( newObj, objMap, paramList )
}
func main_TransRequest_FromMapMain( newObj *main_TransRequest, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["model"], false, nil); !ok {
       return false,nil,"model:" + mess.(string)
    } else {
       newObj.model = conv.(string)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["q"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil}}); !ok {
       return false,nil,"q:" + mess.(string)
    } else {
       newObj.q = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["format"], false, nil); !ok {
       return false,nil,"format:" + mess.(string)
    } else {
       newObj.format = conv.(string)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["target"], false, nil); !ok {
       return false,nil,"target:" + mess.(string)
    } else {
       newObj.target = conv.(string)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["source"], false, nil); !ok {
       return false,nil,"source:" + mess.(string)
    } else {
       newObj.source = conv.(string)
    }
    return true, newObj, nil
}
// 105: DeclConstr
func (self *main_TransRequest) Initmain_TransRequest(_env *LnsEnv) {
    self.model = "nmt"
    self.q = NewLnsList([]LnsAny{})
    self.format = "text"
    self.target = "en"
    self.source = "ja"
}


// declaration Class -- Translation
type main_TranslationMtd interface {
    ToMap() *LnsMap
    Get_model(_env *LnsEnv) string
    Get_translatedText(_env *LnsEnv) string
}
type main_Translation struct {
    translatedText string
    model string
    FP main_TranslationMtd
}
func main_Translation2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*main_Translation).FP
}
type main_TranslationDownCast interface {
    Tomain_Translation() *main_Translation
}
func main_TranslationDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(main_TranslationDownCast)
    if ok { return work.Tomain_Translation() }
    return nil
}
func (obj *main_Translation) Tomain_Translation() *main_Translation {
    return obj
}
func Newmain_Translation(_env *LnsEnv, arg1 string, arg2 string) *main_Translation {
    obj := &main_Translation{}
    obj.FP = obj
    obj.Initmain_Translation(_env, arg1, arg2)
    return obj
}
func (self *main_Translation) Initmain_Translation(_env *LnsEnv, arg1 string, arg2 string) {
    self.translatedText = arg1
    self.model = arg2
}
func (self *main_Translation) Get_translatedText(_env *LnsEnv) string{ return self.translatedText }
func (self *main_Translation) Get_model(_env *LnsEnv) string{ return self.model }
func (self *main_Translation) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["translatedText"] = Lns_ToCollection( self.translatedText )
    obj.Items["model"] = Lns_ToCollection( self.model )
    return obj
}
func (self *main_Translation) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func main_Translation__fromMap_4_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return main_Translation_FromMap( arg1, paramList )
}
func main_Translation__fromStem_5_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return main_Translation_FromMap( arg1, paramList )
}
func main_Translation_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := main_Translation_FromMapSub(obj,false, paramList);
    return conv,mess
}
func main_Translation_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &main_Translation{}
    newObj.FP = newObj
    return main_Translation_FromMapMain( newObj, objMap, paramList )
}
func main_Translation_FromMapMain( newObj *main_Translation, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["translatedText"], false, nil); !ok {
       return false,nil,"translatedText:" + mess.(string)
    } else {
       newObj.translatedText = conv.(string)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["model"], false, nil); !ok {
       return false,nil,"model:" + mess.(string)
    } else {
       newObj.model = conv.(string)
    }
    return true, newObj, nil
}

// declaration Class -- TranslationData
type main_TranslationDataMtd interface {
    ToMap() *LnsMap
    Get_translations(_env *LnsEnv) *LnsList
}
type main_TranslationData struct {
    translations *LnsList
    FP main_TranslationDataMtd
}
func main_TranslationData2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*main_TranslationData).FP
}
type main_TranslationDataDownCast interface {
    Tomain_TranslationData() *main_TranslationData
}
func main_TranslationDataDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(main_TranslationDataDownCast)
    if ok { return work.Tomain_TranslationData() }
    return nil
}
func (obj *main_TranslationData) Tomain_TranslationData() *main_TranslationData {
    return obj
}
func Newmain_TranslationData(_env *LnsEnv, arg1 *LnsList) *main_TranslationData {
    obj := &main_TranslationData{}
    obj.FP = obj
    obj.Initmain_TranslationData(_env, arg1)
    return obj
}
func (self *main_TranslationData) Initmain_TranslationData(_env *LnsEnv, arg1 *LnsList) {
    self.translations = arg1
}
func (self *main_TranslationData) Get_translations(_env *LnsEnv) *LnsList{ return self.translations }
func (self *main_TranslationData) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["translations"] = Lns_ToCollection( self.translations )
    return obj
}
func (self *main_TranslationData) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func main_TranslationData__fromMap_3_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return main_TranslationData_FromMap( arg1, paramList )
}
func main_TranslationData__fromStem_4_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return main_TranslationData_FromMap( arg1, paramList )
}
func main_TranslationData_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := main_TranslationData_FromMapSub(obj,false, paramList);
    return conv,mess
}
func main_TranslationData_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &main_TranslationData{}
    newObj.FP = newObj
    return main_TranslationData_FromMapMain( newObj, objMap, paramList )
}
func main_TranslationData_FromMapMain( newObj *main_TranslationData, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToListSub( objMap.Items["translations"], false, []Lns_ToObjParam{Lns_ToObjParam{
            main_Translation_FromMapSub, false,nil}}); !ok {
       return false,nil,"translations:" + mess.(string)
    } else {
       newObj.translations = conv.(*LnsList)
    }
    return true, newObj, nil
}

// declaration Class -- Conf
type main_ConfMtd interface {
    ToMap() *LnsMap
    Get_token(_env *LnsEnv) string
}
type main_Conf struct {
    token string
    FP main_ConfMtd
}
func main_Conf2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*main_Conf).FP
}
type main_ConfDownCast interface {
    Tomain_Conf() *main_Conf
}
func main_ConfDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(main_ConfDownCast)
    if ok { return work.Tomain_Conf() }
    return nil
}
func (obj *main_Conf) Tomain_Conf() *main_Conf {
    return obj
}
func Newmain_Conf(_env *LnsEnv, arg1 string) *main_Conf {
    obj := &main_Conf{}
    obj.FP = obj
    obj.Initmain_Conf(_env, arg1)
    return obj
}
func (self *main_Conf) Initmain_Conf(_env *LnsEnv, arg1 string) {
    self.token = arg1
}
func (self *main_Conf) Get_token(_env *LnsEnv) string{ return self.token }
func (self *main_Conf) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["token"] = Lns_ToCollection( self.token )
    return obj
}
func (self *main_Conf) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func main_Conf__fromMap_3_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return main_Conf_FromMap( arg1, paramList )
}
func main_Conf__fromStem_4_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return main_Conf_FromMap( arg1, paramList )
}
func main_Conf_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := main_Conf_FromMapSub(obj,false, paramList);
    return conv,mess
}
func main_Conf_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &main_Conf{}
    newObj.FP = newObj
    return main_Conf_FromMapMain( newObj, objMap, paramList )
}
func main_Conf_FromMapMain( newObj *main_Conf, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["token"], false, nil); !ok {
       return false,nil,"token:" + mess.(string)
    } else {
       newObj.token = conv.(string)
    }
    return true, newObj, nil
}

func Lns_main_init(_env *LnsEnv) {
    if init_main { return }
    init_main = true
    main__mod__ = "@main"
    Lns_InitMod()
    Lns_OrgDoc_init(_env)
    Lns_HttpIF_init(_env)
    Lns_Util_init(_env)
    main_LimitSize = 100 * 1000
    main_LimitTxtNum = 128
}
func init() {
    init_main = false
}
