[toc]
# GM命令
## 起始命令

```
.level 79    -------------------------升级
.modify tp 300  ---------------------改天赋点
.modify money 9999999999--------------加钱
.additem 21876 4  -------------------4个包包
.learn 34091  -----------------------300骑术
.learn 54197  ---------------------寒冷气候飞行
.learn 42777  ----------------------幽灵虎
.learn 32345 -------------------------凤凰
.explore 1  ---------------------探索所有地图
.maxskill   ---------------------所有技能熟练度增加满点
.setskill 762 300 300 ---------------------提昇騎乘到300
```

```
货币
 
.additem 47241 2147483647 凯旋
.additem 45624 2147483647 征服
.additem 49426 2147483647 冰霜
.additem 40753 2147483647 勇气
.additem 40752 2147483647 英雄
.additem 29434 2147483647 正义
 
.additem 43016 2147483647 达拉然烹饪
.additem 41596 2147483647 达拉然珠宝
.additem 44990 2147483647 勇士徽记
.additem 20559 100   阿拉希
.additem 20558 100   战歌
.additem 20560 100   奥山
.modify honor 2147483647 荣誉值
 
特殊技能，单独做宏技能
 
.respawn
刷新周围怪物
 
.revive
复活当前目标（复活自己需/G启用工会频道后使用）
 
.cooldown 
立即使目标的魔法冷却
 
.start 脱离卡死
```

```
护甲精通
 
.learn 9077   皮甲 
.learn 8737   锁甲 
.learn 750    板甲 
武器技能
 
.learn 201  单手剑
.learn 202  双手剑
.learn 198  单手锤
.learn 199  双手锤
.learn 196  单手斧 
.learn 197  双手斧 
.learn 1180  匕首 
.learn 9116  盾牌
.learn 227  法杖
.learn 5009  魔杖 
.learn 200  长柄武器
.learn 15590  拳套武器
.learn 264  弓 
.learn 5011  弩 
.learn 266  枪械 
.learn 2567  投掷 
.learn 674 双持
.learn 46917 泰坦之握
 
.learn 1804 开锁
.learn 1809 初级开锁
.learn 1810 中级开锁
.learn 6460 高级开锁 
=====================
```
## 常用命令
- 服务器控制
    - .reload all //重新载入所有配置
    - .gps              显示角色或生物的坐标(x,y,z) 地图标号和地区
    - .announce 信息内容        //发布公告 
    - .notify 信息内容                    // 发送全局公告信息，屏幕中间
    - .send items 玩家名字 "感谢对本服的支持" "感谢对本服的支" 25:1 // 邮寄物品
    - .server restart 60 //服务器重启60秒
    - .server restart cancel //取消服务器重启

    - .achievement add 成就ID或者shift选中//添加成就
    - .revive //复活。复活自己需/G启用工会频道后使用
- 玩家相关
- .respawn //立即刷新身边最近的生物或游戏目标
    - .appear 玩家名                   移动到目标位置。 
    - .summon 玩家名                   召唤目标到你的位置。
    - .kick $name //踢出玩家
    - .respawn // 立刻刷新附近生物 
    - .revive // 复活选择的玩家
    - .morp 模型ID变身
    - .demorph 取消变身
    - .cooldown //立即使目标的魔法冷却
    - .reset stats 重置目标所有状态
    - .character level 80 //改变角色级别
- 角色属性控制
    - .modify money 999999 给自己加金钱
    - .modify aspeed 10 改变玩家的速度
    - .respawn //立即刷新身边最近的生物或游戏目标
    - .debug bg //开启：即可单人进战场关闭：恢复正常战场人数
    -  .npc say 内容 //编辑NPC说话，命令后加你想表述的话语
    -  .npc yell 内容 //编辑NPC大喊，命令后加你想表述的话语
    - .revive        复活，先选中自
    - .die          杀死目标
    - .gm on/off       GM模式开启，打怪怪不打你 无敌状态
    - .gm visible on/off //使GM是否对玩家可见
    - .modify hp 999999999 增加目标的HP
    - .modify mana 999999999          修改法力值，使用同上
    - .go 16222.1 16252.1 12.5872 1             传送自己到GM岛
    - .bank            打开仓库 
    - .maxskill   (所有技能熟练度增加满点)
- 查询物品（少用，最好精确点。会消耗服务器资源）
    - .lookup item 名称
    - .lookup item 物品名字           // 查询物品编号
    - .lookup quest 任务名字          // 查询任务编号
    - .lookup creature 生物名字       // 查询生物编号
    - . lookup object 物体名字         // 查询物体编号
- 刷物品
    - .additem ID
    - .additem 23162 36格包
    - .add 25596 偷窥之哨
    - .add 49623 影之哀伤
    - .add 46017 奶锤
    - .add 34334 橙弓
    - .add 32837 蛋刀主
    - .add 32838 蛋刀副
    - .add 13262 灰烬使者
    - .add 17182 萨弗拉斯，炎魔拉格纳罗斯之手
    - .add 19019 雷霆之怒，逐风者的祝福之剑 
    - .add 22589 罗克迪洛尔，上古守护者的手杖

- 活动控制
    - .event start 活动ID
- NPC相关操作
    - .npc add 生物Id                 // 当前位置刷新一个生物C
    - .npc add temp 生物Id            // 当前位置临时刷新一个生物，重启后消失
    - .npc delete                     // 删除选择生物
    - .npc move                       // 移动选择生物至当前位置
    - .npc info                       // 查询选择生物的详细信息
    - .npc delete item 物品ID         // 从选择NPC的销售菜单内删除物品
    - .npc add item 物品ID         // 增加售卖物品
    - .npc playemote 3    //使所选生物做 #emoteid 所指的表情动作 
    - .npc set spawntime 3600  //调整所选生物的刷新时间#秒.
    - .npc changelevel 80 //改变生物等级
    - .npc add 32541  训练假人
    - .npc add 32666  训练假人
- 物体控制
    - .gobject add 物品ID             // 在当前位置刷新一个世界物体
    - .gobject add temp 物体ID        // 在当前位置临时刷新一个世界物体，重启后消失
    - .gobject delete 刷新ID          // 删除世界物体刷新
    - .gobject near                   // 查询当前距离最新的世界物体信息
- 任务控制
    - .quest add 任务Id               // 增加任务
    - .quest remove 任务Id            // 删除任务
    - .quest complete 任务Id          // 完成任务
- 技能相关
    - .learn 31700 一个大黄蜂（可以内陆飞行）
    - .learn all_myclass      学习自己的全部技能
    - .unlearn 31700    遗忘你所选择角色的魔法
    - .learn 40733 神圣之盾 无敌蛋壳 无时间
    - .learn 40683 激怒 buff魔法 1.3秒施放 持续20秒伤害+50%攻击速度30%
    - .learn 40497 混乱冲锋 5-35码冲锋 附近造成20000+火伤 15秒冷却
    - .learn 29968 无限距离无冷却闪现术
    - .learn 34973 寒冰屏障
    - .learn 44140    鱼叉冲浪      
    - .learn 32345  凤凰
    - .learn 17244 精神控制任何生物
    - .learn 41431 符文护盾 提高施法和攻击速度100%，免疫法术打断
    - .Learn 39869 狂野怨灵 伤害提高500%移动速度提高200%
    - .learn 39811 伤害免疫:火焰、冰霜、暗影、自然、秘法    
    - .learn 39804 伤害免疫:魔法     
    - .learn 34310 物理免疫         
    - .learn 34311 伤害免疫:物理
    - .learn 40936 无CD超范围战争践踏
    - .learn 40636 大范围群恐 
    - .learn 41080 吞噬灵魂 秒杀
    - .learn 70063 霜之哀伤烈怒 
    - .learn 17244  控制  学BOSS技能

