-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: bigagent
-- ------------------------------------------------------
-- Server version	5.7.43-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `agent_config`
--

DROP TABLE IF EXISTS `agent_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `agent_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '配置唯一标识符',
  `title` varchar(255) NOT NULL COMMENT '配置标题',
  `status` varchar(15) NOT NULL COMMENT '配置状态',
  `role_name` varchar(255) NOT NULL COMMENT '操作角色',
  `details` varchar(255) NOT NULL COMMENT '备注',
  `ranges` varchar(255) NOT NULL COMMENT '操作范围',
  `auth_name` varchar(255) NOT NULL COMMENT '授鉴类型',
  `data_name` varchar(255) NOT NULL COMMENT '数据类型',
  `protocol` varchar(255) NOT NULL COMMENT '网络协议',
  `host` varchar(255) NOT NULL COMMENT '主机信息',
  `port` int(11) NOT NULL COMMENT '端口信息',
  `path` varchar(255) NOT NULL COMMENT '路径信息',
  `token` varchar(255) NOT NULL COMMENT '认证物料token',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `solt_name` varchar(100) NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `collection_frequency` varchar(100) DEFAULT NULL COMMENT '采集频率',
  `times` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8 COMMENT='agent配置信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `agent_config`
--

LOCK TABLES `agent_config` WRITE;
/*!40000 ALTER TABLE `agent_config` DISABLE KEYS */;
INSERT INTO `agent_config` VALUES (1,'22','有效','11','11','全部范围','token','grpc_cmdb1_stand1','http','example.com:111',8080,'/api','example_token','2024-12-12 02:08:58','2024-12-12 07:29:10','1','2024-12-12 15:29:11',NULL,0),(2,'22','有效','','11','','token','grpc_cmdb1_stand1','','example.com:111',0,'','example_token','2024-12-12 03:13:13','2024-12-12 07:29:12','1','2024-12-12 15:29:12',NULL,0),(11,'test3','有效','','test3','','token','stand1','','192.168.163.2:9000',0,'','123456','2024-12-11 07:42:08','2024-12-11 08:57:29','','2024-12-11 16:57:29',NULL,0),(12,'test4','有效','','test4','','token','stand1','','192.168.168.8:9897',0,'','123456','2024-12-11 07:44:59','2024-12-11 08:58:30','','2024-12-11 16:58:31',NULL,0),(13,'test4','有效','','test4','','token','stand1','','192.168.162.4:7672',0,'','123456','2024-12-11 07:53:24','2024-12-11 09:04:10','slot1','2024-12-11 17:04:11',NULL,0),(14,'test5','有效','','test5','','token','stand1','','192.168.163.88:8865',0,'','123456','2024-12-11 08:24:58','2024-12-12 01:59:21','2','2024-12-12 09:59:22',NULL,0),(15,'test6','有效','','test6','','token','stand1','','192.168.162.44:9987',0,'','123456','2024-12-11 08:31:02','2024-12-12 02:04:41','1','2024-12-12 10:04:41',NULL,0),(16,'111','有效','','111','','token','stand1','','192.168.11.11:1111',0,'','123456','2024-12-12 07:31:48','2024-12-20 09:19:58','1','2024-12-20 17:19:58',NULL,0),(17,'222','有效',' ','222',' ','token','stand1',' ','192.168.22.22:2222',0,' ','123456','2024-12-13 01:51:50','2024-12-20 09:20:43','1','2024-12-20 17:20:44',NULL,0),(18,'test9','有效','','test9','','token','stand1','','192.168.163.99',0,'','123456','2024-12-17 05:58:12','2024-12-20 09:20:52','1','2024-12-20 17:20:53',NULL,0),(19,'test10','有效','','test10','','token','stand1','','192.177.22.22:8888',0,'','123456','2024-12-17 15:00:21','2024-12-20 09:20:58','3','2024-12-20 17:20:58',NULL,0),(20,'test11','有效','','test11','','token','stand1','','192.111.11.111:1111',0,'','123456','2024-12-20 09:18:16','2024-12-20 09:21:01','1','2024-12-20 17:21:02',NULL,0),(21,'test11','有效','','test11','','token','stand1','','192.111.11.111:1111',0,'','123456','2024-12-20 09:18:28','2024-12-20 09:21:06','1','2024-12-20 17:21:07',NULL,0),(22,'test11','有效','','test11','','token','stand1','','192.177.11.11:8888',0,'','123456','2024-12-20 09:20:29','2024-12-20 09:21:08','1','2024-12-20 17:21:09',NULL,0),(23,'test12','有效','','test12','','token','stand1','','192.168.166.222:9999',0,'','123456','2024-12-20 09:21:29','2024-12-27 06:09:22','1','2024-12-27 14:09:22',NULL,0),(24,'test-演示','有效','','test-演示','','token','stand1','','192.168.163.22:8081',0,'','123456','2024-12-24 05:44:26','2024-12-27 06:09:24','1','2024-12-27 14:09:24',NULL,0),(25,'test110','有效','','test110','','token','stand1','','192.168.0.83:9090',0,'','123456','2024-12-25 02:33:22','2024-12-27 06:09:26','1','2024-12-27 14:09:26',NULL,0),(26,'TEST120','有效','','TEST120','','token','stand1','','192.168.11.11:5555',0,'','123456','2024-12-27 05:55:17','2025-01-06 08:16:42','1',NULL,'2s',7),(27,'test999','有效','','test999','','token','stand1','','192.168.44.22:9999',0,'','123456','2025-01-06 08:08:26','2025-01-06 08:25:42','1',NULL,'3s',3);
/*!40000 ALTER TABLE `agent_config` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `agent_info`
--

DROP TABLE IF EXISTS `agent_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `agent_info` (
  `uuid` char(36) NOT NULL COMMENT '机器唯一标识符',
  `net_ip` varchar(15) NOT NULL COMMENT '通信IPv4地址',
  `hostname` varchar(255) NOT NULL COMMENT '主机名',
  `ipv4_first` varchar(15) NOT NULL COMMENT '首个IPv4地址',
  `active` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'agent是否在线（1: 是，0: 否）',
  `status` varchar(255) NOT NULL COMMENT '机器当前状态',
  `action_detail` varchar(255) NOT NULL COMMENT 'agent动作描述',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `grpc_port` varchar(100) NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `platform` varchar(100) DEFAULT NULL,
  `kernel` varchar(100) DEFAULT NULL,
  `os` varchar(100) DEFAULT NULL,
  `machine_type` varchar(100) DEFAULT NULL,
  `arch` varchar(100) DEFAULT NULL,
  `disk_use` varchar(100) DEFAULT NULL,
  `memory_use` varchar(100) DEFAULT NULL,
  `cpu_use` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='agent基础信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `agent_info`
--

LOCK TABLES `agent_info` WRITE;
/*!40000 ALTER TABLE `agent_info` DISABLE KEYS */;
INSERT INTO `agent_info` VALUES ('03560274-043C-050D-8C06-800700080009','192.168.0.83','DESKTOP-CAE4LNE','192.168.0.83',1,'部分数据推送异常','27','2024-12-30 07:19:21','2025-01-06 08:28:08','5678',NULL,'windows','10.0.22621.4601 (WinBuild.160101.0800)','Windows 11 Enterprise10.0','physical','x86_64','{\"C:\":\"87.12%\",\"D:\":\"51.36%\"}','85.00%','20%'),('20230208-BCF4-D4AE-7B1F-BCF4D4AE7B20','127.0.0.1','DESKTOP-FEEVI6T','192.168.31.242',0,'部分数据推送异常','当前配置[19]','2024-12-17 14:47:48','2024-12-17 16:14:59','5678',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `agent_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `roleId` varchar(255) DEFAULT NULL,
  `permissions` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('admin',NULL,'admin','admin',NULL,NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'bigagent'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-01-06 16:41:36
