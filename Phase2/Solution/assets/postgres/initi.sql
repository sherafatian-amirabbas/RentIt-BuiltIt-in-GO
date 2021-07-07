CREATE TABLE plants(
   Id           INT PRIMARY KEY NOT NULL,
   Name         TEXT    NOT NULL,
   Description  TEXT    NOT NULL,
   PricePerDay  NUMERIC(18,2)  NOT NULL
);

INSERT INTO plants VALUES(1, 'eq1', 'desc1', 10.50);
INSERT INTO plants VALUES(2, 'eq2', 'desc2', 16.65);


CREATE TABLE plantorders(
   Id           INT PRIMARY KEY NOT NULL,
   PlantId      INT    NOT NULL,
   StartDate    DATE   NOT NULL,
   EndDate      DATE   NOT NULL
);

INSERT INTO plantorders VALUES(1, 1, '2020-03-10', '2020-03-15');
INSERT INTO plantorders VALUES(2, 1, '2020-03-15', '2020-03-20');
INSERT INTO plantorders VALUES(3, 1, '2020-03-21', '2020-03-25');