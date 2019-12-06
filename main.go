// define a "packackage"- a container for the files here to communicate in a sandbox.
package main

// simple go import syntax
import (
  "fmt"
)

type Person struct {
    FirstName string
    LastName  string
    Home      Address 
}

type Address struct {
  address1  string
  address2  string
  city      string
  state     string
  country   string
  zipcode   string
}

type PersonCombineParams struct {
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

  p2.Home = Address{
    address1: "Gopher",
    address2: "World",
  }
  
  // shortcut "gopher" assignment syntax follows the order that properties were defined, Use Address instance with all "zero values" 
  p := Person{"Matt", "Bob", Address{}}
  
  // shortcut "gopher" assignment syntax follows the order that properties were defined
  p3 := Person{"Matt", "Bob", Address{}}

  fmt.Println(p3.Greeting())

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