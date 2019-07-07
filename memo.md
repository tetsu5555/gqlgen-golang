```
mutation createTodo {
  createTodo(input:{text:"todo", userId:"T5577006791947779410"}) {
    user {
      id
      name
    }
    text
    done
  }
}

mutation createUser {
  createUser(input:{id:"1", name:"todo", password: "hellow world"}) {
    id
   	name
    password
  }
}

query findTodos {
  	todos {
      text
      done
      user {
        id
        name
        password
      }
    }
}

query findUsers {
  	users {
      id
      name
      password
    }
}

query findUser {
  user(ID: "T867466523082153551") {
    id
    name
  }
}
```