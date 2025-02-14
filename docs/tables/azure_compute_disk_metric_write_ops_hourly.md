# Table: azure_compute_disk_metric_write_ops_hourly

Azure Monitor metrics provide data about the performance of your systems. The `azure_compute_disk_metric_write_ops_hourly` table provides metric statistics at 1 hour intervals for the most recent 60 days.

## Examples

### Basic info

```sql
select
  name,
  timestamp,
  minimum,
  maximum,
  average,
  sample_count
from
  azure_compute_disk_metric_write_ops_hourly
order by
  name,
  timestamp;
```

### Operations Over 10 Bytes average

```sql
select
  name,
  timestamp,
  round(minimum::numeric,2) as min_write_ops,
  round(maximum::numeric,2) as max_write_ops,
  round(average::numeric,2) as avg_write_ops,
  sample_count
from
  azure_compute_disk_metric_write_ops_hourly
where
  average > 10
order by
  name,
  timestamp;
```
