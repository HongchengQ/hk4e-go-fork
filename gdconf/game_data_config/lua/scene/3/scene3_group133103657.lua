-- 基础信息
local base_info = {
	group_id = 133103657
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 657001, monster_id = 28060601, pos = { x = 336.337, y = 312.213, z = 1864.159 }, rot = { x = 0.000, y = 29.735, z = 0.000 }, level = 27, drop_tag = "走兽", pose_id = 1, area_id = 6 },
	{ config_id = 657002, monster_id = 28060601, pos = { x = 339.142, y = 311.828, z = 1859.908 }, rot = { x = 0.000, y = 282.397, z = 0.000 }, level = 27, drop_tag = "走兽", pose_id = 2, area_id = 6 },
	{ config_id = 657003, monster_id = 21010201, pos = { x = 334.897, y = 311.844, z = 1860.470 }, rot = { x = 0.000, y = 97.864, z = 0.000 }, level = 27, drop_tag = "丘丘人", pose_id = 9012, area_id = 6 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
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
		monsters = { 657001, 657002, 657003 },
		gadgets = { },
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