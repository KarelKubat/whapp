package privacysettings

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.PrivacySettings, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	p := ev.(*events.PrivacySettings)
	s := "PrivacySettings:"

	newsettings := ""
	newsettings = tools.Setting(newsettings, "GroupAdd", string(p.NewSettings.GroupAdd))
	newsettings = tools.Setting(newsettings, "LastSeen", string(p.NewSettings.LastSeen))
	newsettings = tools.Setting(newsettings, "Status", string(p.NewSettings.Status))
	newsettings = tools.Setting(newsettings, "Profile", string(p.NewSettings.Profile))
	newsettings = tools.Setting(newsettings, "ReadReceipts", string(p.NewSettings.ReadReceipts))
	s = tools.Setting(s, "NewSettings", newsettings)

	s = tools.Setting(s, "GroupAddchanged", tools.BoolStringIfTrue(p.GroupAddChanged))
	s = tools.Setting(s, "LastSeenChanged", tools.BoolStringIfTrue(p.LastSeenChanged))
	s = tools.Setting(s, "StatusChanged", tools.BoolStringIfTrue(p.StatusChanged))
	s = tools.Setting(s, "ProfileChanged", tools.BoolStringIfTrue(p.ProfileChanged))
	s = tools.Setting(s, "ReceiptReadChanged", tools.BoolStringIfTrue(p.ReadReceiptsChanged))

	fmt.Println(s)

	return nil
}
