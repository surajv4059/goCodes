package main
import (
	"fmt"
	"strings"
)
func main(){
	conferenceName := "Go Conference";
	const conferenceTickets = 50;
	var remainingTickets = 50;

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets,remainingTickets,conferenceName); 

	fmt.Printf("Welcome to %v booking application\n",conferenceName);
	fmt.Printf("We have total of %v tickets and %v are still avaialable. \n",conferenceTickets,remainingTickets);
	fmt.Println("Get your tickets here to attend");
	

	//ARRAY
	//var bookings = [50]string{}

	// var bookings [50]string
	

	// Slices

	//var bookings2 = []string{}
	// bookings2 := []string{}

	var bookings2 []string

	//Loops

	for {

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	fmt.Println("please enter your first name :")
	fmt.Scan(&firstName)

	fmt.Println("please enter your last name :")
        fmt.Scan(&lastName)

	fmt.Println("please enter your email :")
        fmt.Scan(&email)
	
	fmt.Println("please enter your userTickets :")
        fmt.Scan(&userTickets)
	
	//adding elements in arrays
	// bookings[0] = firstName + " " + lastName
	// bookings[1] = "nana"
	
	
	//adding elemets in slices
	bookings2 =  append(bookings2, firstName + " " + lastName)

	// fmt.Println(bookings);
	// fmt.Println(bookings2);


	if userTickets < remainingTickets{

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you. %v %v booked %v tickets. You will recieve a confirmation on %v \n", firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v \n",remainingTickets,conferenceName);
	

	firstNames := []string{}
	for _ , booking := range bookings2 {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("The first names of bookings are: %v\n", firstNames);

	if remainingTickets == 0 {
		//end program
		fmt.Println("our conference is booked out. come back next year")
		break
	}

	
	
 
	} else {
		 fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n",remainingTickets,userTickets)
		
	}	
}
}
