SELECT * FROM "sale_product" AS SP
JOIN "sale" AS S ON S.id = SP.id
JOIN "remainder" AS R ON R.branch_id = S.branch_id;

SELECT R.quantity FROM "sale_product" AS SP
LEFT JOIN "sale" AS S ON S.id = SP.sale_id
LEFT JOIN "remainder"AS R ON R.branch_id = S.branch_id;

SELECT 
  B.name,
  SUM(SP.quantity)AS quantity,
  SUM("price") AS price
FROM "sale_product" AS SP
JOIN "sale" AS S ON S.id = SP.sale_id
JOIN "branch" AS B ON B.id = S.branch_id
GROUP BY B.name;