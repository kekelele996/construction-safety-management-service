package service
import (
 "errors"
 "testing"
 "safetyplatform/internal/model"
 "safetyplatform/internal/repository"
)
func expectNoPanicSf(t *testing.T,name string,fn func() error) error { t.Helper(); var err error; func(){ defer func(){ if r:=recover(); r!=nil { t.Fatalf("%s panicked: %v",name,r) } }(); err=fn() }(); return err }
func TestMissingIncidentAndInspection(t *testing.T){
 db:=newSfDB(t)
 ir:=repository.NewSafetyIncidentRepository(db); ur:=repository.NewUserRepository(db)
 isvc:=NewSafetyIncidentService(ir,ur,sfLogger())
 sr:=repository.NewSafetyInspectionRepository(db); iir:=repository.NewInspectionItemRepository(db)
 svc:=NewSafetyInspectionService(db,sr,iir,ur,sfLogger())
 err:=expectNoPanicSf(t,"Assign",func() error { _,err:=isvc.Assign(999); return err })
 if !errors.Is(err,repository.ErrNotFound){t.Fatalf("Assign=%v",err)}
 err=expectNoPanicSf(t,"Execute",func() error { _,err:=svc.Execute(999,[]model.InspectionItem{}); return err })
 if !errors.Is(err,repository.ErrNotFound){t.Fatalf("Execute=%v",err)}
}
