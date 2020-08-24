logger.info(args.author.id)
discord.sendMessage(args.channel.id, args.message.content)

args.args.forEach(function (val, i) {
    discord.sendMessage(args.channel.id, val)
})