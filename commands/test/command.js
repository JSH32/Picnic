logger.info(args.author.id)
discord.sendMessage(args.channel.id, args.message.content)

class Rectangle {
    constructor(height, width) {
        this.height = height;
        this.width = width;
    }
}

let rec = new Rectangle(5, 7)

args.args.forEach((val) => {
    discord.sendMessage(args.channel.id, val)
})

discord.sendMessage(args.channel.id, rec.width)