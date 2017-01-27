create table applog (
	id serial primary key,
	pid varchar(255) unique not null,
	firstname text not null,
	lastname text not null,
	email text unique not null,
	passwordhash varchar(255) not null
);