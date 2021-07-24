--./OrgDoc.lns
local _moduleObj = {}
local __mod__ = '@OrgDoc'
local _lune = {}
if _lune6 then
   _lune = _lune6
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

local ItemKind = {}
_moduleObj.ItemKind = ItemKind
ItemKind._val2NameMap = {}
function ItemKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ItemKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ItemKind._from( val )
   if ItemKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ItemKind.__allList = {}
function ItemKind.get__allList()
   return ItemKind.__allList
end

ItemKind.Text = 0
ItemKind._val2NameMap[0] = 'Text'
ItemKind.__allList[1] = ItemKind.Text
ItemKind.Headline = 1
ItemKind._val2NameMap[1] = 'Headline'
ItemKind.__allList[2] = ItemKind.Headline
ItemKind.Table = 2
ItemKind._val2NameMap[2] = 'Table'
ItemKind.__allList[3] = ItemKind.Table
ItemKind.Verb = 3
ItemKind._val2NameMap[3] = 'Verb'
ItemKind.__allList[4] = ItemKind.Verb
ItemKind.Exp = 4
ItemKind._val2NameMap[4] = 'Exp'
ItemKind.__allList[5] = ItemKind.Exp
ItemKind.Block = 5
ItemKind._val2NameMap[5] = 'Block'
ItemKind.__allList[6] = ItemKind.Block
ItemKind.List = 6
ItemKind._val2NameMap[6] = 'List'
ItemKind.__allList[7] = ItemKind.List
ItemKind.Rule = 7
ItemKind._val2NameMap[7] = 'Rule'
ItemKind.__allList[8] = ItemKind.Rule
ItemKind.Blank = 8
ItemKind._val2NameMap[8] = 'Blank'
ItemKind.__allList[9] = ItemKind.Blank
ItemKind.Name = 9
ItemKind._val2NameMap[9] = 'Name'
ItemKind.__allList[10] = ItemKind.Name
ItemKind.Descript = 10
ItemKind._val2NameMap[10] = 'Descript'
ItemKind.__allList[11] = ItemKind.Descript
ItemKind.Emphasis = 11
ItemKind._val2NameMap[11] = 'Emphasis'
ItemKind.__allList[12] = ItemKind.Emphasis
ItemKind.Keyword = 12
ItemKind._val2NameMap[12] = 'Keyword'
ItemKind.__allList[13] = ItemKind.Keyword
ItemKind.Comment = 13
ItemKind._val2NameMap[13] = 'Comment'
ItemKind.__allList[14] = ItemKind.Comment


local Item = {}
_moduleObj.Item = Item
function Item.setmeta( obj )
  setmetatable( obj, { __index = Item  } )
end
function Item.new(  )
   local obj = {}
   Item.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Item:__init(  )

end


local Headline = {}
_moduleObj.Headline = Headline
function Headline.setmeta( obj )
  setmetatable( obj, { __index = Headline  } )
end
function Headline.new(  )
   local obj = {}
   Headline.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Headline:__init(  )

end


local TableRow = {}
_moduleObj.TableRow = TableRow
function TableRow.setmeta( obj )
  setmetatable( obj, { __index = TableRow  } )
end
function TableRow.new(  )
   local obj = {}
   TableRow.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function TableRow:__init(  )

end


local Table = {}
_moduleObj.Table = Table
function Table.setmeta( obj )
  setmetatable( obj, { __index = Table  } )
end
function Table.new(  )
   local obj = {}
   Table.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Table:__init(  )

end


local Verb = {}
_moduleObj.Verb = Verb
function Verb.setmeta( obj )
  setmetatable( obj, { __index = Verb  } )
end
function Verb.new(  )
   local obj = {}
   Verb.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Verb:__init(  )

end


local Exp = {}
_moduleObj.Exp = Exp
function Exp.setmeta( obj )
  setmetatable( obj, { __index = Exp  } )
end
function Exp.new(  )
   local obj = {}
   Exp.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Exp:__init(  )

end


local Block = {}
_moduleObj.Block = Block
function Block.setmeta( obj )
  setmetatable( obj, { __index = Block  } )
end
function Block.new(  )
   local obj = {}
   Block.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Block:__init(  )

end


local ListItem = {}
_moduleObj.ListItem = ListItem
function ListItem.setmeta( obj )
  setmetatable( obj, { __index = ListItem  } )
end
function ListItem.new(  )
   local obj = {}
   ListItem.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function ListItem:__init(  )

end


local Descript = {}
_moduleObj.Descript = Descript
function Descript.setmeta( obj )
  setmetatable( obj, { __index = Descript  } )
end
function Descript.new(  )
   local obj = {}
   Descript.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Descript:__init(  )

end


local Keyword = {}
_moduleObj.Keyword = Keyword
function Keyword.setmeta( obj )
  setmetatable( obj, { __index = Keyword  } )
end
function Keyword.new(  )
   local obj = {}
   Keyword.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Keyword:__init(  )

end


local Comment = {}
_moduleObj.Comment = Comment
function Comment.setmeta( obj )
  setmetatable( obj, { __index = Comment  } )
end
function Comment.new(  )
   local obj = {}
   Comment.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Comment:__init(  )

end


local Document = {}
_moduleObj.Document = Document
function Document.setmeta( obj )
  setmetatable( obj, { __index = Document  } )
end
function Document.new(  )
   local obj = {}
   Document.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Document:__init(  )

end


local Writer = require( "writer" )
_moduleObj.Writer = Writer

local TransCtrl = {}
_moduleObj.TransCtrl = TransCtrl
function TransCtrl.new( srcTxtList, src2TransIdMap )
   local obj = {}
   TransCtrl.setmeta( obj )
   if obj.__init then obj:__init( srcTxtList, src2TransIdMap ); end
   return obj
end
function TransCtrl:__init(srcTxtList, src2TransIdMap) 
   self.src2TransIdMap = src2TransIdMap
   self.srcTxtList = srcTxtList
   self.transTxtList = {}
end
function TransCtrl:get( index )

   do
      local transId = self.src2TransIdMap[index]
      if transId ~= nil then
         return self.transTxtList[transId]
      end
   end
   
   return self.srcTxtList[index]
end
function TransCtrl.setmeta( obj )
  setmetatable( obj, { __index = TransCtrl  } )
end
function TransCtrl:set_transTxtList( transTxtList )
   self.transTxtList = transTxtList
end


local function getTxt( doc, id, transCtl )

   
   local src = transCtl:get( id )
   local txt = ""
   local startIndex = 1
   while true do
      local index = src:find( "SyM_%d", startIndex )
      if  nil == index then
         local _index = index
      
         txt = txt .. src:sub( startIndex )
         break
      end
      
      local pos, endPos = src:find( "_%d+", index )
      if  nil == pos or  nil == endPos then
         local _pos = pos
         local _endPos = endPos
      
         error( "not found _%d+" )
      end
      
      local expId = tonumber( src:sub( pos + 1, endPos ) )
      if  nil == expId then
         local _expId = expId
      
         error( "illegal num -- " .. src:sub( pos + 1, endPos ) )
      end
      
      local exp = doc:GetExpList(  )[math.floor(expId)]
      txt = string.format( "%s%s%s%s%s", txt, src:sub( startIndex, index - 1 ), exp:get_Delimit(), transCtl:get( exp:get_TxtId() ), exp:get_Delimit())
      startIndex = endPos + 1
   end
   
   return (txt:gsub( "\\n", "" ) )
end
_moduleObj.getTxt = getTxt

return _moduleObj
