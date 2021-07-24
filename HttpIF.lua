--./HttpIF.lns
local _moduleObj = {}
local __mod__ = '@HttpIF'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
if not _lune6 then
   _lune6 = _lune
end
local Response = {}
_moduleObj.Response = Response
function Response.setmeta( obj )
  setmetatable( obj, { __index = Response  } )
end
function Response.new( httpStatus, header, body )
   local obj = {}
   Response.setmeta( obj )
   if obj.__init then
      obj:__init( httpStatus, header, body )
   end
   return obj
end
function Response:__init( httpStatus, header, body )

   self.httpStatus = httpStatus
   self.header = header
   self.body = body
end
function Response:get_httpStatus()
   return self.httpStatus
end
function Response:get_header()
   return self.header
end
function Response:get_body()
   return self.body
end


local HttpClient = require( "HttpClient" )
_moduleObj.HttpClient = HttpClient

return _moduleObj
