package missions

import (
	"context"
	"errors"

	"github.com/5aradise/sca-manager/internal/models"
)

var (
	ErrMissionAlreadyAssigned = errors.New("mission is already assigned to a cat")
	ErrCatAlreadyAssigned     = errors.New("cat is already assigned to mission")
)

func (s *service) CreateMission(ctx context.Context, mission models.Mission) (models.Mission, error) {
	if err := IsValidTargets(mission.Targets); err != nil {
		return models.Mission{}, err
	}

	if mission.CatID != 0 {
		_, err := s.ms.GetMissionByCatId(ctx, mission.CatID)
		if err == nil {
			return models.Mission{}, ErrCatAlreadyAssigned
		}
	}

	newMission, err := s.ms.CreateMission(ctx, mission)
	if err != nil {
		return models.Mission{}, err
	}

	newMission.Targets = make([]models.Target, 0, len(mission.Targets))
	for _, target := range mission.Targets {
		newTarget, err := s.ts.CreateTarget(ctx, newMission.ID, target)
		if err != nil {
			s.ms.DeleteMissionById(context.Background(), newMission.ID)
			return models.Mission{}, err
		}
		newMission.Targets = append(newMission.Targets, newTarget)
	}
	return newMission, nil
}

func (s *service) GetMission(ctx context.Context, id int32) (models.Mission, error) {
	return s.ms.GetMissionById(ctx, id)
}

func (s *service) ListMissions(ctx context.Context) ([]models.Mission, error) {
	return s.ms.ListMissions(ctx)
}

func (s *service) MarkMissionAsCompleted(ctx context.Context, id int32) (models.Mission, error) {
	mission, err := s.ms.UpdateMissionCompletionById(ctx, id, true)
	if err != nil {
		return models.Mission{}, err
	}

	mission, err = s.ms.GetMissionById(ctx, mission.ID)
	if err != nil {
		return models.Mission{}, err
	}

	return mission, nil
}

func (s *service) AssignCat(ctx context.Context, missionID, catID int32) (models.Mission, error) {
	_, err := s.ms.GetMissionByCatId(ctx, catID)
	if err == nil {
		return models.Mission{}, ErrCatAlreadyAssigned
	}

	mission, err := s.ms.UpdateMissionCatById(ctx, missionID, catID)
	if err == nil {
		return models.Mission{}, err
	}

	mission, err = s.ms.GetMissionById(ctx, mission.ID)
	if err != nil {
		return models.Mission{}, err
	}

	return mission, nil
}

func (s *service) DeleteMission(ctx context.Context, id int32) error {
	mission, err := s.ms.GetMissionById(ctx, id)
	if err != nil {
		return nil
	}

	if mission.CatID != 0 {
		return ErrMissionAlreadyAssigned
	}
	return s.ms.DeleteMissionById(ctx, id)
}
