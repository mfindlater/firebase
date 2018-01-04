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

type AdditionalUserInfo struct {
	*js.Object
	ProviderID string     `js:"providerId"`
	Profile    *js.Object `js:"profile"`
	Username   string     `js:"username"`
	IsNewUser  bool       `js:"isNewUser"`
}

type UserCredential struct {
	*js.Object
	User               *User               `js:"user"`
	Credential         *AuthCredential     `js:"credential"`
	OperationType      string              `js:"operationType"`
	AdditionalUserInfo *AdditionalUserInfo `js:"additionalUserInfo"`
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
	VerificationID string `js:"verificationId"`
}

func (c *ConfirmationResult) Confirm(verificationCode string) (*UserCredential, error) {
	o, err := (&Promise{Object: c.Call("confirm", verificationCode)}).ConvertWithResult()
	return &UserCredential{Object: o}, err
}

type ApplicationVerifier struct {
	*js.Object
	Type string `js:"type"`
}

func (a *ApplicationVerifier) Verify() (string, error) {
	o, err := (&Promise{Object: a.Call("verify")}).ConvertWithResult()
	return o.String(), err
}

type ActionCodeInfo struct {
	*js.Object
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

func (a *Auth) ApplyActionCode(code string) error {
	return (&Promise{Object: a.Call("applyActionCode", code)}).Convert()
}

func (a *Auth) CheckActionCode(code string) (*ActionCodeInfo, error) {
	o, err := (&Promise{Object: a.Call("checkActionCode", code)}).ConvertWithResult()
	return &ActionCodeInfo{Object: o}, err
}

func (a *Auth) ConfirmPasswordReset(code string, newPassword string) error {
	return (&Promise{Object: a.Call("confirmPasswordReset", code, newPassword)}).Convert()
}

func (a *Auth) CreateUserAndRetrieveDataWithEmailAndPassword(email string, password string) (*UserCredential, error) {
	o, err := (&Promise{Object: a.Call("ceateUserAndRetrieveDataWithEmailAndPassword", email, password)}).ConvertWithResult()
	return &UserCredential{Object: o}, err
}

func (a *Auth) CreateUserWithEmailAndPassword(email string, password string) (*User, error) {
	o, err := (&Promise{Object: a.Call("createUserWithEmailAndPassword", email, password)}).ConvertWithResult()
	u := &User{}
	u.UserInfo.Object = o
	return u, err
}

func (a *Auth) FetchProvidersForEmail(email string) (*js.Object, error) {
	return (&Promise{Object: a.Call("fetchProvidersForEmail", email)}).ConvertWithResult()
}

func (a *Auth) OnAuthStateChanged(next func(*User), args ...interface{}) *js.Object {
	return a.Call("onAuthStateChanged", next, args)
}

func (a *Auth) OnIDTokenChanged(next func(*User), args ...interface{}) *js.Object {
	return a.Call("onIdTokenChanged", next, args)
}

func (a *Auth) SendPasswordResetEmail(email, actionCodeSettings string) error {
	return (&Promise{Object: a.Call("sendPasswordResetEmail", email, actionCodeSettings)}).Convert()
}

func (a *Auth) SetPersistence(persistence Persistence) error {
	return (&Promise{Object: a.Call("setPersistence", persistence)}).Convert()
}

func (a *Auth) SignInAndRetrieveDataWithCredential(credential *AuthCredential) (*UserCredential, error) {
	o, err := (&Promise{Object: a.Call("signInAndRetrieveDataWithCredential", credential)}).ConvertWithResult()
	return &UserCredential{Object: o}, err
}

func (a *Auth) SignInAndRetrieveDataWithCustomToken(token string) (*UserCredential, error) {
	o, err := (&Promise{Object: a.Call("signInAndRetrieveDataWithCustomToken", token)}).ConvertWithResult()
	return &UserCredential{Object: o}, err
}

func (a *Auth) SignInAndRetrieveDataWithEmailAndPassword(email, password string) (*UserCredential, error) {
	o, err := (&Promise{Object: a.Call("signInAndRetrieveDataWithEmailAndPassword", email, password)}).ConvertWithResult()
	return &UserCredential{Object: o}, err
}

func (a *Auth) SignInAnonymously() (*User, error) {
	o, err := (&Promise{Object: a.Call("signInAnonymously")}).ConvertWithResult()
	u := &User{}
	u.UserInfo = &UserInfo{Object: o}
	return u, err
}

func (a *Auth) SignInAnonymouslyAndRetrieveData() (*UserCredential, error) {
	o, err := (&Promise{Object: a.Call("signInAnonymouslyAndRetrieveData")}).ConvertWithResult()
	return &UserCredential{Object: o}, err
}

func (a *Auth) SignInWithCustomToken(token string) (*User, error) {
	o, err := (&Promise{Object: a.Call("signInWithCustomToken", token)}).ConvertWithResult()
	u := &User{}
	u.UserInfo = &UserInfo{Object: o}
	return u, err
}

func (a *Auth) SignInWithEmailAndPassword(email, password string) (*User, error) {
	o, err := (&Promise{Object: a.Call("signInWithEmailAndPassword", email, password)}).ConvertWithResult()
	u := &User{}
	u.UserInfo = &UserInfo{Object: o}
	return u, err
}

func (a *Auth) SignInWithPhoneNumber(phoneNumber string, applicationVerifier *ApplicationVerifier) (*ConfirmationResult, error) {
	o, err := (&Promise{Object: a.Call("signInWithPhoneNumber", phoneNumber, applicationVerifier)}).ConvertWithResult()
	return &ConfirmationResult{Object: o}, err
}

func (a *Auth) SignOut() error {
	return (&Promise{Object: a.Call("signOut")}).Convert()
}

func (a *Auth) UseDeviceLanguage() {
	a.Call("useDeviceLanguage")
}

func (a *Auth) VerifyPasswordResetCode(code string) (string, error) {
	o, err := (&Promise{Object: a.Call("verifyPasswordResetCode", code)}).ConvertWithResult()
	return o.String(), err
}
