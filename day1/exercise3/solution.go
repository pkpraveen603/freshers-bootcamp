package main

import "fmt"

type GetSalary interface{
	CalculateSalary() int
}
type FullTimeEmp struct{
	name string
	daysOfWork int
}
func (emp *FullTimeEmp)CalculateSalary() int{
	return emp.daysOfWork*500
}

type ContractorEmp struct{
	name string
	daysOfWork int
}
func (emp *ContractorEmp)CalculateSalary() int{
	return emp.daysOfWork*100
}
type FreeLancerEmp struct{
	name string
	hoursOfWork int
}
func (emp *FreeLancerEmp)CalculateSalary() int{
	return emp.hoursOfWork*10
}

func main(){
	var employee GetSalary
	ftEmployee := FullTimeEmp{
		"Employee_FT1",
		28 }
	contractorEmployee := ContractorEmp{
		"Contractor_1",
		28 }
	freelanceEmployee := FreeLancerEmp{
		"FreeLancer_1",
		100 }
	employee = &ftEmployee
	fmt.Printf("Salary of %s is: %d for %d days of work. \n",ftEmployee.name,employee.CalculateSalary(),ftEmployee.daysOfWork )
	employee = &contractorEmployee
	fmt.Printf("Salary of %s is: %d for %d days of work. \n",contractorEmployee.name,employee.CalculateSalary(),contractorEmployee.daysOfWork )
	employee = &freelanceEmployee
	fmt.Printf("Salary of %s is: %d for %d hours of work. \n",freelanceEmployee.name,employee.CalculateSalary(),freelanceEmployee.hoursOfWork )
}