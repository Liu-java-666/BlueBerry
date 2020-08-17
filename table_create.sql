-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.24 - MySQL Community Server (GPL)
-- 服务器OS:                        Win32
-- HeidiSQL 版本:                  10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for blueberry
CREATE DATABASE IF NOT EXISTS `blueberry` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;
USE `blueberry`;

-- Dumping structure for table blueberry.action_log
CREATE TABLE IF NOT EXISTS `action_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL COMMENT '执行操作的用户',
  `to_user_id` int(11) NOT NULL COMMENT '被操作的用户',
  `action` enum('focus','blacklist','gift','pay','dynamic_like','dynamic_comment','dynamic_comment_like','video_like','video_comment','video_comment_like') NOT NULL DEFAULT 'focus' COMMENT '执行的操作（关注，黑名单，充值，送礼物，点赞动态，评论动态，点赞动态评论，点赞视频，评论视频，点赞视频评论）',
  `type` int(11) NOT NULL COMMENT '操作类型：1执行，0取消',
  `coins` int(11) NOT NULL DEFAULT '0' COMMENT '货币（>0为收入，<0为支出）',
  `cdate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '日期时间',
  `description` varchar(50) DEFAULT NULL COMMENT '描述',
  `extra` int(11) NOT NULL DEFAULT '0' COMMENT '额外参数，比如动态ID',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=348 DEFAULT CHARSET=utf8 COMMENT='用户操作日志';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.blacklist
CREATE TABLE IF NOT EXISTS `blacklist` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '执行拉黑操作的用户',
  `to_user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '被拉黑的用户',
  `cdate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '拉黑时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_to_user_id` (`to_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='拉黑列表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.captcha
