package proffixrest

//Tests disable because of bug in remote REST API
//
////TestClient_GetList
//func TestClient_GetList(t *testing.T) {
//
//	ctx := context.Background()
//
//	//Connect
//	pxrest, err := ConnectTest(ctx)
//
//	//Check error. Should be nil
//	if err != nil {
//		t.Errorf("Expected no error for Connect. Got '%v'", err)
//	}
//
//	_, headers, status, err := pxrest.GetList(ctx, "ADR_Adress-Etiketten Kontakte.labx", nil)
//
//	//Check error. Should be nil
//	if err != nil {
//		t.Errorf("Expected no error for GET List Request. Got '%v'", err)
//	}
//
//	//Check status. Should be 201
//	if status != 201 {
//		t.Errorf("Expected Status 201. Got '%v'", status)
//	}
//
//	//Check Content Type. Should be pdf
//	if headers.Get("Content-Type") != "application/pdf" {
//		t.Errorf("Expected Content-Type PDF. Got '%v'", status)
//	}
//
//	//Check Logout
//	statuslogout, err := pxrest.Logout(ctx)
//
//	//Check error. Should be nil
//	if err != nil {
//		t.Errorf("Expected no error for Logout Request. Got '%v'", err)
//	}
//
//	//Check HTTP Status of Logout. Should be 204
//	if statuslogout != 204 {
//		t.Errorf("Expected HTTP Status Code 204. Got '%v'", err)
//	}
//}
