import app/database
import gleam/bool
import gleam/json
import wisp

pub type Context {
  Context(db: database.Connection)
}

pub fn middleware(
  req: wisp.Request,
  handle_request: fn(wisp.Request) -> wisp.Response,
) {
  let req = wisp.method_override(req)
  use <- wisp.log_request(req)
  use <- wisp.rescue_crashes
  use req <- wisp.handle_head(req)

  use <- default_responses

  handle_request(req)
}

fn default_responses(handle_request: fn() -> wisp.Response) -> wisp.Response {
  let response = handle_request()

  use <- bool.guard(when: response.body != wisp.Empty, return: response)

  case response.status {
    404 | 405 ->
      json.object([#("msg", json.string("Not Found"))])
      |> json.to_string_builder
      |> wisp.json_response(response.status)

    400 | 422 ->
      json.object([#("msg", json.string("Bad request"))])
      |> json.to_string_builder
      |> wisp.json_response(response.status)

    500 ->
      json.object([#("msg", json.string("Internal server error"))])
      |> json.to_string_builder
      |> wisp.json_response(response.status)

    _ ->
      json.object([#("msg", json.string("Something"))])
      |> json.to_string_builder
      |> wisp.json_response(response.status)
  }
}
