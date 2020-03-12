create table yawarakasa(
  user_id varchar(255) NOT NULL,
  year int(4) NOT NULL,
  month int(2) NOT NULL,
  day int(2) NOT NULL,
  yawarakasa int(255),
  PRIMARY KEY (UserID, year, month, day)
);
