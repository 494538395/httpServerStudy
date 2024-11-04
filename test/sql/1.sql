-- 切换到目标库 service_admin_111，如果该库不存在，先创建它
USE service_admin_297;

-- 拷贝 level_data 表结构及数据
INSERT INTO service_admin_297.level_data SELECT * FROM service_admin_281.level_data;

-- 拷贝 level_order 表结构及数据
INSERT INTO service_admin_297.level_order SELECT * FROM service_admin_281.level_order;

-- 拷贝 level_type 表结构及数据
INSERT INTO service_admin_297.level_type SELECT * FROM service_admin_281.level_type;

-- 验证数据是否成功导入
SELECT 'level_data row count:', COUNT(*) FROM service_admin_297.level_data;
SELECT 'level_order row count:', COUNT(*) FROM service_admin_297.level_order;
SELECT 'level_type row count:', COUNT(*) FROM service_admin_297.level_type;
