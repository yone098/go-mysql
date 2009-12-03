/*
	THIS IS NOT DONE AT ALL! USE AT YOUR OWN RISK!
    This source code was made referring to Eden's and Peter's source code.
    Eden http://github.com/eden/mysqlgo
    Peter http://github.com/phf/go-sqlite
*/
package mysql

/*
	Error in the database interface itself, *not*
	the database system we talk to.
*/
type InterfaceError struct {
	message string;
}

/*
	Textual description of the error.
	Implements os.Error interface.
*/
func (self InterfaceError) String() string {
	return self.message;
}

/*
	Error in the database system we talk to.
	MySQL has basic and extended status codes
	in addition to textual messages.
*/
type DatabaseError struct {
	message string;
	basic int;
	extended int;
}

/*
	Textual description of the error.
	Implements os.Error interface.
*/
func (self DatabaseError) String() string {
	return self.message;
}

/*
	Basic SQLite status code. These are plain
	integers.
*/
func (self DatabaseError) Basic() int {
	return self.basic;
}

/*
	Extended SQLite status code. These are OR'd
	together from various bits and pieces on top
	of basic status codes.
*/
func (self DatabaseError) Extended() int {
	return self.extended;
}
