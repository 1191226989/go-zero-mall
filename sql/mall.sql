-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- 主机： 192.168.1.7:3306
-- 生成日期： 2022-05-02 14:43:26
-- 服务器版本： 5.7.37
-- PHP 版本： 8.0.15

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `mall`
--

-- --------------------------------------------------------

--
-- 表的结构 `order`
--

CREATE TABLE `order` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `uid` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
  `pid` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '产品ID',
  `amount` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '订单金额',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '订单状态',
  `order_no` varchar(255) NOT NULL DEFAULT '' COMMENT '订单号',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `pay`
--

CREATE TABLE `pay` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `uid` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
  `oid` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '订单ID',
  `amount` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '产品金额',
  `source` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '支付方式',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '支付状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `product`
--

CREATE TABLE `product` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '产品名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '产品描述',
  `stock` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '产品库存',
  `amount` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '产品金额',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '产品状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户姓名',
  `gender` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户性别',
  `mobile` varchar(255) NOT NULL DEFAULT '' COMMENT '用户电话',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转储表的索引
--

--
-- 表的索引 `order`
--
ALTER TABLE `order`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_uid` (`uid`),
  ADD KEY `idx_pid` (`pid`);

--
-- 表的索引 `pay`
--
ALTER TABLE `pay`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_uid` (`uid`),
  ADD KEY `idx_oid` (`oid`);

--
-- 表的索引 `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_mobile_unique` (`mobile`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `order`
--
ALTER TABLE `order`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `pay`
--
ALTER TABLE `pay`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `product`
--
ALTER TABLE `product`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
