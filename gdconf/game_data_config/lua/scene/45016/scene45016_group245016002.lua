-- 基础信息
local base_info = {
	group_id = 245016002
}

-- Trigger变量
local defs = {
	group_id = 245016002,
	fundation_id = 70350145,
	challange_group_id = 245016001
}

-- DEFS_MISCS
local towerPrebuild = 
{
 [2004]= 1,
 [2017]= 3,
 [2019]= 3,
 [2007]= 10, 
 [2012]= 10,
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
	{ config_id = 2001, gadget_id = 70350145, pos = { x = 77.748, y = -9.498, z = -32.524 }, rot = { x = 0.000, y = 269.870, z = 0.000 }, level = 1 },
	{ config_id = 2002, gadget_id = 70350145, pos = { x = 81.224, y = -9.494, z = -30.187 }, rot = { x = 0.000, y = 269.870, z = 0.000 }, level = 1 },
	{ config_id = 2003, gadget_id = 70350145, pos = { x = 74.323, y = -9.487, z = -30.184 }, rot = { x = 0.000, y = 269.870, z = 0.000 }, level = 1 },
	{ config_id = 2004, gadget_id = 70350145, pos = { x = 110.345, y = -9.670, z = -35.969 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2005, gadget_id = 70350145, pos = { x = 101.675, y = -9.668, z = -35.994 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2006, gadget_id = 70350145, pos = { x = 103.208, y = -9.634, z = -21.634 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2007, gadget_id = 70350145, pos = { x = 106.208, y = -9.634, z = -24.634 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2008, gadget_id = 70350145, pos = { x = 109.208, y = -9.634, z = -27.634 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2009, gadget_id = 70350145, pos = { x = 109.208, y = -9.635, z = -21.634 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2010, gadget_id = 70350145, pos = { x = 103.208, y = -9.634, z = -27.634 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2011, gadget_id = 70350145, pos = { x = 103.146, y = -9.628, z = 7.727 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2012, gadget_id = 70350145, pos = { x = 106.146, y = -9.634, z = 4.727 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2013, gadget_id = 70350145, pos = { x = 109.146, y = -9.623, z = 1.727 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2014, gadget_id = 70350145, pos = { x = 109.146, y = -9.631, z = 7.727 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2015, gadget_id = 70350145, pos = { x = 103.146, y = -9.634, z = 1.727 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2016, gadget_id = 70350145, pos = { x = 139.086, y = -13.401, z = -14.589 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2017, gadget_id = 70350145, pos = { x = 136.086, y = -13.401, z = -11.587 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2018, gadget_id = 70350145, pos = { x = 133.086, y = -13.398, z = -14.588 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2019, gadget_id = 70350145, pos = { x = 118.337, y = -9.772, z = 5.134 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2020, gadget_id = 70350145, pos = { x = 118.337, y = -9.752, z = 1.134 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2021, gadget_id = 70350145, pos = { x = 118.337, y = -9.757, z = 9.134 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2048, gadget_id = 70350145, pos = { x = 103.208, y = -9.634, z = -7.238 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2049, gadget_id = 70350145, pos = { x = 106.208, y = -9.634, z = -10.238 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2050, gadget_id = 70350145, pos = { x = 109.208, y = -9.634, z = -13.238 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2051, gadget_id = 70350145, pos = { x = 109.208, y = -9.635, z = -7.238 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2052, gadget_id = 70350145, pos = { x = 103.208, y = -9.634, z = -13.238 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2053, gadget_id = 70350145, pos = { x = 106.013, y = -9.670, z = -35.969 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2055, gadget_id = 70350145, pos = { x = 139.086, y = -9.752, z = 1.398 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 2056, gadget_id = 70350145, pos = { x = 133.086, y = -9.795, z = 1.399 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 点位
points = {
	{ config_id = 2022, pos = { x = 77.748, y = -9.498, z = -32.524 }, rot = { x = 0.000, y = 269.870, z = 0.000 } },
	{ config_id = 2023, pos = { x = 81.224, y = -9.494, z = -30.187 }, rot = { x = 0.000, y = 269.870, z = 0.000 } },
	{ config_id = 2024, pos = { x = 74.323, y = -9.487, z = -30.184 }, rot = { x = 0.000, y = 269.870, z = 0.000 } },
	{ config_id = 2025, pos = { x = 110.276, y = -9.669, z = -35.969 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2026, pos = { x = 101.728, y = -9.670, z = -36.055 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2027, pos = { x = 103.208, y = -9.634, z = -21.634 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2028, pos = { x = 106.208, y = -9.634, z = -24.634 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2029, pos = { x = 109.208, y = -9.634, z = -27.634 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2030, pos = { x = 109.208, y = -9.635, z = -21.634 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2031, pos = { x = 103.208, y = -9.634, z = -27.634 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2032, pos = { x = 103.146, y = -9.628, z = 7.727 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2033, pos = { x = 106.146, y = -9.634, z = 4.727 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2034, pos = { x = 109.146, y = -9.623, z = 1.727 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2035, pos = { x = 109.146, y = -9.631, z = 7.727 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2036, pos = { x = 103.146, y = -9.634, z = 1.727 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2037, pos = { x = 139.086, y = -13.401, z = -14.589 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2038, pos = { x = 136.086, y = -13.401, z = -11.587 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2039, pos = { x = 133.086, y = -13.398, z = -14.588 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2040, pos = { x = 118.337, y = -9.772, z = 5.134 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2041, pos = { x = 118.337, y = -9.752, z = 1.134 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2042, pos = { x = 118.337, y = -9.757, z = 9.134 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2043, pos = { x = 103.208, y = -9.634, z = -7.281 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2044, pos = { x = 106.208, y = -9.634, z = -10.281 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2045, pos = { x = 109.208, y = -9.634, z = -13.280 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2046, pos = { x = 109.140, y = -9.634, z = -7.281 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2047, pos = { x = 103.143, y = -9.629, z = -13.245 }, rot = { x = 0.000, y = 180.000, z = 0.000 } },
	{ config_id = 2054, pos = { x = 106.012, y = -9.670, z = -35.969 }, rot = { x = 0.000, y = 270.000, z = 0.000 } },
	{ config_id = 2057, pos = { x = 139.129, y = -9.757, z = 1.456 }, rot = { x = 0.000, y = 0.000, z = 0.000 } },
	{ config_id = 2058, pos = { x = 133.095, y = -9.795, z = 1.417 }, rot = { x = 0.000, y = 0.000, z = 0.000 } }
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

require "V2_6/TowerDefense_Gear_V3.0"