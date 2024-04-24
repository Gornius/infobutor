# Infobutor
Distribute messages across many communication channels using a single, centralized HTTP interface!

## Purpose
When hosting services, it's convenient to set up sending messages on various events, eg. cronjobs and backups. However it quickly becomes cumbersome, as you need to set up API keys and other credentials in every script you use. Additionally when your API key changes or you simply want to change the way you send out messages, you need to change every script on every server.

This program allows you to set up a single sink of all your messages, and redistribute it the way you want - including multiple channels at once.

## Example use case

For example you can create a discord chat with all your logs about backups. Then you set up a sender in infobutor using this chat via WebHook. Then you can plug that sender into an infobutor channel, to which you send logs. Finally make backup script send message to infobutor channel whether backup was successful or not.

Then you do the same for 12 different servers.

Now you decide you also want to receive telegram messages. With infobutor, you just add new telegram sender and plug it into your infobutor channel.

Need to also send messages to another discord channel on another server? No problem - just create another Discord sender and plug it into your channel.

Using multiple channels you can send non-important messages to different places.

## Already available senders
- Discord
- Telegram
- Local file

## Planned senders
- E-mail
- Queued E-mail
