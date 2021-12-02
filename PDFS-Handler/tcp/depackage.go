package tcp

func depackage(byteStream []byte, pg *Package) bool {
	pg.Op = string(byteStream[0])
	usernameLength := int(byteStream[1])

	pg.username = string(byteStream[2 : 2+usernameLength])
	if pg.Op == NEW_USER_OP || pg.Op == DEL_USER_OP ||pg.Op == LOGIN_OP {
		passwdLength :=  int(byteStream[2+usernameLength])
		if passwdLength == 0{
			return false
		}
		pg.passwd = string(byteStream[3+usernameLength : 3+usernameLength+passwdLength])
		if pg.passwd == ""{
			return false
		}
	} else if pg.Op == CHANGE_PASSWD_OP {
		passwdLength :=  int(byteStream[2+usernameLength])
		if passwdLength == 0{
			return false
		}
		pg.passwd = string(byteStream[3+usernameLength : 3+usernameLength+passwdLength])
		if pg.passwd == ""{
			return false
		}
		newpasswdLength := int(byteStream[3+usernameLength+passwdLength])
		if newpasswdLength == 0{
			return false
		}
		pg.newpasswd = string(byteStream[4+usernameLength+passwdLength:4+usernameLength+passwdLength+newpasswdLength])
		if pg.newpasswd == ""{
			return false
		}
	} else if pg.Op == WRITE_OP || pg.Op == READ_OP || pg.Op == DEL_OP {
		filenameLength := int(byteStream[2+usernameLength])
		if filenameLength == 0{
			return false
		}
		pg.filename = string(byteStream[3+usernameLength : 3+usernameLength+filenameLength])
		if pg.filename == ""{
			return false
		}
		pathLength:= int(byteStream[3+usernameLength+filenameLength])
		if pathLength == 0{
			return false
		}
		pg.path = string(byteStream[4+usernameLength+filenameLength:4+usernameLength+filenameLength+pathLength])
		if pg.path == ""{
			return false
		}
	} else if pg.Op == NEW_PATH_OP || pg.Op == DEL_PATH_OP || pg.Op == ASK_FILES_OP {
		pathLength:= int(byteStream[2+usernameLength])
		if pathLength == 0{
			return false
		}
		pg.path = string(byteStream[3+usernameLength : 3+usernameLength+pathLength])
		if pg.path == ""{
			return false
		}
	}
	return true
}
