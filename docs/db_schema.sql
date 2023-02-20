CREATE TABLE `user` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` varchar(50) NOT NULL,
	`name` varchar(50) Not NULL,
    `password` varchar(50) Not NUll,
    `email` varchar(50) Not NULL,
    PRIMARY KEY (id)
);