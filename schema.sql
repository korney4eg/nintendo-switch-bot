PRAGMA encoding="UTF-8";
CREATE TABLE IF NOT EXISTS [games] (
  [game_id] integer PRIMARY KEY,
  [change_date] DATE NULL,
  [url] nvarchar(40) NOT NULL,
  [image_url] nvarchar(70),
  [paid_subscription_required_b] bool,
  [cloud_saves_b] bool,
  [valid_til] date NULL,
  [demo_availability] bool,
  [age_rating_sorting_i] byte NULL,
  [publisher] nvarchar(30) NULL,
  [excerpt] nvarchar(300) NULL,
  [date_from] date,
  [price_discount_percentage_f] float,
  [title_ru] nvarchar(70) NULL,
  [sorting_title] nvarchar(70) NOT NULL,
  [players_to] byte,
  [price_regular_rur] float NOT NULL,
  [price_lowest_rur] float NOT NULL,
  [version] int64
);
CREATE TABLE IF NOT EXISTS [categories] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [name] vnarchar(15),
  [russian_name] nvarchar(20) NULL,
  game_id integer NOT NULL,
  FOREIGN KEY(game_id) REFERENCES games(game_id)
);
CREATE TABLE IF NOT EXISTS [languages] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [lang_name] vnarchar(10),
  game_id integer NOT NULL,
  FOREIGN KEY(game_id) REFERENCES games(game_id)
);
