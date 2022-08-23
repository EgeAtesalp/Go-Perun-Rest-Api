USE `go-perun-rest-api`;

CREATE TABLE `user_session` (
    `alias` VARCHAR(50) NOT NULL,
    `ip_addr` VARCHAR(50) NULL DEFAULT NULL,
    `port` INT NULL DEFAULT NULL,
    PRIMARY KEY (`alias`)
)
COLLATE='utf8mb4_unicode_ci';

CREATE TABLE `smart_contract_hashes` (
    `name` VARCHAR(50) NOT NULL DEFAULT 'default',
    `asset_holder` VARCHAR(100) NULL DEFAULT NULL,
    `adjudicator` VARCHAR(100) NULL DEFAULT NULL,
    PRIMARY KEY (`name`)
)
COLLATE='utf8mb4_unicode_ci';

CREATE TABLE `simple_data_protocol` (
	`ID` INT NOT NULL AUTO_INCREMENT,
	`sender_alias` VARCHAR(50) NOT NULL,
	`receiver_alias` VARCHAR(50) NOT NULL,
	`msg_id` INT NOT NULL DEFAULT 0,
	`paymentchannel` BIGINT NOT NULL DEFAULT 0,
	`msg_id_ref` INT NOT NULL DEFAULT 0,
	`data` JSON NULL,
	`last_updated` TIMESTAMP NOT NULL,
	PRIMARY KEY (`ID`),
	INDEX `sender_alias` (`sender_alias`),
	INDEX `receiver_alias` (`receiver_alias`),
	INDEX `msg_id` (`msg_id`),
	INDEX `paymentchannel` (`paymentchannel`),
	INDEX `msg_id_ref` (`msg_id_ref`)
)
COLLATE='utf8mb4_unicode_ci';

CREATE TABLE `users` (
	`alias` VARCHAR(50) NOT NULL,
	`id` VARCHAR(50) NOT NULL,
	`secret` VARCHAR(50) NOT NULL,
	`secret_key` TEXT NOT NULL,
	`private_key` TEXT NOT NULL,
	PRIMARY KEY (`alias`),
	INDEX `id` (`id`)
)
COLLATE='utf8mb4_unicode_ci';
