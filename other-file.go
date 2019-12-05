package main

import (
  "fmt"
)

func (person Person) Wave() string {
  return fmt.Sprintf("%s is waving at you!", person.FirstName)
}

// Lowercase functions are only hidden from OTHER PACKAGES but are still callable from anywhere withing the current package.
func (person Person) privateFunc() string {
  return fmt.Sprintf("Can't call me from outside this file!")
}