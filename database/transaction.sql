-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 12, 2023 at 07:00 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bank_sampah`
--
-- CREATE DATABASE IF NOT EXISTS `bank_sampah` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
-- USE `bank_sampah`;

DELIMITER $$
--
-- Procedures
--
DROP PROCEDURE IF EXISTS `Deposit`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `Deposit` (IN `nik_param` INT, IN `amount_param` DECIMAL(19,2))  BEGIN
    -- Update the user's balance
    UPDATE users SET saldo = saldo + amount_param WHERE NIK = nik_param;
    
    -- Record the transaction in the transaction_history table
    -- INSERT INTO transaction_history (UID, TransactionType, Amount)
    -- SELECT UID, 'deposit', amount_param FROM user_data WHERE NIK = nik_param;
    
    SELECT 'Deposit successful' AS Result;
END$$

DROP PROCEDURE IF EXISTS `InsertUser`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `InsertUser` (IN `NIK_param` INT, IN `Name_param` VARCHAR(255), IN `Password_param` VARCHAR(255), IN `Gender_param` ENUM('pria','wanita'), IN `Address_param` VARCHAR(255), IN `Role_param` ENUM('super-admin','admin','user'), IN `Balance_param` DECIMAL(19,2))  BEGIN
    DECLARE generatedUID CHAR(36);
    SET generatedUID = LOWER(UUID());

    INSERT INTO user_data (UID, NIK, Name, Password, Gender, Address, Role, Balance)
    VALUES (generatedUID, NIK_param, Name_param, Password_param, Gender_param, Address_param, Role_param, Balance_param);
END$$

DROP PROCEDURE IF EXISTS `Withdraw`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `Withdraw` (IN `nik_param` INT, IN `amount_param` DECIMAL(19,2))  BEGIN
    DECLARE current_balance DECIMAL(19, 2);
    
    -- Get the current balance of the user
    SELECT Balance INTO current_balance FROM user_data WHERE NIK = nik_param;
    
    -- Check if the withdrawal amount is valid
    IF current_balance >= amount_param THEN
        -- Update the user's balance
        UPDATE user_data SET Balance = Balance - amount_param WHERE NIK = nik_param;
        
        -- Record the transaction in the transaction_history table
        INSERT INTO transaction_history (UID, TransactionType, Amount)
        SELECT UID, 'withdraw', amount_param FROM user_data WHERE NIK = nik_param;
        
        SELECT 'Transaction successful' AS Result;
    ELSE
        SELECT 'Insufficient balance' AS Result;
    END IF;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `transaction_history`
--
-- Creation: Aug 12, 2023 at 04:59 AM
-- Last update: Aug 12, 2023 at 04:59 AM
--

-- DROP TABLE IF EXISTS `transaction_history`;
CREATE TABLE IF NOT EXISTS `transaction_history` (
  `TransactionID` int(11) NOT NULL AUTO_INCREMENT,
  `UID` char(36) DEFAULT NULL,
  `TransactionType` enum('withdraw','deposit') NOT NULL,
  `Amount` decimal(19,2) NOT NULL,
  `TransactionDate` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`TransactionID`),
  KEY `UID` (`UID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

--
-- RELATIONSHIPS FOR TABLE `transaction_history`:
--

--
-- Dumping data for table `transaction_history`
--
-- --------------------------------------------------------

--
-- Table structure for table `user_data`
--
-- Creation: Aug 12, 2023 at 04:14 AM
-- Last update: Aug 12, 2023 at 04:55 AM
--

