package service

import (
	"log/slog"
	"os"
	"testing"

	"safetyplatform/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newSfDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil { t.Fatalf("open sqlite: %v", err) }
	if err := db.AutoMigrate(&model.User{}, &model.SafetyIncident{}, &model.SafetyInspection{}, &model.InspectionItem{}, &model.WorkerCertification{}, &model.SafetyTraining{}, &model.AuditLog{}); err != nil { t.Fatalf("migrate: %v", err) }
	return db
}

func sfLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
}
