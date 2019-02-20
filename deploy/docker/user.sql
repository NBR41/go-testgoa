INSERT INTO `myinventory`.`user` (`id`,`nickname`,`email`,`activated`,`admin`,`salt`,`password`,`create_ts`,`update_ts`)
VALUES
(NULL, "admin", "admin@myinventory.com", 1,1, '', '', NOW(), NOW()),
(NULL, "new", "new@myinventory.com", 0, 0, '', '', NOW(), NOW()),
(NULL, "user", "user@myinventory.com", 1, 0, '', '', NOW(), NOW()),
(NULL, "nobooks",  "nobooks@myinventory.com", 1, 0, '', '', NOW(), NOW());
