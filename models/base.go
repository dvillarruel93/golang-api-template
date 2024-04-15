package models

import (
	helper "empty-api-struct/helper/pointer"
	"empty-api-struct/helper/uuid"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"time"
)

const (
	defaultCreatedBy = "diego.villarruel@somemail.com"
)

// ModelBase holds common base fields most models should have.
type ModelBase struct {
	ID         *string    `json:"id,omitempty" gorm:"type:string;size:20;primaryKey;"`
	CreatedBy  *string    `json:"created_by" gorm:"<-:create;type:string;size:255;not null"`
	CreatedAt  *time.Time `json:"created_at" gorm:"<-:create;type:time;not null;"`
	ModifiedBy *string    `json:"modified_by" gorm:"type:string; size:255; not null"`
	ModifiedAt *time.Time `json:"modified_at" gorm:"type:time; autoUpdateTime; not null"`
	DeletedBy  *string    `json:"-" gorm:"type:string; size:255; default null;"`
	DeletedAt  *time.Time `json:"-" gorm:"type:time; default:'0000-00-00 00:00:00.000'; not null;"`
}

// BeforeCreate sets the fields that are required to be created by each table in the database
func (m *ModelBase) BeforeCreate(tx *gorm.DB) (err error) {
	if _, ok := tx.Statement.Clauses["ON CONFLICT"]; ok {
		idSet := m.SetID(uuid.GenerateUUID())
		if !idSet {
			log.Debug("ID already set. Triggering upsert.")
			m.DeletedBy = nil
			m.DeletedAt = &time.Time{}
		}
	} else {
		m.ID = helper.ToStringPtr(uuid.GenerateUUID())
	}
	m.CreatedBy = helper.ToStringPtr(defaultCreatedBy)
	m.ModifiedBy = helper.ToStringPtr(defaultCreatedBy)
	return nil
}

func (m *ModelBase) SetID(id string) (ok bool) {
	if m.ID != nil {
		return false
	}
	m.ID = &id
	return true
}
