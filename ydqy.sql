/*
	客户信息表
*/
create table customer(
  customerId int(10)  PRIMARY KEY AUTO_INCREMENT,
  customerName varchar(16),
  cPhone varchar(16),
  cAdr varchar(128),
  cProductId int(10),
  buyDate Date,
  cSum float,
  deposit float,
  tail float,
  remark varchar(1000),
  createDate date,
  updateDate date
);
