CREATE TABLE customers(
   ID           INT PRIMARY KEY NOT NULL,
   Name         TEXT    NOT NULL
);

INSERT INTO customers VALUES(1, 'Customer 1');
INSERT INTO customers VALUES(2, 'Customer 2');



CREATE TABLE plants(
   ID           INT PRIMARY KEY NOT NULL,
   Name         TEXT    NOT NULL,
   Description  TEXT    NOT NULL,
   PricePerDay  NUMERIC(18,2)  NOT NULL
);

INSERT INTO plants VALUES(1, 'eq1', 'desc1', 10.60);
INSERT INTO plants VALUES(2, 'eq2', 'desc2', 16.65);



CREATE TABLE plantOrders(
   ID           INT PRIMARY KEY NOT NULL,
   PlantId      INT    NOT NULL,
   CustomerID   INT    NOT NULL,
   StartDate    DATE   NOT NULL,
   EndDate      DATE   NOT NULL,
   Status       INT    NOT NULL,
   InvoiceID    INT
);

INSERT INTO plantOrders VALUES(1, 1, 1, '2020-03-10', '2020-03-16', 0, null);
INSERT INTO plantOrders VALUES(2, 1, 1, '2020-03-15', '2020-03-20', 1, null);
INSERT INTO plantOrders VALUES(3, 1, 2, '2020-03-21', '2020-03-25', 3, 1);
INSERT INTO plantOrders VALUES(4, 2, 2, '2020-03-26', '2020-03-28', 2, null);



CREATE TABLE cancellationRequests(
   ID                INT PRIMARY KEY NOT NULL,
   PlantOrderId      INT    NOT NULL,
   SubmissionDate    DATE   NOT NULL,
   Status            INT    NOT NULL
);

INSERT INTO cancellationRequests VALUES(1, 4, '2020-03-25', 1);



CREATE TABLE invoices(
   ID                INT PRIMARY KEY NOT NULL,
   PlantOrderId      INT    NOT NULL,
   Price             NUMERIC(18,2)  NOT NULL,
   Status            INT    NOT NULL,
   RemittanceID      INT
);

INSERT INTO invoices VALUES(1, 3, 42.4, 0, 1);
INSERT INTO invoices VALUES(2, 2, 422.4, 1, null);



CREATE TABLE remittances(
   ID                INT PRIMARY KEY NOT NULL,
   InvoiceId         INT    NOT NULL,
   ReferenceNumber   TEXT,
   Status            INT    NOT NULL
);

INSERT INTO remittances VALUES(1, 1, '23424', 1);