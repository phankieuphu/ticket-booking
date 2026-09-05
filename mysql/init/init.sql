create database if not exists booking_system;
use booking_system;

create table if not exists users (
   id int auto_increment primary key,
   username varchar(255) not null unique, 
  -- username is indexed for faster lookups
   password_hash varchar(255) not null,
   created_at timestamp default current_timestamp
);
-- create index idx_username on users(username);


create table if not exists user_profile (
   user_id int not null primary key,
   first_name varchar(255) not null,
   last_name varchar(255) not null,
   address varchar(255) null,
   phone varchar(20) not null,
   email varchar(320) not null unique,
   created_at timestamp default current_timestamp,
   foreign key (user_id) references users(id) on delete cascade
);
-- create index idx_email on user_profile(email);


create table if not exists roles (
   id int auto_increment primary key,
   role_name varchar(50) not null unique,
   created_at timestamp default current_timestamp
);
-- create index idx_role_name on roles(role_name);


create table if not exists user_roles (
   user_id int not null,
   role_id int not null,
   created_at timestamp default current_timestamp,
   primary key (user_id, role_id),
   foreign key (user_id) references users(id) on delete cascade,
   foreign key (role_id) references roles(id) on delete cascade
);


create table if not exists permissions (
   id int auto_increment primary key,
   permission_name varchar(50) not null unique,
   created_at timestamp default current_timestamp
   resource varchar(255) not null,
   action varchar(50) not null, // enum
   description varchar(255) null
);
-- create index idx_permission_name on permissions(permission_name);


create table if not exists role_permissions (
   role_id int not null,
   permission_id int not null,
   created_at timestamp default current_timestamp,
   primary key (role_id, permission_id),
   foreign key (role_id) references roles(id) on delete cascade,
   foreign key (permission_id) references permissions(id) on delete cascade
);
-- create index idx_role_id on role_permissions(role_id);
