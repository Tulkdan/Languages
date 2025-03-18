import gleam/dynamic.{type Dynamic}
import gleam/json

pub type Transaction {
  Transaction(valor: Int, tipo: String, descricao: String)
}

pub fn decode_transaction(
  json: Dynamic,
) -> Result(Transaction, dynamic.DecodeErrors) {
  let decoder =
    dynamic.decode3(
      Transaction,
      dynamic.field("valor", dynamic.int),
      dynamic.field("tipo", dynamic.string),
      dynamic.field("descricao", dynamic.string),
    )

  decoder(json)
}

pub type TransactionResponse {
  TransactionResponse(limite: Int, saldo: Int)
}

pub fn transaction_response(transaction: TransactionResponse) -> json.Json {
  json.object([
    #("limite", json.int(transaction.limite)),
    #("saldo", json.int(transaction.saldo)),
  ])
}

pub type UserBalance {
  UserBalance(limite: Int, saldo: Int)
}
