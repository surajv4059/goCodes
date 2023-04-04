package main
import "fmt"

func main(){
	conferenceName := "Go Conference";
	const conferenceTickets = 50;
	var remainingTickets = 50;

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets,remainingTickets,conferenceName); 

	fmt.Printf("Welcome to %v booking application\n",conferenceName);
	fmt.Printf("We have total of %v tickets and %v are still avaialable. \n",conferenceTickets,remainingTickets);
	fmt.Println("Get your tickets here to attend");
	
	//var bookings = [50]string{}
	var bookings [50]string
	
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
	
	bookings[0] = firstName + " " + lastName
	bookings[1] = "nana"


	fmt.Println(bookings);

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you. %v %v booked %v tickets. You will recieve a confirmation on %v \n", firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v \n",remainingTickets,conferenceName);
}
