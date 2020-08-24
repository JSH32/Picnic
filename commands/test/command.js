logger.info(message.userid)
discord.sendMessage(message.channelid, message.content)

message.args.forEach(function (val, i) {
    discord.sendMessage(message.channelid, val)
})