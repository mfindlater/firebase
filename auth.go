package firebase

import (
	"github.com/gopherjs/gopherjs/js"
)

type Persistence string

type Operation string

const (
	Local   Persistence = "LOCAL"
	None    Persistence = "NONE"
	Session Persistence = "SESSION"

	PasswordReset Operation = "PASSWORD_RESET"
	VerifyEmail   Operation = "VERIFY_EMAIL"
	RecoverEmail  Operation = "RECOVERY_EMAIL"
)

type UserInfo struct {
	*js.Object
	DisplayName string `js:"displayName"`
	Email       string `js:"email"`
	PhoneNumber string `js:"phoneNumber"`
	PhotoURL    string `js:"photoURL"`
	ProviderID  string `js:"provierId"`
	UID         string `js:"uid"`
}

type UserMetadata struct {
	*js.Object
	CreationTime   string `js:"creationTime"`
	LastSignInTime string `js:"lastSignInTime"`
}

type User struct {
	*UserInfo
	EmailVerified bool          `js:"emailVerified"`
	IsAnonymous   bool          `js:"isAnonymous"`
	Metadata      *UserMetadata `js:"metadata"`
	ProviderData  []UserInfo    `js:"providerData"`
	RefreshToken  string        `js:"refreshToken"`
}

type UserCredential struct {
	*js.Object
}

type AuthCredential struct {
	*js.Object
	ProviderID string `js:"providerId"`
}

type OAuthCredential struct {
	*AuthCredential
	AccessToken string `js:"accessToken"`
	IDToken     string `js:"idToken"`
	Secret      string `js:"secret"`
}

type ConfirmationResult struct {
	*js.Object
}

type ApplicationVerifier struct {
	*js.Object
}

type AuthCodeInfo struct {
	Operation Operation `js:"operation"`
	Data      struct {
		Email     string `js:"email"`
		FromEmail string `js:"fromEmail"`
	} `js:"data"`
}

type Auth struct {
	*js.Object
	Persistence  Persistence `js:"Persistence"`
	App          *App        `js:"app"`
	CurrentUser  *User       `js:"currentUser"`
	LanguageCode string      `js:"languageCode"`
}

func (a *Auth) ApplyActionCode(code string) *Promise {
	return &Promise{Object: a.Call("applyActionCode", code)}
}

func (a *Auth) CheckActionCode(code string) *Promise {
	return &Promise{Object: a.Call("checkActionCode", code)}
}

func (a *Auth) ConfirmPasswordReset(code string, newPassword string) *Promise {
	return &Promise{Object: a.Call("confirmPasswordReset", code, newPassword)}
}

func (a *Auth) CreateUserAndRetrieveDataWithEmailAndPassword(email string, password string) *Promise {
	return &Promise{Object: a.Call("ceateUserAndRetrieveDataWithEmailAndPassword", email, password)}
}

func (a *Auth) CreateUserWithEmailAndPassword(email string, password string) *Promise {
	return &Promise{Object: a.Call("createUserWithEmailAndPassword", email, password)}
}

func (a *Auth) FetchProvidersForEmail(email string) *Promise {
	return &Promise{Object: a.Call("fetchProvidersForEmail", email)}
}

func (a *Auth) OnAuthStateChanged(args ...interface{}) *Promise {
	return &Promise{Object: a.Call("onAuthStateChanged", args)}
}

func (a *Auth) OnIDTokenChanged(args ...interface{}) *Promise {
	return &Promise{Object: a.Call("onIdTokenChanged", args)}
}

func (a *Auth) SendPasswordResetEmail(email, actionCodeSettings string) *Promise {
	return &Promise{Object: a.Call("sendPasswordResetEmail", email, actionCodeSettings)}
}

func (a *Auth) SetPersistence(persistence Persistence) *Promise {
	return &Promise{Object: a.Call("setPersistence", persistence)}
}

func (a *Auth) SignInAndRetrieveDataWithCredential(credential *AuthCredential) *Promise {
	return &Promise{Object: a.Call("signInAndRetrieveDataWithCredential", credential)}
}

func (a *Auth) SignInAndRetrieveDataWithCustomToken(token string) *Promise {
	return &Promise{Object: a.Call("signInAndRetrieveDataWithCustomToken", token)}
}

func (a *Auth) SignInAndRetrieveDataWithEmailAndPassword(email, password string) *Promise {
	return &Promise{Object: a.Call("signInAndRetrieveDataWithEmailAndPassword", email, password)}
}

func (a *Auth) SignInAnonymously() *Promise {
	return &Promise{Object: a.Call("signInAnonymously")}
}

func (a *Auth) SignInAnonymouslyAndRetrieveData() *Promise {
	return &Promise{Object: a.Call("signInAnonymouslyAndRetrieveData")}
}

func (a *Auth) SignInWithCustomToken(token string) *Promise {
	return &Promise{Object: a.Call("signInWithCustomToken", token)}
}

func (a *Auth) SignInWithEmailAndPassword(email, password string) *Promise {
	return &Promise{Object: a.Call("signInWithEmailAndPassword")}
}

func (a *Auth) SignInWithPhoneNumber(phoneNumber string, applicationVerifier *ApplicationVerifier) *Promise {
	return &Promise{Object: a.Call("signInWithPhoneNumber")}
}

func (a *Auth) SignOut() *Promise {
	return &Promise{Object: a.Call("signOut")}
}

func (a *Auth) UseDeviceLanguage() {
	a.Call("useDeviceLanguage")
}

func (a *Auth) VerifyPasswordResetCode(code string) *Promise {
	return &Promise{Object: a.Call("verifyPasswordResetCode", code)}
}
