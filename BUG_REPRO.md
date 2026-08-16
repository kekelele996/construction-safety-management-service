# BUG_REPRO

## Bug 是什么
事件指派/关闭状态检查写反；待整改列表把 resolved/closed 算进去；事件状态默认值写错。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestIncidentStateAndPending 失败：指派/关闭状态流转被拒、待整改列表多出已结束事件。
