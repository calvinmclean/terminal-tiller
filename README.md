# Terminal Tiller

Terminal Tiller is a simple idle game to grow, harvest, and sell crops from your terminal!

Built with tools from [charm.sh](https://charm.sh)!

![screenshot](screenshot.png)

## Install

```shell
go install github.com/calvinmclean/terminal-tiller@latest
```

## Play
After installing, simply run:

```shell
terminal-tiller
```

This will create a new save file in `~/.terminal-tiller`

See the on-screen usage instructions and start planting! When you close the game, the state is saved and your crops will continue growing!

## Conveniently Get Status Notifications

The `terminal-tiller -status` command will output a convenient summary for your farm without having to fully open it:

```
--------------------------------------------------------------
9 plants can be harvested, earning 135g
27 plants are still growing. Come back in 1m55s to harvest
13 plots are readed to be sowed
You currently have 729g
--------------------------------------------------------------
```

Add this command to your `~/.zshrc`, `~/.bashrc`, or `~/.config/fish/config.fish` to conveniently and non-disruptively receive the status message when you open a new terminal session.
