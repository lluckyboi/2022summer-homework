package dao

//
//func TestSelectUserByUserName(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//	username:="sario"
//	rows:=sqlmock.NewRows([]string{"id","username","User_mail"}).AddRow(1,"sario","15982730952@qq.com")
//	mock.ExpectQuery("select id,username,User_mail from user where username=?").WithArgs(username).WillReturnRows(rows)
//	resp, err := SelectUserByUserName(username)
//	if err != nil {
//		t.Fatal("get personals info err:", err)
//	} else {
//		fmt.Println(resp)
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Fatalf("there were unfulfilled expectations: %s", err)
//	}
//}
//
//func TestInsertUser(t *testing.T) {
//	type args struct {
//		user model.User
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := InsertUser(tt.args.user); (err != nil) != tt.wantErr {
//				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestUpdateWinCount(t *testing.T) {
//	type args struct {
//		userid int
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := UpdateWinCount(tt.args.userid); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateWinCount() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestSelectWinCount(t *testing.T) {
//	type args struct {
//		userid int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := SelectWinCount(tt.args.userid); got != tt.want {
//				t.Errorf("SelectWinCount() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
