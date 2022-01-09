package zone

type ZoneActionType int

const (
	ZoneActionOpenAll  ZoneActionType = ZoneActionType(iota)
	ZoneActionCloseAll ZoneActionType = ZoneActionType(iota)
	ZoneActionOpenOne  ZoneActionType = ZoneActionType(iota)
	ZoneActionCloseOne ZoneActionType = ZoneActionType(iota)
)

type ZoneAction struct {
	ActionType ZoneActionType
	Zone       int
}

func HandleZoneActions(zoneActionChan chan ZoneAction) {
	closeAllZones()
	for {
		zoneAction := <-zoneActionChan
		switch zoneAction.ActionType {
		case ZoneActionOpenAll:
			openAllZones()
		case ZoneActionCloseAll:
			closeAllZones()
		case ZoneActionOpenOne:
			if zoneAction.Zone < 1 || zoneAction.Zone > ZoneCount {
				continue
			}
			openZone(zoneAction.Zone)
		case ZoneActionCloseOne:
			if zoneAction.Zone < 1 || zoneAction.Zone > ZoneCount {
				continue
			}
			closeZone(zoneAction.Zone)
		}
	}
}
