# BUG_REPRO

## Bug 是什么
严重等级分布按 category 分组；近 30 天趋势时间条件写反导致空；事件列表排序写反；等级默认值写错。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestIncidentDistributionAndTrend 失败：分布分组错、趋势为空。