## 核心端新增命令
| 命令                              | 语法实例                  | 说明                                                         |
| --------------------------------- | ------------------------- | ------------------------------------------------------------ |
| .rl all                           | .rl all                   | 重载核心所有自定义数据表                                     |
| .rl item                          | .rl item                  | 重载item_template                                            |
| ._add xp                          | ._add xp 1000             | 增加1000经验值                                               |
| ._add token                       | ._add token 1000          | 增加1000积分                                                 |
| ._add sps                         | ._add sps 1000            | 增加1000斗气点                                               |
| ._add tps                         | ._add tps 10              | 增加额外10点天赋                                             |
| ._add cps                         | ._add cps 5               | 增加额外5点商业技能点                                        |
| ._add pos                         | ._add pos 加基森          | 采集当前坐标到坐标模板表，备注名为加基森                     |
| ._add skill XXX                   | ._add skill 1             | 获得_skill中skillId                                          |
| ._add rankvalue                   | ._add rankvalue 10        | 增加10点自定义等级值                                         |
| ._add spiritpower                 | ._add spiritpower 30      | 增加30点灵力值                                               |
| ._add maxspiritpower              | ._add maxspiritpower      | 10	增加10点灵力值上限                                     |
| ._add talismanvalue               | ._add talismanvalue 20    | 增加20点法宝值                                               |
| ._add faker                       | ._add faker               | 上线一个在线假人                                             |
| ._add aura                        | ._add aura 571 23948      | 571地图所有玩家增加23948光环。更多参数游戏中输入._add aura查询 |
| ._rew ip rewId                    | ._rew ip rewId 1          | 按IP全服发放奖励模板ID为1的奖励                              |
| ._rew ip item                     | ._rew ip item 25          | 按IP全服发放25这个物品                                       |
| ._rew rewId                       | ._rew rewId 1             | 发放奖励模板ID为1的奖励                                      |
| ._set viplevel                    | ._set viplevel 5          | 提升会员等级至5                                              |
| ._set faction                     | ._set faction 3           | 设置玩家自定义阵营为3                                        |
| ._set prefix                      | ._set prefix 我是一个好人 | 设置玩家名字前缀为我是一个好人                               |
| ._set suffix                      | ._set suffix 杀人狂       | 设置玩家名字后缀为杀人狂                                     |
| ._set scale                       | ._set scale 3             | 设置目标生物模型大小，直接保存数据库                         |
| ._set display                     | ._set display 100         | 设置目标生物模型编号为100，直接保存数据库                    |
| ._set name                        | ._set name XXX            | 设置目标生物名字为XXX，直接保存数据库                        |
| ._set subname                     | ._set subname SSS         | 设置目标生物子名字为SSS，直接保存数据库                      |
| ._set npcflag                     | ._set npcflag 1           | 设置目标生物npcflag为1，直接保存数据库                       |
| ._reset sps                       | ._reset sps               | 重置斗气加点                                                 |
| ._reset daily                     | ._reset daily             | 重置日常任务                                                 |
| ._reset weekly                    | ._reset weekly            | 重置周常任务                                                 |
| ._vendor                          | ._vendor 54               | 打开npc_vendor中entry为54的售卖列表                          |
| ._tele                            | ._tele 77                 | 传送至坐标模板ID为77的位置                                   |
| .jk                               | .jk                       | 死亡状态下解卡，传送至炉石点                                 |
| ._summon                          | ._summon 3 100            | 召唤一个ID为3的生物，持续100秒。更多参数请在游戏中输入._summon查询 |
| ._db_clear                        | ._db_clear                | 清空账号库与角色库，慎用！！！！                             |
| .pl                               | .pl                       | 查询IP、玩家、假人在线数量                                   |
| ._loot                            | ._loot                    | 一键拾取1000码内的所有掉落                                   |
| .cdk create                       | .cdk create 100 2         | 宣传奖励	数据中自动生成 100个奖励模板ID为2、备注名为宣传奖励的兑换码 |
| .cdk out                          | .cdk out                  | 核心文件夹下输出兑换码的文档                                 |
| ._group summon                    | ._group summon            | 队伍召唤                                                     |
| ._group revive                    | ._group revive            | 群体复活                                                     |
| ._heal                            | ._heal 3 999              | 给1000码内ID为3的生物增加血量999。ICC绿龙，可以直接做个NPC调用这个命令 |
| ._level                           | ._level 70                | 当玩家等级低于70时，将玩家提升至70级                         |
| ._morpet XXX                      | ._morpet 33               | 改变当前宠物的模型ID，会保存在数据库                         |
| .backup a                         | .backup a                 | 备份账号库                                                   |
| .backup c                         | .backup c                 | 备份角色库                                                   |
| .backup w                         | .backup w                 | 备份世界库，导出的sql文件用Navicat导入会很慢，请用其他工具   |
| .sql itemdisplayinfo              | .sql itemdisplayinfo      | itemdisplayinfo.dbc到_dbc_itemdisplayinfo                    |
| .sql itemextendedcost             | .sql itemextendedcost     | itemextendedcost.dbc到_dbc_itemextendedcost                  |
| .sql spell                        | .sql spell                | spell.dbc到_dbc_spell                                        |
| .sql spellitemenchantment         | .sql spellitemenchantment | spellitemenchantment到_dbc_spellitemenchantment              |
| .sql spellicon                    | .sql spellicon            | spellicon.dbc到_dbc_spellicon                                |
| .dbc item                         | .dbc item                 | item_template到item.dbc                                      |
| .dbc itemdisplayinfo              | .dbc itemdisplayinfo      | _dbc_itemdisplayinfo到itemdisplayinfo.dbc                    |
| .dbc itemextendedcost             | .dbc itemextendedcost     | _dbc_itemextendedcost到itemextendedcost.dbc                  |
| .dbc spell                        | .dbc spell                | _dbc_spell到spell.dbc                                        |
| .dbc spellitemenchantment         | .dbc spellitemenchantment | _dbc_spellitemenchantment到spellitemenchantment.dbc          |
| .dbc spellicon                    | .dbc spellicon            | _dbc_spellicon到spellicon.dbc                                |
| define ARENA_DUEL_PERMIN_REW_ID   | 1001                      | //_rew 每分钟奖励                                            |
| define ARENA_DUEL_END_REW_ID      | 1002                      | //_rew 活动结束每人奖励                                      |
| define ARENA_DUEL_WIN_REW_ID      | 1003                      | //_rew 决斗获胜奖励                                          |
| define ARENA_DUEL_LOS_REW_ID      | 1004                      | //_rew 决斗失败奖励                                          |
| define ARENA_DUEL_CHAMPION_REW_ID | 1005                      | //_rew 冠军奖励                                              |
| ._black on                        | ._black on/off            | 开启，关闭黑市                                               |
| ._add rune [spellid]              | ._add rune [spellid]      | 获取/gc中的铭文                                              |
| req                               | 需求表                    | 升级的ReqId不超过4个长度 最大9999                            |
| rew                               | 奖励表                    | 售卖的RewId不超过3个长度 最大999                             |
| reload                            | /reload all               | 重载                                                         |
| ## 附魔位置说明                   |                           |                                                              |

```
0.主FM位置，官方FM卷轴FM的位置

1.临时FM位置，例如毒药

2.宝石1

3.宝石2

4.宝石3

5.宝石奖励

6.棱彩宝石

7.随机FM位置1

8.随机FM位置2

9.随机FM位置3

10.随机FM位置4

11.随机FM位置5
```
## 新增技能说明

