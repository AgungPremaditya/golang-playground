CREATE TABLE movie(
  id serial primary key,
  title varchar(255) not null,
  rating integer,
  details text
);