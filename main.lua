--./main.lns
local _moduleObj = {}
local __mod__ = '@main'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
function _lune._Set_or( setObj, otherSet )
   for val in pairs( otherSet ) do
      setObj[ val ] = true
   end
   return setObj
end
function _lune._Set_and( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if not otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_has( setObj, val )
   return setObj[ val ] ~= nil
end
function _lune._Set_sub( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_len( setObj )
   local total = 0
   for val in pairs( setObj ) do
      total = total + 1
   end
   return total
end
function _lune._Set_clone( setObj )
   local obj = {}
   for val in pairs( setObj ) do
      obj[ val ] = true
   end
   return obj
end

function _lune._toSet( val, toKeyInfo )
   if type( val ) == "table" then
      local tbl = {}
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal = _lune._toBool( mem )
         if mapKey == nil or mapVal == nil then
            if mapKey == nil then
               return nil
            end
            if keySub == nil then
               return nil, mapKey
            end
            return nil, string.format( "%s.%s", mapKey, keySub)
         end
         tbl[ mapKey ] = mapVal
      end
      return tbl
   end
   return nil
end

function _lune.nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
end

function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
end

function _lune._toStem( val )
   return val
end
function _lune._toInt( val )
   if type( val ) == "number" then
      return math.floor( val )
   end
   return nil
end
function _lune._toReal( val )
   if type( val ) == "number" then
      return val
   end
   return nil
end
function _lune._toBool( val )
   if type( val ) == "boolean" then
      return val
   end
   return nil
end
function _lune._toStr( val )
   if type( val ) == "string" then
      return val
   end
   return nil
end
function _lune._toList( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toValInfo = toValInfoList[ 1 ]
      for index, mem in ipairs( val ) do
         local memval, mess = toValInfo.func( mem, toValInfo.child )
         if memval == nil and not toValInfo.nilable then
            if mess then
              return nil, string.format( "%d.%s", index, mess )
            end
            return nil, index
         end
         tbl[ index ] = memval
      end
      return tbl
   end
   return nil
end
function _lune._toMap( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toKeyInfo = toValInfoList[ 1 ]
      local toValInfo = toValInfoList[ 2 ]
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal, valSub = toValInfo.func( mem, toValInfo.child )
         if mapKey == nil or mapVal == nil then
            if mapKey == nil then
               return nil
            end
            if keySub == nil then
               return nil, mapKey
            end
            return nil, string.format( "%s.%s", mapKey, keySub)
         end
         tbl[ mapKey ] = mapVal
      end
      return tbl
   end
   return nil
end
function _lune._fromMap( obj, map, memInfoList )
   if type( map ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( map[ memInfo.name ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s.%s", memInfo.name, key) or memInfo.name
      end
      obj[ memInfo.name ] = val
   end
   return true
end

function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if ifType == class then
               return true
            end
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

if not _lune6 then
   _lune6 = _lune
end
local OrgDoc = _lune.loadModule( 'OrgDoc' )
local HttpIF = _lune.loadModule( 'HttpIF' )
local Util = _lune.loadModule( 'Util' )

local LimitSize = 100 * 1000
local LimitTxtNum = 128
local function doc2org( doc, transCtrl )

   for __index, item in pairs( doc:GetItemList(  ) ) do
      
      local id = item:get_Id()
      do
         local _switchExp = item:get_Kind()
         if _switchExp == OrgDoc.ItemKind.Text then
            print( OrgDoc.getTxt( doc, id, transCtrl ) )
         elseif _switchExp == OrgDoc.ItemKind.Headline then
            local headline = doc:GetHeadlineList(  )[id]
            print( string.format( "\n\n%s %s", string.rep( "*", headline:get_Level() ), OrgDoc.getTxt( doc, headline:get_TxtId(), transCtrl )) )
            do
               local customId = headline:get_CustomId()
               if customId ~= nil then
                  print( string.format( [==[:PROPERTIES:
:CUSTOM_ID: %s
:END:
]==], customId) )
               end
            end
            
         elseif _switchExp == OrgDoc.ItemKind.Table then
            local table = doc:GetTableList(  )[id]
            local itemNum = #table:get_RowList()[1]:get_TxtIdList()
            print( string.format( "|%s-|", string.rep( "-|", itemNum - 1 )) )
            for __index, tableRow in pairs( table:get_RowList() ) do
               if #tableRow:get_TxtIdList() == 0 then
                  print( string.format( "|%s-|", string.rep( "-+", itemNum - 1 )) )
               else
                
                  io.stdout:write( "| " )
                  for __index, textId in pairs( tableRow:get_TxtIdList() ) do
                     io.stdout:write( OrgDoc.getTxt( doc, textId, transCtrl ) )
                     io.stdout:write( " | " )
                  end
                  
                  io.stdout:write( "\n" )
               end
               
            end
            
         elseif _switchExp == OrgDoc.ItemKind.Verb then
            local verb = doc:GetVerbList(  )[id]
            print( OrgDoc.getTxt( doc, verb:get_TxtId(), transCtrl ) )
         elseif _switchExp == OrgDoc.ItemKind.Exp then
         elseif _switchExp == OrgDoc.ItemKind.Block or _switchExp == OrgDoc.ItemKind.Name then
            local block = doc:GetBlockList(  )[id]
            print( OrgDoc.getTxt( doc, block:get_TxtId(), transCtrl ) )
         elseif _switchExp == OrgDoc.ItemKind.List then
            local list = doc:GetListList(  )[id]
            print( string.format( "%s- %s", string.rep( " ", (list:get_Level() - 1 ) * 2 ), OrgDoc.getTxt( doc, list:get_TxtId(), transCtrl )) )
         elseif _switchExp == OrgDoc.ItemKind.Descript then
            local list = doc:GetDescriptList(  )[id]
            print( string.format( "%s- %s :: %s", string.rep( " ", (list:get_Level() - 1 ) * 2 ), OrgDoc.getTxt( doc, list:get_TermId(), transCtrl ), OrgDoc.getTxt( doc, list:get_TxtId(), transCtrl )) )
         elseif _switchExp == OrgDoc.ItemKind.Rule then
            print( "\n-----" )
         elseif _switchExp == OrgDoc.ItemKind.Blank then
            print( "" )
         elseif _switchExp == OrgDoc.ItemKind.Emphasis then
            local emphasis = doc:GetExpList(  )[id]
            local txt = OrgDoc.getTxt( doc, emphasis:get_TxtId(), transCtrl ):gsub( "^%s+", "" ):gsub( "%s+$", "" )
            print( string.format( "%s%s%s", emphasis:get_Delimit(), txt, emphasis:get_Delimit()) )
         elseif _switchExp == OrgDoc.ItemKind.Keyword then
            local keyword = doc:GetKeywordList(  )[id]
            print( string.format( "#+%s: %s", keyword:get_Keyword(), OrgDoc.getTxt( doc, keyword:get_TxtId(), transCtrl )) )
         elseif _switchExp == OrgDoc.ItemKind.Comment then
            local comment = doc:GetCommentList(  )[id]
            print( string.format( "# %s", OrgDoc.getTxt( doc, comment:get_TxtId(), transCtrl )) )
         end
      end
      
   end
   
end

local function isOnlyAscii( txt )

   if txt:find( "[^%g%s]+" ) then
      return false
   end
   
   return true
end

local TransRequest = {}
setmetatable( TransRequest, { ifList = {Mapping,} } )
function TransRequest.new(  )
   local obj = {}
   TransRequest.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function TransRequest:__init() 
   self.model = "nmt"
   self.q = {}
   self.format = "text"
   self.target = "en"
   self.source = "ja"
end
function TransRequest:addTxt( txt )

   table.insert( self.q, txt )
   return #self.q
end
function TransRequest.setmeta( obj )
  setmetatable( obj, { __index = TransRequest  } )
end
function TransRequest:get_model()
   return self.model
end
function TransRequest:get_q()
   return self.q
end
function TransRequest:get_format()
   return self.format
end
function TransRequest:get_target()
   return self.target
end
function TransRequest:get_source()
   return self.source
end
function TransRequest:_toMap()
  return self
end
function TransRequest._fromMap( val )
  local obj, mes = TransRequest._fromMapSub( {}, val )
  if obj then
     TransRequest.setmeta( obj )
  end
  return obj, mes
end
function TransRequest._fromStem( val )
  return TransRequest._fromMap( val )
end

function TransRequest._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "model", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "q", func = _lune._toList, nilable = false, child = { { func = _lune._toStr, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "format", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "target", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "source", func = _lune._toStr, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local Translation = {}
setmetatable( Translation, { ifList = {Mapping,} } )
function Translation.setmeta( obj )
  setmetatable( obj, { __index = Translation  } )
end
function Translation.new( translatedText, model )
   local obj = {}
   Translation.setmeta( obj )
   if obj.__init then
      obj:__init( translatedText, model )
   end
   return obj
end
function Translation:__init( translatedText, model )

   self.translatedText = translatedText
   self.model = model
end
function Translation:get_translatedText()
   return self.translatedText
end
function Translation:get_model()
   return self.model
end
function Translation:_toMap()
  return self
end
function Translation._fromMap( val )
  local obj, mes = Translation._fromMapSub( {}, val )
  if obj then
     Translation.setmeta( obj )
  end
  return obj, mes
end
function Translation._fromStem( val )
  return Translation._fromMap( val )
end

function Translation._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "translatedText", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "model", func = _lune._toStr, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local TranslationData = {}
setmetatable( TranslationData, { ifList = {Mapping,} } )
function TranslationData.setmeta( obj )
  setmetatable( obj, { __index = TranslationData  } )
end
function TranslationData.new( translations )
   local obj = {}
   TranslationData.setmeta( obj )
   if obj.__init then
      obj:__init( translations )
   end
   return obj
end
function TranslationData:__init( translations )

   self.translations = translations
end
function TranslationData:get_translations()
   return self.translations
end
function TranslationData:_toMap()
  return self
end
function TranslationData._fromMap( val )
  local obj, mes = TranslationData._fromMapSub( {}, val )
  if obj then
     TranslationData.setmeta( obj )
  end
  return obj, mes
end
function TranslationData._fromStem( val )
  return TranslationData._fromMap( val )
end

function TranslationData._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "translations", func = _lune._toList, nilable = false, child = { { func = Translation._fromMap, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local function createTransReq( doc )

   
   local ignoreTxtIdSet = {}
   for __index, block in pairs( doc:GetBlockList(  ) ) do
      ignoreTxtIdSet[block:get_TxtId()]= true
   end
   
   for __index, verb in pairs( doc:GetVerbList(  ) ) do
      ignoreTxtIdSet[verb:get_TxtId()]= true
   end
   
   
   local request = TransRequest.new()
   local src2convId = {}
   
   for index, txt in pairs( doc:GetTextList(  ) ) do
      if not _lune._Set_has(ignoreTxtIdSet, index ) and not isOnlyAscii( txt ) then
         src2convId[index] = request:addTxt( (txt:gsub( "\\n", " " ) ) )
      end
      
   end
   
   
   local reqList = {}
   local workReq = TransRequest.new()
   local size = 0
   local count = 0
   for __index, txt in pairs( request:get_q() ) do
      workReq:addTxt( txt )
      count = count + 1
      size = size + #txt
      if size > LimitSize or count >= LimitTxtNum then
         table.insert( reqList, Util.Json.Obj2Txt( workReq:_toMap(  ) ) )
         size = 0
         count = 0
         workReq = TransRequest.new()
      end
      
   end
   
   if #workReq:get_q() > 0 then
      table.insert( reqList, Util.Json.Obj2Txt( workReq:_toMap(  ) ) )
   end
   
   
   return reqList, OrgDoc.TransCtrl.new(doc:GetTextList(  ), src2convId)
end

local Conf = {}
setmetatable( Conf, { ifList = {Mapping,} } )
function Conf.setmeta( obj )
  setmetatable( obj, { __index = Conf  } )
end
function Conf.new( token )
   local obj = {}
   Conf.setmeta( obj )
   if obj.__init then
      obj:__init( token )
   end
   return obj
end
function Conf:__init( token )

   self.token = token
end
function Conf:get_token()
   return self.token
end
function Conf:_toMap()
  return self
end
function Conf._fromMap( val )
  local obj, mes = Conf._fromMapSub( {}, val )
  if obj then
     Conf.setmeta( obj )
  end
  return obj, mes
end
function Conf._fromStem( val )
  return Conf._fromMap( val )
end

function Conf._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "token", func = _lune._toStr, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local function translate( doc, conf )

   local requestBodyList, transCtrl = createTransReq( doc )
   
   local header = {["Authorization"] = "Bearer " .. conf:get_token(), ["Content-Type"] = "application/json; charset=utf-8"}
   
   local transList = {}
   for __index, requestBody in pairs( requestBodyList ) do
      local err, resp = HttpIF.HttpClient.Req( "https://translation.googleapis.com/language/translate/v2", "POST", header, requestBody )
      if err ~= nil then
         print( err )
         return nil
      end
      
      if resp ~= nil then
         
         if resp:get_httpStatus() == 200 then
            local map = Util.Json.Txt2Map( _lune.unwrapDefault( resp:get_body(), "") )
            do
               local translation = TranslationData._fromStem( _lune.nilacc( map, nil, 'item', 'data') )
               if translation ~= nil then
                  for __index, trans in pairs( translation:get_translations() ) do
                     
                     table.insert( transList, trans:get_translatedText() )
                  end
                  
               end
            end
            
         else
          
            print( string.format( "httpStatus = %d", resp:get_httpStatus()) )
            print( string.format( "body = %s", resp:get_body()) )
            return nil
         end
         
      else
         print( "resp is nil" )
         return nil
      end
      
   end
   
   
   transCtrl:set_transTxtList( transList )
   return transCtrl
end

local Mode = {}
Mode._val2NameMap = {}
function Mode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Mode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Mode._from( val )
   if Mode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Mode.__allList = {}
function Mode.get__allList()
   return Mode.__allList
end

Mode.Org = 'org'
Mode._val2NameMap['org'] = 'Org'
Mode.__allList[1] = Mode.Org
Mode.MkReq = 'mkreq'
Mode._val2NameMap['mkreq'] = 'MkReq'
Mode.__allList[2] = Mode.MkReq
Mode.Trans = 'trans'
Mode._val2NameMap['trans'] = 'Trans'
Mode.__allList[3] = Mode.Trans


local function __main( argList )

   local function printUsage( code )
   
      print( string.format( "usage: %s orgfile conffile [-v]", argList[1]) )
      os.exit( code )
   end
   
   if #argList == 1 then
      print( string.format( "usage: %s orgfile conffile [-v]", argList[1]) )
      return 1
   end
   
   
   local argIndex = 1
   local function getNextOp(  )
   
      if #argList <= argIndex then
         return nil
      end
      
      argIndex = argIndex + 1
      return argList[argIndex]
   end
   local function getNextOpNonNil(  )
   
      do
         local nextOp = getNextOp(  )
         if nextOp ~= nil then
            return nextOp
         end
      end
      
      printUsage( 1 )
   end
   
   local path = ""
   local conf = nil
   local enableLog = false
   local mode = Mode.Org
   
   while true do
      local arg = getNextOp(  )
      if  nil == arg then
         local _arg = arg
      
         break
      end
      
      if arg:find( "^-" ) then
         do
            local _switchExp = arg
            if _switchExp == "-c" then
               local confPath = getNextOpNonNil(  )
               local confTxt = Util.readFile( confPath )
               if  nil == confTxt then
                  local _confTxt = confTxt
               
                  print( "failed to open -- ", confPath )
                  printUsage( 1 )
               end
               
               do
                  local _exp = Conf._fromStem( Util.Json.Txt2Map( confTxt ) )
                  if _exp ~= nil then
                     conf = _exp
                  else
                     print( "load error" )
                     printUsage( 1 )
                  end
               end
               
            elseif _switchExp == "-v" then
               enableLog = true
            elseif _switchExp == "-m" then
               local nextOp = getNextOpNonNil(  )
               do
                  local _exp = Mode._from( nextOp )
                  if _exp ~= nil then
                     mode = _exp
                  else
                     print( string.format( "illegal mode -- %s", nextOp) )
                     printUsage( 1 )
                  end
               end
               
            end
         end
         
      else
       
         path = arg
      end
      
   end
   
   
   local doc = OrgDoc.Writer.loadOrg( path, enableLog )
   do
      local _switchExp = mode
      if _switchExp == Mode.Org then
         doc2org( doc, OrgDoc.TransCtrl.new(doc:GetTextList(  ), {}) )
      elseif _switchExp == Mode.MkReq then
         local _
         local requestBodyList, _1 = createTransReq( doc )
         for __index, requestBody in pairs( requestBodyList ) do
            print( requestBody )
         end
         
      elseif _switchExp == Mode.Trans then
         if conf ~= nil then
            do
               local transCtrl = translate( doc, conf )
               if transCtrl ~= nil then
                  doc2org( doc, transCtrl )
               else
                  printUsage( 1 )
               end
            end
            
         else
            print( "no conf" )
            printUsage( 1 )
         end
         
      end
   end
   
   
   return 0
end
_moduleObj.__main = __main

return _moduleObj
