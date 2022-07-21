-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema order_service
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema order_service
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `order_service` ;
USE `order_service` ;

-- -----------------------------------------------------
-- Table `order_service`.`orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `order_service`.`orders` (
  `id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `status` VARCHAR(100) NULL,
  `exchange` VARCHAR(45) NULL,
  `request_coin` VARCHAR(45) NULL,
  `request_chain` VARCHAR(45) NULL,
  `provide_coin` VARCHAR(45) NULL,
  `provide_chain` VARCHAR(45) NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`, `user_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `order_service`.`deposites`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `order_service`.`deposites` (
  `id` INT NOT NULL,
  `order_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `exchange` VARCHAR(40) NULL,
  `volume` VARCHAR(50) NULL,
  `fullfilled` TINYINT NULL,
  `address` VARCHAR(250) NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_deposites_orders1_idx` (`order_id` ASC, `user_id` ASC) VISIBLE,
  CONSTRAINT `fk_deposites_orders1`
    FOREIGN KEY (`order_id` , `user_id`)
    REFERENCES `order_service`.`orders` (`id` , `user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `order_service`.`withdrawals`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `order_service`.`withdrawals` (
  `id` VARCHAR(50) NOT NULL,
  `order_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `exchange` VARCHAR(45) NULL,
  `coin` VARCHAR(45) NULL,
  `chain` VARCHAR(45) NULL,
  `total` VARCHAR(45) NULL,
  `address` VARCHAR(45) NULL,
  `fee` VARCHAR(45) NULL,
  `exchange_fee` VARCHAR(45) NULL,
  `executed` VARCHAR(45) NULL,
  `tx_id` VARCHAR(250) NULL,
  `status` VARCHAR(45) NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_withdrawals_orders1_idx` (`order_id` ASC, `user_id` ASC) VISIBLE,
  CONSTRAINT `fk_withdrawals_orders1`
    FOREIGN KEY (`order_id` , `user_id`)
    REFERENCES `order_service`.`orders` (`id` , `user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `order_service`.`exchange_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `order_service`.`exchange_orders` (
  `id` VARCHAR(50) NOT NULL,
  `order_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `exchange` VARCHAR(45) NULL,
  `symbol` VARCHAR(45) NULL,
  `side` VARCHAR(45) NULL,
  `funds` VARCHAR(45) NULL,
  `size` VARCHAR(45) NULL,
  `fee` VARCHAR(45) NULL,
  `fee_currency` VARCHAR(45) NULL,
  `status` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_exchange_orders_orders1_idx` (`order_id` ASC, `user_id` ASC) VISIBLE,
  CONSTRAINT `fk_exchange_orders_orders1`
    FOREIGN KEY (`order_id` , `user_id`)
    REFERENCES `order_service`.`orders` (`id` , `user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
