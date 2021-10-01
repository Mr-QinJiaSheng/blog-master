CREATE TABLE article (
  id int NOT NULL PRIMARY KEY,
  user_id int DEFAULT '0' ,
  title text NOT NULL DEFAULT '',
  category_id int NOT NULL,
  tag text NOT NULL DEFAULT '',
  remark text NOT NULL DEFAULT '',
  "desc" text NOT NULL,
  pv int DEFAULT '0',
  created timestamp DEFAULT NULL,
  updated timestamp DEFAULT NULL,
  status int DEFAULT '1',
  review int DEFAULT '0' 
);

COMMENT ON COLUMN article.id IS 'ID';
COMMENT ON COLUMN article.user_id  IS '用户ID';
COMMENT ON COLUMN article.title IS '标题';
COMMENT ON COLUMN article.category_id IS '分类ID';
COMMENT ON COLUMN article.tag IS 'Tag';
COMMENT ON COLUMN article.remark IS '摘要';
COMMENT ON COLUMN article."desc" IS '详情';
COMMENT ON COLUMN article.pv  IS 'px';
COMMENT ON COLUMN article.created IS '创建时间';
COMMENT ON COLUMN article.updated IS '更新时间';
COMMENT ON COLUMN article.status IS '1可用，2禁用，3删除';
COMMENT ON COLUMN article.review IS '评论';

ALTER TABLE article ADD recommend int NOT NULL DEFAULT '0';
COMMENT ON COLUMN article.recommend IS '是否顶置，0否；1是，默认否';
ALTER TABLE article ADD COLUMN "like"  int NOT NULL DEFAULT 0 ;
COMMENT ON COLUMN article."like" IS '点赞数量';
ALTER TABLE article ADD COLUMN "html" text NOT NULL;
COMMENT ON COLUMN article."html" IS 'html内容';
-- --------------------------------------------------------

--
-- 表的结构 category
--

CREATE TABLE category (
  id int NOT NULL PRIMARY KEY,
  name text DEFAULT NULL,
  pid int DEFAULT '0',
  sort int DEFAULT '0',
  status int DEFAULT '1'
);
COMMENT ON COLUMN category.id IS '主键';
COMMENT ON COLUMN category.pid IS '父ID';
COMMENT ON COLUMN category.sort IS '排序';
COMMENT ON COLUMN category.status IS '状态1正常，2删除';
-- --------------------------------------------------------

--
-- 表的结构 log
--

CREATE TABLE log (
  id int NOT NULL PRIMARY KEY,
  ip varchar(50) NOT NULL,
  city varchar(50) NOT NULL,
  "create" timestamp NOT NULL,
  user_agent text NOT NULL,
  page text DEFAULT NULL,
  uri text NOT NULL
);

COMMENT ON COLUMN log.id IS 'ID';
-- --------------------------------------------------------

--
-- 表的结构 message
--

CREATE TABLE message (
  id int NOT NULL PRIMARY KEY,
  name text NOT NULL DEFAULT '',
  review text NOT NULL DEFAULT '',
  reply text NOT NULL,
  site text NOT NULL,
  created timestamp DEFAULT NULL,
  updated timestamp DEFAULT NULL,
  status int DEFAULT '1'
);
COMMENT ON COLUMN message.id IS 'ID';
COMMENT ON COLUMN message.name IS '名字';
COMMENT ON COLUMN message.review IS '评论';
COMMENT ON COLUMN message.reply IS '回复';
COMMENT ON COLUMN message.site IS '网址';
COMMENT ON COLUMN message.created IS '创建时间';
COMMENT ON COLUMN message.updated IS '回复时间';
COMMENT ON COLUMN message.status IS '1可用，2禁用';

-- --------------------------------------------------------

--
-- 表的结构 review
--

CREATE TABLE review (
  id int NOT NULL PRIMARY KEY,
  name text NOT NULL DEFAULT '',
  review text NOT NULL DEFAULT '',
  reply text NOT NULL,
  site text NOT NULL,
  created timestamp DEFAULT NULL,
  updated timestamp DEFAULT NULL,
  status int DEFAULT '1',
  article_id int DEFAULT '1' 
);

COMMENT ON COLUMN review.id IS 'ID';
COMMENT ON COLUMN review.name IS '名字';
COMMENT ON COLUMN review.review IS '评论';
COMMENT ON COLUMN review.reply IS '回复';
COMMENT ON COLUMN review.site IS '网址';
COMMENT ON COLUMN review.created IS '创建时间';
COMMENT ON COLUMN review.updated IS '回复时间';
COMMENT ON COLUMN review.status IS '1可用，2禁用';

-- --------------------------------------------------------

--
-- 表的结构 user
--

