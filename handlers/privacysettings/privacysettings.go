package privacysettings

import (
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type privacysettings struct {
	pr *events.PrivacySettings
}

func New(p *events.PrivacySettings) *privacysettings {
	return &privacysettings{
		pr: p,
	}
}

func (p *privacysettings) String() string {
	s := "PrivacySettings:"

	newsettings := ""
	newsettings = tools.Setting(newsettings, "GroupAdd", string(p.pr.NewSettings.GroupAdd))
	newsettings = tools.Setting(newsettings, "LastSeen", string(p.pr.NewSettings.LastSeen))
	newsettings = tools.Setting(newsettings, "Status", string(p.pr.NewSettings.Status))
	newsettings = tools.Setting(newsettings, "Profile", string(p.pr.NewSettings.Profile))
	newsettings = tools.Setting(newsettings, "ReadReceipts", string(p.pr.NewSettings.ReadReceipts))
	s = tools.Setting(s, "NewSettings", newsettings)

	s = tools.Setting(s, "GroupAddchanged", tools.BoolStringIfTrue(p.pr.GroupAddChanged))
	s = tools.Setting(s, "LastSeenChanged", tools.BoolStringIfTrue(p.pr.LastSeenChanged))
	s = tools.Setting(s, "StatusChanged", tools.BoolStringIfTrue(p.pr.StatusChanged))
	s = tools.Setting(s, "ProfileChanged", tools.BoolStringIfTrue(p.pr.ProfileChanged))
	s = tools.Setting(s, "ReceiptReadChanged", tools.BoolStringIfTrue(p.pr.ReadReceiptsChanged))
	return s
}
