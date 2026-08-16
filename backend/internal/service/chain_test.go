package service
import (
 "testing"
 "time"
 "safetyplatform/internal/repository"
)
func TestIncidentStateAndPending(t *testing.T){
 db:=newSfDB(t)
 ir:=repository.NewSafetyIncidentRepository(db); ur:=repository.NewUserRepository(db)
 svc:=NewSafetyIncidentService(ir,ur,sfLogger())
 inc,err:=svc.Report(1,"坠落","",time.Now(),"A","","minor","",[]string{},[]string{})
 if err!=nil{t.Fatalf("Report: %v",err)}
 if _,err:=svc.Assign(inc.ID);err!=nil{t.Fatalf("Assign: %v",err)}
 if _,err:=svc.SubmitRectification(inc.ID,"加固",nil);err!=nil{t.Fatalf("Rectify: %v",err)}
 if _,err:=svc.Close(inc.ID);err!=nil{t.Fatalf("Close: %v",err)}
 pend,_:=svc.PendingRectification()
 for _,p:=range pend{ if p.ID==inc.ID { t.Fatal("resolved incident should not be pending") } }
}
