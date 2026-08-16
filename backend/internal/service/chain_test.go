package service
import (
 "testing"
 "time"
 "safetyplatform/internal/model"
 "safetyplatform/internal/repository"
)
func TestIncidentDistributionAndTrend(t *testing.T){
 db:=newSfDB(t)
 ir:=repository.NewSafetyIncidentRepository(db)
 now:=time.Now()
 mk:=func(sev string){ db.Create(&model.SafetyIncident{Title:"x",OccurredAt:now,SeverityLevel:sev,Category:"其他",Status:"reported",ReporterID:1,InvolvedUserIDs:model.JSONList([]string{}),PhotoURLs:model.JSONList([]string{})}) }
 mk("minor"); mk("major")
 dist,err:=ir.SeverityDistribution(); if err!=nil{t.Fatalf("dist: %v",err)}
 if len(dist)!=2{t.Fatalf("dist grouped wrong: %v",dist)}
 trend,err:=ir.Trend30(); if err!=nil{t.Fatalf("trend: %v",err)}
 if len(trend)==0{t.Fatalf("trend should not be empty")}
}
