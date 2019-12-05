// define a "packackage"- a container for the files here to communicate in a sandbox.
package main

// simple go import syntax
import (
  "fmt"
)


type Person struct {
    FirstName, LastName string
}

func (person Person) Greeting() string {
  return fmt.Sprintf("Dear %s %s", person.FirstName, person.LastName)
}


func main() {

  const ok = 42

  // Instatiate a Person Type and set some properties
  var p2 = new(Person)
  p2.FirstName = "Freddy"
  p2.LastName = "jacobs"
  
  // shortcut "gopher" assignment syntax
  p := Person{"Matt", "Bob"}

  // Note- In Go you can't instatiate Types to a "const" variable.
  // Compile Error - (value of type *Person) of not a constant 
  // const p3 = new(Person) 

  fmt.Println(p.Greeting())
  fmt.Println(p2.Greeting())
  
  fmt.Println(p2.Wave())
  fmt.Println(p2.privateFunc())

  fmt.Println(ok)
  fmt.Println("Hello World")
}