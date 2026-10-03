package domain

// ContactType identifica o tipo do complemento da pessoa.
type ContactType string

const (
	ContactResidential ContactType = "residential"
	ContactCommercial  ContactType = "commercial"
	ContactGuardian    ContactType = "guardian" // responsável (obrigatório para menores)
	ContactParent1     ContactType = "parent_1" // filiação 1 (mãe)
	ContactParent2     ContactType = "parent_2" // filiação 2 (pai)
)

// Contact é o complemento da pessoa. Todos os tipos compartilham os mesmos
// campos; apenas o comercial preenche Work.
type Contact struct {
	Type         ContactType
	ContactName  string // nome do parente/contato
	Relationship string // parentesco, usado em guardian e parent_*
	Address      Address
	Email        string
	Phone        string
	Notes        string
	Work         *WorkInfo // somente no tipo commercial
}

// Address é o endereço de um complemento.
type Address struct {
	ZipCode  string // CEP
	Street   string
	Number   string
	District string // bairro
	City     string // município
	State    string // UF
	Country  string
}

// WorkInfo reúne os dados exclusivos do complemento comercial.
type WorkInfo struct {
	CompanyID          string // chave para o futuro cadastro de empresa
	Occupation         string // profissão
	Schedule           string // horário de trabalho
	Department         string // setor de trabalho
	RegistrationNumber string // matrícula
}
