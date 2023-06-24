# SoundCloud RPC

SoundCloud RPC is a small program written in Go that sits in your system tray and displays your current SoundCloud song on Discord using Rich Presence.

## Installation

To install SoundCloud RPC, you'll need to download the appropriate binary for your platform from releases. You'll then want to place it in a new folder, and create a config.json file. Your file should look something like [this](https://github.com/connordennison/soundcloud-rpc/blob/main/config.json.example).

## How to obtain a client_id and auth_token
SoundCloud shut down the old developer API. This makes use of the internal (v2) API. You can obtain credentials by inspecting a network request. To do this DevTools, click "Network", and reload the page. In the "Filter URLs" box enter `me?client_id=` and click on the request with the type of JSON. In the "Headers" tab you'll see the client ID, and in the "Request" tab you'll find the auth_token.

## Running
Launch the executable in the same folder as your config.json.
