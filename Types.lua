--./Types.lns
local _moduleObj = {}
local __mod__ = '@Types'
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
local InStream = {}
_moduleObj.InStream = InStream
function InStream.setmeta( obj )
  setmetatable( obj, { __index = InStream  } )
end
function InStream.new(  )
   local obj = {}
   InStream.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function InStream:__init(  )

end


local OutStream = {}
_moduleObj.OutStream = OutStream
function OutStream.setmeta( obj )
  setmetatable( obj, { __index = OutStream  } )
end
function OutStream.new(  )
   local obj = {}
   OutStream.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function OutStream:__init(  )

end


local LuaInStream = {}
setmetatable( LuaInStream, { ifList = {InStream,} } )
_moduleObj.LuaInStream = LuaInStream
function LuaInStream:readStream( mode )

   do
      local bin = self.stream:read( mode )
      if bin ~= nil then
         return bin, ""
      end
   end
   
   return nil, "err"
end
function LuaInStream:read( size )

   return self:readStream( size )
end
function LuaInStream:readAll(  )

   return self:readStream( "*a" )
end
function LuaInStream.setmeta( obj )
  setmetatable( obj, { __index = LuaInStream  } )
end
function LuaInStream.new( stream )
   local obj = {}
   LuaInStream.setmeta( obj )
   if obj.__init then
      obj:__init( stream )
   end
   return obj
end
function LuaInStream:__init( stream )

   self.stream = stream
end


local LuaOutStream = {}
setmetatable( LuaOutStream, { ifList = {OutStream,} } )
_moduleObj.LuaOutStream = LuaOutStream
function LuaOutStream:write( bin )

   local _
   local _1, err = self.stream:write( bin )
   if err ~= nil then
      return err
   end
   
   return ""
end
function LuaOutStream.setmeta( obj )
  setmetatable( obj, { __index = LuaOutStream  } )
end
function LuaOutStream.new( stream )
   local obj = {}
   LuaOutStream.setmeta( obj )
   if obj.__init then
      obj:__init( stream )
   end
   return obj
end
function LuaOutStream:__init( stream )

   self.stream = stream
end


return _moduleObj
