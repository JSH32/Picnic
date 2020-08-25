logger.info(args.author.id)
discord.sendMessage(args.channel.id, args.message.content)

args.args.forEach((val) => {
    discord.sendMessage(args.channel.id, val)
})

let res = http.post("https://waifu.pics/api/many/sfw", JSON.stringify({
    exclude: []
}))

let v = new discord.embed()
    .setTitle("sdf")
    .setDescription("lmao")

logger.info(v.title)

discord.sendEmbed(args.channel.id, v)
// v.title = "fsd"

// v.send(args.channel.id)