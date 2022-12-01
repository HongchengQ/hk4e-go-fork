-- 基础信息
local base_info = {
	group_id = 144001083
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 83001, gadget_id = 70380274, pos = { x = 1011.717, y = 120.074, z = 889.556 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, isOneoff = true, arguments = { 17 }, area_id = 102, talk_state = 6800217 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

-- 废弃数据
garbages = {
	triggers = {
		{ config_id = 1083002, name = "GADGETTALK_DONE_83002", event = EventType.EVENT_GADGETTALK_DONE, source = "6800217", condition = "", action = "action_EVENT_GADGETTALK_DONE_83002" }
	}
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { },
		gadgets = { 83001 },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================