logger.info(args.author.id)
discord.sendMessage(args.channel.id, args.message.content)

args.args.forEach((val) => {
    discord.sendMessage(args.channel.id, val)
})