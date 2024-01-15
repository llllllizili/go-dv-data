golang 1.20



## 每周构建平均成功率
```sql
SELECT
	convert((
	SELECT
		(
		SELECT
			count(*) AS 构建状态
		FROM
			pipeline_data
		where
			status = 'success'
			and DATE_SUB(CURDATE(), INTERVAL 8 DAY) <= date(updata_at)) / ((
		SELECT
			count(*) AS 构建状态
		FROM
			pipeline_data
		where
			status = 'success'
			and DATE_SUB(CURDATE(), INTERVAL 8 DAY) <= date(updata_at)) + (
		SELECT
			count(*) AS 构建状态
		FROM
			pipeline_data
		where
			status != 'success'
			and DATE_SUB(CURDATE(), INTERVAL 8 DAY) <= date(updata_at))) * 100 as aaa),
	decimal(3)) as 成功率
```

## CICD季度增长
```sql
SELECT 
    t1.project,
    t1.total_duration AS cur_quarter_total_duration,
    t2.total_duration AS pre_quarter_total_duration,
    ROUND((t1.total_duration/3 - t2.total_duration/3) / t2.total_duration/3 * 100, 2) AS growth_rate
FROM
    (SELECT project, SUM(duration) AS total_duration, QUARTER(updata_at) AS quarter 
     FROM pipeline_data
     WHERE QUARTER(updata_at) = QUARTER(CURDATE())
     GROUP BY project, QUARTER(updata_at)
    ) AS t1
LEFT JOIN
    (SELECT project, SUM(duration) AS total_duration, QUARTER(updata_at) AS quarter
     FROM pipeline_data
     WHERE QUARTER(updata_at) = QUARTER(DATE_SUB(CURDATE(), INTERVAL 1 QUARTER)) 
     GROUP BY project, QUARTER(updata_at)
    ) AS t2
ON t1.project = t2.project
```