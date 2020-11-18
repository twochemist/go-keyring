//with Go Keyring save Password in Linux 

/*For compile
export GOARCH=amd64
export GOOS=darwing
go build -o ./build/win/secret-api
export GOOS=window
*/

func login(server string) error {
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter Username: ")
user,_:=reader.ReadString('\n')
user =strings.Replace(user, "\n", "", -1)
if len(user)<2{
fmt.Println("no username entered")
return nil
}	

service := "secret-api"
pass :=""
pass, err := keyring.Get(service,user)
if err != nil{
fmt.Print("Enter Password: ")
bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
if err != nil {
fmt.Println("no password entered.")
}
pass = string(bytePassword)
	
err = keyring.Set(service, user, pass)
if err != nil{
fmt.Println("could not save to keychain.")
return nil
}
}

fmt.Println("we have a password")
fmt.Println(pass)
//Here it could the HTTP_Operation return nil
}