CREATE TABLE "user" (
  id int NOT NULL PRIMARY KEY,
  name text DEFAULT NULL,
  password text DEFAULT NULL,
  email text DEFAULT NULL,
  created timestamp DEFAULT NULL,
  status int DEFAULT '1'
);
COMMENT ON COLUMN "user".id IS 'ID';
COMMENT ON COLUMN "user".name IS '用户名';
COMMENT ON COLUMN "user".status IS '状态';

CREATE TABLE setting (
  name text NOT NULL DEFAULT '',
  value text NOT NULL,
  PRIMARY KEY (name)
);

COMMENT ON COLUMN setting.name IS '标题';
COMMENT ON COLUMN setting.value IS '详情';

CREATE TABLE wechat_user (
  id int NOT NULL,
  openid varchar(32) NOT NULL,
  nickname text DEFAULT NULL,
  sex int DEFAULT '0',
  city varchar(32) DEFAULT NULL,
  country varchar(32) DEFAULT NULL,
  province varchar(32) DEFAULT NULL,
  language varchar(32) DEFAULT NULL,
  headimgurl text DEFAULT NULL,
  subscribe_time int DEFAULT NULL,
  created timestamp DEFAULT NULL,
  PRIMARY KEY (id)
);

COMMENT ON COLUMN wechat_user.id IS 'ID';
COMMENT ON COLUMN wechat_user.openid IS 'open ID';
COMMENT ON COLUMN wechat_user.nickname IS '用户名';
COMMENT ON COLUMN wechat_user.sex IS '性别';
COMMENT ON COLUMN wechat_user.city IS '城市';
COMMENT ON COLUMN wechat_user.country IS '国家';
COMMENT ON COLUMN wechat_user.province IS '省份';
COMMENT ON COLUMN wechat_user.language IS '语种';
COMMENT ON COLUMN wechat_user.headimgurl IS '头像';
COMMENT ON COLUMN wechat_user.subscribe_time IS '订阅时间';


CREATE TABLE menu (
  id int NOT NULL,
  title text  DEFAULT '',
  target text   DEFAULT '',
  url text  DEFAULT '',
  sort int DEFAULT '100',
  pid int DEFAULT '0',
  PRIMARY KEY (id)
);

COMMENT ON COLUMN menu.id IS '主键';
COMMENT ON COLUMN menu.title IS '栏目名称';
COMMENT ON COLUMN menu.target IS '链接打开方式';
COMMENT ON COLUMN menu.url IS '链接URL';
COMMENT ON COLUMN menu.sort IS '排序';
COMMENT ON COLUMN menu.pid IS '顶级栏目';

CREATE TABLE link (
  id int NOT NULL,
  title text DEFAULT '',
  url text  DEFAULT '',
  sort int DEFAULT '100',
  PRIMARY KEY (id)
);

COMMENT ON COLUMN link.id IS '主键';
COMMENT ON COLUMN link.title IS '名称';
COMMENT ON COLUMN link.url IS 'URL';
COMMENT ON COLUMN link.sort IS '排序';
--
-- 转储表的索引
--

--
-- 表的索引 article
--
 
--
-- 表的索引 log
 --
-- 表的索引 message
--
 
--
-- 表的索引 review
--
 
--
-- 表的索引 user
--
 
--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT article
--

INSERT INTO "user" VALUES (1, 'user', '5c0b8081c10ad236fa004adfe685867f', '491126240@qq.com', '2020-2-11 12:22:55', 1);

INSERT INTO category (id, name, pid, sort, status) VALUES
(1, '随笔', 0, 100, 1);

INSERT INTO article (id, user_id, title, category_id, tag, remark, "desc", pv, created, updated, status, review, recommend, "like") VALUES
(1, 1, '这是我的第一篇博客', 1, '博客，Go Blog', '这是我的第一篇博客', '## 这是我的第一篇博客', 1, '2020-02-12 23:07:52', '2020-02-13 05:03:22', 1, 0, 0, 0);

INSERT INTO menu (id, title, target, url, sort, pid) VALUES
(1, '首页', 'none', '/', 100, 0);
INSERT INTO menu (id, title, target, url, sort, pid) VALUES
(2, '留言', 'none', '/message.html', 100, 0);

INSERT INTO link (id, title, url, sort) VALUES
(1, 'Go Blog', 'http://go-blog.cn', 100);
INSERT INTO link (id, title, url, sort) VALUES
(2, 'LeeChan''Blog', 'http://leechan.online', 100);

INSERT INTO setting (name, value) VALUES
('image', '/static/uploads/2020021121190681.png'),
('name', 'Go Blog'),
('notice', '欢迎来到使用 Go Blog 。'),
('remark', '一个使用 Beego 开发的博客系统'),
('tag', '一个使用 Beego 开发的博客系统'),
('title', 'Go Blog'),
('keyword', 'Go,Go Blog,Go Blog社区,社区,博客系统'),
('description', 'Go Blog 一个使用 Beego 开发的博客系统');
