
--
-- Table structure for table `_admin_action`
--

CREATE TABLE `_admin_action` (
  `id` int(11) NOT NULL,
  `action` varchar(225) NOT NULL,
  `display_name` text NOT NULL,
  `type` varchar(45) NOT NULL,
  `enable` int(11) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--
-- Table structure for table `_admin_api_log`
--

CREATE TABLE `_admin_api_log` (
  `id` int(11) NOT NULL,
  `method` varchar(10) NOT NULL,
  `requestPath` varchar(50) NOT NULL,
  `urlQueryParams` varchar(100) DEFAULT NULL,
  `requestAgent` varchar(500) NOT NULL,
  `requestIp` varchar(30) NOT NULL,
  `requestParameter` varchar(500) NOT NULL,
  `responseParameter` varchar(250) NOT NULL,
  `permission` tinyint(1) NOT NULL,
  `createdTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedOn` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `currentUser` int(64) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Table structure for table `_admin_role`
--

CREATE TABLE `_admin_role` (
  `id` int(11) NOT NULL,
  `role_name` varchar(225) NOT NULL,
  `enable` int(11) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Table structure for table `_admin_permission`
--

CREATE TABLE `_admin_permission` (
  `id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  `action_id` int(11) NOT NULL,
  `enable` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



--
-- Table structure for table `_admin_user`
--

CREATE TABLE `_admin_user` (
  `id` int(11) NOT NULL,
  `name` varchar(225) NOT NULL,
  `email` varchar(225) NOT NULL,
  `role` varchar(225) NOT NULL,
  `picture` text NOT NULL,
  `hd` text NOT NULL,
  `verified_email` tinyint(1) NOT NULL,
  `auth_id` varchar(256) NOT NULL,
  `enable` int(11) DEFAULT '1',
  `firebase_auth_id` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--
-- Table structure for table `_api_request_agent`
--

CREATE TABLE `_api_request_agent` (
  `id` int(11) NOT NULL,
  `requestAgent` varchar(250) NOT NULL,
  `createdAt` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Table structure for table `_user`
--

CREATE TABLE `_user` (
  `id` int(11) NOT NULL,
  `email` varchar(64) DEFAULT NULL COMMENT 'user email id',
  `name` varchar(64) DEFAULT NULL COMMENT 'signup name given',
  `mobile` varchar(13) DEFAULT NULL COMMENT 'user entered mobile number',
  `creationTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'user creation time'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--
-- Table structure for table `_config`
--

CREATE TABLE `_config` (
  `id` int(11) NOT NULL,
  `configType` varchar(45) NOT NULL COMMENT 'config type -Application details',
  `key` varchar(45) NOT NULL COMMENT 'Name of the config',
  `value` varchar(256) NOT NULL COMMENT 'Value of config',
  `description` varchar(300) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--
-- Indexes for table `_admin_action`
--
ALTER TABLE `_admin_action`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `action` (`action`);

--
-- Indexes for table `_admin_api_log`
--
ALTER TABLE `_admin_api_log`
  ADD PRIMARY KEY (`id`),
  ADD KEY `_admin_user_fk_idx` (`currentUser`);

--
-- Indexes for table `_admin_permission`
--
ALTER TABLE `_admin_permission`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `role_id` (`role_id`,`action_id`),
  ADD KEY `__admin_permission_action` (`action_id`);

--
-- Indexes for table `_admin_role`
--
ALTER TABLE `_admin_role`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `role_name` (`role_name`);

--
-- Indexes for table `_admin_user`
--
ALTER TABLE `_admin_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `email_2` (`email`,`role`),
  ADD KEY `fk_admin_user_role` (`role`);

--
-- Indexes for table `_api_request_agent`
--
ALTER TABLE `_api_request_agent`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `_user`
--
ALTER TABLE `_user`
  ADD PRIMARY KEY (`id`);
  
--
-- Indexes for table `_config`
--
ALTER TABLE `_config`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for table `_admin_action`
--
ALTER TABLE `_admin_action`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `_admin_api_log`
--
ALTER TABLE `_admin_api_log`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `_admin_permission`
--
ALTER TABLE `_admin_permission`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `_admin_role`
--
ALTER TABLE `_admin_role`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `_admin_user`
--
ALTER TABLE `_admin_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `_api_request_agent`
--
ALTER TABLE `_api_request_agent`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `_user`
--
ALTER TABLE `_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for table `_admin_permission`
--
ALTER TABLE `_admin_permission`
  ADD CONSTRAINT `fk_admin_permission_action` FOREIGN KEY (`action_id`) REFERENCES `_admin_action` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `fk_admin_permission_role` FOREIGN KEY (`role_id`) REFERENCES `_admin_role` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `_admin_user`
--
ALTER TABLE `_admin_user`
  ADD CONSTRAINT `fk_admin_user_role` FOREIGN KEY (`role`) REFERENCES `_admin_role` (`role_name`);

--
-- AUTO_INCREMENT for table `_config`
--
ALTER TABLE `_config`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;