insert into users(username, password_hash) values ('phankieuphu', '1234');

insert into user_profile(user_id, email, first_name, last_name, phone, address) values ( 1, 'phankieuphu@gmail.com','Phan','Patrick', '123-456-7890', '123 Main St');

insert into roles(role_name) values ('admin');
insert into roles(role_name) values ('user');

insert into permissions(permission_name) values ('read');
insert into permissions(permission_name) values ('write');  

insert into user_roles(user_id, role_id) values (1, 1);
insert into role_permissions(role_id, permission_id) values (1, 1);
insert into role_permissions(role_id, permission_id) values (1, 2);
  