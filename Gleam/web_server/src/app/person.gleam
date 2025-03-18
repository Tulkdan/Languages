import gleam/dynamic.{type Dynamic}

pub type Person {
  Person(name: String, nickname: String, birthdate: String, stack: List(String))
}

pub fn decode_person(json: Dynamic) -> Result(Person, dynamic.DecodeErrors) {
  let decoder =
    dynamic.decode4(
      Person,
      dynamic.field("name", dynamic.string),
      dynamic.field("nickname", dynamic.string),
      dynamic.field("birthdate", dynamic.string),
      dynamic.field("stack", dynamic.list(dynamic.string)),
    )

  decoder(json)
}