```
光环类型

102 攻击强度提高X点

131 远程攻击强度提高X点

180 法术强度提高X点

168 所有伤害提高X%

169 爆击几率提高X%

 

复合值B 

1000 对玩家

1001 对生物

1002 对BOSS

1003 对小怪

 

命令

产生兑换码

.cdk create 个数 rewId 备注 

.cdk out 导出兑换码.txt

._group summon 群体召唤

._group revive 群体复活

._heal 生物编号 血量  给生物增加血量
```
## 全部游戏命令
```
A
.account                        // 查询账号等级
.account create 账号 密码       // 创建账号
.account delete 账号            // 删除账号
.account set gmlevel 账号 3 -1  // 设置账号GM权限
.account onlinelist             // 查询账号在线列表
.additem 物品Id 数量            // 刷物品到背包
.additemset 套装ID              // 刷新套装全件到背包
.account password 老密码  新密码   新密码   // 修改账号密码
.account set password 账号  新密码  新密码  // 设置账号密码
.aura 1 语法: .aura #spellid    为目标添加ID为 #spellid 的光环 
.announce 3 语法： .announce 需要广播的消息    广播 $需要广播的消息 给所有在线的玩家. 
B
.bank                           // 进入银行
.baninfo ip                     // 查看被封Ip信息
.baninfo character              // 查看被封角色信息
.baninfo account                // 查看被封账号信息
.ban 3 语法： ban <account|ip|character> $NameOrIp $bantime $reason    封禁相应IP、账号或角色,如果玩家在线,则先踢下线.  $bantime: 封禁时间值，也可以用类似"4d20h3s"表达.  
baninfo 3 语法： baninfo <account|ip|character>    查看某一封禁的详细信息 
banlist 3 语法： banlist <account|ip|character> $NameOrIp    按要求显示封禁列表
----------------------------
.ban account 禁封账号
.ban character 角色
.ban ip IP地址
.baninfo account 被封账号信息
.baninfo character
.baninfo ip
.banlist account
.banlist character
.banlist ip
.bank 打开银行
----------------------------
C
.commands                       // 查看所有可使用的GM命令
.cooldown 技能ID                // 迅速冷却技能CD
.cast 3 语法: .cast #spellid      向目标施放一个#spellid的魔法，没有目标则向你施放 
.cast back 3 语法: .cast back #spellid    使目标向你的人物施放一个#spellid代表的魔法  
.cast self 3 语法: .cast self #spellid    使目标向他自己施放一个#spellid代表的魔法 
.cast target 3 语法: .cast target #spellid      使目标向他的目标施放一个#spellid代表的魔法 
----------------------------
.cast 施放法术
.cast back 使目标对你施放法术
.cast dist 施法距离
.cast self 对自己施放法术
.cast target 对目标施放法术
.character customize 使玩家重新设定外观与名字
.character delete  删除角色
.character level 角色级别
.character rename 使玩家重新命名
.character reputation 角色声望
.combatstop 
.commands 命令列表
.cooldown 使法术冷却时间结束
----------------------------
D
.debug arena                    // 开启单人竞技场
.debug bg                       // 开启单人战场
.debug play sound 声音ID        // 播放声音，例如BOSS喊话
.die                            // 杀死选择生物
.dismount                       // 下马，即脱离坐骑状态
.damage 数字  要打10000 就输入 .damage 10000
.debug anim 2 语法: .debug anim #emoteid    使你施放一个#emoteid代表的动作表情  
.debug getvalue 3 语法: .debug getvalue #field #isInt    获取目标的#field 值，无目标则为你自己。如果输出为 整数则#isInt 取值 1   
.debug playsound 1 语法: .debug playsound #soundid   播放#soundid代表的声音，该声音只为你播放，其他玩家无法听到。 警告：客户端可能存在5000多种声音……  
.debug setvalue 3 语法: .debug setvalue #field #value #isInt   设定目标的#field 值，无目标则为你自己。如果设定为 整数则#isInt 取值 1。  
.debug standstate 2 语法: .debug standstate #emoteid   将你的角色站立时的表情动作改为 #emoteid所指.  
.debug update 3 语法: .debug update #field #value   更新目标的 #field 值为 #value.  如果省略#value则显示 #field的内容.  
.delticket 3 语法： .delticket all    .delticket #num     .delticket $character_name    all：将删除服务器上所有的ticket, $character_name 将删除该角色名的所有ticket, #num 将删除ID为 #num 的ticket.  
.demorph 1 语法： .demorph    恢复使用变形效果变形的角色.  
.die 1 语法： .die    杀死你所选择角色或生物. 如果没选择对象,将杀死你.  
.dismount 0 语法： .dismount    如果你有骑坐骑,将下马.  
.distance 3 语法： .distance    显示你的角色到你选定生物的距离. 
----------------------------
.damage 造成伤害
.debug anim
.debug arena
.debug bg 调试战场可以一个人进入
.debug getvalue 
.debug Mod32Value
.debug play cinematic
.debug play movie
.debug play sound
.debug setvalue
.debug update
.delticket
.demorph 取消变形效果
.die 使目标立即死亡
.dismount 卸载任何坐骑
.distance 距离
----------------------------
E
.event activelist               // 显示当前游戏事件列表
.event start 事件Id             // 执行游戏事件
.event stop  事件Id             // 结束游戏事件
.explorecheat 1       地图作弊,打开全部地图魔兽

F
----------------------------
.flusharenapoints
----------------------------
G
.gm on/off                      // 进入/退出GM模式
.gm fly on/off                  // 进入/退出飞行模式
.go creature 生物刷新Id         // 传送至生物的身边
.go object   物体刷新Id         // 传送至物品的身边
.go X坐标 Y坐标 Z坐标 地图Id    // 传送至指定坐标
.gps                            // 显示当前位置的完整信息
.gobject add 物品ID             // 在当前位置刷新一个世界物体
.gobject add temp 物体ID        // 在当前位置临时刷新一个世界物体，重启后消失
.gobject delete 刷新ID          // 删除世界物体刷新
.gobject near                   // 查询当前距离最新的世界物体信息
.guild create 工会Leader名字 工会名称 // 创建一个工会
.guild delete 工会名称          // 删除一个工会
.guild uninvite 玩家名字        // 从工会提出指定名字的玩家
.gobject add
.gobject delete
.gobject move
.gobject near
.gobject setphase
.gobject target
.gobject turn
.gm on/off 语法: .gm on/off    启用、禁用 GM 模式  
.gm list 0 语法: .gm list  显示在线GM列表  
.gm visible 1 语法: .gm visible on/off    使GM是否对玩家可见  
.go creature 2 语法: .go creature #creature_guid    传送你到#creature_guid代表的生物位置.    .gocreature #creature_name  传送你到名为#creature_name的生物未位置   .gocreature id #creature_id  传送你到 #creature_id为模板的生物位置，如果不止一个，则传送到最早刷出的生物身边  
.go graveyard 2 语法: .go graveyard #graveyardId     传送你到#graveyardId 所代表的墓地。  
.go grid 1 语法: .go grid #gridX #gridY [#mapId]    传送你到 #mapId所代表的地图、#gridX #gridY所构成的网格的中心，省略#mapId 则为当前地图.  
.go object 1 语法: .go object #object_guid    传送你到 #object_guid所代表的物体位置  
.go trigger 2 语法: .go trigger #trigger_id    传送你到 #trigger_id所代表的传送门.如果传送门正在运作将传送你到该传送门的目的地  
.go xy 1 语法: .go xy #x #y [#mapid]    传送你到地图id为#mapid 、坐标为(#x,#y) 的地平面或水平面上，省略 #mapid则为当前地图  
.go xyz 1 语法: .go xyz #x #y #z [#mapid]    传送你到地图id为#mapid 、坐标为(#x,#y,#z)的 地平面或水平面上，省略 #mapid则为当前地图  
.go zonexy 1 语法: .go zonexy #x #y [#zone]   传送你到区域id为#zone 、坐标为(#x,#y) 地平面或水平面上，省略#zone则为当前区域，可以用 .lookup area $namepart 查询相应#zone  
.gobject add 2 语法: .gobject add #id <spawntimeSecs>   从物体末班中添加一个 #id代表的物体到当前世界位置.    Spawntimesecs为刷新时间，非必须.    注意: 此为 .gameobject的翻版  
.gobject delete 2 语法: .gobject delete #go_guid   删除当前世界中 #go_guid所代表的物体  
.gobject move 2 语法: .gobject move #goguid [#x #y #z]    移动#goguid代表的物体到 (#x,#y,#z)的坐标位置 ，省略坐标则为你的当前位置  
.gobject near  3 语法: .gobject near  [#distance]    列出与玩家距离#distance 的所有物体，包括物体guidID and 坐标. 省略 #distance则为默认值10  
.gobject target 2 语法: .gobject target [#go_id|#go_name_part]   锁定 以 #go_id 为模板或名字包含 #go_name_part 所刷出的最近的物体，省略#go_id 和 #go_name_part 则锁定最近的一个物体.  
.gobject turn 2 语法: .gobject turn #goguid    设定 #goguid 所代表的物体朝向为当前你的朝向  
.goname 1 语法： .goname $charactername    传送到角色名为$charactername的玩家那. 用角色名或者点击玩家的头像,例如 组队时候的头像.  
.gps 1 语法： .gps    显示你所选择角色或生物的位置信息. 位置信息包括 X轴, Y轴, Z轴, 和朝向, MapId and ZoneId  
.guid 3 语法： .guid    显示你所选择角色的GUID.  
.guild create 2 语法: .guild create $GuildLeaderName $GuildName  创建一个以 $GuildName 为名的公会，会长为 $GuildLeaderName.  
.guild delete 2 语法: .guild delete $GuildName   删除名为$GuildName的公会.  
.guild invite 2 语法: .guild invite $CharacterName $GuildName   添加名为$CharacterName的角色进入公会 $GuildName.  
.guild rank 2 语法: .guild rank $CharacterName #Rank    设定角色$CharacterName 的公会等级为 #Rank   
guild uninvite 2 语法: .guild uninvite $CharacterName    从公会中将角色 $CharacterName移除 
.guid 工会
.guild create 创建
.guild delete 删除
.guild invite 邀请
.guild rank 
.guild uninvite 踢出
H
.honor add 数量                 // 增加荣誉点数
.hover off            取消腾空
.hover                角色悬空模式.     1 允许悬空  0 禁止悬空.'


I
.instance listbinds             // 查询玩家已绑定副本列表
.instance savedata              // 保存副本数据到数据库
.instance stats                 // 查询副本相关信息
.instance unbind all            // 重置全部副本
----------------------------
.itemmove
----------------------------
J
.jk                             // 迅速脱离战斗状态

K
.kick                           // 将选择的角色踢出游戏
.kan 观看装备
L
.learn all my talents           // 学习本职业全部天赋
.learn all my spells            // 学习本职业全部技能 (不推荐使用，出现各类BUG)
.learn all my pettalents        // 学习全部宠物天赋
.learn all_gm                   //学习GM技能
.learn all_myclass              //学习本职业所有技能
.learn all_myspells             //学习所有职业技能
.levelup 等级                   // 提升等级
.lookup item 物品名字           // 查询物品编号
.lookup quest 任务名字          // 查询任务编号
.lookup creature 生物名字       // 查询生物编号
.lookup object 物体名字         // 查询物体编号
.lookup area 查找地区
.lookup event 事件
.lookup itemset 套装
M
.morp 模型ID变身
.demorph 取消变身
.maxskill                       // 技能满熟练度(武器、防御等等)
.modify hp 数值                 // 调整角色HP
.modify mana 数值               // 调整角色MANA
.modify damage                 //调整你的伤害 Min/Max 
.modify money 数值              // 调整角色金钱
.modify speed 数值              // 调整角色奔跑速度
.modify fly 数值                // 调整角色飞行速度
.modify swim 数值               // 调整角色游泳速度
.modify gender male/female      // 调整角色性别
.modify tp 数值                 // 调整角色天赋点数量
.modify rage # # 将目标的怒气值设置为#/#
.modify energy # # 将目标的能量值设置为#/#（上面2个必须上限和当前血量一起修改） 
.modify runicpower # # 将目标的符文能量设置为#/#
.modify speed # 将目标移动速度设定为初始速度的#倍（#=0.1-10）
.modify bwalk # 将目标移后退度设定为初始速度的#倍（#=0.1-10）
.modify swim # 将目标游泳速度设定为初始速度的#倍（#=0.1-10）
.modify fly # 将目标飞行速度设定为初始速度的#倍（#=0.1-10）
.modify aspeed # 将目标所有速度设定为初始速度的#倍（#=0.1-10）
.modify scale # 将目标体积设定为初始状态的#倍（#=0.00-3）
.modify mount # 为目标召唤一种坐骑（#=1-69）
.modify drunk # 修改目标的醉酒度为#（#=0-100）
.modify faction # 修改你所选择生物的阵营为factionid,flagid为flagid.
.mute 禁言
N
.notify 信息                    // 发送全局公告信息，屏幕中间
.npc add 生物Id                 // 当前位置刷新一个生物
.npc add temp 生物Id            // 当前位置临时刷新一个生物，重启后消失
.npc delete                     // 删除选择生物
.npc move                       // 移动选择生物至当前位置
.npc info                       // 查询选择生物的详细信息
.npc delete item 物品ID         // 从选择NPC的销售菜单内删除物品
.npc add item 物品ID         // 增加售卖物品
.npc playemote 3 语法: .npc playemote #emoteid    使所选生物做 #emoteid 所指的表情动作 
.npc spawntime 2 语法: .npc spawntime #time     调整所选生物的刷新时间#time. 
O 

P
.pet create                     // 将选择生物设置为宠物

Q
.quest add 任务Id               // 增加任务
.quest remove 任务Id            // 删除任务
.quest complete 任务Id          // 完成任务

R
.respawn                        // 立刻刷新附近生物
.revive                         // 复活选择的玩家
.reload all 重新载入
.reload all_area 重载所有地图
.reload all_locales
.reload all_loot 重载所有LOOT（BOSS掉落）
.reload all_quest 重载所有任务
.reload all_spell 重载所有技能
.reload config 重载所有设置
.reload npc_trainer             // 在线加载npc_trainer表
.reload npc_vendor              // 在线加载npc_vendor表
.reload config                  // 在线加载配置文件，仅适合部分配置项
.repairitems                    // 修理选择玩家的全部物品
.reset stats 重置目标所有状态
.reset honor  重置目标荣誉点数
.reset level  重置目标等级
.reset spells  重置目标所有技能
.reset talents 重置目标所有天赋
.reset achievements 重置成就
.reset all 所有
.repairitems 修理装备
S
.save                           // 保存角色数据
.saveall                        // 保存在线所有角色数据
.server info                    // 查看服务器信息
.start                          // 返回炉石旅店
.send items 玩家名字 "感谢对本服的支持" "感谢对本服的支" 25:1 // 邮寄物品
.server restart 60 //服务器重启60秒
.server restart cancel //取消服务器重启
.send items 寄送装备
.send mail 邮件
.send message 信息
.send money 金钱
----------------------------
.server corpses
.server exit 服务器退出
.server idlerestart 
.server idlerestart cancel
.server idleshutdown
.server idleshutdown cancel
.server info 服务器信息
.server motd
.server plimit 
.server restart 服务器重启
.server restart cancel 取消服务器重启
.server set loglevel
.server set motd
.server shutdown
.server shutdown cancel
----------------------------
T
.titles add #ID                 // 增加一个称号
.titles remove #ID              // 删除一个称号

U
.unaura 3 语法: .unaura #spellid   移除目标身上#spellid 所指的光环  
.unban 3 语法 ： unban <account|ip|character> $NameOrIp    取消账号或IP的封禁  
.unlearn 3 语法： .unlearn #startspell #endspell    遗忘你所选择角色的魔法ID从 #startspell 到 #endspell范围内的所有魔法. 如果没有提供#endspell, 则只遗忘id为 #startspell的魔法.  
.unmute 1 语法: .unmute $playerName    取消对玩家 $playerName的禁言. 

V // 无相关命令

W
.wchange 3 语法： .wchange #weathertype #status    设置当前的天气为 #weathertype, 天气状况 #status.    #weathertype 等于1为下雨, 2为下雪, 以及 3为沙尘暴. #status 等于0为禁止天气, 已经 1 为允许.  
.whispers 3 用法： .whispers on|off    允许/禁止 Gm接受其他玩家的悄悄话. 默认时候使用mangosd.conf的设置.  
.wp 2 使用WP命令:    每个路径命令都有自己的描述  
.wp add 2 语法: .wp add [#creature_guid] 为#creature_guid所指生物增加一个路径点，省略#creature_guid 则必须选定一个生物  
.wp export 3 语法: .wp export [#creature_guid] $filename  将#creature_guid所指生物的路径 输出到文件$filename中，省略#creature_guid 则必须选定一个生物  
.wp import 3 语法: .wp import $filename  为所选目标生物导入名为$filename的路径文件 
.wp modify 2 语法: .wp modify [#creature_guid or Select a Creature] add :添加一个确定的路径点  visual：路径点可见  waittime $time：等候时间    emote ID：动作表情    spell ID：魔法技能    text1| text2| text3| text4| text5 <text>：文本    model1 ID：模型1    model2 ID：模型2   move：移动到玩家位置    del：删除    每次只允许用一个参数 
.wp show 2 语法: .wp show [#creature_guid or Select a Creature] on：开启   first：首个路径点   last：最后一个路径点  off：关闭   info：路径点信息  你必须先开启路径显示，然后选择一个可见路径点，才能使用 info参数显示详细信息  

X // 无相关命令
Y // 无相关命令
Z // 无相关命令
```
# 其他内容
## 常用坐标

