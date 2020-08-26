const http = require("http")
const discord = require("discord")

let res = http.get("https://waifu.pics/api/sfw")

discord.sendMessage(args.channel.id, JSON.parse(res.data.text).url)