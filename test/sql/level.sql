-- 切换到目标库 service_admin_297
USE service_admin_297;

-- 删除 level_data 表（如果存在）
DROP TABLE IF EXISTS service_admin_297.level_data;

-- 拷贝 level_data 表结构
CREATE TABLE service_admin_297.level_data LIKE service_admin_281.level_data;

-- 拷贝 level_data 表数据
INSERT INTO service_admin_297.level_data SELECT * FROM service_admin_281.level_data;

-- 删除 level_order 表（如果存在）
DROP TABLE IF EXISTS service_admin_297.level_order;

-- 拷贝 level_order 表结构
CREATE TABLE service_admin_297.level_order LIKE service_admin_281.level_order;

-- 拷贝 level_order 表数据
INSERT INTO service_admin_297.level_order SELECT * FROM service_admin_281.level_order;

-- 验证数据是否成功导入
SELECT 'level_data row count:', COUNT(*) FROM service_admin_297.level_data;
SELECT 'level_order row count:', COUNT(*) FROM service_admin_297.level_order;