-- DROP TABLE IF EXISTS `user_data`;
CREATE TABLE IF NOT EXISTS `user_data` (
  `UID` char(36) NOT NULL,
  `NIK` int(11) DEFAULT NULL,
  `Name` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Gender` enum('pria','wanita') NOT NULL,
  `Address` varchar(255) DEFAULT NULL,
  `Role` enum('super-admin','admin','user') NOT NULL,
  `Balance` decimal(19,2) NOT NULL,
  PRIMARY KEY (`UID`),
  UNIQUE KEY `NIK` (`NIK`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELATIONSHIPS FOR TABLE `user_data`:
--

--
-- Dumping data for table `user_data`
--

--
-- Metadata
--
USE `phpmyadmin`;

--
-- Metadata for table transaction_history
--

--
-- Metadata for table user_data
--

--
-- Metadata for database bank_sampah
--
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 12, 2023 at 07:00 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bank_sampah`
--
-- CREATE DATABASE IF NOT EXISTS `bank_sampah` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
-- USE `bank_sampah`;

DELIMITER $$
--
-- Procedures
--
-- DROP PROCEDURE IF EXISTS `Deposit`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `Deposit` (IN `nik_param` INT, IN `amount_param` DECIMAL(19,2))  BEGIN
    -- Update the user's balance
    UPDATE user_data SET Balance = Balance + amount_param WHERE NIK = nik_param;
    
    -- Record the transaction in the transaction_history table
    INSERT INTO transaction_history (UID, TransactionType, Amount)
    SELECT UID, 'deposit', amount_param FROM user_data WHERE NIK = nik_param;
    
    SELECT 'Deposit successful' AS Result;
END$$

-- DROP PROCEDURE IF EXISTS `InsertUser`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `InsertUser` (IN `NIK_param` INT, IN `Name_param` VARCHAR(255), IN `Password_param` VARCHAR(255), IN `Gender_param` ENUM('pria','wanita'), IN `Address_param` VARCHAR(255), IN `Role_param` ENUM('super-admin','admin','user'), IN `Balance_param` DECIMAL(19,2))  BEGIN
    DECLARE generatedUID CHAR(36);
    SET generatedUID = LOWER(UUID());

    INSERT INTO user_data (UID, NIK, Name, Password, Gender, Address, Role, Balance)
    VALUES (generatedUID, NIK_param, Name_param, Password_param, Gender_param, Address_param, Role_param, Balance_param);
END$$

-- DROP PROCEDURE IF EXISTS `Withdraw`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `Withdraw` (IN `nik_param` INT, IN `amount_param` DECIMAL(19,2))  BEGIN
    DECLARE current_balance DECIMAL(19, 2);
    
    -- Get the current balance of the user
    SELECT Balance INTO current_balance FROM user_data WHERE NIK = nik_param;
    
    -- Check if the withdrawal amount is valid
    IF current_balance >= amount_param THEN
        -- Update the user's balance
        UPDATE user_data SET Balance = Balance - amount_param WHERE NIK = nik_param;
        
        -- Record the transaction in the transaction_history table
        INSERT INTO transaction_history (UID, TransactionType, Amount)
        SELECT UID, 'withdraw', amount_param FROM user_data WHERE NIK = nik_param;
        
        SELECT 'Transaction successful' AS Result;
    ELSE
        SELECT 'Insufficient balance' AS Result;
    END IF;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `transaction_history`
--
-- Creation: Aug 12, 2023 at 04:59 AM
-- Last update: Aug 12, 2023 at 04:59 AM
--

-- DROP TABLE IF EXISTS `transaction_history`;
CREATE TABLE IF NOT EXISTS `transaction_history` (
  `TransactionID` int(11) NOT NULL AUTO_INCREMENT,
  `UID` char(36) DEFAULT NULL,
  `TransactionType` enum('withdraw','deposit') NOT NULL,
  `Amount` decimal(19,2) NOT NULL,
  `TransactionDate` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`TransactionID`),
  KEY `UID` (`UID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

--
-- RELATIONSHIPS FOR TABLE `transaction_history`:
--

--
-- Dumping data for table `transaction_history`
--

INSERT DELAYED IGNORE INTO `transaction_history` (`TransactionID`, `UID`, `TransactionType`, `Amount`, `TransactionDate`) VALUES
(1, 'd3528c50-3876-11ee-9b5d-0068eb7a4490', 'deposit', '4000.00', '2023-08-12 04:53:49'),
(2, 'd3528c50-3876-11ee-9b5d-0068eb7a4490', 'deposit', '4000.00', '2023-08-12 04:55:33');

-- --------------------------------------------------------

--
-- Table structure for table `user_data`
--
-- Creation: Aug 12, 2023 at 04:14 AM
-- Last update: Aug 12, 2023 at 04:55 AM
--

-- DROP TABLE IF EXISTS `user_data`;
CREATE TABLE IF NOT EXISTS `user_data` (
  `UID` char(36) NOT NULL,
  `NIK` int(11) DEFAULT NULL,
  `Name` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Gender` enum('pria','wanita') NOT NULL,
  `Address` varchar(255) DEFAULT NULL,
  `Role` enum('super-admin','admin','user') NOT NULL,
  `Balance` decimal(19,2) NOT NULL,
  PRIMARY KEY (`UID`),
  UNIQUE KEY `NIK` (`NIK`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELATIONSHIPS FOR TABLE `user_data`:
--

--
-- Dumping data for table `user_data`
--

--
-- Metadata
--
USE `phpmyadmin`;

--
-- Metadata for table transaction_history
--

--
-- Metadata for table user_data
--

--
-- Metadata for database bank_sampah
--
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
