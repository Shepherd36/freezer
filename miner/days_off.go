// SPDX-License-Identifier: ice License 1.0

package miner

import (
	"context"
	"fmt"
	stdlibtime "time"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"

	"github.com/ice-blockchain/freezer/model"
	"github.com/ice-blockchain/freezer/tokenomics"
	messagebroker "github.com/ice-blockchain/wintr/connectors/message_broker"
	"github.com/ice-blockchain/wintr/log"
	"github.com/ice-blockchain/wintr/time"
)

func didANewDayOffJustStart(now *time.Time, usr *user) *DayOffStarted {
	if usr == nil ||
		usr.MiningSessionSoloStartedAt.IsNil() ||
		usr.MiningSessionSoloEndedAt.IsNil() ||
		usr.MiningSessionSoloLastStartedAt.IsNil() ||
		usr.BalanceLastUpdatedAt.IsNil() ||
		usr.MiningSessionSoloEndedAt.Before(*now.Time) ||
		usr.MiningSessionSoloLastStartedAt.Add(usr.maxMiningSessionDuration()).After(*now.Time) {
		return nil
	}
	naturalEndedAt := usr.MiningSessionSoloLastStartedAt.Add(usr.maxMiningSessionDuration())
	startedAt := time.New(naturalEndedAt.Add((now.Sub(naturalEndedAt) / cfg.MiningSessionDuration.Max) * cfg.MiningSessionDuration.Max))
	if usr.BalanceLastUpdatedAt.After(*startedAt.Time) {
		return nil
	}

	return &DayOffStarted{
		StartedAt:                   startedAt,
		EndedAt:                     time.New(startedAt.Add(cfg.MiningSessionDuration.Max)),
		UserID:                      usr.UserID,
		ID:                          fmt.Sprintf("%v~%v", usr.UserID, startedAt.UnixNano()/cfg.MiningSessionDuration.Max.Nanoseconds()),
		RemainingFreeMiningSessions: usr.calculateRemainingFreeMiningSessions(now),
		MiningStreak:                model.CalculateMiningStreak(now, usr.MiningSessionSoloStartedAt, usr.MiningSessionSoloEndedAt, cfg.MiningSessionDuration.Max),
	}
}

func dayOffStartedMessage(ctx context.Context, event *DayOffStarted) *messagebroker.Message {
	valueBytes, err := json.MarshalContext(ctx, event)
	log.Panic(errors.Wrapf(err, "failed to marshal %#v", event))

	return &messagebroker.Message{
		Headers: map[string]string{"producer": "freezer"},
		Key:     event.UserID,
		Topic:   cfg.MessageBroker.Topics[5].Name,
		Value:   valueBytes,
	}
}

func (u *user) maxMiningSessionDuration() stdlibtime.Duration {
	if u == nil || u.MiningBoostLevelIndex == nil {
		return cfg.MiningSessionDuration.Max
	}

	return stdlibtime.Duration((*cfg.miningBoostLevels.Load())[int(*u.MiningBoostLevelIndex)].MiningSessionLengthSeconds) * stdlibtime.Second
}

func (u *user) calculateRemainingFreeMiningSessions(now *time.Time) uint64 {
	if u == nil {
		return 0
	}
	start, end := u.MiningSessionSoloLastStartedAt, u.MiningSessionSoloEndedAt
	if end.IsNil() || now.After(*end.Time) {
		return 0
	}

	if maxMiningSession := u.maxMiningSessionDuration(); maxMiningSession > cfg.MiningSessionDuration.Max {
		latestMiningSession := tokenomics.CalculateMiningSession(now, start, end, maxMiningSession)

		if latestMiningSession == nil || end.Before(*latestMiningSession.EndedAt.Time) {
			return 0
		}

		return uint64(end.Sub(*latestMiningSession.EndedAt.Time) / cfg.MiningSessionDuration.Max)
	}

	return uint64(end.Sub(*now.Time) / cfg.MiningSessionDuration.Max)
}
