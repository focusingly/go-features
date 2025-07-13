#include "call.h"

static inline User print_user(User user) {
  if (user == NULL) {
    printf("null user!\n");
    return user;
  }

  printf("user name is: %s\n", user->username);
  printf("user name age: %d\n", user->age);
  printf("user gender is: %s\n", user->gender ? "male" : "female");
  fflush(stdout);

  return user;
}

static inline void callFromGo(int id) {
  user u = {.age = 12, .username = "hello world", .gender = 1};
  cgoBridgeHandler(id, &u);
}
