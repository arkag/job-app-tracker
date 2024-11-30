CREATE DATABASE IF NOT EXISTS careers;
USE careers;
DROP TABLE IF EXISTS jobapps;
CREATE TABLE jobapps (
  id       INT AUTO_INCREMENT NOT NULL,
  status   VARCHAR(128) NOT NULL,
  title    VARCHAR(255) NOT NULL,
  company  VARCHAR(255) NOT NULL,
  url      VARCHAR(255) NOT NULL,
  source   VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO jobapps
  (status, title, company, url, source)
VALUES
  ('Applied', 'DevOps Engineer', 'Google', 'https://google.com/careersorsomethingidk', 'LinkedIn'),
  ('Rejected', 'Site Reliability Engineer I', 'Netflix', 'https://netflix.com/careersorsomethingidk', 'GlassDoor'),
  ('Interviewing', 'Cloud Engineer III', 'Microsoft', 'https://microsoft.com/careersorsomethingidk', 'Otta'),
  ('Offered', 'Sr. DevOps Engineer', 'Amazon', 'https://amazon.com/careersorsomethingidk', 'Indeed');