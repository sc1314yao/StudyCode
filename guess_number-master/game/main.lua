local skynet = require "skynet"
local socket = require "skynet.socket"

local function accept(clientfd, addr)
    skynet.newservice("agent", clientfd, addr) -- 临时创建 actor, 可以提前创建  pool 式管理方式
end

skynet.start(function ()
    local listenfd = socket.listen("0.0.0.0", 8888)
    skynet.uniqueservice("redis")
    skynet.uniqueservice("hall")
    socket.start(listenfd, accept)
end)

-- 1. git clone 
-- 2. ubuntu  centos: make linux
--    mac:  make macosx
-- 3. config
-- 4. ./skynet config