```
辛特兰 Hinterlands: 
.go 724.846 -3996.11 149.735 0 瑟拉丹 
.go 362.039 -3796.03 153.998 0 隐匿石
.go -69.8514 -4536 7.28923 0 望海崖(错误) 
.go -678.757 -4018.61 238.351 0 辛萨罗
.go -295.384 -3459.12 194.005 0 祖尔祭坛
.go -464.208 -2837.23 110.073 0 沙德拉洛 
.go 266.941 -2751.41 122.544 0 奎尔丹尼小屋
.go -35.7245 -2479.51 120.423 0 祖瓦沙
.go 327.814 -1959.99 197.724 0 鹰巢山


Loch Modan: 洛克莫丹 
.go -5335.61 -2982.58 332.669 0 塞尔萨玛
.go -5201.86 -3136.59 298.761 0 洛克湖 
.go -4939.1 -3423.74 305.595 0 洛克湖中心 
.go -4871.78 -4025.77 313.141 0 莫格罗什要塞 
.go -5675.42 -4244.87 407.002 0 旅行者营地
.go -5755.53 -3998.42 330.436 0 铁环挖掘场 
.go -4771.99 -3329.01 345.504 0 巨石水坝(掉入湖里) 
.go -4231.86 -2361.37 204.069 0 丹奥加兹





赤脊山 Red Ridge Mountains: 
.go -9319 -1937.94 58.0698 0 止水湖边 
.go -9282.98 -2269.64 69.39 0 湖畔镇 
.go -9168.66 -2726.31 90.0426 0 奥瑟尔伐木场 
.go -9385.46 -3039.27 139.437 0 石堡 
.go -8687.39 -2330.38 155.916 0 撕裂者营地


灼热峡谷 Searing Gorge: 
.go -6892.24 -1342.38 239.913 0 大熔炉
.go -7317.34 -1072.33 277.069 0 黑石山
.go -6986.92 -1705.54 241.667 0 煤渣挖掘场


银松森林 Silverpine Forest: 
.go -126.954 815.624 66.0224 0 安伯米尔
.go -416.466 1543.87 17.5941 0 焚木村 
.go -202.557 1666.88 79.7641 0 影牙城堡
.go -577.865 1807.08 8.2492 0 南流海岸 坐标:38,80
.go -757.376 1527.28 17.2465 0 格雷迈恩之墙
.go -987.449 1585.69 53.4298 0 无尽之海 坐标:44,95 
.go 374.222 1083.9 106.509 0 埃利姆矿洞
.go 507.784 1611.33 124.921 0 瑟伯切尔
.go 134.209 1496.64 114.394 0 奥森农场
.go 675.698 974.873 34.8849 0 破旧渡口
.go 762.653 909.072 31.2142 0 洛丹米尔湖(在破旧渡口旁的湖中)
.go 731.866 727.793 37.0975 0 芬里斯岛
.go 960.45 689.611 59.7365 0 芬里斯城堡
.go 982.34 201.239 34.9509 0 黎明岛
.go 1414.28 1073.22 52.4649 0 玛尔丁果园
.go 1285.69 1242.33 52.6914 0 伊瓦农场 
.go 964.877 1238.75 48.0979 0 瓦尔甘牧场 
.go 1035.91 1540.85 30.525 0 亡者农场 
.go 1293.65 1957.71 19.5619 0 粘丝洞
.go 873.391 1852.5 5.0548 0 北流谷


荆棘谷 Strangelthon Vale: 
.go -14692.4 506.162 1.78241 0 蛮荒海岸 
.go -14740.7 -432.482 4.00624 0 哈圭罗岛 
.go -14406.6 419.353 22.3907 0 藏宝海湾
.go -14178.2 712.03 29.1868 0 加尼罗哨站
.go -13274.4 769.951 2.45505 0 血帆营地
.go -13382.6 2.10815 21.8683 0 朱布瓦尔
.go -13152.9 342.729 52.1328 0 竞技场
.go -12352.8 211.452 4.5846 0 格罗姆高营地
.go -11977.4 332.254 3.20626 0 巴拉尔废墟
.go -12258.9 621.105 -68.3247 0 暗礁海(在水底,注意呼吸) 坐标:24,26
.go -11683.1 925.209 3.64735 0 祖丹亚废墟 
.go -11693.9 702.532 49.9689 0 祖昆达
.go -11586.5 -657.662 32.9941 0 库尔森的营地
.go -11845 -1199.29 77.2075 0 祖尔格拉布(错误) 
.go -11712.7 -1758.67 22.4509 0 祖尔格拉布(水中) 
.go -11988 -1484.99 78.9735 0 祖尔格拉布(错误) 
.go -13010.2 -1617.82 146.263 0 荆棘谷(未开发的位置) 坐标 59,43
.go -11311.5 -195.19 76.3198 0 反抗军营地 坐标 37,3


悲伤沼泽 Swamp of Sorrows: 
.go -10103.4 -2431.61 28.4491 0 迷雾谷
.go -10452.5 -3263.59 20.1782 0 斯通纳德 
.go -9980.38 -3568.28 22.0569 0 农田避难所(精英怪) 
.go -10022.2 -4266.67 7.26064 0 芦荟海滩 
.go -10303.5 -3972.28 20.2882 0 泪水之池(精英怪) 
.go -10349.1 -3849.67 -25.6078 0 阿塔拉卡神庙(错误) 


Trisfal Glades: 提瑞斯法林地
.go 1871.14 1587.91 91.2143 0 丧钟镇
.go 2268.03 1333.63 34.7835 0 索利丹农场
.go 2538.92 1407.01 5.69061 0 耳语海岸 坐标:35,44
.go 2803.27 847.119 111.841 0 阿加曼德磨坊
.go 2861.67 398.526 21.1504 0 加伦鬼屋
.go 2955.79 99.8215 3.32947 0 北部海岸 坐标:64,29
.go 2685.13 -198.851 31.4095 0 澈水湖
.go 2563.98 -51.7975 31.7441 0 冈瑟尔的居所
.go 2260.64 289.021 34.1291 0 布瑞尔 
.go 2146.99 658.485 33.59 0 炉灰庄园
.go 1832.44 236.426 60.4171 0 洛丹伦废墟
.go 1628.3 239.925 64.5006 0 幽暗城中心 
.go 1586.48 239.562 -52.149 0 贸易区 
.go 1614.68 643.289 37.0547 0 下水道 
.go 1658.95 303.76 -42.6923 0 军事区 
.go 1786.82 47.9279 -29.1457 0 魔法区
.go 1466.11 49.6445 -62.2932 0 盗贼区
.go 1410.31 430.512 -80.3588 0 炼金房
.go 2032.01 -432.954 35.4329 0 巴尼尔农场
.go 2795.02 -753.797 138.036 0 血色修道院--耳语花园 
.go 2922.59 -740.071 153.983 0 血色修道院--休息区
.go 2853.92 -718.217 148.381 0 血色修道院入口台阶
.go 3040.8 -552.374 122.216 0 血色十字军哨岗
.go 1716.02 -788.217 56.844 0 亡灵壁垒


西瘟疫之地 Western Plaguelands: 
.go 1386.47 -1518.8 72.4034 0 安多哈尔废墟
.go 1064.09 -1718.04 61.1348 0 悔恨岭
.go 981.477 -1821.84 80.4872 0 乌瑟尔之墓 
.go 1738.52 -2319.93 59.5751 0 盖罗恩农场
.go 1487.77 -1884.87 59.2039 0 嚎哭鬼屋
.go 1835.04 -1568.1 59.1818 0 达尔松之泪 
.go 1782.26 -1211.93 59.4057 0 费尔斯通农场
.go 2423.42 -1646.44 60.5098 0 北山伐木场(错误)
.go 2918.72 -1439.39 150.782 0 碧炉谷--玛登霍尔德城堡


Westfall: 西部荒野
.go -10510 1046.89 60.518 0 哨兵岭
.go -11018.4 1513.69 43.0152 0 月溪镇(中心) 
.go -10644.8 1681.3 42.0338 0 阿历克斯顿农场 
.go -10513.9 2075.23 12.1819 0 长滩(金海岸矿洞旁)
.go -10171.8 1195.41 36.4345 0 萨丁农场 
.go -11399.2 1947.85 10.1451 0 西部荒野灯塔
.go -9903.53 1245.26 42.0563 0 法布隆南瓜农场


湿地 Wetlands: 
.go -4053.99 -3450.62 283.383 0 格瑞姆巴扎
.go -3465.16 -3727.56 64.5778 0 龙喉大门
.go -3256.88 -2718.36 9.41205 0 绿带草地 
.go -2849.21 -2220.06 31.3835 0 铁须之墓 
.go -2955 -1022.21 10.0919 0 失落的舰队 
.go -3754.19 -1087.3 -1.71875 0 米奈希尔海湾(水中) 
.go -3740.29 -755.08 10.9643 0 米奈希尔港 
.go -2605.21 -2341.09 83.3551 0 丹莫德(有怪) 
.go -2336.47 -2509.82 85.2212 0 萨多尔大桥 
.go -3522.96 -1848.58 25.1502 0 维尔加挖掘场 

Kalimdor 

灰谷 Ashenvale: 
.go 2867.03 -2595.67 219.911 1 萨维亚 
.go 2745.85 -378.33 108.253 1 阿斯特兰纳
.go 3155 -3702 121 1 大树荫(错误)
.go 4577.61 419.307 33.7161 1 黑海岸--主宰之剑 
.go 2439.16 -3500.08 98.5954 1 卡加希亚要塞(战歌伐木场)
.go 2757.59 -2967.58 143.882 1 萨提纳尔 
.go 2454.38 -2943.27 124 1 夜歌森林(萨提纳尔的外面)
.go 2114.86 -2998.32 111.396 1 冥火岭 
.go 1626.91 -3057.36 89.4942 1 屠魔峡谷 
.go 1775.1 -2679.19 111.666 1 杂丹尼尔兽穴 
.go 1669.74 -1714.14 97.9564 1 沸伦河 
.go 2673.51 -1859.72 188.112 1 林中树居

艾莎拉 Aszhara: 
.go 3466.36 -6656.29 -58.5217 1 (别去,在深水中,可能会死亡) 
.go 3546.8 -5287.96 109.935 1 辛玛洛神殿 

贫瘠之地 Barrens: 
.go -1719.08 -3824.99 12.0836 1 商旅海岸 
.go -3120.86 -2327.89 93.1243 1 巨人旷野 
.go -4095.7 -2305.74 124.914 1 巴尔莫丹
.go -4319.38 -2110.38 80.8662 1 剃刀高地
.go -4619.15 -1850.91 86.0563 1 升降梯
.go -2388.95 -1918.82 96.7422 1 ★千万别用!!!可能会失去你的角色
.go -964.776 -2039.74 81.3491 1 甜水绿洲
.go -1330.17 -3120.07 91.6667 1 死水绿洲 
.go -943.935 -3715.49 11.8385 1 棘齿城
.go -456.263 -2652.7 95.615 1 十字路口
.go 110.197 -1891.39 93.5444 1 遗忘之池
.go 1059.54 -3003.53 91.6441 1 淤泥沼泽

黑海岸 Dark Shore: 
.go 4995.94 82.9197 54.3857 1 古树之林
.go 5028.14 534.745 7.28397 1 暮光海岸 
.go 6439.28 614.957 5.98831 1 码头下面 
.go 6735.43 6.71422 42.7028 1 巴莎兰 
.go 6931.74 -569.077 44.8192 1 碧泉河 
.go 7177.46 -761.607 59.6101 1 奥萨拉克斯之塔 
.go 7373.38 -938.331 32.6196 1 玛塞斯特拉废墟 
.go 7742.92 -769.867 5.22102 1 薄雾海(NPC旁)

凄凉之地 Desolace: 
.go -439.192 1708.22 125.856 1 雷斧堡垒
.go -242.347 764.848 98.7113 1 萨格隆
.go -939.787 1091.4 93.8119 1 科尔卡村
.go -1754.33 967.89 92.5626 1 玛格拉姆村
.go -1879.28 1745.49 78.8892 1 玛诺洛克集会所
.go -2222.47 2522.4 68.4424 1 吉尔吉斯村
.go -1482.87 2855.86 112.854 1 玛拉顿(副本)
.go -592.792 2592.84 15.467 1 萨瑟里斯海岸
.go -1305.19 1837.56 55.731 1 科多兽坟场 
.go -1156.34 1894.49 86.2854 1 幽灵哨岗(错误)

Felwood: 费伍德森林
.go 6200 -1035 387 1 铁木森林
.go 3652.24 928.308 7.01517 1 佐拉姆海岸-灰谷

Feralas: 菲拉斯
.go -2871.76 1885.29 52.6501 1 梦境之树
.go -2858.35 2611.48 58.3777 1 鸦风废墟
.go -4522.22 2038.54 50.1436 1 被遗忘的海岸
.go -5566.04 1449.82 20.1135 1 伊斯迪尔废墟
.go -4369.68 242.294 25.4133 1 莫沙彻营地
.go -4517.1 -780.415 -40.736 1 底地荒野---萨兰钠尔 

杜隆塔尔 Durotar: 
.go -1124.19 -5535.02 8.62076 1 回音群岛
.go -636.469 -4243.52 38.1339 1 试炼谷
.go -844.586 -4924.29 20.8708 1 森金村
.go -1028.63 -4599.8 25.5756 1 科卡尔峭壁 
.go -250.245 -5070.41 22.5875 1 提拉加德城堡 
.go 242.548 -5151.46 1.60441 1 流亡海岸
.go 304.762 -4734.97 9.30458 1 剃刀岭 
.go 114.769 -3758.95 17.8907 1 怒水河
.go 925.127 -4038.29 -13.338 1 雷霆山
.go 1484.36 -4417.03 24.4709 1 奥格瑞玛---力量谷
.go 1525.73 -4968.13 17.1397 1 刃拳海湾(奥格瑞玛附近)
.go 918.715 -5115.69 2.85835 1 死眼海岸 

尘泥沼泽 Dustwallow Marsh: 
.go -3719.26 -2530.63 69.58 1 树荫旅店(坐标 29,49) 
.go -3922.24 -2839.21 44.6212 1 废弃哨塔(坐标 35,54) 
.go -4336.82 -3018.67 33.1744 1 火焰洞穴(精英怪)
.go -4197.56 -2873.76 44.6771 1 黑龙谷(坐标 36,62) 
.go -4354.46 -3275.34 46.0475 1 石缒废墟(精英怪) 
.go -4006.19 -3777.83 40.6804 1 比吉尔的飞艇残骸
.go -3129.38 -2864.51 34.8711 1 嵌墙村 
.go -2855.96 -3422.66 36.7473 1 北点哨塔
.go -3459.39 -4130.3 16.3786 1 警戒哨岗
.go -3012.72 -4345.51 6.83608 1 恐惧海岸
```

