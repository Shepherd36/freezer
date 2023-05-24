// SPDX-License-Identifier: ice License 1.0

package extrabonusnotifier

import (
	stdlibtime "time"

	"github.com/ice-blockchain/wintr/time"
)

func isExtraBonusAvailable(
	currentTime, extraBonusStartDate *time.Time,
	extraBonusIndicesDistribution map[uint16]map[uint16]uint16,
	usr *user,
) bool {
	const notifyHourStart, notifyHourEnd = 10, 20
	var (
		utcOffset                       = stdlibtime.Duration(usr.UTCOffset) * cfg.ExtraBonuses.UTCOffsetDuration
		location                        = stdlibtime.FixedZone(utcOffset.String(), int(utcOffset.Seconds()))
		now                             = currentTime.In(location)
		extraBonusStartDateWithLocation = extraBonusStartDate.In(location)
		currentExtraBonusIndex          = 1 + uint16(now.Sub(extraBonusStartDateWithLocation)/cfg.ExtraBonuses.Duration)
		chunkNumber                     = uint16(usr.ID % int64(len(extraBonusIndicesDistribution)))
		_, dayHasExtraBonus             = extraBonusIndicesDistribution[chunkNumber][currentExtraBonusIndex]
	)

	if !usr.ExtraBonusLastClaimAvailableAt.IsNil() {
		usr.extraBonusIndex = 1 + uint16(usr.ExtraBonusLastClaimAvailableAt.In(location).Sub(extraBonusStartDateWithLocation)/cfg.ExtraBonuses.Duration)
	}

	if !dayHasExtraBonus ||
		usr.extraBonusIndex >= currentExtraBonusIndex ||
		now.Hour() < notifyHourStart || now.Hour() > notifyHourEnd ||
		(!usr.ExtraBonusStartedAt.IsNil() && currentTime.Before(usr.ExtraBonusStartedAt.Add(cfg.ExtraBonuses.Duration-cfg.ExtraBonuses.AvailabilityWindow))) ||
		(!usr.ExtraBonusLastClaimAvailableAt.IsNil() && currentTime.Before(usr.ExtraBonusLastClaimAvailableAt.Add(cfg.ExtraBonuses.Duration-cfg.ExtraBonuses.AvailabilityWindow))) ||
		now.Before(stdlibtime.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), notifyHourStart, 0, 0, 0, location).
			Add(stdlibtime.Duration(chunkNumber)*(cfg.ExtraBonuses.AvailabilityWindow-cfg.ExtraBonuses.ClaimWindow)/stdlibtime.Duration(len(extraBonusIndicesDistribution)))) || //nolint:lll // .
		now.After(stdlibtime.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), notifyHourEnd, 0, 0, 0, location)) {
		return false
	}
	usr.ExtraBonusLastClaimAvailableAt = currentTime
	usr.ExtraBonusDaysClaimNotAvailable = currentExtraBonusIndex - usr.extraBonusIndex - 1
	usr.extraBonusIndex = currentExtraBonusIndex

	return true
}