-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema order_service
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema order_service
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `order_service` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `order_service` ;

-- -----------------------------------------------------
-- Table `order_service`.`orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `order_service`.`orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `status` VARCHAR(100) NULL DEFAULT NULL,
  `exchange` VARCHAR(256) NULL DEFAULT NULL,
  `base_coin` VARCHAR(45) NULL DEFAULT NULL,
  `base_chain` VARCHAR(45) NULL DEFAULT NULL,
  `quote_coin` VARCHAR(45) NULL DEFAULT NULL,
  `quote_chain` VARCHAR(45) NULL DEFAULT NULL,
  `side` VARCHAR(45) NULL DEFAULT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  `broken` TINYINT NULL DEFAULT NULL,
  `break_reason` TEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id`, `user_id`))
ENGINE = InnoDB
AUTO_INCREMENT = 24
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `order_service`.`deposites`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `order_service`.`deposites` (
  `id` INT NULL DEFAULT NULL,
  `order_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `exchange` VARCHAR(256) NULL DEFAULT NULL,
  `volume` VARCHAR(50) NULL DEFAULT NULL,
  `fullfilled` TINYINT NULL DEFAULT NULL,
  `address` VARCHAR(1024) NULL DEFAULT NULL,
  `tag` VARCHAR(50) NULL DEFAULT NULL,
  `tx_id` VARCHAR(1024) NULL DEFAULT NULL,
  PRIMARY KEY (`order_id`, `user_id`),
  INDEX `fk_deposites_orders1_idx` (`order_id` ASC, `user_id` ASC) VISIBLE,
  CONSTRAINT `fk_deposites_orders1`
    FOREIGN KEY (`order_id` , `user_id`)
    REFERENCES `order_service`.`orders` (`id` , `user_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `order_service`.`exchange_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `order_service`.`exchange_orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ex_id` VARCHAR(100) NULL DEFAULT NULL,
  `order_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `exchange` VARCHAR(256) NULL DEFAULT NULL,
  `symbol` VARCHAR(45) NULL DEFAULT NULL,
  `side` VARCHAR(45) NULL DEFAULT NULL,
  `funds` VARCHAR(45) NULL DEFAULT NULL,
  `size` VARCHAR(45) NULL DEFAULT NULL,
  `fee` VARCHAR(45) NULL DEFAULT NULL,
  `fee_currency` VARCHAR(45) NULL DEFAULT NULL,
  `status` VARCHAR(45) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_exchange_orders_orders1_idx` (`order_id` ASC, `user_id` ASC) VISIBLE,
  CONSTRAINT `fk_exchange_orders_orders1`
    FOREIGN KEY (`order_id` , `user_id`)
    REFERENCES `order_service`.`orders` (`id` , `user_id`))
ENGINE = InnoDB
AUTO_INCREMENT = 10
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `order_service`.`withdrawals`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `order_service`.`withdrawals` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `w_id` VARCHAR(100) NULL DEFAULT NULL,
  `order_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `exchange` VARCHAR(256) NULL DEFAULT NULL,
  `coin` VARCHAR(45) NULL DEFAULT NULL,
  `chain` VARCHAR(45) NULL DEFAULT NULL,
  `total` VARCHAR(45) NULL DEFAULT NULL,
  `address` VARCHAR(1024) NULL DEFAULT NULL,
  `tag` VARCHAR(50) NULL DEFAULT NULL,
  `fee` VARCHAR(45) NULL DEFAULT NULL,
  `exchange_fee` VARCHAR(45) NULL DEFAULT NULL,
  `executed` VARCHAR(45) NULL DEFAULT NULL,
  `tx_id` VARCHAR(1024) NULL DEFAULT NULL,
  `status` VARCHAR(45) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_withdrawals_orders1_idx` (`order_id` ASC, `user_id` ASC) VISIBLE,
  CONSTRAINT `fk_withdrawals_orders1`
    FOREIGN KEY (`order_id` , `user_id`)
    REFERENCES `order_service`.`orders` (`id` , `user_id`))
ENGINE = InnoDB
AUTO_INCREMENT = 42
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
