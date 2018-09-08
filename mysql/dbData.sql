INSERT INTO `_admin_role` (`id`, `role_name`, `enable`) VALUES (1, 'admin', '1'), (2, 'user', '1');

INSERT INTO `_admin_action` (`id`, `action`, `display_name`, `type`, `enable`) VALUES
(NULL, 'admin/grid', 'Admin Grid Permission', 'admin', 1),
(NULL, 'admin/edit', 'Admin User Edit Details', 'admin', 1),
(NULL, 'admin/add', 'Admin Add New User', 'admin', 1),
(NULL, 'role/view', 'Admin User Role Details', 'role', 1),
(NULL, 'role/view/id','Admin User Role Details By ID', 'role', 1),
(NULL, 'role/edit', 'Admin User Edit Details', 'role', 1),
(NULL, 'role/add', 'Admin User to Add New Roles', 'role', 1),
(NULL, 'customer/grid', 'Customer Grid', 'customer', 1),
(NULL, 'customer/edit', 'Customer Edit Personal Details', 'customer', 1),
(NULL, 'customer/add', 'Customer Add New User', 'customer', 1),
(NULL, 'customer/view', 'Customer view Details', 'customer', 1),
(NULL, 'customer/export/csv', 'Customers Export CSV', 'customer', 1);

INSERT INTO `_admin_permission` (`id`, `role_id`, `action_id`, `enable`) VALUES 
(NULL,  '1', (SELECT id FROM _admin_action where action = "admin/grid"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "admin/edit"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "admin/add"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "role/view"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "role/view/id"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "role/edit"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "role/add"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "customer/grid"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "customer/edit"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "customer/add"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "customer/view"), '1'),
(NULL,  '1', (SELECT id FROM _admin_action where action = "customer/export/csv"), '1');


INSERT INTO `_config` (`id`, `configType`, `key`, `value`, `description`) VALUES
(NULL, 'admin/request/s3/region', 'S3_REGION', '', 'S3 BUCKET region for Store Import file'),
(NULL, 'admin/request/s3/bucket', 'S3_BUCKET', '', 'S3 BUCKET Name');