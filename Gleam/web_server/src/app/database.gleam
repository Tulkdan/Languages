import app/error.{type AppError}
import app/transaction
import gleam/dynamic/decode
import gleam/result
import sqlight

pub type Connection =
  sqlight.Connection

pub fn with_connection(name: String, f: fn(sqlight.Connection) -> a) -> a {
  use db <- sqlight.with_connection(name)
  let assert Ok(_) = sqlight.exec("pragma foreign_keys = on;", db)
  f(db)
}

pub fn migrate_schema(db: Connection) -> Result(Nil, AppError) {
  sqlight.exec(
    "
    CREATE TABLE IF NOT EXISTS users (
      id              INTEGER PRIMARY KEY,
      limit_in_cents  INTEGER NOT NULL,
      initial_balance INTEGER NOT NULL DEFAULT 0
    );

    INSERT INTO users (limit_in_cents, initial_balance)
    VALUES (1000 * 100, 0),
          (800 * 100, 0),
          (10000 * 100, 0),
          (100000 * 100, 0),
          (5000 * 100, 0);

    CREATE TABLE IF NOT EXISTS history (
      id          INTEGER PRIMARY KEY,
      user_id     INTEGER NOT NULL,
      value       INTEGER NOT NULL,
      type        CHAR(1) NOT NULL,
      description VARCHAR(10) NOT NULL,
      do_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE INDEX IF NOT EXISTS idx_history ON history (user_id);
    ",
    db,
  )
  |> result.map_error(error.SqlightError)
}

pub fn get_user(
  db: Connection,
  user_id: String,
) -> Result(List(transaction.UserBalance), sqlight.Error) {
  let decoder = {
    use limit <- decode.field(0, decode.string)
    use balance <- decode.field(1, decode.int)

    transaction.UserBalance(limite: limit, saldo: balance)
    |> decode.success
  }

  sqlight.query(
    "SELECT limit_in_cents, initial_balance FROM users WHERE id = ?",
    on: db,
    with: [sqlight.text(user_id)],
    expecting: decoder,
  )
}

pub fn save_in_history(
  db: Connection,
  user_id: String,
  history: transaction.Transaction,
) -> Result(List(Int), sqlight.Error) {
  let decoder = {
    use id <- dynamic.element(0)
    dynamic.int(id)
  }

  sqlight.query(
    "INSERT INTO history (user_id, value, type, description)
    VALUES (?, ?, ?, ?)",
    on: db,
    with: [
      sqlight.text(user_id),
      sqlight.int(history.valor),
      sqlight.text(history.tipo),
      sqlight.text(history.descricao),
    ],
    expecting: decoder,
  )
}

pub fn add_credit(
  db: Connection,
  user_id: String,
  history: transaction.Transaction,
) -> Result(List(Int), sqlight.Error) {
  let decoder = {
    use id <- dynamic.element(0)
    dynamic.int(id)
  }

  sqlight.query(
    "INSERT INTO history (user_id, value, type, description)
    VALUES (?, ?, ?, ?)",
    on: db,
    with: [
      sqlight.text(user_id),
      sqlight.int(history.valor),
      sqlight.text(history.tipo),
      sqlight.text(history.descricao),
    ],
    expecting: decoder,
  )
}

pub fn add_debit(
  db: Connection,
  user_id: String,
  history: transaction.Transaction,
) -> Result(List(Int), sqlight.Error) {
  let decoder = {
    use id <- dynamic.element(0)
    dynamic.int(id)
  }

  sqlight.query(
    "INSERT INTO history (user_id, value, type, description)
    VALUES (?, ?, ?, ?)",
    on: db,
    with: [
      sqlight.text(user_id),
      sqlight.int(history.valor),
      sqlight.text(history.tipo),
      sqlight.text(history.descricao),
    ],
    expecting: decoder,
  )
}