## 全13种商业技能相关代码

```
.learn 代码  加满 就是 下面的setskill  比如采矿就是 .setskill 186 375 375
＝＝铭文＝＝
.learn 45363   然后用命令升级到450   
.setskill 773 450 450

＝＝采矿＝＝
初级：.learn 2575
中级：.learn 2576
高级：3564
专家：10248
大师：29354
.setskill 186 375 375

＝＝炼金＝＝
初级：.learn 2259
中级：3101
高级：3464
专家：11611
大师：28596
.setskill 171 375 375

＝＝锻造＝＝
初级：.learn 2018
中级：.learn 3100
高级：3538
专家：9785
大师：29844
.setskill 164 375 375

＝＝裁缝＝＝
初级：.learn 3908
中级：3909
高级：3910
专家：12180
大师：26790
.setskill 197 375 375

＝＝烹饪＝＝
初级：2550
中级： .learn 3102
高级：3413
专家：18260
大师：33359
.setskill 185 375 375

＝＝制皮＝＝
初级：.learn 2108
中级：.learn 3104
高级：3811
专家：10662
大师：32549
.setskill 165 375 375

＝＝钓鱼＝＝
初级： .learn 7620
中级：7731
高级：7732
专家：18248
大师：33095
.setskill 356 375 375

＝＝附魔＝＝
初级： .learn 7411
中级：7412
高级：7413
专家：13920
大师：28029
.setskill 333 375 375

＝＝珠宝加工＝＝
初级： .learn 25229
中级：25230
高级：28894
专家：28895
大师：28897
.setskill 755 375 375

＝＝工程学＝＝
初级： .learn 4036
中级：4037
高级：4038
专家：12656
大师：30350
.setskill 202 375 375
 
＝＝急救＝＝
初级：.learn 3273
中级：3274
高级：7924
专家：10846
大师：27028
.setskill 129 375 375

＝＝*药＝＝
初级：.learn 2366
中级：2368
高级：3570
专家：.learn 11993
大师：28695
.setskill 182 375 375

＝＝剥皮＝＝
初级：.learn 8613
中级：8617
高级：8618
专家：10768
大师：32678
.setskill 393 375 375
```

