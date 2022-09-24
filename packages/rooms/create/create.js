const createClient = require('redis').createClient;
require('dotenv').config()

async function main(args) {
    const redis = createClient({
        url: process.env.DATABASE_URL
    });
    await redis.connect();
    room_template = {}
    room_template.id = args.id
    room_template.players = []
    room_template.creation_time = Date.now()
    room_template.bomb_start_time = 0
    room_template.words = []
    room_template.language = args.language
    room_template.current_word = ""
    room_template.password = args.password
    room_template.started = false
    room_template.finished = false
    await redis.set(args.id, JSON.stringify(room_template) )
    redis.disconnect();
    return {}
}
main()
exports.main = main;