-- SELECT * FROM calls WHERE
--     call.func().name() = 'inc_getAdsUnionIdNew' AND
--     call.args() > 1 AND call.arg(1).isConstant() AND call.arg(0).string().contains('$union[')
-- LIMIT 20 ORDER BY args_count DESC

SELECT deps
FROM (
         SELECT *
         FROM funcs
         WHERE func.name() = 'inc_getAdsPriceLists'
     )
WHERE path.end().name().contains('rpc')
  AND path.length() = 6
  AND path.contains(
          SELECT *
          FROM funcs
          WHERE func.name() = 'rpcMcGet'
	  )
