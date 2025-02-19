package missions

import (
	"context"

	"github.com/5aradise/sca-manager/internal/models"
)

type (
	missionStorage interface {
		CreateMission(ctx context.Context, mission models.Mission) (models.Mission, error)
		GetMissionById(ctx context.Context, id int32) (models.Mission, error)
		GetMissionByCatId(ctx context.Context, id int32) (models.Mission, error)
		ListMissions(ctx context.Context) ([]models.Mission, error)
		UpdateMissionCatById(ctx context.Context, id int32, catID int32) (models.Mission, error)
		UpdateMissionCompletionById(ctx context.Context, id int32, isCompleted bool) (models.Mission, error)
		DeleteMissionById(ctx context.Context, id int32) error
	}

	targetStorage interface {
		CreateTarget(ctx context.Context, missionID int32, target models.Target) (models.Target, error)
		UpdateTargetCompletionById(ctx context.Context, id int32, isCompleted bool) (models.Target, error)
		UpdateTargetNotesById(ctx context.Context, id int32, notes string) (models.Target, error)
		DeleteTargetById(ctx context.Context, id int32) error
	}

	service struct {
		ms missionStorage
		ts targetStorage
	}
)

func New(ms missionStorage, ts targetStorage) *service {
	return &service{ms, ts}
}
