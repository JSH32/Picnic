logger.info(args.author.id)
discord.sendMessage(args.channel.id, args.message.content)

args.args.forEach((val) => {
    discord.sendMessage(args.channel.id, val)
})

// let res = http.post("https://waifu.pics/api/many/sfw", JSON.stringify({
//     exclude: []
// }))

new discord.embed({
    title: "test",
    description: "test"
}).send(args.channel.id)
