-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: mariadb:3306
-- Generation Time: Mar 20, 2023 at 01:52 PM
-- Server version: 10.6.4-MariaDB-1:10.6.4+maria~focal
-- PHP Version: 8.0.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `WRAPP`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` char(36) DEFAULT NULL,
  `name` varchar(300) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`user_id`, `name`, `created_at`) VALUES
('00000000-0000-0000-0000-000000000000', '', '0000-00-00 00:00:00'),
('00000000-0000-0000-0000-000000000000', 'testtest', '0000-00-00 00:00:00'),
('5ea00161-dd91-479c-614d-a80dd55732e4', 'testtest', '2023-03-18 11:43:20'),
('5ea00161-dd91-479c-614d-a80dd55732e4', 'testtest', '2023-03-18 11:43:20'),
('6ea00161-dd91-479c-614d-a80dd55732e4', 'testtest', '2023-03-18 11:43:20'),
('6ea00161-dd91-479c-614d-a80dd55732e4', 'testtest', '2023-03-18 11:43:20'),
('6ea00161-dd91-479c-614d-a80dd55732e4', 'testtest', '2023-03-18 11:43:20'),
('7ea00161-dd91-479c-614d-a80dd55732e4', 'testtest', '2023-03-18 11:43:20'),
('7ea00161-dd91-479c-614d-a80dd55732e4', 'testtest', '2023-03-18 11:43:20');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
