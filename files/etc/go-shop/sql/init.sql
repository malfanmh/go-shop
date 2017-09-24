-- phpMyAdmin SQL Dump
-- version 4.2.7.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Sep 24, 2017 at 04:42 AM
-- Server version: 5.6.20
-- PHP Version: 5.5.15

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `go_shop`
--

-- --------------------------------------------------------

--
-- Table structure for table `coupons`
--

CREATE TABLE IF NOT EXISTS `coupons` (
`id` int(11) NOT NULL,
  `code` varchar(16) NOT NULL,
  `description` text NOT NULL,
  `type` enum('disc_amount','disc_percentage','','') NOT NULL,
  `value` float NOT NULL,
  `quantity` int(11) NOT NULL DEFAULT '0',
  `used` int(11) DEFAULT '0',
  `valid_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `valid_until` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `tnc` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(16) NOT NULL DEFAULT '''system''',
  `updated_at` timestamp NULL DEFAULT NULL,
  `updated_by` varchar(16) NOT NULL,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `deleted_by` varchar(16) NOT NULL,
  `status` enum('created','deleted','','') NOT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=3 ;

--
-- Dumping data for table `coupons`
--

INSERT INTO `coupons` (`id`, `code`, `description`, `type`, `value`, `quantity`, `used`, `valid_at`, `valid_until`, `tnc`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`, `deleted_by`, `status`) VALUES
(1, 'bMmtkR', 'Coupon 10k', 'disc_amount', 10000, 40, 0, '2017-09-22 17:00:00', '2017-12-22 17:00:00', 'for all product', '2017-09-23 13:53:12', '''system''', NULL, '', '0000-00-00 00:00:00', '', 'created'),
(2, 'qeJlsk', 'Coupon 30%', 'disc_percentage', 30, 50, 0, '2017-09-22 17:00:00', '2017-12-22 17:00:00', 'for all product', '2017-09-23 16:55:02', '''system''', NULL, '', '0000-00-00 00:00:00', '', 'created');

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE IF NOT EXISTS `customers` (
`id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `phone` varchar(16) NOT NULL,
  `email` varchar(128) NOT NULL,
  `address` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(16) NOT NULL DEFAULT 'system',
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `deleted_by` varchar(16) NOT NULL,
  `status` enum('created','deleted') NOT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=18 ;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `name`, `phone`, `email`, `address`, `created_at`, `created_by`, `deleted_at`, `deleted_by`, `status`) VALUES
(11, 'muhammad alfan miftakhul hudha', '081286312989', 'muhammadalfan.mh@gmail.com', 'Jalan mampang Prapatan XIV no 77 , tegalparang , mampang prapatan , jakarta selatan , jakarta 12190', '2017-09-24 01:39:58', 'system', '0000-00-00 00:00:00', '', 'created'),
(12, 'muhammad alfan miftakhul hudha', '081286312989', 'muhammadalfan.mh@gmail.com', 'Jalan mampang Prapatan XIV no 77 , tegalparang , mampang prapatan , jakarta selatan , jakarta 12190', '2017-09-24 01:41:45', 'system', '0000-00-00 00:00:00', '', 'created'),
(13, 'muhammad alfan miftakhul hudha', '081286312989', 'muhammadalfan.mh@gmail.com', 'Jalan mampang Prapatan XIV no 77 , tegalparang , mampang prapatan , jakarta selatan , jakarta 12190', '2017-09-24 01:49:54', 'system', '0000-00-00 00:00:00', '', 'created'),
(14, 'muhammad alfan miftakhul hudha', '081286312989', 'muhammadalfan.mh@gmail.com', 'Jalan mampang Prapatan XIV no 77 , tegalparang , mampang prapatan , jakarta selatan , jakarta 12190', '2017-09-24 01:54:48', 'system', '0000-00-00 00:00:00', '', 'created'),
(15, 'muhammad alfan miftakhul hudha', '081286312989', 'muhammadalfan.mh@gmail.com', 'Jalan mampang Prapatan XIV no 77 , tegalparang , mampang prapatan , jakarta selatan , jakarta 12190', '2017-09-24 01:56:40', 'system', '0000-00-00 00:00:00', '', 'created'),
(16, 'muhammad alfan miftakhul hudha', '081286312989', 'muhammadalfan.mh@gmail.com', 'Jalan mampang Prapatan XIV no 77 , tegalparang , mampang prapatan , jakarta selatan , jakarta 12190', '2017-09-24 02:09:12', 'system', '0000-00-00 00:00:00', '', 'created'),
(17, 'muhammad alfan miftakhul hudha', '081286312989', 'muhammadalfan.mh@gmail.com', 'Jalan mampang Prapatan XIV no 77 , tegalparang , mampang prapatan , jakarta selatan , jakarta 12190', '2017-09-24 02:37:16', 'system', '0000-00-00 00:00:00', '', 'created');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE IF NOT EXISTS `orders` (
`id` int(11) NOT NULL,
  `code` varchar(32) NOT NULL,
  `coupon` varchar(32) DEFAULT '''''',
  `payment_type` varchar(16) DEFAULT '''bank_transfer''',
  `payment_proof` tinyint(1) NOT NULL DEFAULT '0',
  `customer_id` int(11) DEFAULT '0',
  `shipping_code` varchar(16) DEFAULT '''''',
  `sub_total` float NOT NULL DEFAULT '0',
  `shipment_cost` float NOT NULL DEFAULT '0',
  `discount_value` float NOT NULL DEFAULT '0',
  `total_amount` float NOT NULL DEFAULT '0',
  `state` enum('created','pending','paid','shipped','success','canceled') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(16) NOT NULL DEFAULT 'system',
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `deleted_by` varchar(16) NOT NULL,
  `status` enum('created','deleted','','') NOT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=39 ;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `code`, `coupon`, `payment_type`, `payment_proof`, `customer_id`, `shipping_code`, `sub_total`, `shipment_cost`, `discount_value`, `total_amount`, `state`, `created_at`, `created_by`, `deleted_at`, `deleted_by`, `status`) VALUES
(22, '5441216137283086', '''''', '''bank_transfer''', 0, 0, '''''', 0, 10000, 0, 0, 'created', '2017-09-24 00:28:17', 'system', '0000-00-00 00:00:00', '', 'created'),
(23, '4505593420189303', 'bMmtkR', 'bank_transfer', 0, 0, '''''', 0, 10000, 0, 0, 'created', '2017-09-24 00:33:55', 'system', '0000-00-00 00:00:00', '', 'created'),
(30, '4505593420189303', 'bMmtkR', '''bank_transfer''', 0, 16, '''''', 900000, 10000, 10000, 880000, 'pending', '2017-09-24 02:09:12', 'system', '0000-00-00 00:00:00', '', 'created'),
(31, '4505593420189303', 'bMmtkR', '''bank_transfer''', 1, 16, '''''', 900000, 10000, 10000, 880000, 'paid', '2017-09-24 02:09:51', 'system', '0000-00-00 00:00:00', '', 'created'),
(37, '4505593420189303', 'bMmtkR', '''bank_transfer''', 0, 16, 'Kbdasdmboasdf', 900000, 10000, 10000, 880000, 'shipped', '2017-09-24 02:40:46', 'system', '0000-00-00 00:00:00', '', 'created'),
(38, '4505593420189303', 'bMmtkR', '''bank_transfer''', 0, 16, 'Kbdasdmboasdf', 900000, 10000, 10000, 880000, 'success', '2017-09-24 02:41:39', 'system', '0000-00-00 00:00:00', '', 'created');

-- --------------------------------------------------------

--
-- Table structure for table `order_details`
--

CREATE TABLE IF NOT EXISTS `order_details` (
`id` int(11) NOT NULL,
  `order_code` varchar(32) DEFAULT '0',
  `product_id` int(11) NOT NULL,
  `price` float NOT NULL,
  `quantity` int(11) NOT NULL,
  `amount` float NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(16) NOT NULL DEFAULT 'system',
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `deleted_by` varchar(16) NOT NULL,
  `status` enum('created','deleted','','') NOT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=5 ;

--
-- Dumping data for table `order_details`
--

INSERT INTO `order_details` (`id`, `order_code`, `product_id`, `price`, `quantity`, `amount`, `created_at`, `created_by`, `deleted_at`, `deleted_by`, `status`) VALUES
(1, '5441216137283086', 4, 450000, 3, 1350000, '2017-09-24 00:28:17', 'system', '0000-00-00 00:00:00', '', 'created'),
(2, '5441216137283086', 5, 450000, 1, 450000, '2017-09-24 00:31:26', 'system', '0000-00-00 00:00:00', '', 'created'),
(3, '4505593420189303', 4, 450000, 1, 450000, '2017-09-24 00:33:55', 'system', '0000-00-00 00:00:00', '', 'created'),
(4, '4505593420189303', 5, 450000, 1, 450000, '2017-09-24 00:34:12', 'system', '0000-00-00 00:00:00', '', 'created');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE IF NOT EXISTS `products` (
`id` int(11) NOT NULL,
  `name` varchar(128) NOT NULL,
  `description` text NOT NULL,
  `img_url` varchar(256) NOT NULL,
  `price` float NOT NULL DEFAULT '0',
  `quantity` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(8) NOT NULL DEFAULT '''system''',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(16) DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(16) DEFAULT NULL,
  `status` varchar(8) NOT NULL DEFAULT 'created'
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=6 ;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `description`, `img_url`, `price`, `quantity`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`, `deleted_by`, `status`) VALUES
(4, 'Shoes - Vans Old skool BW', 'saize : 41-44 , color : black & white', '', 450000, 15, '2017-09-24 00:27:49', '''system''', '2017-09-24 00:27:49', NULL, NULL, NULL, 'created'),
(5, 'Shoes - Vans Era', 'saize : 41-44 , color : black & white', '', 450000, 25, '2017-09-24 00:28:50', '''system''', '2017-09-24 00:28:50', NULL, NULL, NULL, 'created');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `coupons`
--
ALTER TABLE `coupons`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `order_details`
--
ALTER TABLE `order_details`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
 ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `coupons`
--
ALTER TABLE `coupons`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=18;
--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=39;
--
-- AUTO_INCREMENT for table `order_details`
--
ALTER TABLE `order_details`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=6;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
