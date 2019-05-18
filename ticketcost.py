def calculate_total_ticket_cost(no_of_adults, no_of_children):
    total_ticket_cost=0
    #Write your logic here
    ticket_amount=(no_of_adults*37550)+(no_of_children*(37550/3))
    tax=ticket_amount+(ticket_amount*(7/100))
    total_ticket_cost=tax-(tax*(10/100))

    return total_ticket_cost


#Provide different values for no_of_adults, no_of_children and test your program
total_ticket_cost=calculate_total_ticket_cost(3,1)
print("Total Ticket Cost:",total_ticket_cost)
