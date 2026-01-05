-- package.cpath = "luaclib/?.so"
-- local so = require "tbl.c"


function Init(args)
    print("call [init] function", args)
end

function Loop()
    print("call [loop] function")
    for k, v in ipairs({1,2,3,4,5}) do
        -- so.echo(v)  -- 等于 调用 5 次  因为每次调用都是一个新的虚拟栈，所以没有必要维护
        print("value = " .. v)
    end
    return 1
end

function Release()
    print("call [release] function")
end
