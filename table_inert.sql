USE `blueberry`;

truncate table action_log;
truncate table blacklist;
truncate table captcha;
truncate table denounce;
truncate table dynamic;
truncate table dynamic_comment;
truncate table dynamic_comment_like;
truncate table dynamic_like;
truncate table gift;
truncate table gift_log;
truncate table image;
truncate table pay_config;
truncate table pay_order;
truncate table photolist;
truncate table room;
truncate table room_seat;
truncate table room_user;
truncate table sentence;
truncate table user;
truncate table user_col_room;
truncate table video;
truncate table voice;

/*句子表*/
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (1,1,'我想你一定很忙#所以只看前三个字就好啦');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (2,1,'你最可爱了。我说的时候来不及思索#我仔细想过之后，还是会这么说');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (3,1,'我喜欢夏天的雨，雨后的光#和任何时候的你');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (4,1,'好好照顾自己#不行就让我来照顾你?');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (5,1,'你今天干嘛打扮成这个样子#好看不说#偏偏是我喜欢的样子?');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (6,1,'脑子真神奇，忙的要死#也要留个小缝儿想你');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (7,1,'最近手头有点紧#想借你的手牵一牵');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (8,1,'你喜欢苹果汁、葡萄汁、西瓜汁#还是我这个精神小伙汁');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (9,1,'给你科普一下鸭子的种类#达克鸭、小黄鸭、扁嘴鸭#我想你了鸭');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (10,1,'喜欢你是一件麻烦的事#那你别喜欢我#可是我偏偏喜欢找麻烦');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (11,1,'你最近真讨人厌，讨人喜欢百看不厌');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (12,1,'和你捉迷藏我一定会输的#因为喜欢一个人是藏也藏不住的');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (13,1,'感觉你今天很奇怪#怪让人心动的');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (14,1,'现在你是小可爱#老了以后就是老可爱#死了之后就是可爱死啦');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (15,1,'你可以笑一下吗？#我的咖啡没加糖');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (16,1,'你应该在淘宝上架因为你也是我的宝贝');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (17,1,'想到你就脸红见你不用腮红');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (18,1,'一生平淡无奇，偏偏遇见了你#我的心便波澜四起');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (19,1,'我还是喜欢你，像小时候吃辣条，不看日期');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (20,1,'我们这里昼夜温差大 我超甜的');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (21,1,'你知道我喜欢吃什么吗?#我喜欢痴痴地望着你');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (22,2,'虽然我不会蒸馒头，但是我会生气#而且是大口大口的');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (23,2,'靠爸妈，你最多是公主，靠男人#你最多是王妃，靠自己，你就是女王');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (24,2,'我从来不喜欢和别人争东西#你喜欢就拿去，前提是你能拿得走');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (25,2,'你不是人民币，做不到人人都喜欢#你活着就要让讨厌你的人越来越不爽。');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (26,2,'等不到的晚安就别等了#挤不进的世界就别挤了');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (27,2,'我是一个很有原则的人#我的原则只有三个字，看心情');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (28,2,'觉得自己做得到和做不到#其实只在一念之间');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (29,2,'活得开心最重要，不管有多少挫折#都要努力冲过去');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (30,2,'只有不停地跑才能追上我的梦');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (31,2,'尊严是自己经营的，别人给不了');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (32,2,'你不能改变过去但你能够改变未来');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (33,2,'我想要的并不多#可我想证明我能够得到的比任何人都多');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (34,2,'有人能让你痛苦，说明你的修行还不够');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (35,2,'没有什么过不去，我不坚信幸福，我坚信你');
insert  into `sentence`(`id`,`sentence_type`,`sentence_text`) values (36,2,'别总是羡慕别人光芒万丈#却忘了自己也会发光');
/*用户表*/
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (1,'19800000001','2020-07-23 10:28:34','小客服',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',2,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (2,'19800000002','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (3,'19800000003','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (4,'19800000004','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (5,'19800000005','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (6,'19800000006','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (7,'19800000007','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (8,'19800000008','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (9,'19800000009','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (10,'19800000010','2020-07-23 10:28:34','预留',1,'1900-01-01','','2020-07-23 10:28:34','0.0.0.0',1,0,0,0,'','','','',0,0,0);

insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (11,'19900000001','2020-07-23 10:28:34','Liiiii',0,'1995-03-24','','2020-07-23 10:28:34','0.0.0.0',3,9,0,0,'想把我唱给你听','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (12,'19900000002','2020-07-23 10:28:35','阿书啊',0,'1998-06-23','','2020-07-23 10:28:35','0.0.0.0',4,10,0,0,'Love is gone','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (13,'19900000003','2020-07-23 10:28:35','全国最闪耀的妞',0,'1997-02-09','','2020-07-23 10:28:35','0.0.0.0',5,11,0,0,'比我优秀的人仍在努力','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (14,'19900000004','2020-07-23 10:28:35','心动乐乐',0,'1992-04-05','','2020-07-23 10:28:35','0.0.0.0',6,12,0,0,'是旺旺不是汪汪','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (15,'19900000005','2020-07-23 10:28:35','勋鹿',1,'1998-07-15','','2020-07-23 10:28:35','0.0.0.0',7,13,5,0,'温温柔柔的男孩纸你们喜欢吗？','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (16,'19900000006','2020-07-23 10:28:35','bear baby',1,'1994-09-18','','2020-07-23 10:28:35','0.0.0.0',8,14,0,0,'实在没必要 跟每个人说','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (17,'19900000007','2020-07-23 10:28:35','林七七',1,'1991-03-03','','2020-07-23 10:28:35','0.0.0.0',9,15,0,0,'有话对你讲.','','','',0,0,0);
insert  into `user`(`id`,`phone_number`,`registration_time`,`nickname`,`sex`,`birthday`,`user_key`,`lastlogon_time`,`lastlogon_ip`,`avatar_id`,`video_id`,`photolist_id`,`certification`,`signature`,`relationship_status`,`friends_purpose`,`hobbies`,`coins`,`coins_used`,`nodisturb`) values (18,'19900000008','2020-07-23 10:28:35','大萝卜卜',1,'1990-06-12','','2020-07-23 10:28:35','0.0.0.0',10,16,0,0,'忍不住对你心动','','','',0,0,0);

/*支付*/
INSERT INTO `pay_config`(money, coins, appid) VALUES(6, 60, '1000000000000001');
INSERT INTO `pay_config`(money, coins, appid) VALUES(68, 680, '1000000000000002');
INSERT INTO `pay_config`(money, coins, appid) VALUES(648, 6480, '1000000000000003');

/*礼物*/
insert  into `gift`(`id`,`price`,`name`) values (1,66,'黄金UFO');
insert  into `gift`(`id`,`price`,`name`) values (2,188,'仙境');
insert  into `gift`(`id`,`price`,`name`) values (3,88,'女神奖杯');
insert  into `gift`(`id`,`price`,`name`) values (4,99,'仙子降临');
insert  into `gift`(`id`,`price`,`name`) values (5,199,'南瓜马车');
insert  into `gift`(`id`,`price`,`name`) values (6,108,'love u');
insert  into `gift`(`id`,`price`,`name`) values (7,88,'仙女权杖');
insert  into `gift`(`id`,`price`,`name`) values (8,188,'花樱');
insert  into `gift`(`id`,`price`,`name`) values (9,199,'海洋之心');
insert  into `gift`(`id`,`price`,`name`) values (10,33,'环环相扣');
insert  into `gift`(`id`,`price`,`name`) values (11,11,'唇唇欲动');
insert  into `gift`(`id`,`price`,`name`) values (12,66,'我是麦霸');
insert  into `gift`(`id`,`price`,`name`) values (13,52,'爱心之翼');
insert  into `gift`(`id`,`price`,`name`) values (14,51,'521');
insert  into `gift`(`id`,`price`,`name`) values (15,99,'一箭钟情');
insert  into `gift`(`id`,`price`,`name`) values (16,0,'是心动吖');
insert  into `gift`(`id`,`price`,`name`) values (17,88,'三生有信');
insert  into `gift`(`id`,`price`,`name`) values (18,66,'浪漫气球');

/*房间*/
insert  into `room`(`id`,`user_id`,`room_type`,`im_group`,`room_name`,`like_num`,`is_open`,`open_time`,`close_time`,`room_cover`) values (1,12,0,'@TGS#14TJZBTGI','',0,1,'2020-07-14 14:30:00',NULL,'');
insert  into `room`(`id`,`user_id`,`room_type`,`im_group`,`room_name`,`like_num`,`is_open`,`open_time`,`close_time`,`room_cover`) values (2,13,1,'@TGS#1PYJZBTGK','',0,1,'2020-07-14 14:20:00',NULL,'');
insert  into `room`(`id`,`user_id`,`room_type`,`im_group`,`room_name`,`like_num`,`is_open`,`open_time`,`close_time`,`room_cover`) values (3,16,0,'@TGS#1K5JZBTGA','',0,1,'2020-07-14 14:10:00',NULL,'');
insert  into `room`(`id`,`user_id`,`room_type`,`im_group`,`room_name`,`like_num`,`is_open`,`open_time`,`close_time`,`room_cover`) values (4,17,1,'@TGS#15DKZBTGZ','',0,1,'2020-07-14 14:00:00',NULL,'');



/*动态*/
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (1,11,'2020-07-31 10:43:33','最近换了个风格',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (2,11,'2020-07-31 10:45:52','每人心中都有一个公主梦',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (3,11,'2020-07-31 10:52:48','得不到的就更加爱',0,'','video','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (4,11,'2020-07-31 10:55:23','我喜欢的你',3,'','voice','',1,'2020-04-02 18:20:00',5);

insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (5,12,'2020-07-31 10:43:33','九年前的自己',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (6,12,'2020-07-31 10:43:33','下课啦 哈哈哈',0,'','video','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (7,12,'2020-07-31 10:43:33','要不停奔跑',0,'','voice','',1,'2020-04-02 18:20:00',3);

insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (8,13,'2020-07-31 10:43:33','没想到3年前的衣服孕晚期还能穿',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (9,13,'2020-07-31 10:43:33','包粽子，妈妈牌粽子',0,'','video','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (10,13,'2020-07-31 10:43:33','想到你...',17,'','voice','',1,'2020-04-02 18:20:00',4);

insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (11,14,'2020-07-31 10:43:33','过六一啦哈哈哈',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (12,14,'2020-07-31 10:43:33','为了吃 我们有无限可能',0,'','video','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (13,14,'2020-07-31 10:43:33','哈哈哈加油',32,'','voice','',1,'2020-04-02 18:20:00',4);

insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (14,15,'2020-07-31 10:43:33','零散的照片',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (15,15,'2020-07-31 10:45:52','假期快乐',0,'','video','',1,'2020-04-02 18:20:00',0);

insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (18,16,'2020-07-31 10:43:33','你说要有光，于是便有了光。',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (19,16,'2020-07-31 10:43:33','最近自拍有点多，多担待哈哈哈',0,'','video','',1,'2020-04-02 18:20:00',0);

insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (21,17,'2020-07-31 10:43:33','落班-​​​​',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (22,17,'2020-07-31 10:43:33','希望外面流浪的动物们 都能健康成长！',0,'','video','',1,'2020-04-02 18:20:00',0);

insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (24,18,'2020-07-31 10:43:33','日常分享之男模…哈哈哈',0,'','image','',1,'2020-04-02 18:20:00',0);
insert  into `dynamic`(`id`,`user_id`,`post_time`,`description`,`sentence_id`,`topic`,`filetype`,`filelist`,`is_audit`,`audit_time`,`voice_second`) values (25,18,'2020-07-31 10:43:33','还是杭州好啊！出去三天伤到了！',0,'','video','',1,'2020-04-02 18:20:00',0);



/*图片*/
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (1,0,'2020-07-31 11:19:45','avatar_0.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (2,1,'2020-07-31 11:19:45','avatar_1.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (3,11,'2020-07-31 11:19:45','avatar_11.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (4,12,'2020-07-31 11:19:45','avatar_12.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (5,13,'2020-07-31 11:19:45','avatar_13.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (6,14,'2020-07-31 11:19:45','avatar_14.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (7,15,'2020-07-31 11:19:45','avatar_15.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (8,16,'2020-07-31 11:19:45','avatar_16.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (9,17,'2020-07-31 11:19:45','avatar_17.jpg','jpg','avatar',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (10,18,'2020-07-31 11:19:45','avatar_18.jpg','jpg','avatar',1,NULL);

insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (11,11,'2020-07-31 11:19:45','dynamic_11_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (12,11,'2020-07-31 11:19:45','dynamic_11_2.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (13,12,'2020-07-31 11:19:45','dynamic_12_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (14,13,'2020-07-31 11:19:45','dynamic_13_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (15,14,'2020-07-31 11:19:45','dynamic_14_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (16,15,'2020-07-31 11:19:45','dynamic_15_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (17,16,'2020-07-31 11:19:45','dynamic_16_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (18,17,'2020-07-31 11:19:45','dynamic_17_1.jpg','jpg','dynamic',1,NULL);
insert  into `image`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`use_type`,`is_audit`,`audit_time`) values (19,18,'2020-07-31 11:19:45','dynamic_18_1.jpg','jpg','dynamic',1,NULL);


/*视频*/
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (1,11,'2020-07-31 11:21:21','dynamic_11_1.mp4','mp4','dynamic_11_1.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (2,12,'2020-07-31 11:21:21','dynamic_12_1.mp4','mp4','dynamic_12_1.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (3,13,'2020-07-31 11:21:21','dynamic_13_1.mp4','mp4','dynamic_13_1.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (4,14,'2020-07-31 11:21:21','dynamic_14_1.mp4','mp4','dynamic_14_1.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (5,15,'2020-07-31 11:21:21','dynamic_15_1.mp4','mp4','dynamic_15_1.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (6,16,'2020-07-31 11:21:21','dynamic_16_1.mp4','mp4','dynamic_16_1.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (7,17,'2020-07-31 11:21:21','dynamic_17_1.mp4','mp4','dynamic_17_1.jpg','jpg',0,'dynamic',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (8,18,'2020-07-31 11:21:21','dynamic_18_1.mp4','mp4','dynamic_18_1.jpg','jpg',0,'dynamic',1,NULL);

insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (9,11,'2020-07-31 11:21:21','detail_11.mp4','mp4','detail_11.jpg','jpg',0,'detail',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (10,11,'2020-07-31 11:21:21','detail_12.mp4','mp4','detail_12.jpg','jpg',0,'detail',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (11,11,'2020-07-31 11:21:21','detail_13.mp4','mp4','detail_13.jpg','jpg',0,'detail',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (12,11,'2020-07-31 11:21:21','detail_14.mp4','mp4','detail_14.jpg','jpg',0,'detail',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (13,11,'2020-07-31 11:21:21','detail_15.mp4','mp4','detail_15.jpg','jpg',0,'detail',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (14,11,'2020-07-31 11:21:21','detail_16.mp4','mp4','detail_16.jpg','jpg',0,'detail',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (15,11,'2020-07-31 11:21:21','detail_17.mp4','mp4','detail_17.jpg','jpg',0,'detail',1,NULL);
insert  into `video`(`id`,`user_id`,`post_time`,`file_name`,`file_type`,`cover_name`,`cover_type`,`rotation`,`use_type`,`is_audit`,`audit_time`) values (16,11,'2020-07-31 11:21:21','detail_18.mp4','mp4','detail_18.jpg','jpg',0,'detail',1,NULL);

/*声音*/
insert into `voice` (`id`, `user_id`, `post_time`, `file_name`, `file_type`, `file_second`) values('1','11','2020-07-30 14:20:10','dynamic_11.mp3','mp3','5');
insert into `voice` (`id`, `user_id`, `post_time`, `file_name`, `file_type`, `file_second`) values('2','12','2020-07-30 14:20:10','dynamic_12.mp3','mp3','3');
insert into `voice` (`id`, `user_id`, `post_time`, `file_name`, `file_type`, `file_second`) values('3','13','2020-07-30 14:20:10','dynamic_13.mp3','mp3','4');
insert into `voice` (`id`, `user_id`, `post_time`, `file_name`, `file_type`, `file_second`) values('4','14','2020-07-30 14:20:10','dynamic_14.mp3','mp3','4');



