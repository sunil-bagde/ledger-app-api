TRUNCATE transactions;

UPDATE "public"."accounts" SET "available_amount" = 0 WHERE "id" = '31b03c53-5193-4f71-90e6-274d7e3e94bd';

UPDATE "public"."accounts" SET "available_amount" = 0 WHERE "id" = 'dfe3c883-7b6c-496c-8f09-024fe3f7fee7';

UPDATE "public"."accounts" SET "available_amount" = 50 WHERE "id" = '31b03c53-5193-4f71-90e6-274d7e3e94bd';



CREATE OR REPLACE FUNCTION calculate_rowhash() RETURNS TRIGGER AS $$
BEGIN
  NEW."RowHash" = digest(
    CONCAT(
      UPPER(COALESCE(NEW."CustomerName", '')), '|',
      COALESCE(CAST(NEW."Price" AS VARCHAR(50)), ''), '|',
      COALESCE(CAST(NEW."Quantity" AS VARCHAR(20)), ''), '|'
    ), 'sha512'
  );
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Transaction table

CREATE TABLE "CustomerTransactiont"
(
    "CustomerTransactionId" SERIAL PRIMARY KEY,
    "CustomerName" VARCHAR(255),
    "Price" NUMERIC(10,2),
    "Quantity" INTEGER,
    "RowHash" BYTEA
);

-- Create a trigger & will get called only INSERT

CREATE TRIGGER update_rowhash
BEFORE INSERT ON "CustomerTransactiont"
FOR EACH ROW EXECUTE FUNCTION calculate_rowhash();


INSERT INTO "public"."CustomerTransactiont" ("CustomerName", "Price", "Quantity") VALUES ('raj', 1, 1);


-- Select as binary , get hash column value
SELECT
	digest(CONCAT(UPPER(COALESCE("CustomerName", '')), '|', COALESCE(CAST("Price" AS VARCHAR(50)), ''), '|', COALESCE(CAST("Quantity" AS VARCHAR(20)), ''), '|'), 'sha512')
FROM
	"CustomerTransactiont";

-- Select As string, get hash column value

SELECT
	encode(digest(CONCAT(UPPER(COALESCE(CAST(sub. "CustomerName" AS VARCHAR), '')), COALESCE(CAST(sub. "Price" AS VARCHAR), ''),
	COALESCE(CAST(sub. "Quantity" AS VARCHAR), '')
	), 'sha512'), 'hex')
FROM (
	SELECT
		"CustomerName",
		"Price",
		"Quantity"
	FROM
		"CustomerTransactiont") AS sub;
