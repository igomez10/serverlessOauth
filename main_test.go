package main

import (
	"testing"

	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/store"
)

var demoClient = &models.Client{
	ID:     "000000",
	Secret: "999999",
	Domain: "http://localhost",
}

var demoUserAlpha = &models.Client{
	ID:     "igomez10",
	Secret: "secret",
	Domain: "http://github.com",
}

var demoUserBeta = User{
	ID:        "userbeta",
	FirstName: "beta",
	LastName:  "ateb",
	password:  "passwordbeta",
}

func TestCreateNewUser(t *testing.T) {

	manager, storage := getBasicSetupWithEmptyStorage()

	IDToTest := "igomez10"

	storage.Set(IDToTest, demoUserAlpha)
	user, err := storage.GetByID(IDToTest)
	if err != nil {
		t.Fatalf("Error retreving user from storage with id %s", IDToTest)
	}

	if user.GetSecret() != demoUserAlpha.GetSecret() {
		errMessage := "Received wrong secret from storage, user secret should be %s but was %s"
		t.Errorf(errMessage, demoUserAlpha.GetSecret(), user.GetSecret())
	}

	cl, err := manager.GetClient("igomez10")
	if err != nil {
		t.Fatal("User could not be retrieved by using manager")
	}

	sec := cl.GetSecret()

	if sec != demoUserAlpha.GetSecret() {
		t.Errorf("Wrong secret received by the manager")
	}
}

func TestGetClientFromClientStorage(t *testing.T) {

	//	manager := getBasicSetupWithDemoEntryStorage()
	IDToTest := demoClient.GetID()
	ExpectedSecret := demoClient.GetSecret()

	manager := getBasicSetupWithDemoEntryStorage()
	cl, err := manager.GetClient(IDToTest)
	if err != nil {
		t.Fatalf(`Error getting client with id "%s" - error: %s`, IDToTest, err)
	}

	if cl.GetSecret() != ExpectedSecret {
		t.Errorf("Wrong secret, expected %s got %s", ExpectedSecret, cl.GetSecret())
	}

	t.Logf("%+v", cl)

}

func getBasicSetupWithDemoEntryStorage() *manage.Manager {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", demoClient)

	manager.MapClientStorage(clientStore)

	return manager
}

func getBasicSetupWithEmptyStorage() (*manage.Manager, *store.ClientStore) {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()

	manager.MapClientStorage(clientStore)

	return manager, clientStore
}