## 宠物相关

```
.npc add 召唤出来 兽王天赋猎人可以抓
.lookup item 宠物名字 
猎人宝宝-Skoll       35189    传说中身上围绕着电狼宝宝！
猎人宝宝-古德利亚   33776     传说中的幽灵虎！很拉风
猎人宝宝-Arcturis   38453       传说中的幽灵熊！也很拉风
猎人宝宝-洛卡纳哈   32517     这只其实我感觉没有上面的帅

点宝宝
#闪电愤怒
.pet learn 33665
#震慑波
.pet learn 75418
#死亡凋零
.pet learn 72110
#破甲
CC
#符文護盾.pet learn 41431
.pet learn 41431
#狂放的憤怒
.pet learn 39869
#加速光环
.pet learn 13589
#低吼
.pet learn 31334
#狂暴光环
.pet learn 41107
#伤害盾
.pet learn 25059
#原始祝福（变老虎）
.pet learn 24255
#战斗命令（加攻速）


级别-代码-名字-地点-坐标

〖豹〗
攻击速度–快，伤害–普通，防御–低，血–少，移动–快，饮食–挑食
3619 幽灵豹(19-20 召唤, 黑海岸) 
2850 断牙 荒芜之地 
10737 希洛塔姆
10741 赛洛塔姆 白狮子
21724 捕鹰者 

〖狼〗
攻击速度–普通，伤害–高，防御–低，血–少，移动–快，饮食–挑食
521 鲁伯斯 暮色森林 
5356 咆哮者 菲拉斯 


〖迅猛龙〗
攻击速度–普通，伤害–高，防御–低，血–少，移动–普通，饮食–挑食
6581 暴掠龙女王 安戈洛环形山 


〖熊〗
攻击速度–慢，伤害–低，防御–普通，血–多，移动–慢，饮食–杂食
12037 乌索洛克 灰谷 
10806 尤西乌斯 (56 精英, 冬泉谷) 

〖猩猩〗
6585 乌卡洛克(52-53 稀有, 安戈洛环形山)


〖乌龟〗
攻击速度–慢，伤害–低，防御–高，血–中，移动–慢，饮食–挑食

14223 暴躁的本希 奥特兰克山脉
```

