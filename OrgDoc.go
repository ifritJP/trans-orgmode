// This code is transcompiled by LuneScript.
package main
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_OrgDoc bool
var OrgDoc__mod__ string
// decl enum -- ItemKind 
type OrgDoc_ItemKind = LnsInt
const OrgDoc_ItemKind__Blank = 8
const OrgDoc_ItemKind__Block = 5
const OrgDoc_ItemKind__Comment = 13
const OrgDoc_ItemKind__Descript = 10
const OrgDoc_ItemKind__Emphasis = 11
const OrgDoc_ItemKind__Exp = 4
const OrgDoc_ItemKind__Headline = 1
const OrgDoc_ItemKind__Keyword = 12
const OrgDoc_ItemKind__List = 6
const OrgDoc_ItemKind__Name = 9
const OrgDoc_ItemKind__Rule = 7
const OrgDoc_ItemKind__Table = 2
const OrgDoc_ItemKind__Text = 0
const OrgDoc_ItemKind__Verb = 3
var OrgDoc_ItemKindList_ = NewLnsList( []LnsAny {
  OrgDoc_ItemKind__Text,
  OrgDoc_ItemKind__Headline,
  OrgDoc_ItemKind__Table,
  OrgDoc_ItemKind__Verb,
  OrgDoc_ItemKind__Exp,
  OrgDoc_ItemKind__Block,
  OrgDoc_ItemKind__List,
  OrgDoc_ItemKind__Rule,
  OrgDoc_ItemKind__Blank,
  OrgDoc_ItemKind__Name,
  OrgDoc_ItemKind__Descript,
  OrgDoc_ItemKind__Emphasis,
  OrgDoc_ItemKind__Keyword,
  OrgDoc_ItemKind__Comment,
})
func OrgDoc_ItemKind_get__allList(_env *LnsEnv) *LnsList{
    return OrgDoc_ItemKindList_
}
var OrgDoc_ItemKindMap_ = map[LnsInt]string {
  OrgDoc_ItemKind__Blank: "ItemKind.Blank",
  OrgDoc_ItemKind__Block: "ItemKind.Block",
  OrgDoc_ItemKind__Comment: "ItemKind.Comment",
  OrgDoc_ItemKind__Descript: "ItemKind.Descript",
  OrgDoc_ItemKind__Emphasis: "ItemKind.Emphasis",
  OrgDoc_ItemKind__Exp: "ItemKind.Exp",
  OrgDoc_ItemKind__Headline: "ItemKind.Headline",
  OrgDoc_ItemKind__Keyword: "ItemKind.Keyword",
  OrgDoc_ItemKind__List: "ItemKind.List",
  OrgDoc_ItemKind__Name: "ItemKind.Name",
  OrgDoc_ItemKind__Rule: "ItemKind.Rule",
  OrgDoc_ItemKind__Table: "ItemKind.Table",
  OrgDoc_ItemKind__Text: "ItemKind.Text",
  OrgDoc_ItemKind__Verb: "ItemKind.Verb",
}
func OrgDoc_ItemKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := OrgDoc_ItemKindMap_[arg1]; ok { return arg1 }
    return nil
}

