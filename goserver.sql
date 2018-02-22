SET SQL_SAFE_UPDATES = 0;
DELETE FROM  tbl_testplan WHERE lg_count = 1;
SET SQL_SAFE_UPDATES = 1;



CREATE DATABASE  IF NOT EXISTS mjolnir_dev;

USE mjolnir_dev;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id varchar(32) NOT NULL ,
  firstname varchar(32) DEFAULT NULL ,
  lastname varchar(32) DEFAULT NULL ,
  email varchar(48) DEFAULT NULL ,
  password varchar(128) DEFAULT NULL,
  active tinyint(1) DEFAULT NULL ,
  login_status tinyint(1) DEFAULT NULL,
  PRIMARY KEY (email)
) ENGINE=InnoDB


INSERT INTO users VALUES ('USER-UUID-01','Admin','Admin','admin@abc.com','$2a$10$.oh/PgGqIvu0iSj6EjxjT./Rug36HQsGLTHZLy51.Kp9TQ75kl1W2',1,0);


DROP TABLE IF EXISTS tbl_team;
CREATE TABLE tbl_team (
  id varchar(32) NOT NULL ,
  name varchar(32)  NOT NULL ,
  PRIMARY KEY (name)
) ENGINE=InnoDB;

INSERT INTO tbl_team VALUES ('TEAM-UUID-01','HMOF'),('TEAM-UUID-02','TC');


DROP TABLE IF EXISTS tbl_users_team;
CREATE TABLE tbl_users_team (
  users_id varchar(32) NOT NULL ,
  team_id varchar(32) NOT NULL
) ENGINE=InnoDB;

INSERT INTO tbl_users_team VALUES ('1','TEAM-UUID-01'),('2','TEAM-UUID-02');




DROP TABLE IF EXISTS tbl_testplan;

CREATE TABLE tbl_testplan (
  id varchar(32) NOT NULL ,
  name varchar(32) DEFAULT NULL ,
  description varchar(128) DEFAULT 'No description',
  lg_count int DEFAULT 1,
  tags json DEFAULT NULL ,
  jmeter_profile_id varchar(32) DEFAULT NULL ,
  user_count  int DEFAULT 0,
  scenario json DEFAULT NULL ,
  duration varchar(32) DEFAULT 0,
  team_id varchar(32) DEFAULT NULL ,
  last_modified_by varchar(32) DEFAULT NULL ,
  last_modified_on  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
  PRIMARY KEY (id)
) ENGINE=InnoDB;

#{ "testplan" : "sc01-baseline-test.jmx","testdata" :[ "file.csv","file2.csv"],"jars" :[ "file.jar","file2.jar"]}
INSERT INTO tbl_testplan(id,name,last_modified_by,jmeter_profile_id,team_id) VALUES ('TP-UUID-0001','Project1:Baseline test','1','PROFILE-UUID-01','TEAM-UUID-01'),('TP-UUID-0002','Project2:Load test','1','PROFILE-UUID-01','TEAM-UUID-01'),('TP-UUID-0003','Project 3','1','PROFILE-UUID-02','TEAM-UUID-02');


DROP TABLE IF EXISTS tbl_testcase;

CREATE TABLE tbl_testcase (
  id varchar(32) NOT NULL ,
  name varchar(32) DEFAULT NULL ,
  description varchar(128) DEFAULT 'No description',
  tags varchar(128) DEFAULT NULL,
  attachment BLOB DEFAULT NULL,
  team_id varchar(32) DEFAULT NULL ,
  last_modified_by varchar(32) DEFAULT NULL ,
  last_modified_on  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
  PRIMARY KEY (id)
) ENGINE=InnoDB;


INSERT INTO tbl_testcase(id,name,tags,attachment,last_modified_by,last_modified_on,team_id) VALUES ('TC-UUID-0001','TC01','tc01','','1',now(),'TEAM-UUID-01'),('TC-UUID-0002','TC02','tc01','','1',now(),'TEAM-UUID-01'),('TC-UUID-0003','TC03','tc01','','1',now(),'TEAM-UUID-02');



DROP TABLE IF EXISTS tbl_jmeter_profile;
CREATE TABLE tbl_jmeter_profile (
  id varchar(32) NOT NULL ,
  name varchar(32) DEFAULT NULL ,
  config TEXT DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB;

INSERT INTO tbl_jmeter_profile VALUES ('PROFILE-UUID-01','Default','{}');
INSERT INTO tbl_jmeter_profile VALUES ('PROFILE-UUID-02','CRM','{}');




