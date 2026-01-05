 -- 主协程
local co = coroutine.create(function (arg1)
    -- 创建的协程
    -- local run, ismain = coroutine.running()
    -- print(run, ismain, arg1)
    local ret1 = arg1+1
    -- 是否满足条件
    local arg2 = coroutine.yield(ret1)  -- 子协程让出
    local ret2 = arg2+1
    return ret2
end)

local co1 = coroutine.running()
local arg1 = 1
local ok, ret1, ret2
ok, ret1 = coroutine.resume(co, arg1) -- 主协程 让出 唤醒 co
print(co1, ok, ret1)
ok, ret2 = coroutine.resume(co, ret1) -- 主协程 让出 唤醒 co
print(co1, ok, ret2)