```
货币
 
.additem 47241 2147483647 凯旋
.additem 45624 2147483647 征服
.additem 49426 2147483647 冰霜
.additem 40753 2147483647 勇气
.additem 40752 2147483647 英雄
.additem 29434 2147483647 正义
 
.additem 43016 2147483647 达拉然烹饪
.additem 41596 2147483647 达拉然珠宝
.additem 44990 2147483647 勇士徽记
.additem 20559 100   阿拉希
.additem 20558 100   战歌
.additem 20560 100   奥山
.modify honor 2147483647 荣誉值
 
特殊技能，单独做宏技能
 
.respawn
刷新周围怪物
 
.revive
复活当前目标（复活自己需/G启用工会频道后使用）
 
.cooldown 
立即使目标的魔法冷却
 
.start 脱离卡死
```

```
护甲精通
 
.learn 9077   皮甲 
.learn 8737   锁甲 
.learn 750    板甲 
武器技能
 
.learn 201  单手剑
.learn 202  双手剑
.learn 198  单手锤
.learn 199  双手锤
.learn 196  单手斧 
.learn 197  双手斧 
.learn 1180  匕首 
.learn 9116  盾牌
.learn 227  法杖
.learn 5009  魔杖 
.learn 200  长柄武器
.learn 15590  拳套武器
.learn 264  弓 
.learn 5011  弩 
.learn 266  枪械 
.learn 2567  投掷 
.learn 674 双持
.learn 46917 泰坦之握
 
.learn 1804 开锁
.learn 1809 初级开锁
.learn 1810 中级开锁
.learn 6460 高级开锁 
=====================
```

```
各种族技能
 
兽人 
.Learn  20575  命令  宠物伤害提高5% 
.Learn  20573  坚韧  抵抗昏迷的几率提高15% 
.Learn  20574  斧专精  +5 
.Learn  20572  血性狂暴  攻击强度提高6点，持续15秒
 
亡灵 
.Learn  7744  亡灵意志  解除并免疫恐惧、魅惑、催眠，持续5秒 
.Learn  20579  暗影抗性  +10 
.Learn  5227  水下呼吸  水下呼吸时间延长300% 
.Learn  20577  食尸
  
牛头人 
.Learn 20550  耐久   生命值上限提高5%  
.Learn 20549  战争践踏  8码内最多5个敌人昏迷2秒  
.Learn 20552  栽培  草药学+15点  
.Learn 20551  自然抗性  +10 
 
巨魔 
.Learn 20555  再生  生命值回复速度+10% 
.Learn un  弓专精  +5 
.Learn 20558  投掷专精  +5 
.Learn 20554  狂暴  攻击和施法速度提高10%到30% 
.Learn 20557  野兽杀手  对野兽的伤害+5% 
 
血精灵 
.Learn  un  奥术亲和  附魔+10 
.Learn  un  奥术洪流  半径8码内敌人沉没2秒 
.Learn  un  法力分流  
.Learn  un  魔法抗性  +5 
 
人类 
.Learn  20598  人类精魂  精神+10% 
.Learn  20597  剑类武器专精  
.Learn  20599  声望获得速度+10% 
.Learn   20600  感知  提高侦测潜行能力，持续20秒 
.Learn  20864  锤类武器专精 
 
矮人 
.Learn  20596  冰霜抗性  +10 
.Learn  2481  寻找财宝  
.Learn  20595  枪械专精  +5 
.Learn  20594  石像形态  护甲提高10%，免疫流血、毒药、疾病，持续8秒 
 
暗夜精灵 
.Learn  20580  影遁  隐形 
.Learn 20583  自然抗性  +10 
.Learn  20585  精灵之魂  死后变成小精灵移动速度+50% 
.Learn  20582  迅捷  躲闪+1% 
 
侏儒 
.Learn  20592  奥术抗性  +10 
.Learn  20593  工程学专精  +15 
.Learn  20591  开阔思维  智力+5% 
.Learn  20589  逃命专家  解除所有移动限制效果 
 
德莱尼 
.Learn un  宝石切割  珠宝加工+5 
.Learn 20579  暗影抗性 +10 
.Learn  un  纳鲁的祝福 
.Learn  un  鼓舞灵气
```

```
精选BT技能
 
 .learn 40733 神圣之盾 (永久持续无敌）
.Learn 39869 狂野怨灵 伤害提高500%移动速度提高200%直到主动取消
.learn 23965 超级圣疗 无CD
.learn 5384 假死 (脱离战斗神技）
.learn 29968 无限距离无冷却闪现术（只有看不到，没有到不了）
 
.learn 41939 狂暴冲锋 60码无限冲锋
.learn 40936 无CD超范围战争践踏
.learn 40636 大范围群恐 
.learn 41080 吞噬灵魂 秒杀
.learn 70063 霜之哀伤烈怒 （颤抖吧，凡人）
.Learn 23452 隐形2分种 （你看不到我，哥也看不到你） 
 
BUFF技能
 
.learn 56525 王者祝福
.learn 25389 真言术：韧
.learn 48469 野性印记
.learn 22818 魔尔达的勇气 +15%耐力 持续2小时
.learn 22888 屠龙者的咆哮 法术致命+10% 近战致命+5% +140点攻击
 
BUFF物品
 
.add 37094 20 耐力符文卷轴8
.add 49632 20 坚韧符文卷轴
.add 49634    野性之鼓
.add 49633    王者之鼓 
=============== 
特别物品
 
 
.add 12947 2   GM 戒指 1000防御等级戒子
.add 23051     GM 弓   触发盾墙
.add 17        GM 衬衫 34力量 技能效果已取消
.add 18582     GM 武器 能拆能合
.add 36942     霜之哀伤 
 
.add 17142 污染者碎片
.add 25596 偷窥之哨
 
.add 49623 影之哀伤
.add 46017 奶锤
.add 34334 橙弓
.add 32837 蛋刀主
.add 32838 蛋刀副
.add 13262 灰烬使者
.add 17182 萨弗拉斯，炎魔拉格纳罗斯之手
.add 19019 雷霆之怒，逐风者的祝福之剑 
.add 22589 罗克迪洛尔，上古守护者的手杖 
.add 49294 20 灰色宝石袋 声望
.modify rep 阵营 声望值
=================================
```

