-- 切换到目标库 service_admin_111，如果该库不存在，先创建它
USE service_admin_297;

-- 清空 template 表
TRUNCATE TABLE service_admin_297.template;
-- 拷贝 template 表结构及数据
INSERT INTO service_admin_297.template SELECT * FROM service_admin_281.template;

-- 清空 config_rule 表
TRUNCATE TABLE service_admin_297.config_rule;
-- 拷贝 config_rule 表结构及数据
INSERT INTO service_admin_297.config_rule SELECT * FROM service_admin_281.config_rule;


-- 验证数据是否成功导入
SELECT 'level_data row count:', COUNT(*) FROM service_admin_297.template;
SELECT 'level_order row count:', COUNT(*) FROM service_admin_297.config_rule;

