
#ifndef CALL_H
#define CALL_H
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  char *username;
  int age;
  int gender;
} user;

typedef user *User;

extern void printSomething();
extern void cgoBridgeHandler(int id, User u);

static inline User print_user(User user);
static inline void callFromGo(int id);

#endif
