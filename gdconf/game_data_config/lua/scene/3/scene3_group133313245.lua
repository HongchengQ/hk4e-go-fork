-- 基础信息
local base_info = {
	group_id = 133313245
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
}

-- 区域
regions = {
	{ config_id = 245001, shape = RegionShape.POLYGON, pos = { x = -342.137, y = 150.000, z = 5661.823 }, height = 300.000, point_array = { { x = -300.265, y = 5181.429 }, { x = -682.766, y = 5667.816 }, { x = -664.095, y = 6142.217 }, { x = -1.508, y = 6126.895 }, { x = -10.731, y = 5186.642 } }, area_id = 32, vision_type_list = { 33130016 } }
}

-- 触发器
triggers = {
}

-- 变量
variables = {
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
		gadgets = { },
		regions = { 245001 },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================