# Discordo

This is a personal fork of [Discordo](https://github.com/ayn2opdiscordo) with several personal modifications.

Color and formatting is slightly modified and unread indicators are shown for any channels that recieve new messages during the duration of your session.

- Lightweight
- Secure
- Configurable
- Cross-platform
- Minimalistic
- Feature-rich
  - Mouse & clipboard support
  - 2-Factor authentication
  - Partial [Discord-flavored markdown](https://support.discord.com/hc/en-us/articles/210298617-Markdown-Text-101-Chat-Formatting-Bold-Italic-Underline-)
  - Unread indicators

## Installation

You will have to build this package from source.

```bash
git clone https://github.com/vmhl87/discordo
cd discordo
go build .
```

### Linux clipboard support

- `xclip` or `xsel` for X11 (`apt install xclip`)
- `wl-clipboard` for Wayland (`apt install wl-clipboard`)

## Usage

1. Run the `discordo` executable with no arguments.

> If you are logging in using an authentication token, provide the `token` command-line flag to the executable (eg: `--token "OTI2MDU5NTQxNDE2Nzc5ODA2.Yc2KKA.2iZ-5JxgxG-9Ub8GHzBSn-NJjNg"`). The token is stored securely in the default OS-specific keyring.

2. Enter your email and password and click on the "Login" button to continue.

## Configuration

The configuration file allows you to configure and customize the behavior, keybindings, and theme of the application.

- Unix: `$XDG_CONFIG_HOME/discordo/config.toml` or `$HOME/.config/discordo/config.toml`
- Darwin: `$HOME/Library/Application Support/discordo/config.toml`
- Windows: `%AppData%/discordo/config.toml`

```toml
mouse = true

timestamps_format = "3:04PM"

messages_limit = 50
editor = "default"

[keys]
focus_guilds_tree = "Ctrl+G"
focus_messages_text = "Ctrl+T"
focus_message_input = "Ctrl+P"
exit_messages_text = "Ctrl+W"
toggle_guild_tree = "Ctrl+B"
select_previous = "Rune[k]"
select_next = "Rune[j]"
select_first = "Rune[g]"
select_last = "Rune[G]"

[keys.guilds_tree]
select_current = "Enter"

[keys.messages_text]
select_reply = "Rune[s]"
reply = "Rune[r]"
reply_send = "Ctrl+R"
reply_mention = "Rune[R]"
delete = "Rune[d]"
yank = "Rune[y]"
open = "Rune[o]"

[keys.message_input]
send = "Enter"
editor = "Ctrl+E"
cancel = "Esc"

[theme]
border = true
border_color = "default"
border_padding = [0, 0, 1, 1]
title_color = "default"
background_color = "default"

[theme.guilds_tree]
auto_expand_channels = true
auto_expand_folders = true
graphics = true

[theme.messages_text]
author_color = "aqua"
reply_indicator = "╭ "
```

## Disclaimer

Automated user accounts or "self-bots" are against Discord's Terms of Service. I am not responsible for any loss caused by using "self-bots" or Discordo.
