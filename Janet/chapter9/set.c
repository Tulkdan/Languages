#include <janet.h>

static Janet cfun_hello(int32_t argc, Janet *argv) {
  janet_fixarity(argc, 0);
  printf("hello world\n");
  return janet_wrap_nil();
}

static JanetReg cfuns[] = {
  {"hello", cfun_hello, "(hello)\n\nprints hello"},
  {NULL, NULL, NULL}
};

JANET_MODULE_ENTRY(JanetTable *env) {
  janet_cfuns(env, "set", cfuns);
}
