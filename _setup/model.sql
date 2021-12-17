create table users(
	user_id serial primary key not null,
	fullname character varying(128) not null,
	phone character varying(13) not null,
	password character varying(32) not null,
	sms_confirm character varying(16) not null,
	is_verified bool default false
);
