CREATE TABLE `authz`  (
                          `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                          `access_key` varchar(64) NOT NULL DEFAULT '' COMMENT 'access key',
                          `secret_key` varchar(64) NOT NULL DEFAULT '' COMMENT 'secret key',
                          `role` varchar(64) NOT NULL DEFAULT '' COMMENT 'role: admin/miner/developer',
                          `created_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
                          `updated_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
                          `deleted_at` int                      DEFAULT NULL COMMENT 'deleted time',

                          PRIMARY KEY (`id`) USING BTREE,
                          KEY `idx_access_key` (`access_key`),
                          KEY `idx_secret_key` (`secret_key`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'authz table';

CREATE TABLE `node` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `miner_addr` varchar(64) NOT NULL DEFAULT '' COMMENT 'miner addr',
                        `domain` varchar(64) NOT NULL DEFAULT '' COMMENT 'domain',
                        `payment_secs` int NOT NULL DEFAULT 0 COMMENT 'seconds to be paid',
                        `created_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
                        `updated_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
                        `deleted_at` int                      DEFAULT NULL COMMENT 'deleted time',

                        PRIMARY KEY (`id`)
                            USING BTREE) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'authz table';
