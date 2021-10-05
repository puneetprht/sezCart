CREATE TABLE `ecom`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(250) NULL,
  `username` VARCHAR(250) NULL,
  `password` VARCHAR(250) NULL,
  `token` VARCHAR(500) NULL,
  `cart_id` INT NULL,
  `created_at` DATETIME NULL DEFAULT NOW(),
  PRIMARY KEY (`id`))
COMMENT = 'to store user information';

CREATE TABLE `ecom`.`items` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(250) NULL,
  `created_at` DATETIME NULL DEFAULT now(),
  PRIMARY KEY (`id`))
COMMENT = 'inventory items to be stored		';


CREATE TABLE `ecom`.`carts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NULL,
  `is_purchased` TINYINT NULL,
  `created_at` DATETIME NULL DEFAULT NOW(),
  PRIMARY KEY (`id`),
  INDEX `FK_users_id_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `FK_users_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `ecom`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
COMMENT = 'Users cart';

CREATE TABLE `ecom`.`cartitems` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cart_id` INT NULL,
  `item_id` INT NULL);

CREATE TABLE `ecom`.`orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cart_id` INT NULL,
  `user_id` INT NULL,
  `created_at` DATETIME NULL DEFAULT NOW(),
  PRIMARY KEY (`id`));
