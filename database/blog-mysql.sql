-- MySQL dump 10.13  Distrib 8.0.23, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: go-blog
-- ------------------------------------------------------
-- Server version	5.7.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ad`
--
use go-blog;

DROP TABLE IF EXISTS `ad`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ad` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) DEFAULT '' COMMENT '名称',
  `image` varchar(255) DEFAULT '' COMMENT '图片链接',
  `url` varchar(255) DEFAULT '' COMMENT 'URL',
  `gid` varchar(100) DEFAULT '' COMMENT '组别ID',
  `group` varchar(100) DEFAULT '' COMMENT '组别名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ad`
--

LOCK TABLES `ad` WRITE;
/*!40000 ALTER TABLE `ad` DISABLE KEYS */;
/*!40000 ALTER TABLE `ad` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) DEFAULT '0' COMMENT '用户ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `category_id` int(11) NOT NULL COMMENT '分类ID',
  `tag` varchar(255) NOT NULL DEFAULT '' COMMENT 'Tag',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '摘要',
  `desc` longtext NOT NULL COMMENT '详情',
  `html` longtext NOT NULL COMMENT 'html',
  `pv` int(11) DEFAULT '0' COMMENT 'px',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `review` int(11) DEFAULT '0' COMMENT '评论',
  `recommend` int(11) NOT NULL DEFAULT '0' COMMENT '是否顶置，0否；1是，默认否',
  `like` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数量',
  `cover` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `other` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
INSERT INTO `article` VALUES (1,1,'欢迎使用Go blog',1,'Go blog','欢迎使用Go blog','欢迎使用Go blog','<p>欢迎使用Go blog</p>\r\n',1,'2021-04-20 17:48:39','2021-04-20 17:48:43',1,0,0,0,'','/','');
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bbs`
--

DROP TABLE IF EXISTS `bbs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bbs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `customer_id` int(11) DEFAULT '0' COMMENT '用户ID',
  `content` longtext NOT NULL COMMENT '内容',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `review` int(11) DEFAULT '0' COMMENT '评论',
  `like` int(11) DEFAULT '0' COMMENT '点赞',
  `topic_id` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bbs`
--

LOCK TABLES `bbs` WRITE;
/*!40000 ALTER TABLE `bbs` DISABLE KEYS */;
/*!40000 ALTER TABLE `bbs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bbs_images`
--

DROP TABLE IF EXISTS `bbs_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bbs_images` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `bbs_id` int(11) NOT NULL COMMENT '动态ID',
  `url` varchar(255) DEFAULT '' COMMENT '图片链接',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bbs_images`
--

LOCK TABLES `bbs_images` WRITE;
/*!40000 ALTER TABLE `bbs_images` DISABLE KEYS */;
/*!40000 ALTER TABLE `bbs_images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bbs_like`
--

DROP TABLE IF EXISTS `bbs_like`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bbs_like` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `customer_id` int(11) DEFAULT '0' COMMENT '用户ID',
  `bbs_id` int(11) DEFAULT '0' COMMENT 'Bbs ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bbs_like`
--

LOCK TABLES `bbs_like` WRITE;
/*!40000 ALTER TABLE `bbs_like` DISABLE KEYS */;
/*!40000 ALTER TABLE `bbs_like` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bbs_review`
--

DROP TABLE IF EXISTS `bbs_review`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bbs_review` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `bbs_id` int(11) NOT NULL COMMENT '动态ID',
  `customer_id` int(11) DEFAULT '0' COMMENT '用户ID',
  `reply_id` int(11) DEFAULT '0' COMMENT '被回复用户ID',
  `content` longtext NOT NULL COMMENT '内容',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，0删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bbs_review`
--

LOCK TABLES `bbs_review` WRITE;
/*!40000 ALTER TABLE `bbs_review` DISABLE KEYS */;
/*!40000 ALTER TABLE `bbs_review` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) DEFAULT NULL,
  `pid` int(11) DEFAULT '0' COMMENT '父ID',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` int(11) DEFAULT '1' COMMENT '状态1正常，2删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'默认分类',0,100,1);
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chat`
--

DROP TABLE IF EXISTS `chat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chat` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `send_id` varchar(500) NOT NULL COMMENT '发送人',
  `receive_id` varchar(500) NOT NULL COMMENT '接收人',
  `type` varchar(500) NOT NULL COMMENT '消息类型',
  `content` varchar(50) NOT NULL COMMENT '消息内容',
  `status` int(11) DEFAULT '1' COMMENT '1未读，2已读禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chat`
--

LOCK TABLES `chat` WRITE;
/*!40000 ALTER TABLE `chat` DISABLE KEYS */;
/*!40000 ALTER TABLE `chat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uid` varchar(50) NOT NULL COMMENT '用户ID',
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `nickname` varchar(255) DEFAULT NULL COMMENT '昵称',
  `image` varchar(255) DEFAULT NULL COMMENT '头像',
  `url` varchar(255) DEFAULT NULL COMMENT '博客地址',
  `signature` varchar(255) DEFAULT NULL COMMENT '个性签名',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(50) DEFAULT NULL COMMENT '电话',
  `wishlist` int(11) DEFAULT '0' COMMENT '收藏',
  `review` int(11) DEFAULT '0' COMMENT '评论',
  `like` int(11) DEFAULT '0' COMMENT '点赞',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，0删除',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '修改时间',
  `integral` int(11) DEFAULT '0' COMMENT '积分',
  `fans` int(11) DEFAULT '0' COMMENT '粉丝数量',
  `focus` int(11) DEFAULT '0' COMMENT '关注数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fans`
--

DROP TABLE IF EXISTS `fans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `fans` (
  `customer_id` int(11) NOT NULL COMMENT '用户ID 关注者 ID',
  `fans_id` int(11) NOT NULL COMMENT '用户ID 被关注者ID',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fans`
--

LOCK TABLES `fans` WRITE;
/*!40000 ALTER TABLE `fans` DISABLE KEYS */;
/*!40000 ALTER TABLE `fans` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `file`
--

DROP TABLE IF EXISTS `file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) DEFAULT '' COMMENT '标题',
  `url` varchar(255) DEFAULT '' COMMENT 'URL',
  `sort` int(11) DEFAULT '100' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `file`
--

LOCK TABLES `file` WRITE;
/*!40000 ALTER TABLE `file` DISABLE KEYS */;
/*!40000 ALTER TABLE `file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `link`
--

DROP TABLE IF EXISTS `link`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `link` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) DEFAULT '' COMMENT '名称',
  `url` varchar(255) DEFAULT '' COMMENT 'URL',
  `sort` int(11) DEFAULT '100' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `link`
--

LOCK TABLES `link` WRITE;
/*!40000 ALTER TABLE `link` DISABLE KEYS */;
/*!40000 ALTER TABLE `link` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `ip` varchar(50) NOT NULL,
  `city` varchar(50) NOT NULL,
  `create` datetime NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `page` varchar(255) DEFAULT NULL,
  `uri` varchar(500) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=611161 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
INSERT INTO `log` VALUES (611078,'127.0.0.1','','2021-04-17 21:59:00','colly - https://github.com/gocolly/colly','index','/'),(611079,'127.0.0.1','','2021-04-17 21:59:00','colly - https://github.com/gocolly/colly','article','/'),(611080,'127.0.0.1','','2021-04-17 22:00:00','colly - https://github.com/gocolly/colly','index','/'),(611081,'127.0.0.1','','2021-04-17 22:00:00','colly - https://github.com/gocolly/colly','article','/'),(611082,'172.27.0.1','','2021-04-17 22:06:52','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611083,'172.27.0.1','','2021-04-17 22:06:52','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/'),(611084,'127.0.0.1','','2021-04-17 22:07:00','colly - https://github.com/gocolly/colly','index','/'),(611085,'127.0.0.1','','2021-04-17 22:07:00','colly - https://github.com/gocolly/colly','article','/'),(611086,'127.0.0.1','','2021-04-17 22:08:00','colly - https://github.com/gocolly/colly','index','/'),(611087,'127.0.0.1','','2021-04-17 22:08:00','colly - https://github.com/gocolly/colly','article','/'),(611088,'127.0.0.1','','2021-04-17 22:09:00','colly - https://github.com/gocolly/colly','index','/'),(611089,'127.0.0.1','','2021-04-17 22:09:00','colly - https://github.com/gocolly/colly','article','/'),(611090,'127.0.0.1','','2021-04-17 22:10:00','colly - https://github.com/gocolly/colly','index','/'),(611091,'127.0.0.1','','2021-04-17 22:10:00','colly - https://github.com/gocolly/colly','article','/'),(611092,'172.27.0.1','','2021-04-17 22:10:23','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611093,'172.27.0.1','','2021-04-17 22:10:23','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/'),(611094,'172.27.0.1','','2021-04-17 22:10:25','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611095,'172.27.0.1','','2021-04-17 22:10:25','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/'),(611096,'127.0.0.1','','2021-04-17 22:11:00','colly - https://github.com/gocolly/colly','index','/'),(611097,'127.0.0.1','','2021-04-17 22:11:00','colly - https://github.com/gocolly/colly','article','/'),(611098,'127.0.0.1','','2021-04-17 22:12:00','colly - https://github.com/gocolly/colly','index','/'),(611099,'127.0.0.1','','2021-04-17 22:12:00','colly - https://github.com/gocolly/colly','article','/'),(611100,'127.0.0.1','','2021-04-17 22:13:00','colly - https://github.com/gocolly/colly','index','/'),(611101,'127.0.0.1','','2021-04-17 22:13:00','colly - https://github.com/gocolly/colly','article','/'),(611102,'127.0.0.1','','2021-04-17 22:14:00','colly - https://github.com/gocolly/colly','index','/'),(611103,'127.0.0.1','','2021-04-17 22:14:00','colly - https://github.com/gocolly/colly','article','/'),(611104,'127.0.0.1','','2021-04-17 22:15:00','colly - https://github.com/gocolly/colly','index','/'),(611105,'127.0.0.1','','2021-04-17 22:15:00','colly - https://github.com/gocolly/colly','article','/'),(611106,'127.0.0.1','','2021-04-17 22:16:00','colly - https://github.com/gocolly/colly','index','/'),(611107,'127.0.0.1','','2021-04-17 22:16:00','colly - https://github.com/gocolly/colly','article','/'),(611108,'127.0.0.1','','2021-04-17 22:17:00','colly - https://github.com/gocolly/colly','index','/'),(611109,'127.0.0.1','','2021-04-17 22:17:00','colly - https://github.com/gocolly/colly','article','/'),(611110,'127.0.0.1','','2021-04-17 22:31:00','colly - https://github.com/gocolly/colly','index','/'),(611111,'127.0.0.1','','2021-04-17 22:31:00','colly - https://github.com/gocolly/colly','article','/'),(611112,'127.0.0.1','','2021-04-17 22:31:00','colly - https://github.com/gocolly/colly','index','/'),(611113,'127.0.0.1','','2021-04-17 22:31:00','colly - https://github.com/gocolly/colly','article','/'),(611114,'172.27.0.1','','2021-04-17 22:31:06','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611115,'172.27.0.1','','2021-04-17 22:31:06','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/'),(611116,'172.27.0.1','','2021-04-17 22:31:33','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611117,'172.27.0.1','','2021-04-17 22:31:33','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/'),(611118,'172.27.0.1','','2021-04-17 22:31:35','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/list.html'),(611119,'172.27.0.1','','2021-04-17 22:31:39','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611120,'172.27.0.1','','2021-04-17 22:31:39','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/'),(611121,'127.0.0.1','','2021-04-17 22:32:00','colly - https://github.com/gocolly/colly','index','/'),(611122,'127.0.0.1','','2021-04-17 22:32:00','colly - https://github.com/gocolly/colly','article','/'),(611123,'127.0.0.1','','2021-04-17 22:32:00','colly - https://github.com/gocolly/colly','article','/list.html'),(611124,'127.0.0.1','','2021-04-17 22:32:00','colly - https://github.com/gocolly/colly','index','/'),(611125,'127.0.0.1','','2021-04-17 22:32:00','colly - https://github.com/gocolly/colly','article','/'),(611126,'172.27.0.1','','2021-04-17 22:32:26','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611127,'172.27.0.1','','2021-04-17 22:32:26','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/'),(611128,'172.27.0.1','','2021-04-17 22:32:28','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/list.html'),(611129,'127.0.0.1','','2021-04-17 22:33:00','colly - https://github.com/gocolly/colly','index','/'),(611130,'127.0.0.1','','2021-04-17 22:33:00','colly - https://github.com/gocolly/colly','article','/'),(611131,'127.0.0.1','','2021-04-17 22:33:00','colly - https://github.com/gocolly/colly','index','/'),(611132,'127.0.0.1','','2021-04-17 22:33:00','colly - https://github.com/gocolly/colly','article','/'),(611133,'127.0.0.1','','2021-04-17 22:33:00','colly - https://github.com/gocolly/colly','article','/list.html'),(611134,'127.0.0.1','','2021-04-17 22:34:00','colly - https://github.com/gocolly/colly','index','/'),(611135,'127.0.0.1','','2021-04-17 22:34:00','colly - https://github.com/gocolly/colly','article','/'),(611136,'127.0.0.1','','2021-04-17 22:34:00','colly - https://github.com/gocolly/colly','index','/'),(611137,'127.0.0.1','','2021-04-17 22:34:00','colly - https://github.com/gocolly/colly','article','/'),(611138,'127.0.0.1','','2021-04-17 22:34:00','colly - https://github.com/gocolly/colly','article','/list.html'),(611139,'127.0.0.1','','2021-04-17 22:35:00','colly - https://github.com/gocolly/colly','index','/'),(611140,'127.0.0.1','','2021-04-17 22:35:00','colly - https://github.com/gocolly/colly','article','/'),(611141,'127.0.0.1','','2021-04-17 22:35:00','colly - https://github.com/gocolly/colly','index','/'),(611142,'127.0.0.1','','2021-04-17 22:35:00','colly - https://github.com/gocolly/colly','article','/'),(611143,'127.0.0.1','','2021-04-17 22:35:00','colly - https://github.com/gocolly/colly','article','/list.html'),(611144,'127.0.0.1','','2021-04-17 22:36:00','colly - https://github.com/gocolly/colly','index','/'),(611145,'127.0.0.1','','2021-04-17 22:36:00','colly - https://github.com/gocolly/colly','article','/'),(611146,'127.0.0.1','','2021-04-17 22:36:00','colly - https://github.com/gocolly/colly','index','/'),(611147,'127.0.0.1','','2021-04-17 22:36:00','colly - https://github.com/gocolly/colly','article','/'),(611148,'127.0.0.1','','2021-04-17 22:36:00','colly - https://github.com/gocolly/colly','article','/list.html'),(611149,'127.0.0.1','','2021-04-17 22:37:00','colly - https://github.com/gocolly/colly','index','/'),(611150,'127.0.0.1','','2021-04-17 22:37:00','colly - https://github.com/gocolly/colly','article','/'),(611151,'127.0.0.1','','2021-04-17 22:37:00','colly - https://github.com/gocolly/colly','index','/'),(611152,'127.0.0.1','','2021-04-17 22:37:00','colly - https://github.com/gocolly/colly','article','/'),(611153,'127.0.0.1','','2021-04-17 22:37:00','colly - https://github.com/gocolly/colly','article','/list.html'),(611154,'172.31.0.1','','2021-04-20 17:46:34','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611155,'172.31.0.1','','2021-04-20 17:46:38','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611156,'172.31.0.1','','2021-04-20 17:46:40','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/list.html'),(611157,'172.31.0.1','','2021-04-20 17:46:43','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611158,'172.31.0.1','','2021-04-20 17:48:42','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','detail','/detail/1.html'),(611159,'172.31.0.1','','2021-04-20 17:48:46','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','index','/'),(611160,'172.31.0.1','','2021-04-20 17:48:49','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36','article','/list.html');
/*!40000 ALTER TABLE `log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) DEFAULT '' COMMENT '栏目名称',
  `target` varchar(255) DEFAULT '' COMMENT '链接打开方式',
  `url` varchar(255) DEFAULT '' COMMENT '链接URL',
  `sort` int(11) DEFAULT '100' COMMENT '排序',
  `pid` int(11) DEFAULT '0' COMMENT '顶级栏目',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (1,'首页','','/',100,0);
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名字',
  `review` varchar(500) NOT NULL DEFAULT '' COMMENT '评论',
  `reply` varchar(500) NOT NULL COMMENT '回复',
  `site` varchar(500) NOT NULL COMMENT '网址',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '回复时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `music`
--

DROP TABLE IF EXISTS `music`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `music` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(500) NOT NULL COMMENT '歌名',
  `url` varchar(500) NOT NULL COMMENT '歌曲URL',
  `cover` varchar(500) NOT NULL COMMENT '封面',
  `author` varchar(50) NOT NULL COMMENT '作者',
  `song_id` int(11) NOT NULL COMMENT '歌曲ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `music`
--

LOCK TABLES `music` WRITE;
/*!40000 ALTER TABLE `music` DISABLE KEYS */;
/*!40000 ALTER TABLE `music` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `notice`
--

DROP TABLE IF EXISTS `notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `send_id` int(11) NOT NULL COMMENT '发送人',
  `receive_id` int(11) NOT NULL COMMENT '接收人',
  `type` varchar(500) NOT NULL COMMENT '消息类型',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `content` varchar(50) NOT NULL COMMENT '消息内容',
  `status` int(11) DEFAULT '1' COMMENT '1未读，2已读',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notice`
--

LOCK TABLES `notice` WRITE;
/*!40000 ALTER TABLE `notice` DISABLE KEYS */;
/*!40000 ALTER TABLE `notice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `review`
--

DROP TABLE IF EXISTS `review`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `review` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名字',
  `review` varchar(500) NOT NULL DEFAULT '' COMMENT '评论',
  `reply` varchar(500) NOT NULL COMMENT '回复',
  `site` varchar(500) NOT NULL COMMENT '网址',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '回复时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用',
  `article_id` int(11) DEFAULT '1' COMMENT '文章ID',
  `customer_id` int(11) DEFAULT '1' COMMENT '用户ID',
  `like` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数量',
  `star` int(11) DEFAULT '0' COMMENT '评分',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `review`
--

LOCK TABLES `review` WRITE;
/*!40000 ALTER TABLE `review` DISABLE KEYS */;
/*!40000 ALTER TABLE `review` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `setting`
--

DROP TABLE IF EXISTS `setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `setting` (
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `value` text NOT NULL COMMENT '详情',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting`
--

LOCK TABLES `setting` WRITE;
/*!40000 ALTER TABLE `setting` DISABLE KEYS */;
INSERT INTO `setting` VALUES ('about_html_code','<h2 id=\"h2-u5173u4E8Eu9713u8679u706Fu4E0B\"><a name=\"关于霓虹灯下\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>关于霓虹灯下</h2><p>在这座小小的城市里，每个人都在为自己的未来努力奋斗着。夜深了，霓虹灯下塞满了一个又一个独一无二的故事。这里，也许就是最好的归宿，在这里你能找到属于自己的片刻的安宁。霓虹灯下，我在等你。</p>\n'),('about_markdown_doc','## 关于霓虹灯下\n\n在这座小小的城市里，每个人都在为自己的未来努力奋斗着。夜深了，霓虹灯下塞满了一个又一个独一无二的故事。这里，也许就是最好的归宿，在这里你能找到属于自己的片刻的安宁。霓虹灯下，我在等你。'),('description','Go Blog 一个使用 Beego 开发的博客系统'),('image','/static/uploads/20201220215135313.png'),('keyword','Go Blog ,博客系统'),('limit','20'),('name',''),('notice','欢迎来到使用 Go Blog 。'),('remark_html_code',''),('remark_markdown_doc',''),('tag','欢迎使用Go Blog！'),('template','leechan'),('title','Go Blog');
/*!40000 ALTER TABLE `setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `topic`
--

DROP TABLE IF EXISTS `topic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `topic` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `content` varchar(255) NOT NULL COMMENT '内容',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `join` int(11) DEFAULT '0' COMMENT '评论',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，0删除',
  `category_id` int(11) DEFAULT '0' COMMENT '主题分类',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `topic`
--

LOCK TABLES `topic` WRITE;
/*!40000 ALTER TABLE `topic` DISABLE KEYS */;
/*!40000 ALTER TABLE `topic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `status` int(11) DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'user','5c0b8081c10ad236fa004adfe685867f','491126240@qq.com','2020-02-11 12:22:55',1);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `visit`
--

DROP TABLE IF EXISTS `visit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `visit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL COMMENT '用户ID',
  `visit` int(11) NOT NULL COMMENT '访问量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `visit`
--

LOCK TABLES `visit` WRITE;
/*!40000 ALTER TABLE `visit` DISABLE KEYS */;
/*!40000 ALTER TABLE `visit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wechat_user`
--

DROP TABLE IF EXISTS `wechat_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wechat_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `openid` varchar(32) NOT NULL COMMENT 'open ID',
  `nickname` varchar(255) DEFAULT NULL COMMENT '用户名',
  `sex` int(11) DEFAULT '0' COMMENT '性别',
  `city` varchar(32) DEFAULT NULL COMMENT '城市',
  `country` varchar(32) DEFAULT NULL COMMENT '国家',
  `province` varchar(32) DEFAULT NULL COMMENT '省份',
  `language` varchar(32) DEFAULT NULL COMMENT '语种',
  `headimgurl` varchar(255) DEFAULT NULL COMMENT '头像',
  `subscribe_time` int(11) DEFAULT NULL COMMENT '订阅时间',
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wechat_user`
--

LOCK TABLES `wechat_user` WRITE;
/*!40000 ALTER TABLE `wechat_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `wechat_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-04-20 17:51:41
