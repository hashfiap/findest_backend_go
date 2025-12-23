-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 23, 2025 at 09:11 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `findest_go`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `transaction_user` (IN `p_userid` INT)   BEGIN
    SELECT *
    FROM transactions
    WHERE userid = p_userid;
END$$

--
-- Functions
--
CREATE DEFINER=`root`@`localhost` FUNCTION `transaction_success` () RETURNS INT(11) DETERMINISTIC RETURN (
    SELECT COUNT(*)
    FROM transactions
    WHERE status = 'success'
)$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `ID` int(3) NOT NULL,
  `UserID` int(3) NOT NULL,
  `Amount` float NOT NULL,
  `Status` varchar(64) NOT NULL,
  `CreatedAt` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`ID`, `UserID`, `Amount`, `Status`, `CreatedAt`) VALUES
(1, 1, 1000000, 'success', '2025-12-23 17:42:29'),
(2, 1, 500000, 'failed', '2025-12-23 17:43:31'),
(3, 3, 500000, 'success', '2025-12-23 17:47:47'),
(4, 2, 10000, 'success', '2025-12-23 17:48:07'),
(5, 2, 10000, 'success', '2025-12-23 17:48:16'),
(6, 4, 2000000, 'success', '2025-12-23 17:48:27'),
(7, 2, 10000, 'failed', '2025-12-23 17:48:36'),
(8, 2, 50000, 'success', '2025-12-23 17:48:42'),
(9, 2, 20000, 'failed', '2025-12-23 17:48:54'),
(10, 3, 10000000, 'success', '2025-12-23 17:49:01'),
(11, 1, 200000, 'failed', '2025-12-23 17:49:17'),
(12, 4, 1000000, 'success', '2025-12-23 17:49:34'),
(13, 4, 1000000, 'failed', '2025-12-23 17:49:43'),
(14, 4, 500000, 'success', '2025-12-24 00:31:35'),
(15, 1, 200000, 'success', '2025-12-24 00:31:57'),
(16, 2, 10000, 'failed', '2025-12-24 00:32:10'),
(17, 3, 3000000, 'success', '2025-12-24 01:51:07'),
(19, 2, 20000, 'success', '2025-12-24 03:01:18'),
(20, 1, 50000, 'failed', '2025-12-24 03:02:06'),
(21, 4, 300000, 'success', '2025-12-24 03:02:55'),
(23, 4, 100000, 'failed', '2025-12-24 03:03:43'),
(24, 1, 1000000, 'success', '2025-12-24 03:04:20'),
(25, 2, 7000000, 'success', '2025-12-24 03:04:48');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `ID` int(3) NOT NULL,
  `Nama` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`ID`, `Nama`) VALUES
(1, 'Hashfi'),
(2, 'Uday'),
(3, 'Ian'),
(4, 'Bimo');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `UserID` (`UserID`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `ID` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `ID` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `user` (`ID`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
