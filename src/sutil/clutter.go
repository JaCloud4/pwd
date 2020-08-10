package sutil

func stor() {

}
//This version asks how many instances expected,
// then loops to capture integer entries only

func CreatesAfile(){
  fmt.Println("Name of File: ")
  reader := bufio.NewReader(os.Stdin)
  input, er := reader.ReadString('\n')
  if er !=nil {log.Fatal(er)}
  input= strings.TrimSpace(input)
  // printfile, fileerror:=sutil.Openfile(input)
  // fmt.Println(printfile,fileerror)
}