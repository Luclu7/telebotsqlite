# Telebot SQLite

This project is just a generic inline telegram bot. It only suppors audios and pictures for now.

You change easily change the content by editing the db.sqlite file, with this structure, in the table *stuff* (yeah, that'll change):

| name        | url                                                                                  | thumburl                                                                            |
|-------------|--------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| photo - tux | https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Tux.svg/1000px-Tux.svg.png | https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Tux.svg/500px-Tux.svg.png |

Please note that pictures must contain "photo" somewhere in the name.
You can put whatever you want in the thumburl if it's not a picture.

I'll add a type column after, but for now it's like that.

Feel free to fork it and do your own stuff with it, it's under the Unlicense!
