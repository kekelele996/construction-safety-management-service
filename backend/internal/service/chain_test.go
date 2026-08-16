package service
import (
 "testing"
 "time"
 "safetyplatform/internal/constants"
 "safetyplatform/internal/model"
 "safetyplatform/internal/repository"
)
func TestInspectionScoreAndStats(t *testing.T){
 db:=newSfDB(t)
 db.Create(&model.User{Phone:"13800000001",PasswordHash:"h",Name:"检查员",Role:"inspector"})
 ir:=repository.NewSafetyInspectionRepository(db); iir:=repository.NewInspectionItemRepository(db); ur:=repository.NewUserRepository(db)
 svc:=NewSafetyInspectionService(db,ir,iir,ur,sfLogger())
 ins,err:=svc.Create("周检","routine","A",time.Now(),1,[]model.InspectionItem{{ItemName:"灭火器"},{ItemName:"电路"}})
 if err!=nil{t.Fatalf("Create: %v",err)}
 got,err:=svc.Execute(ins.ID,[]model.InspectionItem{{ID:1,Passed:true},{ID:2,Passed:false}})
 if err!=nil{t.Fatalf("Execute: %v",err)}
 if got.TotalScore!=50 || got.IssueCount!=1 || got.Status!=constants.InspectionFailed {t.Fatalf("score/issue/status wrong: %+v",got)}
 st,err:=svc.Stats(); if err!=nil{t.Fatalf("Stats: %v",err)}
 if st["completed_rate"]!=0.0 {t.Fatalf("completed_rate=%v",st["completed_rate"])}
}
