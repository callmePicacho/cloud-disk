create table if not exists repository_pool
(
	id int unsigned auto_increment
		primary key,
	identity varchar(36) null comment '记录的唯一标识',
	hash varchar(32) null comment '文件的唯一标识',
	name varchar(255) null comment '文件名称',
	ext varchar(30) null comment '文件扩展名',
	size int null comment '文件大小',
	path varchar(255) null comment '文件路径',
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null
)
comment '公共文件存储池' charset=utf8;

create table if not exists share_basic
(
	id int unsigned auto_increment
		primary key,
	identity varchar(36) null,
	user_identity varchar(36) null comment '对应用户的唯一标识',
	repository_identity varchar(36) null comment '公共池中文件的唯一标识',
	expired_time int null comment '失效时间，单位秒,【0-永不失效】',
	click_num int default 0 null comment '点击次数',
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null
)
comment '文件共享' charset=utf8;

create table if not exists user_basic
(
	id int unsigned auto_increment
		primary key,
	identity varchar(36) null comment '用户唯一标识',
	name varchar(60) null,
	password varchar(32) null,
	email varchar(100) null,
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null
)
comment '用户信息' charset=utf8;

create table if not exists user_repository
(
	id int unsigned auto_increment
		primary key,
	identity varchar(36) null,
	parent_id int null comment '父级文件层级, 0-【文件夹】',
	user_identity varchar(36) null comment '对应用户的唯一标识',
	repository_identity varchar(36) null comment '公共池中文件的唯一标识',
	ext varchar(255) null comment '文件或文件夹类型',
	name varchar(255) null comment '用户定义的文件名',
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null
)
comment '用户存储池' charset=utf8;

