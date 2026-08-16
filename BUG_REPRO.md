# BUG_REPRO

## Bug 是什么
检查执行时 passed/issues 计数写反；得分漏乘 100；本月通过率漏乘 100；检查状态默认值写错。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestInspectionScoreAndStats 失败：得分/问题数/状态不正确。
