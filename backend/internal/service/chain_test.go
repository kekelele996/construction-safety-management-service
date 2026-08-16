package service
import (
 "errors"
 "testing"
 "safetyplatform/internal/repository"
)
func TestErrorChainSentinels(t *testing.T){
 db:=newSfDB(t)
 ir:=repository.NewSafetyIncidentRepository(db); sr:=repository.NewSafetyInspectionRepository(db)
 if _,err:=ir.FindByID(999);!errors.Is(err,repository.ErrNotFound){t.Fatalf("incident FindByID=%v",err)}
 if _,err:=sr.FindByID(999);!errors.Is(err,repository.ErrNotFound){t.Fatalf("inspection FindByID=%v",err)}
}
