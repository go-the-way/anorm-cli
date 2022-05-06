package bundle

var tableXML = `<?xml version="1.0" encoding="UTF-8"?>
<batis-mapper binding="table">
    <select id="SelectTableList">
		SELECT 
		  t.TABLE_NAME,
		  IFNULL(t.TABLE_COMMENT, '') as TABLE_COMMENT
		FROM
		  information_schema.TABLES AS t 
		WHERE t.TABLE_SCHEMA = '{{.DBName}}'
		{{.Where}}
    </select>
    <select id="SelectTableColumnList">
		SELECT
		  t.TABLE_NAME,
		  t.COLUMN_KEY,
		  t.COLUMN_NAME,
		  t.IS_NULLABLE = 'NO' AS NOT_NULL,
		  IFNULL(t.COLUMN_DEFAULT, '__NULL__') AS COLUMN_DEFAULT,
		  IF(
			LOCATE('(', t.COLUMN_TYPE) = 0,
			t.COLUMN_TYPE,
			SUBSTRING(
			  t.COLUMN_TYPE,
			  1,
			  LOCATE('(', t.COLUMN_TYPE) - 1
			)
		  ) AS COLUMN_TYPE,
			t.COLUMN_TYPE as COLUMN_DATA_TYPE,
		  IFNULL(t.COLUMN_COMMENT, '') AS COLUMN_COMMENT,
		  t.EXTRA = 'auto_increment' as AUTO_INCREMENT
		FROM
		  information_schema.COLUMNS AS t
		WHERE t.TABLE_SCHEMA = '{{.}}'
		ORDER BY t.TABLE_NAME ASC,
		  t.ORDINAL_POSITION ASC
    </select>
    <select id="SelectTableIndexList">
		select t.TABLE_NAME,
		   t.INDEX_NAME,
		   t.NON_UNIQUE = 0                                        as INDEX_UNIQUE,
		   group_concat(t.COLUMN_NAME order by t.SEQ_IN_INDEX asc) as INDEX_COLUMN
	from information_schema.STATISTICS as t
	where t.TABLE_SCHEMA = '{{.}}'
	  and t.INDEX_NAME != 'PRIMARY'
	group by t.TABLE_NAME, t.INDEX_NAME
    </select>
</batis-mapper>`