CREATE TABLE IF NOT EXISTS `captcha` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `phone_number` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '电话号码',
  `captcha` varchar(8) COLLATE utf8_unicode_ci NOT NULL COMMENT '验证码',
  `generation_time` int(11) NOT NULL DEFAULT '0' COMMENT '验证码生成时间',
  `expire_time` int(11) NOT NULL DEFAULT '0' COMMENT '过期时间',
  `is_used` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被使用，0未被使用，1已被使用',
  PRIMARY KEY (`id`),
  KEY `idx_phone_number` (`phone_number`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='手机验证码表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.certification
CREATE TABLE IF NOT EXISTS `certification` (
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `post_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `front_img` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '正面图',
  `back_img` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '背面图',
  `is_audit` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态 0 审核 1通过 -1 失败',
  `audit_time` timestamp NULL DEFAULT NULL COMMENT '审核时间',
  PRIMARY KEY (`user_id`),
  KEY `post_time` (`post_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='实名认证表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.denounce
CREATE TABLE IF NOT EXISTS `denounce` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '举报用户Id',
  `to_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '被举报用户Id',
  `post_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '举报类型',
  `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '举报内容',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='举报';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.dynamic
CREATE TABLE IF NOT EXISTS `dynamic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `post_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `description` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述',
  `sentence_id` int(11) NOT NULL DEFAULT '0',
  `topic` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'filetype=voice:句子类型 (1 土味情话 2 霸气宣言 3 自己说了算)',
  `filetype` enum('image','video','voice') CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'image' COMMENT '文件类型',
  `filelist` varchar(512) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '文件列表：1,2,...n 对应于(image/voice/video)主键',
  `is_audit` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态 0 审核 1通过 -1 失败',
  `audit_time` timestamp NULL DEFAULT NULL COMMENT '审核时间',
  `voice_second` int(11) NOT NULL DEFAULT '0' COMMENT '声音时长/秒',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_audit_time` (`audit_time`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='动态视频表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.dynamic_comment
CREATE TABLE IF NOT EXISTS `dynamic_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `dynamic_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '动态Id',
  `postuser_id` int(11) NOT NULL DEFAULT '0' COMMENT '动态发布者Id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '评论内容',
  `cdate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
  PRIMARY KEY (`id`),
  KEY `idx_dynamic_id` (`dynamic_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='动态评论表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.dynamic_comment_like
CREATE TABLE IF NOT EXISTS `dynamic_comment_like` (
  `comment_id` int(11) NOT NULL DEFAULT '0' COMMENT '评论Id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '点赞的人',
  `commentuser_id` int(11) NOT NULL DEFAULT '0' COMMENT '评论者Id',
  `dynamic_id` int(11) NOT NULL DEFAULT '0' COMMENT '动态id',
  PRIMARY KEY (`comment_id`,`user_id`),
  KEY `postuser_id` (`commentuser_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='动态评论点赞表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.dynamic_like
CREATE TABLE IF NOT EXISTS `dynamic_like` (
  `dynamic_id` int(11) NOT NULL DEFAULT '0' COMMENT '动态Id(联合主键) ',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '点赞的人(联合主键) ',
  `postuser_id` int(11) NOT NULL DEFAULT '0' COMMENT '动态发布者Id',
  `like_type` varchar(20) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'dynamic' COMMENT '点赞类型 dynamic(动态) work(作品)',
  `create_time` int(11) NOT NULL COMMENT '时间',
  PRIMARY KEY (`dynamic_id`,`user_id`),
  KEY `postuser_id` (`postuser_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='动态点赞表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.feedback
CREATE TABLE IF NOT EXISTS `feedback` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `post_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `image_id` int(11) NOT NULL DEFAULT '0' COMMENT '图片id，对应image表主键',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '描述',
  `address` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '联系地址',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='意见反馈';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.focuslist
CREATE TABLE IF NOT EXISTS `focuslist` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '执行关注操作的用户',
  `to_user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '被关注的用户',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_to_user_id` (`to_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='关注列表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.focus_notice
CREATE TABLE IF NOT EXISTS `focus_notice` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '执行关注操作的用户',
  `to_user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '被关注的用户',
  `cdate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '关注时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_to_user_id` (`to_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='关注通知表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.gift
CREATE TABLE IF NOT EXISTS `gift` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '礼物ID',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '礼物价格',
  `name` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '礼物名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='礼物表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.gift_log
CREATE TABLE IF NOT EXISTS `gift_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL COMMENT '送出的用户',
  `to_user_id` int(11) NOT NULL COMMENT '接收的用户',
  `gift_id` int(11) NOT NULL COMMENT '礼物ID',
  `coins` int(11) NOT NULL DEFAULT '0' COMMENT '花费货币',
  `cdate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '送礼时间',
  `scene` enum('dynamic','room','message') NOT NULL DEFAULT 'dynamic' COMMENT '送礼场景（动态，房间，私信）',
  `scene_id` int(11) NOT NULL DEFAULT '0' COMMENT '场景标志（动态id，房间id，用户id）',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='礼物日志';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.image
CREATE TABLE IF NOT EXISTS `image` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `post_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `file_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '文件名',
  `file_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'jpg' COMMENT '文件类型（jpg,png,url）',
  `use_type` enum('avatar','dynamic','photo','certification','feedback') COLLATE utf8mb4_bin NOT NULL DEFAULT 'avatar' COMMENT '图片用途（头像，动态，照片，身份认证，意见反馈）',
  `is_audit` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态 0 审核 1通过 -1 失败',
  `audit_time` timestamp NULL DEFAULT NULL COMMENT '审核时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='图片表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.match_log
CREATE TABLE IF NOT EXISTS `match_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '发起用户Id',
  `to_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '被匹配用户Id',
  `cdate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '匹配时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_audit_time` (`cdate`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='匹配日志表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.pay_config
CREATE TABLE IF NOT EXISTS `pay_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `money` int(11) NOT NULL DEFAULT '0' COMMENT '人民币（元）',
  `coins` int(11) NOT NULL DEFAULT '0' COMMENT '货币',
  `appid` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT 'AppStore商品ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='充值配置';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.pay_order
CREATE TABLE IF NOT EXISTS `pay_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `orderid` varchar(128) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '订单号',
  `money` int(11) NOT NULL COMMENT '人民币（元）',
  `coins` int(11) NOT NULL COMMENT '货币',
  `postdate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '订单状态：0未完成，1成功，-1失败',
  `finishdate` timestamp NULL DEFAULT NULL COMMENT '完成时间',
  PRIMARY KEY (`id`),
  KEY `orderid` (`orderid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='充值订单';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.photolist
CREATE TABLE IF NOT EXISTS `photolist` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `post_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
  `photolist` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '相片列表：1,2,...n 对应photo表主键',
  `is_audit` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态 0 审核 1通过 -1 失败',
  `audit_time` timestamp NULL DEFAULT NULL COMMENT '审核时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_audit_time` (`audit_time`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='照片列表表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.room
CREATE TABLE IF NOT EXISTS `room` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '房间ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `room_type` int(11) NOT NULL DEFAULT '0' COMMENT '0-单人房，1-多人房',
  `im_group` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT 'IM群组',
  `room_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '房间名',
  `like_num` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
  `is_open` tinyint(4) NOT NULL DEFAULT '0' COMMENT '房间状态 0 审核 1开启 -1 关闭',
  `open_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开房时间',
  `close_time` timestamp NULL DEFAULT NULL COMMENT '关房时间',
  `room_cover` varchar(100) COLLATE utf8mb4_bin NOT NULL COMMENT '房间封面',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='房间表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.room_seat
CREATE TABLE IF NOT EXISTS `room_seat` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `room_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '房间id',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `room_id` (`room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='房间上座申请列表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.room_user
CREATE TABLE IF NOT EXISTS `room_user` (
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `room_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '房间id',
  `im_group` varchar(200) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT 'IM群组',
  PRIMARY KEY (`user_id`),
  KEY `room_id` (`room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='房间用户表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.sentence
CREATE TABLE IF NOT EXISTS `sentence` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sentence_type` int(11) NOT NULL COMMENT '句子类型：1土味情话 2霸气宣言',
  `sentence_text` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '句子内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- Data exporting was unselected.

-- Dumping structure for table blueberry.user
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `phone_number` varchar(32) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `registration_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用户注册时间',
  `nickname` char(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '昵称',
  `sex` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1男,0女',
  `birthday` date NOT NULL DEFAULT '1900-01-01' COMMENT '生日',
  `user_key` varchar(32) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '用户key',
  `lastlogon_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次登陆时间',
  `lastlogon_ip` varchar(16) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '0.0.0.0' COMMENT '最后一次登陆ip',
  `avatar_id` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '用户头像，关联avatar表主键',
  `video_id` int(11) NOT NULL DEFAULT '0' COMMENT '主页视频',
  `photolist_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户照片列表，关联photolist表主键',
  `certification` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否实名认证',
  `signature` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '签名',
  `relationship_status` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '情感状态',
  `friends_purpose` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '交友意向',
  `hobbies` varchar(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '兴趣爱好，多个用逗号分隔',
  `coins` int(11) NOT NULL DEFAULT '0' COMMENT '货币',
  `coins_used` int(11) NOT NULL DEFAULT '0' COMMENT '已使用货币',
  `nodisturb` int(11) NOT NULL DEFAULT '0' COMMENT '勿扰开关 0 关闭 1 打开',
  PRIMARY KEY (`id`),
  KEY `idx_phone_number` (`phone_number`),
  KEY `registration_time` (`registration_time`),
  KEY `lastlogon_time` (`lastlogon_time`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.user_col_room
CREATE TABLE IF NOT EXISTS `user_col_room` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `room_id` int(11) NOT NULL COMMENT '房间ID',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- Data exporting was unselected.

-- Dumping structure for table blueberry.video
CREATE TABLE IF NOT EXISTS `video` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `post_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `file_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '文件名',
  `file_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'mp4' COMMENT '文件类型',
  `cover_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '封面文件名',
  `cover_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT 'jpg' COMMENT '封面文件类型',
  `rotation` int(11) NOT NULL DEFAULT '0' COMMENT '旋转角度',
  `use_type` enum('dynamic','certification','feedback','detail') COLLATE utf8mb4_bin NOT NULL DEFAULT 'dynamic' COMMENT '视频用途（动态，身份认证，意见反馈）',
  `is_audit` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态 0 审核 1通过 -1 失败',
  `audit_time` timestamp NULL DEFAULT NULL COMMENT '审核时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='视频表';

-- Data exporting was unselected.

-- Dumping structure for table blueberry.voice
CREATE TABLE IF NOT EXISTS `voice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `post_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `file_name` varchar(200) COLLATE utf8mb4_bin NOT NULL COMMENT '文件名',
  `file_type` varchar(50) COLLATE utf8mb4_bin NOT NULL COMMENT '文件类型',
  `file_second` int(11) NOT NULL COMMENT '语音时长 秒',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- Data exporting was unselected.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
