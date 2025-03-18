import sqlight

pub type AppError {
  SqlightError(sqlight.Error)
}
