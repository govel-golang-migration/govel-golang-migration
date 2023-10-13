package main

import "errors"

func UpCreateMembersTable() error {
	// uncomment to test
	//fmt.Print("up create users table")
	//return nil
	return errors.New("up create members table error")
}

func DownCreateMembersTable() error {
	// uncomment to test
	//fmt.Print("up create users table")
	//return nil
	return errors.New("down create members table error")
}
