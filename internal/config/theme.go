package config

import "github.com/rivo/tview"

type (
	Theme struct {
		Border        bool   `toml:"border"`
		BorderPadding [4]int `toml:"border_padding"`

		GuildsTree   GuildsTreeTheme   `toml:"guilds_tree"`
		MessagesText MessagesTextTheme `toml:"messages_text"`
	}

	GuildsTreeTheme struct {
		AutoExpandFolders bool `toml:"auto_expand_folders"`
		Graphics          bool `toml:"graphics"`
	}

	MessagesTextTheme struct {
		ReplyIndicator string `toml:"reply_indicator"`
	}
)

func defaultTheme() Theme {
	return Theme{
		Border:        true,
		BorderPadding: [...]int{0, 0, 1, 1},

		GuildsTree: GuildsTreeTheme{
			AutoExpandFolders: true,
			Graphics:          true,
		},
		MessagesText: MessagesTextTheme{
			ReplyIndicator: string(tview.BoxDrawingsLightArcDownAndRight) + " ",
		},
	}
}
