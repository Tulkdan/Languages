import app/database
import app/person
import app/transaction
import app/web.{type Context}
import gleam/http.{Get, Post}
import gleam/io
import gleam/json
import gleam/result
import wisp.{type Request, type Response}

pub fn handle_request(req: Request, ctx: Context) -> Response {
  use req <- web.middleware(req)

  case wisp.path_segments(req) {
    ["clientes", id, "transacoes"] -> transactions(req, ctx, id)
    // ["clientes", id, "extrato"] -> transactions(req, ctx, id)
    [] -> home(req)
    ["pessoas"] -> people(req)

    //["person", id] -> show_person(req)
    _ -> wisp.not_found()
  }
}

fn transactions(req: Request, ctx: Context, user_id: String) -> Response {
  use <- wisp.require_method(req, Post)
  use json <- wisp.require_json(req)

  let result = {
    use transaction_data <- result.try(transaction.decode_transaction(json))

    let assert Ok(user_balance) = database.get_user(ctx.db, user_id)

    io.debug(user_balance)

    // let assert Ok(_) =
    //   database.save_in_history(ctx.db, user_id, transaction_data)

    // let assert Ok(_) = case transaction_data {
    //   transaction.Transaction(tipo: "c", _, _) ->
    //     database.add_credit(ctx.db, user_id, transaction_data)
    //   transaction.Transaction(tipo: "d", _, _) ->
    //     database.add_debit(ctx.db, user_id, transaction_data)
    //   _ -> Error("")
    // }

    transaction.TransactionResponse(limite: 0, saldo: transaction_data.valor)
    |> transaction.transaction_response
    |> json.to_string_builder
    |> Ok
  }

  case result {
    Ok(json) -> wisp.json_response(json, 201)
    _ -> wisp.unprocessable_entity()
  }
}

fn home(req: Request) -> Response {
  use <- wisp.require_method(req, Get)

  json.object([#("hello", json.string("world"))])
  |> json.to_string_builder
  |> wisp.json_response(200)
}

fn people(req: Request) -> Response {
  case req.method {
    // Get -> list_people()
    Post -> save_person(req)
    _ -> wisp.method_not_allowed([Get, Post])
  }
}

fn save_person(req: Request) -> Response {
  use json <- wisp.require_json(req)

  let result = {
    use person <- result.try(person.decode_person(json))

    json.object([#("id", json.string(person.name))])
    |> json.to_string_builder
    |> Ok
  }

  case result {
    Ok(json) -> wisp.json_response(json, 201)
    _ -> wisp.unprocessable_entity()
  }
}
