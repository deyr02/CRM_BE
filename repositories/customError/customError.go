package customerror

type FieldExist struct{}

func (m *FieldExist) Error() string {
	return "Field already exists"
}

type SystemField struct{}

func (m *SystemField) Error() string {
	return "System Field can not be modified"
}

type NewSystemFiled struct{}

func (m *NewSystemFiled) Error() string {
	return "Can not be added as a system field"
}

//-----------------------------------------
//-----------------User Role---------------
//-----------------------------------------

type NoRecordFound struct{}

func (m *NoRecordFound) Error() string {

	return "NO record found"
}

type UserRoleExist struct{}

func (m *UserRoleExist) Error() string {
	return "User Role exist"
}

type SystemRole struct{}

func (m *SystemRole) Error() string {
	return "System Role can not be modified"
}

//-----------------------------------------
//-----------------User Role---------------
//-----------------------------------------
type UserNameTaken struct{}

func (m *UserNameTaken) Error() string {
	return "Username taken"
}

type EmailAddressTaken struct{}

func (m *EmailAddressTaken) Error() string {
	return "Email already taken"
}

type SystemUser struct{}

func (m *SystemUser) Error() string {
	return "System Users can not be modified"
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "Wrong Userneame or Password"
}
