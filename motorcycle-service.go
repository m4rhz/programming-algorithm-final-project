package main

import "fmt"

const (
	MaxSpareParts   int = 2023
	MaxCustomers    int = 2023
	MaxTransactions int = 2023
)

type SparePart struct {
	ID       int
	Name     string
	Quantity int
}

type Customer struct {
	ID    int
	Name  string
	Phone string
	Email string
	Date  [3]int
}

type Transaction struct {
	ID           int
	CustomerID   int
	SparePartID  int
	Quantity     int
	ServiceRate  float64
	TotalService float64
	Date         [3]int
}

var spareParts [MaxSpareParts]SparePart
var customers [MaxCustomers]Customer
var transactions [MaxTransactions]Transaction

var sparePartCount int
var customerCount int
var transactionCount int

func main() {
	running := true
	for running {
		fmt.Println("\nService Motor")
		fmt.Println("1. Add Spare Part")
		fmt.Println("2. Edit Spare Part")
		fmt.Println("3. Add Customer")
		fmt.Println("4. Edit Customer")
		fmt.Println("5. Add Transaction")
		fmt.Println("6. Edit Transaction")
		fmt.Println("7. Delete Spare Part")
		fmt.Println("8. Delete Customer")
		fmt.Println("9. Delete Transaction")
		fmt.Println("10. Search Customers by Date")
		fmt.Println("11. Search Customers by Spare Part")
		fmt.Println("12. Print Spare Parts")
		fmt.Println("13. Print Customers")
		fmt.Println("14. Print Transactions")
		fmt.Println("0. Exit")

		fmt.Print("\nEnter your choice: ")
		var choice int
		fmt.Scan(&choice)
		if choice == 1 {
			addSparePart()
		} else if choice == 2 {
			editSparePart()
		} else if choice == 3 {
			addCustomer()
		} else if choice == 4 {
			editCustomer()
		} else if choice == 5 {
			addTransaction()
		} else if choice == 6 {
			editTransaction()
		} else if choice == 7 {
			deleteSparePart()
		} else if choice == 8 {
			deleteCustomer()
		} else if choice == 9 {
			deleteTransaction()
		} else if choice == 10 {
			searchCustomersByDate()
		} else if choice == 11 {
			searchCustomersBySparePart()
		} else if choice == 12 {
			printSpareParts()
		} else if choice == 13 {
			printCustomers()
		} else if choice == 14 {
			printTransactions()
		} else if choice == 0 {
			fmt.Println("Exiting...")
			running = false
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addSparePart() {
	fmt.Println("\nAdd Spare Part")
	if sparePartCount >= MaxSpareParts {
		fmt.Println("Maximum number of spare parts reached.")
	} else {
		var id, quantity int
		var name string

		fmt.Print("Enter ID: ")
		fmt.Scan(&id)
		fmt.Print("Enter Name: ")
		fmt.Scan(&name)
		fmt.Print("Enter Quantity: ")
		fmt.Scan(&quantity)

		sparePart := SparePart{ID: id, Name: name, Quantity: quantity}
		spareParts[sparePartCount] = sparePart
		sparePartCount++
		fmt.Println("Spare part added successfully.")
	}
}

func editSparePart() {
	fmt.Println("\nEdit Spare Part")
	if sparePartCount == 0 {
		fmt.Println("No spare parts available.")
	} else {
		var id, quantity int
		var name string

		fmt.Print("Enter ID: ")
		fmt.Scan(&id)

		found := false
		for i := 0; i < sparePartCount; i++ {
			if spareParts[i].ID == id {
				fmt.Print("Enter Name: ")
				fmt.Scan(&name)
				fmt.Print("Enter Quantity: ")
				fmt.Scan(&quantity)
				spareParts[i].Name = name
				spareParts[i].Quantity = quantity
				fmt.Println("Spare part edited successfully.")
				found = true
			}
		}
		if !found {
			fmt.Println("Spare part not found.")
		}
	}
}

func addCustomer() {
	fmt.Println("\nAdd Customer")
	if customerCount >= MaxCustomers {
		fmt.Println("Maximum number of customers reached.")
	} else {
		var id int
		var name, phone, email string
		var day, month, year int

		fmt.Print("Enter ID: ")
		fmt.Scan(&id)
		fmt.Print("Enter Name: ")
		fmt.Scan(&name)
		fmt.Print("Enter Phone: ")
		fmt.Scan(&phone)
		fmt.Print("Enter Email: ")
		fmt.Scan(&email)
		fmt.Print("Enter Date (Day Month Year): ")
		fmt.Scan(&day, &month, &year)

		customer := Customer{ID: id, Name: name, Phone: phone, Email: email}
		customer.Date = [3]int{day, month, year}
		customers[customerCount] = customer
		customerCount++
		fmt.Println("Customer added successfully.")
	}

}

func editCustomer() {
	fmt.Println("\nEdit Customer")
	if customerCount == 0 {
		fmt.Println("No customers available.")
	} else {
		var id int
		var name, phone, email string
		var day, month, year int

		fmt.Print("Enter ID: ")
		fmt.Scan(&id)

		found := false
		for i := 0; i < customerCount; i++ {
			if customers[i].ID == id && !found {
				fmt.Print("Enter Name: ")
				fmt.Scan(&name)
				fmt.Print("Enter Phone: ")
				fmt.Scan(&phone)
				fmt.Print("Enter Email: ")
				fmt.Scan(&email)
				fmt.Print("Enter Date (Day Month Year): ")
				fmt.Scan(&day, &month, &year)
				customers[i].Name = name
				customers[i].Phone = phone
				customers[i].Email = email
				customers[i].Date = [3]int{day, month, year}
				fmt.Println("Customer edited successfully.")
				found = true
			}
		}
		if !found {
			fmt.Println("Customer not found.")
		}
	}
}

func addTransaction() {
	fmt.Println("\nAdd Transaction")
	if transactionCount >= MaxTransactions {
		fmt.Println("Maximum number of transactions reached.")
	} else {
		var id, customerID, sparePartID, quantity int
		var serviceRate float64
		var day, month, year int

		fmt.Print("Enter ID: ")
		fmt.Scan(&id)
		fmt.Print("Enter Customer ID: ")
		fmt.Scan(&customerID)
		fmt.Print("Enter Spare Part ID: ")
		fmt.Scan(&sparePartID)
		fmt.Print("Enter Quantity: ")
		fmt.Scan(&quantity)
		fmt.Print("Enter Service Rate: ")
		fmt.Scan(&serviceRate)
		fmt.Print("Enter Date (Day Month Year): ")
		fmt.Scan(&day, &month, &year)

		transaction := Transaction{ID: id, CustomerID: customerID, SparePartID: sparePartID, Quantity: quantity, ServiceRate: serviceRate}
		transaction.Date = [3]int{day, month, year}
		transactions[transactionCount] = transaction
		transactionCount++
		fmt.Println("Transaction added successfully.")
	}
}

func editTransaction() {
	fmt.Println("\nEdit Transaction")
	if transactionCount == 0 {
		fmt.Println("No transactions available.")
	} else {
		var id, customerID, sparePartID, quantity int
		var serviceRate float64
		var day, month, year int

		fmt.Print("Enter ID: ")
		fmt.Scan(&id)

		found := false
		for i := 0; i < transactionCount; i++ {
			if transactions[i].ID == id && !false {
				fmt.Print("Enter Customer ID: ")
				fmt.Scan(&customerID)
				fmt.Print("Enter Spare Part ID: ")
				fmt.Scan(&sparePartID)
				fmt.Print("Enter Quantity: ")
				fmt.Scan(&quantity)
				fmt.Print("Enter Service Rate: ")
				fmt.Scan(&serviceRate)
				fmt.Print("Enter Date (Day Month Year): ")
				fmt.Scan(&day, &month, &year)
				transactions[i].CustomerID = customerID
				transactions[i].SparePartID = sparePartID
				transactions[i].Quantity = quantity
				transactions[i].ServiceRate = serviceRate
				transactions[i].Date = [3]int{day, month, year}
				fmt.Println("Transaction edited successfully.")
				found = true
			}
		}
		if !found {
			fmt.Println("Transaction not found.")
		}
	}
}

func deleteSparePart() {
	fmt.Println("\nDelete Spare Part")
	if sparePartCount == 0 {
		fmt.Println("No spare parts available.")
	} else {
		var id int
		fmt.Print("Enter ID: ")
		fmt.Scan(&id)
		found := false
		for i := 0; i < sparePartCount; i++ {
			if spareParts[i].ID == id {
				found = true
				for j := i; j < sparePartCount-1; j++ {
					spareParts[j] = spareParts[j+1]
				}
				spareParts[sparePartCount-1] = SparePart{}
				sparePartCount--
				fmt.Println("Spare part deleted successfully.")
			}
		}
		if !found {
			fmt.Println("Spare part not found.")
		}
	}
}

func deleteCustomer() {
	fmt.Println("\nDelete Customer")
	if customerCount == 0 {
		fmt.Println("No customers available.")
	} else {
		var id int
		fmt.Print("Enter ID: ")
		fmt.Scan(&id)
		found := false
		for i := 0; i < customerCount; i++ {
			if customers[i].ID == id {
				found = true
				for j := i; j < customerCount-1; j++ {
					customers[j] = customers[j+1]
				}
				customers[customerCount-1] = Customer{}
				customerCount--
				fmt.Println("Customer deleted successfully.")
			}
		}
		if !found {
			fmt.Println("Customer not found.")
		}
	}
}

func deleteTransaction() {
	fmt.Println("\nDelete Transaction")
	if transactionCount == 0 {
		fmt.Println("No transactions available.")
	} else {
		var id int
		fmt.Print("Enter ID: ")
		fmt.Scan(&id)
		found := false
		for i := 0; i < transactionCount; i++ {
			if transactions[i].ID == id {
				found = true
				for j := i; j < transactionCount-1; j++ {
					transactions[j] = transactions[j+1]
				}
				transactions[transactionCount-1] = Transaction{}
				transactionCount--
				fmt.Println("Transaction deleted successfully.")
			}
		}
		if !found {
			fmt.Println("Transaction not found.")
		}
	}
}

func searchCustomersByDate() {
	fmt.Println("\nSearch Customers by Date")
	var day, month, year int
	fmt.Print("Enter Date (Day Month Year): ")
	fmt.Scan(&day, &month, &year)

	fmt.Println("Customers who performed services in the specified period:")
	for i := 0; i < transactionCount; i++ {
		t := transactions[i]
		if t.Date[0] == day && t.Date[1] == month && t.Date[2] == year {
			customer := getCustomerByID(t.CustomerID)
			fmt.Printf("Customer ID: %d, Name: %s\n", customer.ID, customer.Name)
		}
	}
}

func searchCustomersBySparePart() {
	fmt.Println("\nSearch Customers by Spare Part")
	var sparePartID int
	fmt.Print("Enter Spare Part ID: ")
	fmt.Scan(&sparePartID)

	fmt.Println("Customers who purchased the specified spare part:")
	for i := 0; i < transactionCount; i++ {
		t := transactions[i]
		if t.SparePartID == sparePartID {
			customer := getCustomerByID(t.CustomerID)
			fmt.Printf("Customer ID: %d, Name: %s\n", customer.ID, customer.Name)
		}
	}
}

func printSpareParts() {
	fmt.Println("\nSpare Parts")
	if sparePartCount == 0 {
		fmt.Println("No spare parts available.")
	} else {
		for i := 1; i < sparePartCount; i++ {
			key := spareParts[i]
			j := i - 1
			for j >= 0 && spareParts[j].Quantity < key.Quantity {
				spareParts[j+1] = spareParts[j]
				j--
			}
			spareParts[j+1] = key
		}

		for i := 0; i < sparePartCount; i++ {
			sparePart := spareParts[i]
			fmt.Printf("ID: %d, Name: %s, Quantity: %d\n", sparePart.ID, sparePart.Name, sparePart.Quantity)
		}
	}
}

func printCustomers() {
	fmt.Println("\nCustomers")
	if customerCount == 0 {
		fmt.Println("No customers available.")
	} else {
		for i := 0; i < customerCount; i++ {
			customer := customers[i]
			fmt.Printf("ID: %d, Name: %s, Phone: %s, Email: %s, Date: %d-%d-%d\n", customer.ID, customer.Name, customer.Phone, customer.Email, customer.Date[0], customer.Date[1], customer.Date[2])
		}
	}
}

func printTransactions() {
	fmt.Println("\nTransactions")
	if transactionCount == 0 {
		fmt.Println("No transactions available.")
	} else {
		for i := 0; i < transactionCount-1; i++ {
			minIndex := i
			for j := i + 1; j < transactionCount; j++ {
				if compareDates(transactions[j].Date, transactions[minIndex].Date) < 0 {
					minIndex = j
				}
			}
			transactions[i], transactions[minIndex] = transactions[minIndex], transactions[i]
		}

		for i := 0; i < transactionCount; i++ {
			transaction := transactions[i]
			fmt.Printf("ID: %d, Customer ID: %d, Spare Part ID: %d, Quantity: %d, Service Rate: %.2f, Date: %d-%d-%d\n", transaction.ID, transaction.CustomerID, transaction.SparePartID, transaction.Quantity, transaction.ServiceRate, transaction.Date[0], transaction.Date[1], transaction.Date[2])
		}
	}
}

func compareDates(date1, date2 [3]int) int {
	if date1[2] != date2[2] {
		return date1[2] - date2[2]
	} else if date1[1] != date2[1] {
		return date1[1] - date2[1]
	} else {
		return date1[0] - date2[0]
	}
}

func getCustomerByID(id int) Customer {
	for i := 0; i < customerCount; i++ {
		if customers[i].ID == id {
			return customers[i]
		}
	}
	return Customer{}
}
