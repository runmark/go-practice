package main

import (
	"errors"
	"fmt"
	"log"
)

type Contract struct {
	ID int64
}

type Customer struct {
	ID       int64
	Contract Contract
}

type Status struct {
	ok bool
}

type DBError struct {
}

func (e *DBError) Error() string {
	return fmt.Sprintf("db temp error")
}

func postHandler(customer Customer) Status {

	err := insert(customer.Contract)
	if err != nil {
		switch {
		default:
			//log.WithError(err).Errorf("unable to serve HTTP POST request for customer %s", customer.ID)
			log.Fatalf("unable to serve HTTP POST request for customer %d: %v", customer.ID, err)
			return Status{ok: false}
		case errors.Is(err, &DBError{}):
			return retry(customer)
		}
	}
	return Status{ok: true}
}

func retry(customer Customer) Status {
	return Status{ok: true}
}

//type QueryError struct {
//	Query string
//	Err   error
//}
//
//func (qe *QueryError) Error() string {
//	return qe.Query + ": " + qe.Err.Error()
//}
//
//func (qe *QueryError) Unwrap() error {
//	return qe.Err
//}

func insert(contract Contract) error {
	err := dbQuery(contract)
	if err != nil {
		//return errors.Wrapf(err, "unable to insert customer contract %s", contract.ID)
		//return &QueryError{Query: fmt.Sprintf("unable to insert customer contract %s", contract.ID), Err: err}
		return fmt.Errorf("unable to insert custom contract %d: %w", contract.ID, err)
	}
	return nil
}

func dbQuery(contract Contract) error {
	return errors.New("unable to commit transaction")
}

func main()  {
	s := postHandler(Customer{ID: 345, Contract: Contract{ID: 345-2}})
	fmt.Println(s)
}