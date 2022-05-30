-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- 主机： 192.168.1.7:3306
-- 生成日期： 2022-05-02 14:43:41
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
-- 数据库： `dtm_barrier`
--

-- --------------------------------------------------------

--
-- 表的结构 `barrier`
--

CREATE TABLE `barrier` (
  `id` bigint(22) NOT NULL,
  `trans_type` varchar(45) DEFAULT '',
  `gid` varchar(128) DEFAULT '',
  `branch_id` varchar(128) DEFAULT '',
  `op` varchar(45) DEFAULT '',
  `barrier_id` varchar(45) DEFAULT '',
  `reason` varchar(45) DEFAULT '' COMMENT 'the branch type who insert this record',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转储表的索引
--

--
-- 表的索引 `barrier`
--
ALTER TABLE `barrier`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `gid` (`gid`,`branch_id`,`op`,`barrier_id`),
  ADD KEY `create_time` (`create_time`),
  ADD KEY `update_time` (`update_time`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `barrier`
--
ALTER TABLE `barrier`
  MODIFY `id` bigint(22) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
