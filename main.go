package main

import(
  "fmt"
  "math/rand"
  "time"
)

func generatepassword(size int) string {
   charset := []string{"1","2","3","4","5","6","7","8","9","0","a",
     "b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r",
     "s","t","u","v","w","x","y","z","A","B","C","D","E","F","G","H",
     "I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}

    password := ""
    for i:=0; i<size; i++ {
      randnum := rand.Intn(62)
      password += charset[randnum]
    }

    return password
}

func main(){
    //to seed a random number with time
    rand.Seed(time.Now().Unix())
    password := generatepassword(16)
    fmt.Println("Password \n",password)
}
