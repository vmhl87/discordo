package cmd

import (
	"context"
	"log"
	"runtime"
	"strings"
	"sort"

	"github.com/ayn2op/discordo/internal/constants"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/httputil/httpdriver"
	"github.com/rivo/tview"
)

func init() {
	api.UserAgent = constants.UserAgent
	gateway.DefaultIdentity = gateway.IdentifyProperties{
		OS:      runtime.GOOS,
		Browser: constants.Name,
		Device:  "",
	}
}

type State struct {
	*state.State
}

func openState(token string) error {
	discordState = &State{
		State: state.New(token),
	}

	// Handlers
	discordState.AddHandler(discordState.onReady)
	discordState.AddHandler(discordState.onMessageCreate)
	discordState.AddHandler(discordState.onMessageDelete)

	discordState.OnRequest = append(discordState.Client.OnRequest, discordState.onRequest)
	return discordState.Open(context.TODO())
}

func (s *State) onRequest(r httpdriver.Request) error {
	req, ok := r.(*httpdriver.DefaultRequest)
	if ok {
		log.Printf("method = %s; url = %s\n", req.Method, req.URL)
	}

	return nil
}

func (s *State) onReady(r *gateway.ReadyEvent) {
	root := mainFlex.guildsTree.GetRoot()
	dmNode := tview.NewTreeNode("Direct Messages")
	mainFlex.guildsTree.DMs = dmNode
	root.AddChild(dmNode)

	cs, err := discordState.Cabinet.PrivateChannels()
	if err != nil {
		log.Println(err)
		return
	}

	sort.Slice(cs, func(i, j int) bool {
		return cs[i].LastMessageID > cs[j].LastMessageID
	})

	for _, c := range cs {
		mainFlex.guildsTree.createChannelNode(dmNode, c)
	}

	dmNode.SetExpanded(false)

	folders := r.UserSettings.GuildFolders
	if len(folders) == 0 {
		for _, g := range r.Guilds {
			mainFlex.guildsTree.createGuildNode(root, g.Guild)
		}
	} else {
		for _, folder := range folders {
			// If the ID of the guild folder is zero, the guild folder only contains single guild.
			if folder.ID == 0 {
				gID := folder.GuildIDs[0]
				g, err := discordState.Cabinet.Guild(gID)
				if err != nil {
					log.Printf("guild %v not found in state: %v\n", gID, err)
					continue
				}

				mainFlex.guildsTree.createGuildNode(root, *g)
			} else {
				mainFlex.guildsTree.createFolderNode(folder)
			}
		}
	}

	mainFlex.guildsTree.SetCurrentNode(root)
	app.SetFocus(mainFlex.guildsTree)
}

func indicate(t *tview.TreeNode) {
	if !strings.HasPrefix(t.GetText(), "[::b]") {
		t.SetText("[::b]" + t.GetText() + "[::B]")
	}
}

func (s *State) onMessageCreate(m *gateway.MessageCreateEvent) {
	if mainFlex.guildsTree.selectedChannelID.IsValid() && mainFlex.guildsTree.selectedChannelID == m.ChannelID {
		mainFlex.messagesText.createMessage(m.Message)
	} else {
		if len(mainFlex.guildsTree.DMs.GetChildren()) != 0 {
			for _, channel := range mainFlex.guildsTree.DMs.GetChildren() {
				if channel.GetReference() == m.ChannelID {
					indicate(channel)
					if !mainFlex.guildsTree.DMs.IsExpanded() {
						indicate(mainFlex.guildsTree.DMs)
					}
				}
			}
		}

		for _, guild := range mainFlex.guildsTree.GetRoot().GetChildren() {
			if guild.GetReference() == m.GuildID {
				if len(guild.GetChildren()) == 0 || !guild.IsExpanded() {
					indicate(guild)
				}

				for _, channel := range guild.GetChildren() {
					if len(channel.GetChildren()) == 0 {
						if channel.GetReference() == m.ChannelID {
							indicate(channel)
						}
					} else {
						for _, chi := range channel.GetChildren() {
							if chi.GetReference() == m.ChannelID {
								indicate(chi)
								if !channel.IsExpanded() {
									indicate(channel)
								}
							}
						}
					}
				}
			}
		}
	}
}

func (s *State) onMessageDelete(m *gateway.MessageDeleteEvent) {
	if mainFlex.guildsTree.selectedChannelID == m.ChannelID {
		mainFlex.messagesText.selectedMessage = -1
		mainFlex.messagesText.Highlight()
		mainFlex.messagesText.Clear()

		mainFlex.messagesText.drawMsgs(m.ChannelID)
	}
}
