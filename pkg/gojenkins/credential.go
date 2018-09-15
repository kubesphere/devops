package gojenkins

const SSHCrenditalStaplerClass = "com.cloudbees.jenkins.plugins.sshcredentials.impl.BasicSSHUserPrivateKey"
const DirectSSHCrenditalStaplerClass = "com.cloudbees.jenkins.plugins.sshcredentials.impl.BasicSSHUserPrivateKey$DirectEntryPrivateKeySource"
const UsernamePassswordCredentialStaplerClass = "com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl"
const SecretTextCredentialStaplerClass = "org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl"
const GLOBALScope = "GLOBAL"

type CreateSshCredentialRequest struct {
	Credentials SshCredential `json:"credentials"`
}

type CreateUsernamePasswordCredentialRequest struct {
	Credentials UsernamePasswordCredential `json:"credentials"`
}

type CreateSecretTextCredentialRequest struct {
	Credentials SecretTextCredential `json:"credentials"`
}

type UsernamePasswordCredential struct {
	Scope        string `json:"scope"`
	Id           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Description  string `json:"description"`
	StaplerClass string `json:"stapler-class"`
}

type SshCredential struct {
	Scope        string           `json:"scope"`
	Id           string           `json:"id"`
	Username     string           `json:"username"`
	Passphrase   string           `json:"passphrase"`
	KeySource    PrivateKeySource `json:"privateKeySource"`
	Description  string           `json:"description"`
	StaplerClass string           `json:"stapler-class"`
}

type SecretTextCredential struct {
	Scope        string `json:"scope"`
	Id           string `json:"id"`
	Secret       string `json:"secret"`
	Description  string `json:"description"`
	StaplerClass string `json:"stapler-class"`
}

type PrivateKeySource struct {
	StaplerClass string `json:"stapler-class"`
	PrivateKey   string `json:"privateKey"`
}

type CredentialResponse struct {
	Id          string `json:"id"`
	TypeName    string `json:"typeName"`
	Fingerprint string `json:"fingerprint"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
}

func NewCreateSshCredentialRequest(id, username, passphrase, privateKey, description string) *CreateSshCredentialRequest {

	keySource := PrivateKeySource{
		StaplerClass: DirectSSHCrenditalStaplerClass,
		PrivateKey:   privateKey,
	}

	sshCredential := SshCredential{
		Scope:        GLOBALScope,
		Id:           id,
		Username:     username,
		Passphrase:   passphrase,
		KeySource:    keySource,
		Description:  description,
		StaplerClass: SSHCrenditalStaplerClass,
	}
	return &CreateSshCredentialRequest{
		Credentials: sshCredential,
	}

}

func NewCreateUsernamePasswordRequest(id, username, password, description string) *CreateUsernamePasswordCredentialRequest {
	credential := UsernamePasswordCredential{
		Scope:        GLOBALScope,
		Id:           id,
		Username:     username,
		Password:     password,
		Description:  description,
		StaplerClass: UsernamePassswordCredentialStaplerClass,
	}
	return &CreateUsernamePasswordCredentialRequest{
		Credentials: credential,
	}
}

func NewCreateSecretTextCredentialRequest(id, secret, description string) *CreateSecretTextCredentialRequest {
	credential := SecretTextCredential{
		Scope:        GLOBALScope,
		Id:           id,
		Secret:       secret,
		Description:  description,
		StaplerClass: SecretTextCredentialStaplerClass,
	}
	return &CreateSecretTextCredentialRequest{
		Credentials: credential,
	}
}

func NewSshCredential(id, username, passphrase, privateKey, description string) *SshCredential {
	keySource := PrivateKeySource{
		StaplerClass: DirectSSHCrenditalStaplerClass,
		PrivateKey:   privateKey,
	}

	return &SshCredential{
		Scope:        GLOBALScope,
		Id:           id,
		Username:     username,
		Passphrase:   passphrase,
		KeySource:    keySource,
		Description:  description,
		StaplerClass: SSHCrenditalStaplerClass,
	}
}

func NewUsernamePasswordCredential(id, username, password, description string) *UsernamePasswordCredential {
	return &UsernamePasswordCredential{
		Scope:        GLOBALScope,
		Id:           id,
		Username:     username,
		Password:     password,
		Description:  description,
		StaplerClass: UsernamePassswordCredentialStaplerClass,
	}
}

func NewSecretTextCredential(id, secret, description string) *SecretTextCredential {
	return &SecretTextCredential{
		Scope:        GLOBALScope,
		Id:           id,
		Secret:       secret,
		Description:  description,
		StaplerClass: SecretTextCredentialStaplerClass,
	}
}