func OrgDoc_ItemKind_getTxt(arg1 LnsInt) string {
    return OrgDoc_ItemKindMap_[arg1];
}
// for 123
func OrgDoc_convExp0_343(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 127
func OrgDoc_convExp0_364(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 117: decl @OrgDoc.getTxt
func OrgDoc_getTxt(_env *LnsEnv, doc OrgDoc_Document,id LnsInt,transCtl *OrgDoc_TransCtrl) string {
    var src string
    src = transCtl.FP.Get(_env, id)
    var txt string
    txt = ""
    var startIndex LnsInt
    startIndex = 1
    for  {
        var index LnsInt
        
        {
            _index := OrgDoc_convExp0_343(Lns_2DDD(_env.GetVM().String_find(src,"SyM_%d", startIndex, nil)))
            if _index == nil{
                txt = txt + _env.GetVM().String_sub(src,startIndex, nil)
                break
            } else {
                index = _index.(LnsInt)
            }
        }
        var pos LnsInt
        var endPos LnsInt
        
        {
            _pos, _endPos := OrgDoc_convExp0_364(Lns_2DDD(_env.GetVM().String_find(src,"_%d+", index, nil)))
            if _pos == nil || _endPos == nil{
                panic("not found _%d+")
            } else {
                pos = _pos.(LnsInt)
                endPos = _endPos.(LnsInt)
            }
        }
        var expId LnsReal
        
        {
            _expId := Lns_tonumber(_env.GetVM().String_sub(src,pos + 1, endPos), nil)
            if _expId == nil{
                panic("illegal num -- " + _env.GetVM().String_sub(src,pos + 1, endPos))
            } else {
                expId = _expId.(LnsReal)
            }
        }
        var exp OrgDoc_Exp
        exp = doc.GetExpList(_env).GetAt((LnsInt)(expId)).(OrgDoc_Exp)
        txt = _env.GetVM().String_format("%s%s%s%s%s", []LnsAny{txt, _env.GetVM().String_sub(src,startIndex, index - 1), exp.Get_Delimit(_env), transCtl.FP.Get(_env, exp.Get_TxtId(_env)), exp.Get_Delimit(_env)})
        startIndex = endPos + 1
    }
    return Lns_car(_env.GetVM().String_gsub(txt,"\\n", "")).(string)
}

// 109: decl @OrgDoc.TransCtrl.get
func (self *OrgDoc_TransCtrl) Get(_env *LnsEnv, index LnsInt) string {
    {
        _transId := self.src2TransIdMap.Get(index)
        if !Lns_IsNil( _transId ) {
            transId := _transId.(LnsInt)
            return self.transTxtList.GetAt(transId).(string)
        }
    }
    return self.srcTxtList.GetAt(index).(string)
}
type OrgDoc_Item interface {
        Get_Id(_env *LnsEnv) LnsInt
        Get_Kind(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_Item( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Item); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Headline interface {
        Get_CustomId(_env *LnsEnv) LnsAny
        Get_Level(_env *LnsEnv) LnsInt
        Get_TxtId(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_Headline( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Headline); ok { 
        return obj
    }
    return nil
}

type OrgDoc_TableRow interface {
        Get_TxtIdList(_env *LnsEnv) *LnsList
}
func Lns_cast2OrgDoc_TableRow( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_TableRow); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Table interface {
        Get_RowList(_env *LnsEnv) *LnsList
}
func Lns_cast2OrgDoc_Table( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Table); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Verb interface {
        Get_TxtId(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_Verb( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Verb); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Exp interface {
        Get_Delimit(_env *LnsEnv) string
        Get_TxtId(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_Exp( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Exp); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Block interface {
        Get_TxtId(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_Block( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Block); ok { 
        return obj
    }
    return nil
}

type OrgDoc_ListItem interface {
        Get_Level(_env *LnsEnv) LnsInt
        Get_TxtId(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_ListItem( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_ListItem); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Descript interface {
        Get_Level(_env *LnsEnv) LnsInt
        Get_TermId(_env *LnsEnv) LnsInt
        Get_TxtId(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_Descript( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Descript); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Keyword interface {
        Get_Keyword(_env *LnsEnv) string
        Get_TxtId(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_Keyword( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Keyword); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Comment interface {
        Get_TxtId(_env *LnsEnv) LnsInt
}
func Lns_cast2OrgDoc_Comment( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Comment); ok { 
        return obj
    }
    return nil
}

type OrgDoc_Document interface {
        GetBlockList(_env *LnsEnv) *LnsList
        GetCommentList(_env *LnsEnv) *LnsList
        GetDescriptList(_env *LnsEnv) *LnsList
        GetExpList(_env *LnsEnv) *LnsList
        GetHeadlineList(_env *LnsEnv) *LnsList
        GetItemList(_env *LnsEnv) *LnsList
        GetKeywordList(_env *LnsEnv) *LnsList
        GetListList(_env *LnsEnv) *LnsList
        GetTableList(_env *LnsEnv) *LnsList
        GetTextList(_env *LnsEnv) *LnsList
        GetVerbList(_env *LnsEnv) *LnsList
}
func Lns_cast2OrgDoc_Document( obj LnsAny ) LnsAny {
    if _, ok := obj.(OrgDoc_Document); ok { 
        return obj
    }
    return nil
}


// declaration Class -- TransCtrl
type OrgDoc_TransCtrlMtd interface {
    Get(_env *LnsEnv, arg1 LnsInt) string
    Set_transTxtList(_env *LnsEnv, arg1 *LnsList)
}
type OrgDoc_TransCtrl struct {
    src2TransIdMap *LnsMap
    srcTxtList *LnsList
    transTxtList *LnsList
    FP OrgDoc_TransCtrlMtd
}
func OrgDoc_TransCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*OrgDoc_TransCtrl).FP
}
type OrgDoc_TransCtrlDownCast interface {
    ToOrgDoc_TransCtrl() *OrgDoc_TransCtrl
}
func OrgDoc_TransCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(OrgDoc_TransCtrlDownCast)
    if ok { return work.ToOrgDoc_TransCtrl() }
    return nil
}
func (obj *OrgDoc_TransCtrl) ToOrgDoc_TransCtrl() *OrgDoc_TransCtrl {
    return obj
}
func NewOrgDoc_TransCtrl(_env *LnsEnv, arg1 *LnsList, arg2 *LnsMap) *OrgDoc_TransCtrl {
    obj := &OrgDoc_TransCtrl{}
    obj.FP = obj
    obj.InitOrgDoc_TransCtrl(_env, arg1, arg2)
    return obj
}
func (self *OrgDoc_TransCtrl) Set_transTxtList(_env *LnsEnv, arg1 *LnsList){ self.transTxtList = arg1 }
// 102: DeclConstr
func (self *OrgDoc_TransCtrl) InitOrgDoc_TransCtrl(_env *LnsEnv, srcTxtList *LnsList,src2TransIdMap *LnsMap) {
    self.src2TransIdMap = src2TransIdMap
    self.srcTxtList = srcTxtList
    self.transTxtList = NewLnsList([]LnsAny{})
}


func Lns_OrgDoc_init(_env *LnsEnv) {
    if init_OrgDoc { return }
    init_OrgDoc = true
    OrgDoc__mod__ = "@OrgDoc"
    Lns_InitMod()
}
func init() {
    init_OrgDoc = false
}
