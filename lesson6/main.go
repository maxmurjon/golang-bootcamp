package main

import "fmt"

func main(){
	var books =make(map[string][]string)
	var newAuthor string
	var authorBook string
	for true{
		davomi:=""
		fmt.Printf("Muallifni kiriting ")
		fmt.Scanln(&newAuthor)

		fmt.Printf("Kitobni kiriting kiriting ")
		fmt.Scanln(&authorBook)
		
		books[newAuthor]=[]string{authorBook}
		for true{
			fmt.Printf("Muallifga oid yana kitob bormi?(ha/yo'q) ")
			fmt.Scanln(&davomi)
			if davomi=="ha"{
				fmt.Printf("Kitobni kiriting kiriting ")
				fmt.Scanln(&authorBook)
				books[newAuthor]=append(books[newAuthor],authorBook)
			}else if davomi=="yo'q"{
				fmt.Println("Ok")
				break
			}
		}
		
		davom:=""
		fmt.Printf("Davom ettiraylikmi? ")
		fmt.Scanln(&davom)
		if davom=="ha"{
			continue
		}else if davom=="yo'q"{
			break
		}else {
			fmt.Println("Nomalum buruq kiritildi shuning uchun davom ettiramiz")
		}
	} 

	muallif:=""
	fmt.Println("Qaysi muallifga oid kitobni qidirmoqdasiz ")
	fmt.Scanln(&muallif)
	getBooksByAuthorName(books,muallif)
	fmt.Println("Qaysi muallifni  o'chirmoqchisiz ")
	fmt.Scanln(&muallif)
	deleteBook(books,muallif)
	

}

func getBooksByAuthorName(map1 map[string][]string,author string){
	for i,v:=range map1[author]{
		fmt.Printf("%v : %v\n",i ,v)
	}
}
func deleteBook(books map[string][]string,author string){
	delete(books,author)
}