```
其他
 
.modify rep 46 100000 斯通纳德兽人 
.modify rep 574 100000 凯尔达隆 
 
.modify rep 83 100000 怪物-无组织 
.modify rep 86 100000 怪物-有组织
.modify rep 549 100000 不知-_- (Leathrworking ??) 
.modify rep 550 100000 不知-_- (Engineering-Goblin ??) 
.modify rep 551 100000 不知-_- (Engineering-Gnome ??) 
.modify rep 569 100000 不知-_- (Blacksmithing - Hammersmithing ??) 
.modify rep 570 100000 不知-_- (Blacksmithing - Axesmithing ??) 
.modify rep 571 100000 不知-_- (Blacksmithing - Swordsmithing ??)头衔
 
PS.代码只能刷头衔，头衔代表的成就是不加的。
未分栏，建议单刷自己喜欢的。
===============================
 
.titles add 85 至高者        首位80 级人物
 
===============================
1.种族
 
.titles add 110 奎尔萨拉斯的 首位80 血精灵
 
.titles add 111 阿古斯的     首位80 德莱尼
 
.titles add 112 卡兹莫丹的   首位80 矮人
 
.titles add 113 诺姆瑞跟的   首位80 侏儒
 
.titles add 114 狮心         首位80 人类
 
.titles add 115 伊露恩勇士   首位80 夜精灵
 
.titles add 116 奥格玛英雄   首位80 兽人
 
.titles add 117 旷野奔驰者   首位80 牛头人
 
.titles add 118 暗矛的       首位80 食人妖
 
.titles add 119 被遗忘者     首位80 亡灵
 
=============================
2.职业
 
.titles add 86 无尽风暴的    首位80 洒满
 
.titles add 87 翡翠梦境的    首位80 德鲁伊
 
.titles add 89 先知          首位80 牧师
 
.titles add 90 煞星          首位80 术士
 
.titles add 91 巡者          首位80 猎人
 
.titles add 92 黯刃          首位80 死亡骑士
 
.titles add 93 大法师        首位80 法师
 
.titles add 94 战争使者      首位80 战士
 
.titles add 95 刺客          首位80 盗贼
 
=================================
 
3.声望
 
.titles add 146 艾克索达的
 
.titles add 147 达纳苏斯的
 
.titles add 148 铁炉堡的
 
.titles add 149 暴风城的
 
.titles add 150 奥格玛的
 
.titles add 151 森金的
 
.titles add 152 银月城的
 
.titles add 153 雷霆崖的
 
.titles add 154 幽暗城的
 
=============================
4.PVP
 
.titles add 42 斗士
.titles add 43 决斗者
.titles add 44 竞争者
.titles add 45 挑战者
.titles add 62 残忍斗士
.titles add 71 复仇斗士
.titles add 72 战斗大师
.titles add 80 野蛮斗士
.titles add 81 搜寻者
.titles add 82 竞技场大师
.titles add 157 致命斗士
.titles add 167 狂烈斗士 
.titles add 163 胜利者
.titles add 169 无情斗士 
.titles add 177 愤怒斗士
 
联盟
.titles add 1 士兵
.titles add 2 下士
.titles add 3 士官
.titles add 4 上士
.titles add 5 士官长
.titles add 6 骑士
.titles add 7 中尉骑士
.titles add 8 上尉骑士
.titles add 9 一等骑士
.titles add 10 少校
.titles add 11 指挥官 
.titles add 12 元帅 
.titles add 13 战场元帅
 
.titles add 14 总元帅
.titles add 48 审判者
.titles add 126 联盟的
 
 
部落
.titles add 15 斥候
.titles add 16 蛮兵
.titles add 17 士官
.titles add 18 资深士官
.titles add 19 一等士官
.titles add 20 石卫士
.titles add 21 血卫士
.titles add 22 军团士兵
.titles add 23 百夫长
.titles add 24 勇士
.titles add 25 中将
.titles add 26 将军
.titles add 27 督军
 
.titles add 28 高阶督军
.titles add 47 征服者
.titles add 127 部落的
 
=================================
5.世界事件
.titles add 74 长老
.titles add 75 火焰看守者
.titles add 76 火焰守护者
.titles add 124 万鬼
.titles add 125 博学大师
.titles add 133 酿酒大师
.titles add 134 狂欢客
.titles add 135 爱情傻瓜
.titles add 137 女监护
.titles add 138 男监护 
.titles add 155 贵族
 
=================================
 
6.专业技能
 
.titles add 83 熟练钓手
.titles add 84 大厨
.titles add 96 宗师级炼金师
.titles add 97 宗师级铁匠
.titles add 98 料理铁人
.titles add 99 宗师级附魔师
.titles add 100 宗师级工程师
.titles add 101 医生
.titles add 102 宗师级钓手
.titles add 103 宗师级草药师
.titles add 104 宗师级雕铭师
.titles add 105 宗师级珠宝设计师
 
.titles add 106 宗师级制皮师
.titles add 107 宗师级矿工
.titles add 108 宗师级剥皮师
.titles add 109 宗师级裁缝 
 
==============================
 
7.其他
 
.titles add 46 圣甲虫领主
.titles add 53 那鲁的勇士
.titles add 63 破碎之日的
.titles add 64 阿达欧之手
.titles add 77 人人崇拜的
.titles add 78 探险家
.titles add 79 外交官
.titles add 120 魔法追星者
.titles add 121 暮光镇压者
.titles add 122 纳克萨玛斯征服者
.titles add 123 北裂境英雄
.titles add 128 完美胜利者
 
.titles add 129 冰冻荒原勇士
.titles add 130 大使
.titles add 131 银白勇士
.titles add 132 塞纳留斯卫士
.titles add 139 黑曜杀戮者 
.titles add 140 夜幕的
.titles add 141 不朽
.titles add 142 不灭
.titles add 143 炸鸡勇者
.titles add 144 血帆上将
.titles add 145 神经病
.titles add 156 十字军 
 
.titles add 158 死神葬送者 
.titles add 159 星穹防卫者 
.titles add 160 奥杜亚征服者 
.titles add 161 奥杜亚勇士 
.titles add 164 唤星者 
.titles add 165 星界巡行者 
.titles add 166 泰坦信使 
.titles add 168 旅人 
.titles add 170 大十字军 
.titles add 171 银白防卫者 
.titles add 172 有耐心的 
 
.titles add 173 黎明曙光 
.titles add 174 坠王煞星
.titles add 175 诛王勇者 
.titles add 176 灰烬裁决者的
==============================
```

```
6.未进入宠物栏的坐骑
 
.add 25596 偷窥之哨
.add 33182 迅捷飞行扫帚
.add 33184 迅捷魔法扫帚
.add 33183 魔法扫帚
.add 37011 魔法扫帚
 
.add 8628 斑点夜刃豹缰绳 60速度 
.add 16339 司令坐骑 100速度
=======================================================
国服新加
 
.learn 66122 
.learn 66123 
.learn 66124 魔法公鸡          TCG卡牌简中五版荣耀战场刮刮卡
.learn 74918 白毛犀牛          第十二版刮刮卡坐骑
.learn 74856 炽焰角鹰兽        魔兽世界卡牌《愤怒之门》版本加入 
.learn 75973 X-53型观光火箭    国服招募活动奖励宠物伙伴
 
```

```
 
其他
 
.add 5976 公会徽章
.add 22999 银色黎明战袍  （天灾入侵世界事件获得）绝版 
.add 28788 保卫者战袍    （TBC 开启黑暗之门任务）绝版
.add 25549 血骑士战袍    （部落圣骑士军马任务获得） 
.add 24344 阿古斯之手战袍（德莱尼种族任务获得）
.add 31279 魔化伊利达雷战袍 （任务物品）
.add 19160 冠军外袍
 烹饪
.add 43478 20 丰盛宴席
.add 43492 20 闹鬼的青鱼
.add 33866 20 风暴牛排
.add 12217 20 龙息红椒
.add 6657 20 美味风蛇
.add 43490 20 美味蛋糕
.add 34832 20 拉姆瑟船长的特酿啤酒（钓鱼技能+10）
 
.add 33924 美味巧克力蛋糕
.add 7676 10 菊花茶(潜行者）
 
.add 8529 诺格弗格药剂
.add 21833 古龙水
.add 21829 香水
 
工程
 
.add 49040 基维斯
.add 40768 随身邮箱
.add 40895 侏儒X光眼镜
.add 48933 虫洞生成器：诺森德
.add 18660 世界放大器
.add 23767 打架机器人
.add 18984 空间撕裂器-永望镇
.add 30542 空间撕裂器-52区
.add 18986 安全传送器 加基森
.add 30544 超级安全传送器：托什雷的基地
 
其他
 
.add 35227 地精天气制造机 - 原型机01-B
.add 45057 发条式货车破坏者
.add 33223 钓鱼椅
.add 33079 鱼人装束
.add 35275 辛多雷宝珠
.add 17696 装满水的天蓝水瓶
 
.add 38578 强者的旗帜
.add 46843 银色北伐军旗帜
.add 45011 暴风城旗帜
.add 45013 雷霆崖旗帜
.add 45014 奥格瑞玛旗帜
.add 45015 森金旗帜
.add 45016 幽暗城旗帜
.add 45017 银月城旗帜
.add 45018 铁炉堡旗帜
.add 45019 诺莫瑞根旗帜
.add 45020 埃索达旗帜
.add 45021 达纳苏斯旗帜
 
.add 5462 达图尔的变形魔棒
.add 33047 贝尔比的电眼浪漫护目镜
.add 18258 戈多克食人魔装
.add 1973 欺诈宝珠
.add 34498 纸飞艇工具包
.add 34499 纸飞机工具包
 
.add 52201 穆拉丁的礼物
.add 52251 吉安娜的吊坠
 
.add 49715 永远娇艳的玫瑰
.add 50471 裂心者
.add 37863 烈酒的遥控器
.add 38506 卡洛斯的著名帽子
```

