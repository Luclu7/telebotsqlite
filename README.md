# Telebot SQLite

This project is just a generic inline telegram bot. It only suppors audios and pictures for now.

You change easily change the content by editing the db.sqlite file (table "entries") with this structure:

| name           | url                                                                                  | thumburl                                                                            | type     | mime              |
|----------------|--------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|----------|-------------------|
| tux            | https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Tux.svg/1000px-Tux.svg.png | https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Tux.svg/500px-Tux.svg.png | photo    | anything you want |
| ultima 6 intro | https://luclu7.fr/pub/ultima6intro.mp3                                               | anything you want                                                                   | audio    | anything you want |
| 42.zip         | http://www.unforgettable.dk/42.zip                                                   | anything you want                                                                   | document | application/zip   |

Please note that the MIME type for documents (file) can only be either "application/pdf" or "application/zip" ([Telebot/Telegram API limitation](https://github.com/tucnak/telebot/blob/v2/inline_types.go#L109)

You must put something (whatever you want) in the _thumburl_ and _mime_ column if it's not respectively a picture or a document (file), at least for now.

As this program uses [go-sqlite3](https://github.com/mattn/go-sqlite3), it uses CGo.

## License
Feel free to fork it and do your own stuff with it, this project is under the Unlicense, see [LICENSE.md](LICENSE.md).
