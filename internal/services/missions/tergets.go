package missions

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/5aradise/sca-manager/internal/models"
)

var (
	ErrTargetNumOutOfRange    = errors.New("target number out of range")
	ErrMissionIsCompleted     = errors.New("mission is completed")
	ErrTargetIsCompleted      = errors.New("target is completed")
	ErrNumOfTargetsExceeds    = fmt.Errorf("number of targets exceeds maximum number (%d)", MaxTargetsCount)
	ErrNumOfTargetsCantBeLess = fmt.Errorf("number of targets can't be less than minimum number (%d)", MinTargetsCount)
)

func (s *service) MarkTargetAsCompleted(ctx context.Context, missionID int32, targetNum int) (models.Mission, error) {
	mission, err := s.ms.GetMissionById(ctx, missionID)
	if err != nil {
		return models.Mission{}, err
	}

	if !(0 <= targetNum && targetNum < len(mission.Targets)) {
		return models.Mission{}, ErrTargetNumOutOfRange
	}

	target := mission.Targets[targetNum]
	target, err = s.ts.UpdateTargetCompletionById(ctx, target.ID, true)
	if err != nil {
		return models.Mission{}, err
	}

	mission.Targets[targetNum] = target
	return mission, nil
}

func (s *service) UpdateTargetNotes(ctx context.Context, missionID int32, targetNum int, notes string) (models.Mission, error) {
	mission, err := s.ms.GetMissionById(ctx, missionID)
	if err != nil {
		return models.Mission{}, err
	}

	if mission.IsCompleted {
		return models.Mission{}, ErrMissionIsCompleted
	}

	if !(0 <= targetNum && targetNum < len(mission.Targets)) {
		return models.Mission{}, ErrTargetNumOutOfRange
	}

	target := mission.Targets[targetNum]
	if target.IsCompleted {
		return models.Mission{}, ErrTargetIsCompleted
	}

	target, err = s.ts.UpdateTargetNotesById(ctx, target.ID, notes)
	if err != nil {
		return models.Mission{}, err
	}

	mission.Targets[targetNum] = target
	return mission, nil
}

func (s *service) AddTarget(ctx context.Context, missionID int32, target models.Target) (models.Mission, error) {
	if err := IsValidTarget(target); err != nil {
		return models.Mission{}, err
	}

	mission, err := s.ms.GetMissionById(ctx, missionID)
	if err != nil {
		return models.Mission{}, err
	}

	if mission.IsCompleted {
		return models.Mission{}, ErrMissionIsCompleted
	}

	if len(mission.Targets) == MaxTargetsCount {
		return models.Mission{}, ErrNumOfTargetsExceeds
	}

	target, err = s.ts.CreateTarget(ctx, mission.ID, target)
	if err != nil {
		return models.Mission{}, err
	}
	mission.Targets = append(mission.Targets, target)

	return mission, nil
}

func (s *service) DeleteTarget(ctx context.Context, missionID int32, targetNum int) (models.Mission, error) {
	mission, err := s.ms.GetMissionById(ctx, missionID)
	if err != nil {
		return models.Mission{}, err
	}

	if mission.IsCompleted {
		return models.Mission{}, ErrMissionIsCompleted
	}

	if len(mission.Targets) == MinTargetsCount {
		return models.Mission{}, ErrNumOfTargetsCantBeLess
	}

	if !(0 <= targetNum && targetNum < len(mission.Targets)) {
		return models.Mission{}, ErrTargetNumOutOfRange
	}

	target := mission.Targets[targetNum]
	err = s.ts.DeleteTargetById(ctx, target.ID)
	if err != nil {
		return models.Mission{}, err
	}

	mission.Targets = slices.Delete(mission.Targets, targetNum, targetNum+1)
	return mission, nil
}
