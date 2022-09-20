-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema exchange-provider
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema exchange-provider
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `exchange-provider` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `exchange-provider` ;

-- -----------------------------------------------------
-- Table `exchange-provider`.`orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `exchange-provider`.`orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `seq` INT NULL DEFAULT NULL,
  `status` VARCHAR(100) NULL DEFAULT NULL,
  `exchange` VARCHAR(512) NULL DEFAULT NULL,
  `base_coin` VARCHAR(45) NULL DEFAULT NULL,
  `base_chain` VARCHAR(45) NULL DEFAULT NULL,
  `quote_coin` VARCHAR(45) NULL DEFAULT NULL,
  `quote_chain` VARCHAR(45) NULL DEFAULT NULL,
  `side` VARCHAR(45) NULL DEFAULT NULL,
  `spread_rate` VARCHAR(45) NULL DEFAULT NULL,
  `spread_vol` VARCHAR(45) NULL DEFAULT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  `failed_code` INT NULL DEFAULT NULL,
  `failed_desc` TEXT NULL DEFAULT NULL,
  `meta_data` JSON NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 2384
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `exchange-provider`.`deposits`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `exchange-provider`.`deposits` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `order_id` INT NOT NULL,
  `user_id` INT NULL DEFAULT NULL,
  `status` VARCHAR(45) NULL DEFAULT NULL,
  `exchange` VARCHAR(512) NULL DEFAULT NULL,
  `coin_id` VARCHAR(45) NULL DEFAULT NULL,
  `chain_id` VARCHAR(45) NULL DEFAULT NULL,
  `volume` VARCHAR(50) NULL DEFAULT NULL,
  `address` VARCHAR(1024) NULL DEFAULT NULL,
  `tag` VARCHAR(50) NULL DEFAULT NULL,
  `tx_id` VARCHAR(1024) NULL DEFAULT NULL,
  `failed_desc` TEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_deposites_orders1_idx` (`order_id` ASC) VISIBLE,
  CONSTRAINT `fk_deposites_orders1`
    FOREIGN KEY (`order_id`)
    REFERENCES `exchange-provider`.`orders` (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 1001
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `exchange-provider`.`exchange_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `exchange-provider`.`exchange_orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ex_id` VARCHAR(100) NULL DEFAULT NULL,
  `order_id` INT NOT NULL,
  `user_id` INT NULL DEFAULT NULL,
  `exchange` VARCHAR(512) NULL DEFAULT NULL,
  `symbol` VARCHAR(45) NULL DEFAULT NULL,
  `side` VARCHAR(45) NULL DEFAULT NULL,
  `funds` VARCHAR(45) NULL DEFAULT NULL,
  `size` VARCHAR(45) NULL DEFAULT NULL,
  `fee` VARCHAR(45) NULL DEFAULT NULL,
  `fee_currency` VARCHAR(45) NULL DEFAULT NULL,
  `status` VARCHAR(45) NULL DEFAULT NULL,
  `failed_desc` TEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_exchange_orders_orders1_idx` (`order_id` ASC) VISIBLE,
  CONSTRAINT `fk_exchange_orders_orders1`
    FOREIGN KEY (`order_id`)
    REFERENCES `exchange-provider`.`orders` (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 1001
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `exchange-provider`.`exchanges`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `exchange-provider`.`exchanges` (
  `id` VARCHAR(512) NOT NULL,
  `name` VARCHAR(45) NULL DEFAULT NULL,
  `configs` LONGTEXT NULL DEFAULT NULL,
  `status` VARCHAR(45) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `exchange-provider`.`pair_deposit_limits`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `exchange-provider`.`pair_deposit_limits` (
  `pair` VARCHAR(50) NOT NULL,
  `min_bc` FLOAT NULL DEFAULT NULL,
  `min_qc` FLOAT NULL DEFAULT NULL,
  PRIMARY KEY (`pair`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `exchange-provider`.`pair_spreads`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `exchange-provider`.`pair_spreads` (
  `pair` VARCHAR(50) NOT NULL,
  `spread` FLOAT NULL DEFAULT NULL,
  PRIMARY KEY (`pair`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `exchange-provider`.`user_fees`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `exchange-provider`.`user_fees` (
  `user_id` INT NOT NULL,
  `fee` FLOAT NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `exchange-provider`.`withdrawals`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `exchange-provider`.`withdrawals` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `w_id` VARCHAR(100) NULL DEFAULT NULL,
  `order_id` INT NULL DEFAULT NULL,
  `user_id` INT NULL DEFAULT NULL,
  `status` VARCHAR(45) NULL DEFAULT NULL,
  `address` VARCHAR(1024) NULL DEFAULT NULL,
  `tag` VARCHAR(50) NULL DEFAULT NULL,
  `exchange` VARCHAR(512) NULL DEFAULT NULL,
  `coin` VARCHAR(45) NULL DEFAULT NULL,
  `chain` VARCHAR(45) NULL DEFAULT NULL,
  `total` VARCHAR(45) NULL DEFAULT NULL,
  `fee` VARCHAR(45) NULL DEFAULT NULL,
  `exchange_fee` VARCHAR(45) NULL DEFAULT NULL,
  `executed` VARCHAR(45) NULL DEFAULT NULL,
  `tx_id` VARCHAR(1024) NULL DEFAULT NULL,
  `failed_desc` TEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_withdrawals_orders1_idx` (`order_id` ASC) VISIBLE,
  CONSTRAINT `fk_withdrawals_orders1`
    FOREIGN KEY (`order_id`)
    REFERENCES `exchange-provider`.`orders` (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 1001
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
